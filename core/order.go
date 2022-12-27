/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-08 10:21:33
 * @LastEditTime: 2022-12-25 17:41:55
 * @FilePath: \Public_GoKwai\core\order.go
 */
package core

import (
	"errors"
	"strings"

	"github.com/9ovn/go-kwai/core/result"

	"github.com/9ovn/go-kwai/core/parameter"
)

type Order struct {
	ApiUrl      string
	Store       *Store
	MethodBase  string
	AccessToken string
}

func (c *Order) Cursor_List(param parameter.Order_CursorList) (result.Order_CursorListData, error) {
	var (
		resultData result.Order_CursorList
		apiRoute   = getFuncNameEx()
		params     = reqParams{
			Method:      c.MethodBase + strings.ReplaceAll(apiRoute, "/", "."),
			AppKey:      c.Store.AppKey,
			AccessToken: c.AccessToken,
			Param:       param,
		}
	)

	resp, _ := Req().SetResult(&resultData).SetFormData(signHmacSha256(params, c.Store.SignSecret)).Post(c.ApiUrl + apiRoute)

	if resp.IsSuccess() {
		if resultData.Result != 1 {
			return resultData.Data, errors.New(resultData.ErrorMsg)
		}
		return resultData.Data, nil
	} else {
		return resultData.Data, resp.Err
	}
}

func (c *Order) Detail(oid int64) (result.Order_DetailData, error) {

	var (
		resultData result.Order_Detail
		apiRoute   = getFuncNameEx()
		params     = reqParams{
			Method:      c.MethodBase + strings.ReplaceAll(apiRoute, "/", "."),
			AppKey:      c.Store.AppKey,
			AccessToken: c.AccessToken,
			Param:       map[string]int64{"oid": oid},
		}
	)

	resp, _ := Req().SetResult(&resultData).SetFormData(signHmacSha256(params, c.Store.SignSecret)).Post(c.ApiUrl + apiRoute)

	if resp.IsSuccess() {
		if resultData.Result != 1 {
			return resultData.Data, errors.New(resultData.ErrorMsg)
		}
		return resultData.Data, nil
	} else {
		return resultData.Data, resp.Err
	}
}
