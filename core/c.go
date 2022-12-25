/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-07 17:38:53
 * @LastEditTime: 2022-12-13 16:35:22
 * @FilePath: \GoKwai\addons\open_kwai\core\c.go
 */
package core

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/imroc/req/v3"
)

type Store struct {
	AppKey     string
	AppSecret  string
	SignSecret string
	MsgSecret  string
}

func Req() *req.Request {

	return req.C(). // Use C() to create a client and set with chainable client settings.
			SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36").
			SetTimeout(1 * time.Minute).R() /*.EnableDumpToFile("./runtime/logs/" + getFuncName(3) + ".log")*/
}

type reqParams struct {
	AccessToken string
	AppKey      string
	Method      string
	Param       interface{}
	SignMethod  string
	Timestamp   string
	Version     string
	Sign        string
}

// 计算HmacSha256值
func signHmacSha256(params reqParams, secret string) map[string]string {

	params.Timestamp = strconv.FormatInt(time.Now().UnixMilli(), 10)
	params.SignMethod = "HMAC_SHA256"
	params.Version = "1"

	var _jps = "{}"
	if j, err := json.Marshal(params.Param); err == nil {
		_jps = string(j)
	}

	m := hmac.New(sha256.New, []byte(secret))
	m.Write([]byte("access_token=" + params.AccessToken + "&appkey=" + params.AppKey + "&method=" + params.Method + "&param=" + _jps + "&signMethod=" + params.SignMethod + "&timestamp=" + params.Timestamp + "&version=" + params.Version + "&signSecret=" + secret))
	params.Sign = base64.StdEncoding.EncodeToString(m.Sum(nil))

	return map[string]string{
		"appkey":       params.AppKey,
		"timestamp":    params.Timestamp,
		"access_token": params.AccessToken,
		"version":      params.Version,
		"param":        _jps,
		"method":       params.Method,
		"sign":         params.Sign,
		"signMethod":   params.SignMethod,
	}
}

// 获取正在运行的函数名
func getFuncName(l int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(l, pc)
	name := runtime.FuncForPC(pc[0]).Name()

	split := strings.Split(name, ".")
	return strings.ToLower(split[len(split)-1])
}

// 获取正在运行的函数名
func getFuncNameEx() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	name := runtime.FuncForPC(pc[0]).Name()

	split := strings.Split(name, ".")
	return strings.ReplaceAll(strings.ToLower(split[len(split)-1]), "_", "/")
}
