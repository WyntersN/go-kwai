/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-12-21 19:09:27
 * @LastEditTime: 2022-12-21 19:09:38
 * @FilePath: \GoKwai\addons\open_kwai\core\result\service.go
 */
package result

type Market_Buyer_Service_Info struct {
	Result    int                           `json:"result"`
	Msg       string                        `json:"msg"`
	Code      string                        `json:"code"`
	Data      Market_Buyer_Service_InfoData `json:"data"`
	RequestID string                        `json:"requestId"`
	SubMsg    string                        `json:"sub_msg"`
	SubCode   string                        `json:"sub_code"`
	Message   string                        `json:"message"`
}

type Market_Buyer_Service_InfoData struct {
	Authorized  bool   `json:"authorized"`
	InService   bool   `json:"inService"`
	StartTime   int64  `json:"startTime"`
	EndTime     int64  `json:"endTime"`
	PackageName string `json:"packageName"`
}
