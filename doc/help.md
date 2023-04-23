
单元测试

testing: warning: no tests to run
go help testfunc

A test function is one named TestXxx (where Xxx does not start with a
lower case letter) and should have the signature,

        func TestXxx(t *testing.T) { ... }


export GOPROXY=https://proxy.golang.com.cn,direct


问：go 下载慢 ，必须设置代理， 云主机 必须设置代理i板瓦工必须设置代理

 配置 GOPROXY 环境变量
export GOPROXY=https://proxy.golang.com.cn,direct
一个全球代理
为 Go 模块而生


问：go 是如何管理依赖的？

https://blog.csdn.net/weixin_39003229/article/details/97638573


## 部署运维

### 环境安装 

go：

export GOROOT=/root/bin/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/root/code

https://gitee.com/wang_cyi/goTryEverything
1.  设置开启启动

window11---任务计划程序--当计算机启动时 --D:\local\goTryEverything.exe


go mod tidy



