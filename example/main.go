/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-12-25 17:23:11
 * @LastEditTime: 2022-12-25 18:07:33
 * @FilePath: \Public_GoKwai\example\main.go
 */
package main

import (
	"github.com/WyntersN/GoKwai/core"
	"github.com/WyntersN/GoKwai/core/parameter"
	kwai "github.com/WyntersN/go-kwai"
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
	if err != nil {
		for _, v := range orderList.OrderList {
			//遍历出所有订单号
			println(v.OrderBaseInfo.Oid)
		}
	} else {
		println(err)
	}
}
