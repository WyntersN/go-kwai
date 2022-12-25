/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-11 13:11:18
 * @LastEditTime: 2022-11-11 13:12:48
 * @FilePath: \GoKwai\addons\open_kwai\core\result\user.go
 */
package result

type User_SellerGet struct {
	Result    int                `json:"result"`
	Msg       string             `json:"msg"`
	Code      string             `json:"code"`
	Data      User_SellerGetData `json:"data"`
	RequestID string             `json:"requestId"`
	SubMsg    string             `json:"sub_msg"`
	SubCode   string             `json:"sub_code"`
}

type User_SellerGetData struct {
	Head     string `json:"head"`
	SellerID int64  `json:"sellerId"`
	OpenID   string `json:"openId"`
	Sex      string `json:"sex"`
	Name     string `json:"name"`
	BigHead  string `json:"bigHead"`
}
