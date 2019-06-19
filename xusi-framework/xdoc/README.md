# ⚡️Xusi General Development Suite

 <a href="https://github.com/Xusixxxx/Xusi"><img align="center" style="margin-right:20px;" src="http://przimj0kd.bkt.clouddn.com/logo-framework.png?e=1560849204&token=KTrMT_fnULmylWtMq0WH4htHUN74vKGMcbY1X_j-:lxR5SHgwPSNZ0XPpfYGPJyO7-8g" title="logo created by @xusixxxx" /></a>

### 介绍（Introduce）
 
 xusi-framework-xdoc用于为项目快捷搭建一个基于xweb的开发文档HTTP Server。

### 代码（Code）
在项目路径下，建立一个名为`xdoc.go`的文件，并写入`main`函数：
```sh
func main() {
    xdoc.Run("9999")
}
```
根据下方使用方法操作后访问对应端口即可看到该项目下所有被支持的开发文档。

### 使用（Using）- xusi-framework-xdoc
****
 - 确保已经在环境变量对`$GOPATH$\bin`进行了配置；
 - 将`xusi-item-app`进行编译，并将最终的文件改名为`xusi`后放置到第一步的目录中；
 - 在`xusi-framework`目录下执行`xusi -build && xdoc`；
 - 浏览器访问[`http://localhost:9999`](http://localhost:9999)即可查看`xusi-framework`完整的`API`文档，并且将是最新的。
 
****

By：@Xusixxxx

如果您希望联系到我，下面会是能够最快捷联系到我们的联系方式💖：

 - Email：xusixxxx@gmail.com
 - Wechat：L-i6ong
 - Tencent QQ：2297612526