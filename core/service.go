/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-12-21 18:59:10
 * @LastEditTime: 2022-12-25 17:41:59
 * @FilePath: \Public_GoKwai\core\service.go
 */
package core

import (
	"errors"
	"strings"

	"github.com/9ovn/go-kwai/core/result"
)

type Service struct {
	ApiUrl      string
	Store       *Store
	MethodBase  string
	AccessToken string
}

func (c *Service) Market_Buyer_Service_Info(buyerOpenId string) (result.Market_Buyer_Service_InfoData, error) {
	var (
		resultData result.Market_Buyer_Service_Info
		apiRoute   = getFuncNameEx()
		params     = reqParams{
			Method:      c.MethodBase + strings.ReplaceAll(apiRoute, "/", "."),
			AppKey:      c.Store.AppKey,
			AccessToken: c.AccessToken,
			Param:       map[string]string{"buyerOpenId": buyerOpenId},
		}
	)

	resp, _ := Req().SetResult(&resultData).SetFormData(signHmacSha256(params, c.Store.SignSecret)).Post(c.ApiUrl + apiRoute)

	if resp.IsSuccess() {
		if resultData.Result != 1 {
			return resultData.Data, errors.New(resultData.Message)
		}
		return resultData.Data, nil
	} else {
		return resultData.Data, resp.Err
	}
}
