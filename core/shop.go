/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-07 18:46:48
 * @LastEditTime: 2022-12-25 17:42:03
 * @FilePath: \Public_GoKwai\core\shop.go
 */
/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-07 18:46:48
 * @LastEditTime: 2022-11-08 11:43:47
 * @FilePath: \GoKwai\addons\kwai\core\shop.go
 */
package core

import (
	"errors"
	"strings"

	"github.com/WyntersN/GoKwai/core/result"
)

type Shop struct {
	ApiUrl      string
	Store       *Store
	MethodBase  string
	AccessToken string
}

/**
 * @summary:
 * @description: 获取店铺信息
 * @param {*} AccessToken
 * @param {string} param
 * @Author: Wynters
 */
func (c *Shop) Info_Get() (result.Shop_InfoGetData, error) {
	var (
		resultData result.Shop_InfoGet
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
			return resultData.Data, errors.New(resultData.ErrorMsg)
		}
		return resultData.Data, nil
	} else {
		return resultData.Data, resp.Err
	}
}
