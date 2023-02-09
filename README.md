# go-webshell

> 使用go语言实现简单的webshell

## 1.介绍

```
使用go实现的webshell, 适配windows和linux编码问题
```

## 2.使用

在Releases中直接下载编译后的版本

查看帮助 , 默认监听端口 9991, 密码 yyds

```
webshell.exe -h
```

![image-20230209153415407](https://note-1301783483.cos.ap-nanjing.myqcloud.com/image/image-20230209153415407.png)

开启webshell

```
webshell.exe
```

执行系统命令

```
http://127.0.0.1:9991/shell?yyds=hostname
```

![image-20230209153821667](https://note-1301783483.cos.ap-nanjing.myqcloud.com/image/image-20230209153821667.png)

```
http://127.0.0.1:9991/shell?yyds=ipconfig
```

![image-20230209153906453](https://note-1301783483.cos.ap-nanjing.myqcloud.com/image/image-20230209153906453.png)

```
http://127.0.0.1:9991/shell?yyds=ahduew
```

![image-20230209153927502](https://note-1301783483.cos.ap-nanjing.myqcloud.com/image/image-20230209153927502.png)

默认访问其他页面, 返回404

```
http://127.0.0.1:9991/
```

![image-20230209153811632](https://note-1301783483.cos.ap-nanjing.myqcloud.com/image/image-20230209153811632.png)

## 3.查杀情况

![image-20230209160108613](https://note-1301783483.cos.ap-nanjing.myqcloud.com/image/image-20230209160108613.png)

## 4.免责声明🧐

本工具仅面向合法授权的企业安全建设行为，如您需要测试本工具的可用性，请自行搭建靶机环境。

在使用本工具进行检测时，您应确保该行为符合当地的法律法规，并且已经取得了足够的授权。请勿对非授权目标进行扫描。

如您在使用本工具的过程中存在任何非法行为，您需自行承担相应后果，我们将不承担任何法律及连带责任。