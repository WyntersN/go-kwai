/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-12-06 16:02:07
 * @LastEditTime: 2022-12-07 15:19:00
 * @FilePath: \GoKwai\addons\open_kwai\core\result\express.go
 */
package result

type Express_Subscribe_Query struct {
	Result    int                           `json:"result"`
	Msg       string                        `json:"msg"`
	ErrorMsg  string                        `json:"error_msg"`
	Code      string                        `json:"code"`
	Data      []Express_Subscribe_QueryData `json:"data"`
	RequestID string                        `json:"requestId"`
	SubMsg    string                        `json:"sub_msg"`
	SubCode   string                        `json:"sub_code"`
}

type Express_Subscribe_QueryData struct {
	AvailableQuantity int `json:"availableQuantity"`
	RecycledQuantity  int `json:"recycledQuantity"`
	SenderAddress     []struct {
		StreetName    string `json:"streetName"`
		DistrictCode  string `json:"districtCode"`
		CityName      string `json:"cityName"`
		DistrictName  string `json:"districtName"`
		CityCode      string `json:"cityCode"`
		CountryCode   string `json:"countryCode"`
		ProvinceCode  string `json:"provinceCode"`
		StreetCode    string `json:"streetCode"`
		DetailAddress string `json:"detailAddress"`
		CountryName   string `json:"countryName"`
		ProvinceName  string `json:"provinceName"`
	} `json:"senderAddress"`
	NetSiteCode        string `json:"netSiteCode"`
	AllocatedQuantity  int    `json:"allocatedQuantity"`
	ExpressCompanyCode string `json:"expressCompanyCode"`
	NetSiteName        string `json:"netSiteName"`
	UsedQuantity       int    `json:"usedQuantity"`
	ExpressCompanyType int    `json:"expressCompanyType"`
	CancelQuantity     int    `json:"cancelQuantity"`
	SettleAccount      string `json:"settleAccount"`
}

type Express_Ebill_Get struct {
	Result    int                     `json:"result"`
	Msg       string                  `json:"msg"`
	ErrorMsg  string                  `json:"error_msg"`
	Code      string                  `json:"code"`
	Data      []Express_Ebill_GetData `json:"data"`
	RequestID string                  `json:"requestId"`
	SubMsg    string                  `json:"sub_msg"`
	SubCode   string                  `json:"sub_code"`
}
type Express_Ebill_GetData struct {
	Data []struct {
		WaybillPackageCode string `json:"waybillPackageCode"`
		ParentWaybillCode  string `json:"parentWaybillCode"`
		Signature          string `json:"signature"`
		WaybillCode        string `json:"waybillCode"`
		PrintData          string `json:"printData"`
		Version            string `json:"version"`
		Key                string `json:"key"`
	} `json:"data"`
	RequestID    string `json:"requestId"`
	BaseResponse struct {
		Result  int    `json:"result"`
		Message string `json:"message"`
		Success bool   `json:"success"`
	} `json:"baseResponse"`
}
