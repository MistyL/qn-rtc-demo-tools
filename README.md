#qlink

###简介
qlink是为连麦服务端做的一个简单的命令行工具，方便对使用demo的用户和连麦房间进行操作

###编译
下载代码，Mac OS可以直接用运行src下的build.sh脚本，cross_build.sh是不同OS下的编译脚本，可以用生成的二进制可运行文件后面加 -h来查看相应的命令

###下载

**建议下载最新版本**

|版本     |支持平台|链接|
|--------|---------|----|
|qlink v1.0|Linux, Windows, Mac OSX|[下载](http://devtools.qiniu.com/qlink.zip)|

###使用
该工具是一个命令行工具，需要指定相关的参数来运行。可以通过 qlink -h 来查看帮助

```
Commands for user:
          query		qlink user query <UserName>
       register		qlink user register -u <UserName> -p <Password> -r <RoomName> [-f <Flag>]
         update		qlink user update -u <UserName> [-p <Password> ] [-f <Flag>]
         delete		qlink user delete <UserName>
          login		qlink user login -u <UserName> -p <Password>

Commands for room:
          query		qlink room query <RoomName>
         create		qlink room create -r <RoomName> -u <UserName>
         delete		qlink room delete <RoomName>
          token		qlink room token -r <RoomName> -u <UserName>
```