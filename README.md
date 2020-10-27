# go package demo

## 使用MODULE管理第三方组件
安装方法
#### 1.首先将go的版本升级为1.11以上

#### 2.设置GO111MODULE

> GO111MODULE
> GO111MODULE有三个值：off, on和auto（默认值）。

> GO111MODULE=off，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
> GO111MODULE=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
> GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
> 当前目录在GOPATH/src之外且该目录包含go.mod文件
> 当前文件在包含go.mod文件的目录下面。

#### 2.在当前目录下，命令行运行 go mod init + 模块名称 初始化模块

```
go mod init hello
```
> 运行完之后，会在当前目录下生成一个go.mod文件，这是一个关键文件，之后的包的管理都是通过这个文件管理。

> 官方说明：除了go.mod之外，go命令还维护一个名为go.sum的文件，其中包含特定模块版本内容的预期加密哈希 
> go命令使用go.sum文件确保这些模块的未来下载检索与第一次下载相同的位，以确保项目所依赖的模块不会出现意外更改，无论是出于恶意、意外还是其他原因。 go.mod和go.sum都应检入版本控制。 
> go.sum 不需要手工维护，所以可以不用太关注。

##  goproxy
```
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```
## 查看配置

```
go env
```