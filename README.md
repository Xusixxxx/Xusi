# ⚡️Xusi General Development Suite : v1.0.0 Alpha(α)

Xusi 是一个新开始的项目，没有准备开始任何商业使用。就如她的名字一样，我们会尝试将她打造成一个通用的开发套件。

 - 首先，这个项目以学习为首要目标，这会是一个重复造轮子的过程。不过，这其中包含了大量的中文注释，如果您是一名比我更“Young”的开发者，那我想这对您应该有用！
 - 再者，目前的Xusi Framework都只是仅仅实现功能，未进行任何性能及代码优化。
 - 作者`Xusixxxx`使用`Golang`这门开发语言的时间不长，并且学习过程的路子比较野生。如果有不符合编码规范的编码内容，或着有来自您的宝贵意见，希望您能够指出，我们将表示由衷感谢！
 
 <a href="https://github.com/Xusixxxx/Xusi"><img align="center" style="margin-right:20px;" src="https://s2.ax1x.com/2019/06/16/VTbJ1S.png" title="logo created by @xusixxxx" /></a>
 
 如果您希望联系到我，下面会是能够最快捷联系到我们的联系方式💖：

 - Email：xusixxxx@gmail.com
 - Wechat：L-i6ong
 - Tencent QQ：2297612526
 
 ### 更新日志
 - [Gitee](https://gitee.com/xusixxxx/Xusi/blob/master/UpdateLog.md)
 - [GitHub](https://github.com/xusixxxx/Xusi/blob/master/UpdateLog.md)
 
 # 安装
请确保安装 [Go Programming Language](https://golang.org/dl/)
```sh
$ go get -u -v github.com/Xusixxxx/Xusi
```
****

# 开发者
如果您想参与到该项目的开发，请通过上方联系方式联系我。

强烈建议在拉取源代码的文件夹命名为`xusixxxx-projects`，如果不这样，可能会产生各式各样的奇怪问题，具体原因暂不清楚。

例如：
 - 自动补全的函数应该为`func(ctx context.Context){}`
 - 但是实际上却会成这样`func(ctx interface{}){}`
 
 如果您知道原因，希望能够告知。感谢！

# 套件

截至目前，Xusi包含了如下开发套件，当然这个在之后会越来越多。


##### [xusi-framework-core](https://github.com/Xusixxxx/Xusi/tree/master/xusi-framework/core)
###### `describe`：Xusi Framework 核心，支持其他套件的正常运作
##
##### [xusi-framework-xdoc](https://github.com/Xusixxxx/Xusi/tree/master/xusi-framework/xdoc)
###### `describe`：Xusi Framework 文档，将静态文件编译到程序中。可对编码进行分析，生成文档。需搭配`xusi-item-app`使用。
##
##### [xusi-framework-xweb](https://github.com/Xusixxxx/Xusi/tree/master/xusi-framework/xweb)
###### `describe`：Xusi Framework WEB，Web开发框架
##

# 额外的资产

Xusi还包含了一些用于支撑一些特定功能实现的程序编码。


##### [xusi-item-app](https://github.com/Xusixxxx/Xusi/tree/master/xusi-item-app)
###### `describe`：Xusi Framework 核心，支持其他套件的正常运作
##

# 忽略资产

如果将代码完整的下载下来，可能还会包含以下的一些目录，这些目录没有实际意义，仅为了便利而用来存储一些无意义内容的文件。

##### [assets](https://github.com/Xusixxxx/Xusi/tree/master/assets)
###### `describe`：一些Xusi Framework或者来自Xusi的项目的资产文件
##
##### [xusi-item-test](https://github.com/Xusixxxx/Xusi/tree/master/xusi-item-test)
###### `describe`：仅用于临时调试，记录代码
##

# License

Xusi source code is licensed under the Apache Licence, Version 2.0 ([http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)).
