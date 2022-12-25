/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-07 18:46:33
 * @LastEditTime: 2022-11-10 13:57:05
 * @FilePath: \GoKwai\addons\open_kwai\core\result\shop.go
 */
package result

type Shop_InfoGet struct {
	Result    int              `json:"result"`
	Msg       string           `json:"msg"`
	ErrorMsg  string           `json:"error_msg"`
	Code      string           `json:"code"`
	Data      Shop_InfoGetData `json:"data"`
	RequestID string           `json:"requestId"`
	SubMsg    string           `json:"sub_msg"`
	SubCode   string           `json:"sub_code"`
}

type Shop_InfoGetData struct {
	ShopName      string `json:"shopName"`
	ShopType      int    `json:"shopType"`
	ShopScoreInfo struct {
		ContentQualifyScoreStr    string `json:"contentQualifyScoreStr"`
		AfterSalesServiceScoreStr string `json:"afterSalesServiceScoreStr"`
		LogisticsServiceScoreStr  string `json:"logisticsServiceScoreStr"`
		ShopExpScoreStr           string `json:"shopExpScoreStr"`
		ProductQualityScoreStr    string `json:"productQualityScoreStr"`
		CustomerServiceScoreStr   string `json:"customerServiceScoreStr"`
	} `json:"shopScoreInfo"`
}
