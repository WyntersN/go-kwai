<!--
 * @Descripttion: 
 * @version: 
 * @Author: Wynters
 * @Date: 2022-12-25 17:18:48
 * @LastEditTime: 2022-12-28 17:31:43
 * @FilePath: \Public_GoKwai\examples\message_decryption\README.md
-->

**消息解密用途**

```go
//初始化快手接口
kwai.NewStore(core.Store{
	AppKey:     "",
	AppSecret:  "",
	SignSecret: "",
	MsgSecret:  "", //必填 消息解密需要用到
})
//声明消息结构  下个版本将加入到kwai.msg中
type msgNoticeForm struct {
	EventID    string      `json:"eventId"`
	MsgID      string      `json:"msgId"`
	BizID      int64       `json:"bizId"`
	UserID     int64       `json:"userId"`
	OpenID     string      `json:"openId"`
	AppKey     string      `json:"appKey"`
	Event      string      `json:"event"`
	Info       string      `json:"info"`
	Status     int         `json:"status"`
	CreateTime int64       `json:"createTime"`
	UpdateTime int64       `json:"updateTime"`
	Operator   interface{} `json:"operator"`
}

if content, err := kwai.Store.AesDecryptCBC([]byte("我是消息体")); err == nil {
	//打印明文内容
	println(content)

	//解析消息数据
	var msgData msgNoticeForm
	json.Unmarshal([]byte(content), &msgData)
	//打印消息唯一标识
	println(msgData.EventID)
}

/////////////////////////////////////////////////////
///下面是iris框架使用例子
func (c *Content) irisPost(){
	//声明消息结构
	type msgNoticeForm struct {
		EventID    string      `json:"eventId"`
		MsgID      string      `json:"msgId"`
		BizID      int64       `json:"bizId"`
		UserID     int64       `json:"userId"`
		OpenID     string      `json:"openId"`
		AppKey     string      `json:"appKey"`
		Event      string      `json:"event"`
		Info       string      `json:"info"`
		Status     int         `json:"status"`
		CreateTime int64       `json:"createTime"`
		UpdateTime int64       `json:"updateTime"`
		Operator   interface{} `json:"operator"`
	}
    if body, err := c.Ctx.GetBody(); err == nil {
        if content, err :=  kwai.Store.AesDecryptCBC(body); err == nil {
			//打印明文内容
			println(content)

			//解析消息数据
			var msgData msgNoticeForm
			json.Unmarshal([]byte(content), &msgData)
			//打印消息唯一标识
			println(msgData.EventID)
        }
    }
}
```