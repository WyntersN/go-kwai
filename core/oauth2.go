/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-07 18:05:43
 * @LastEditTime: 2022-11-22 14:20:13
 * @FilePath: \GoKwai\addons\open_kwai\core\oauth2.go
 */
package core

import "github.com/WyntersN/go-kwai/core/result"

type Oauth2 struct {
	ApiUrl string
	Store  *Store
}

func (o *Oauth2) Access_Token(code string) result.Oauth2_CodeToAccessToken {
	var (
		resultData result.Oauth2_CodeToAccessToken
		params     = map[string]string{
			"grant_type": "code",
			"app_id":     o.Store.AppKey,
			"app_secret": o.Store.AppSecret,
			"code":       code,
		}
	)

	Req().SetResult(&resultData).SetQueryParams(params).Get(o.ApiUrl + getFuncName(2))

	return resultData
}

func (o *Oauth2) Refresh_Token(refresh_token string) result.Oauth2_CodeToAccessToken {
	var (
		resultData result.Oauth2_CodeToAccessToken
		params     = map[string]string{
			"grant_type":    "refresh_token",
			"app_id":        o.Store.AppKey,
			"app_secret":    o.Store.AppSecret,
			"refresh_token": refresh_token,
		}
	)

	Req().SetResult(&resultData).SetQueryParams(params).Get(o.ApiUrl + getFuncName(2))

	return resultData
}
