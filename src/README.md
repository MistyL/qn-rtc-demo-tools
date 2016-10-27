#qlink

###简介
qlink是为连麦服务端做的一个简单的命令行工具，方便对使用demo的用户和连麦房间进行操作

###下载

**建议下载最新版本**

|版本     |支持平台|链接|
|--------|---------|----|
|qlink v1.0|Linux, Windows, Mac OSX|[下载](http://devtools.qiniu.com/qfetch-v1.4.zip)|

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