# Instant Messaging System
###### readme
### 功能
1. 在线网友显示
2. channel 群聊
3. 私聊
4. 给离线网友发消息不会丢失
5. 能传文件、<font color = "#dd00dd">图片</font>
6. ~~能摸鱼🐟~~
----
### 客户端
数据收发显示前端的内容。
### 服务器
给客户端发送结果。

----
### <font color = "cc2175">模块划分</font>
**main.go**：程序的入口  
**server.go**：
- 包含一个服务器类型，包含Start方法。
- 有处理连接请求的函数 HandlerConnection

**user.go**：
- 将客户端抽象到服务器端，任何发送给客户端channel的数据将会被

----
### 版本
**version v0.1**
* 增加了广播用户上线的功能

----
### BUG
* ~~客户端连接数超过4个服务器不再有响应~~
  * 需要Ctrl+C刷新控制台才会显示，实际上是有响应但没显示出来
* ~~客户端只能显示自己上线信息~~
  * 有一个广播消息的Goroutine忘记循环了，导致它只广播一遍就退出不广播了。
  <img src="D:\GoLand 2022.2.3\GoProj\Instant Messaging System\Bug2.png">
  * 特别要注意Goroutine的<font color = "#dd0000">**循环要求**</font>、<font color = "#dd0000">**阻塞要求**</font>。