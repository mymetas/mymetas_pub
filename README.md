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
    * go get https://github.com/mymetas/mymetas_pub.git