<!--
 * @Descripttion: 
 * @version: 
 * @Author: Wynters
 * @Date: 2022-12-25 17:18:48
 * @LastEditTime: 2022-12-28 15:59:54
 * @FilePath: \Public_GoKwai\README.md
-->
## <a name="Get-Started">Get Started</a>

**安装**

首先需要安装[Go]（需要1.16+版本），然后可以使用下面的Go命令进行安装 go-kwai:
``` sh
go get github.com/WyntersN/go-kwai
```

**导入**

导入 go-kwai 到你的代码中:

```go
import  "github.com/WyntersN/go-kwai"
```

**基本用途**

```go
//初始化快手接口
kwai.NewStore(core.Store{
	AppKey:     "",
	AppSecret:  "",
	SignSecret: "",
	MsgSecret:  "",
})

//code 兑换 快手小店token
tokenData := kwai.Store.Oauth2().Access_Token("code")

//获取订单列表
orderList, err := kwai.Store.Order(tokenData.AccessToken).Cursor_List(parameter.Order_CursorList{})
if err == nil {
	for _, v := range orderList.OrderList {
	//遍历出所有订单号
		println(v.OrderBaseInfo.Oid)
	}
} else {
	println(err.Error())
}
```
## 对照参考 快手小店官方API https://open.kwaixiaodian.com/docs/api

* 如调用 订单API->订单详情v2（open.order.detail）
* 则只需要使用命令 kwai.Store.Order("token").Detail("oid")
* 其中 kwai.Store(可以比作为open)后的Order().Detail() 与 官方中的order.detail是一一对应的，所以调用快手小店API只需要参考官网中的接口即可灵活使用本SDK。

# 有什么问题可以提交issues
