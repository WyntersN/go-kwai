<!--
 * @Descripttion: 
 * @version: 
 * @Author: Wynters
 * @Date: 2022-12-25 17:18:48
 * @LastEditTime: 2022-12-25 18:31:38
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
import kwai "github.com/WyntersN/go-kwai"
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
	if err != nil {
		for _, v := range orderList.OrderList {
			//遍历出所有订单号
			println(v.OrderBaseInfo.Oid)
		}
	} else {
		println(err)
	}
```
# 有什么问题可以提交issues
