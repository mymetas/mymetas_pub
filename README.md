# mymetas_pub
## 1. ganache start [以太坊开发（二）以太坊开发环境](https://blog.mymetas.top/posts/2022-02-15-1-eth_learn_2_eth_dev_env/#%E4%BA%94ganache-cli-%E4%BD%BF%E7%94%A8/)
```
# ganache-cli -d
# geth attach http://localhost:8545      //另一个终端登录到控制台

```
## 2. 几个测试 
* add one cmd 
cmd: <https://gfcdn.johng.cn/pages/viewpage.action?pageId=35357521/>

```
go run main.go -h
go run main.go httpWeb3ClientVersion
```

* add web3_clientVersion
```
go run main.go httpWeb3ClientVersion
```

* add httpRpcMsg
```
go run main.go httpRpcWeb3ClientVersion
```

* add EthAccounts
    * Req/Rsp 定义从 cmd文件移动到 util
    * 出现GOPATH问题（同包文件不能引用结构体Req/Rsp）
    * go get github.com/mymetas/mymetas_pub
      - ./pkg/mod/cache/download/github.com/mymetas/mymetas_pub
    * go.mod 文件启用了module模式，因此不是 GOPATH, GOROOT 的包管理模式
    * [彻底搞懂golang的GOROOT和GOPATH](https://blog.csdn.net/qq_38151401/article/details/105729884)
      - GOROOT 是系统代码，相当于JDK
      - GOPATH 是用户代码包
      - go.mod启用了module包管理模式，就不用GOPATH方式了
```
go run main.go httpRpcEthAccounts
```
*  rpc dial
```
go run main.go  ethRpcDial
```
* add ethclient
```
go run main.go ethClient
```