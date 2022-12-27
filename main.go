/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-07 17:51:31
 * @LastEditTime: 2022-12-25 17:40:27
 * @FilePath: \Public_GoKwai\main.go
 */
package kwai

import (
	"github.com/WyntersN/go-kwai/core"
)

var Store *core.Store

func NewStore(data core.Store) {
	Store = core.NewStore(&core.Store{
		AppKey:     data.AppKey,
		AppSecret:  data.AppSecret,
		SignSecret: data.SignSecret,
		MsgSecret:  data.MsgSecret,
	})
}
