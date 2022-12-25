/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-08 13:54:42
 * @LastEditTime: 2022-11-08 14:06:04
 * @FilePath: \GoKwai\addons\open_kwai\core\parameter\order.go
 */
package parameter

type Order_CursorList struct {
	OrderViewStatus uint   `json:"orderViewStatus"`
	PageSize        uint   `json:"pageSize"`
	CpsType         uint   `json:"cpsType"`
	Cursor          string `json:"cursor"`
	EndTime         int64  `json:"endTime"`
	BeginTime       int64  `json:"beginTime"`
	QueryType       uint   `json:"queryType"`
	Sort            uint   `json:"sort"`
}
