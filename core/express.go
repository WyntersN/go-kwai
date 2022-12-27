/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-08 11:18:22
 * @LastEditTime: 2022-12-08 18:16:40
 * @FilePath: \GoKwai\addons\open_kwai\core\express.go
 */
package core

import (
	"errors"
	"strings"

	"github.com/9ovn/go-kwai/core/parameter"
	"github.com/9ovn/go-kwai/core/result"
)

type Express struct {
	ApiUrl      string
	Store       *Store
	MethodBase  string
	AccessToken string
}

func (c *Express) Printer_Element_Query() (result.Order_DetailData, error) {

	var (
		resultData result.Order_Detail
		apiRoute   = getFuncNameEx()
		params     = reqParams{
			Method:      c.MethodBase + strings.ReplaceAll(apiRoute, "/", "."),
			AppKey:      c.Store.AppKey,
			AccessToken: c.AccessToken,
			Param:       "{}",
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

func (c *Express) Standard_Template_List_Get(expressCompanyCode string) (result.Order_DetailData, error) {

	var (
		resultData result.Order_Detail
		apiRoute   = getFuncNameEx()
		params     = reqParams{
			Method:      c.MethodBase + strings.ReplaceAll(apiRoute, "/", "."),
			AppKey:      c.Store.AppKey,
			AccessToken: c.AccessToken,
			Param:       map[string]string{"expressCompanyCode": expressCompanyCode},
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

func (c *Express) Subscribe_Query(expressCompanyCode string) ([]result.Express_Subscribe_QueryData, error) {

	var (
		resultData result.Express_Subscribe_Query
		apiRoute   = getFuncNameEx()
		params     = reqParams{
			Method:      c.MethodBase + strings.ReplaceAll(apiRoute, "/", "."),
			AppKey:      c.Store.AppKey,
			AccessToken: c.AccessToken,
			Param:       map[string]string{"expressCompanyCode": expressCompanyCode},
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

func (c *Express) Ebill_Get(parame parameter.Express_GetEbillOrderRequest) ([]result.Express_Ebill_GetData, error) {

	var (
		resultData result.Express_Ebill_Get
		apiRoute   = getFuncNameEx()
		params     = reqParams{
			Method:      c.MethodBase + strings.ReplaceAll(apiRoute, "/", "."),
			AppKey:      c.Store.AppKey,
			AccessToken: c.AccessToken,
			Param:       parame,
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

func (c *Express) Ebill_Cancel(expressCompanyCode, waybillCode string) (result.CurrencyResult, error) {

	var (
		resultData result.CurrencyResult
		apiRoute   = getFuncNameEx()
		params     = reqParams{
			Method:      c.MethodBase + strings.ReplaceAll(apiRoute, "/", "."),
			AppKey:      c.Store.AppKey,
			AccessToken: c.AccessToken,
			Param:       map[string]string{"expressCompanyCode": strings.ToUpper(expressCompanyCode), "waybillCode": waybillCode},
		}
	)

	resp, _ := Req().SetResult(&resultData).SetFormData(signHmacSha256(params, c.Store.SignSecret)).Post(c.ApiUrl + apiRoute)

	if resp.IsSuccess() {
		if resultData.Result != 1 {
			return resultData, errors.New(resultData.ErrorMsg)
		}
		return resultData, nil
	} else {
		return resultData, resp.Err
	}
}
