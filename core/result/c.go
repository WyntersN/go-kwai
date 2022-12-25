/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-12-08 18:06:57
 * @LastEditTime: 2022-12-08 18:07:07
 * @FilePath: \GoKwai\addons\open_kwai\core\result\c.go
 */
package result

type CurrencyResult struct {
	Result    int         `json:"result"`
	Msg       string      `json:"msg"`
	ErrorMsg  string      `json:"error_msg"`
	Code      string      `json:"code"`
	Data      interface{} `json:"data"`
	RequestID string      `json:"requestId"`
	SubMsg    string      `json:"sub_msg"`
	SubCode   string      `json:"sub_code"`
}
