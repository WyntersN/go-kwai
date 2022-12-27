/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-11 13:05:24
 * @LastEditTime: 2022-12-25 17:42:08
 * @FilePath: \Public_GoKwai\core\user.go
 */
package core

import (
	"errors"
	"strings"

	"github.com/WyntersN/go-kwai/core/result"
)

type User struct {
	ApiUrl      string
	Store       *Store
	MethodBase  string
	AccessToken string
}

func (c *User) Seller_Get() (result.User_SellerGetData, error) {
	var (
		resultData result.User_SellerGet
		apiRoute   = getFuncNameEx()
		params     = reqParams{
			Method:      c.MethodBase + strings.ReplaceAll(apiRoute, "/", "."),
			AppKey:      c.Store.AppKey,
			AccessToken: c.AccessToken,
		}
	)

	resp, _ := Req().SetResult(&resultData).SetFormData(signHmacSha256(params, c.Store.SignSecret)).Post(c.ApiUrl + apiRoute)

	if resp.IsSuccess() {
		if resultData.Result != 1 {
			return resultData.Data, errors.New(resultData.SubMsg)
		}
		return resultData.Data, nil
	} else {
		return resultData.Data, resp.Err
	}
}
