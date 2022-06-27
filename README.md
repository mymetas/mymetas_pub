# mymetas_pub

## 1. add one cmd 
cmd: <https://gfcdn.johng.cn/pages/viewpage.action?pageId=35357521/>

```
go run main.go -h
go run main.go httpWeb3ClientVersion
```

## 2, add web3_clientVersion
```
go run main.go httpWeb3ClientVersion
```

## 3, add httpRpcMsg
```
go run main.go httpRpcWeb3ClientVersion
```

## 4, add EthAccounts
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
## 5, add ethclient
```
go run main.go ethClient
```
