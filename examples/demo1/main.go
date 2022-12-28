/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-12-28 15:41:44
 * @LastEditTime: 2022-12-28 15:48:01
 * @FilePath: \Public_GoKwai\examples\demo1\main.go
 */
package main

import (
	"github.com/WyntersN/go-kwai"
	"github.com/WyntersN/go-kwai/core"
	"github.com/WyntersN/go-kwai/core/parameter"
)

func main() {
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
}
