/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-07 17:20:08
 * @LastEditTime: 2022-12-21 19:07:41
 * @FilePath: \GoKwai\addons\open_kwai\core\d.go
 */
package core

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"strings"
)

const API_BASE_URL = "https://openapi.kwaixiaodian.com/"
const API_METHOD_BASE = "open"

func NewStore(conf *Store) *Store {
	return &Store{
		AppKey:     conf.AppKey,
		AppSecret:  conf.AppSecret,
		SignSecret: conf.SignSecret,
		MsgSecret:  conf.MsgSecret,
	}
}

func aaa() {
	///*
	// 	Go 代码
	// 使用AES/CBC解密模式进行解密，最后使用PKCS5Padding进行填充
	// //传入两个参数
	// message:带解密的密文 字节型
	// privateKey: 加密密钥 文本型

}

func (c *Store) AesDecryptCBC(base64Content []byte) (string, error) {

	var message = strings.TrimSpace(string(base64Content))

	if message == "" || c.MsgSecret == "" {
		return "", errors.New("decode param is blank")
	}
	key, err := base64.StdEncoding.DecodeString(c.MsgSecret)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", err
	}
	iv := make([]byte, aes.BlockSize)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText = pkcs5UnPadding(cipherText)
	return string(cipherText), nil
}

// 去除填充
func pkcs5UnPadding(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	if length == 0 {
		return ciphertext
	}
	unpadding := length - int(ciphertext[length-1])

	if unpadding <= 0 {
		return ciphertext
	}
	return ciphertext[:unpadding]
}

/**
 * @summary: 快手小店授权类接口
 * @description: 主要用于快手code兑换token
 * @Author: Wynters
 */
func (c *Store) Oauth2() *Oauth2 {
	var Method = getFuncName(2)
	return &Oauth2{
		ApiUrl: API_BASE_URL + Method + "/",
		Store:  c,
	}
}

/**
 * @summary: 快手小店用户类接口
 * @description:获取用户公开信息、获取商家信息
 * @Author: Wynters
 */
func (c *Store) User(AccessToken string) *User {
	var Method = getFuncName(2)
	return &User{
		ApiUrl:      API_BASE_URL + API_METHOD_BASE + "/" + Method + "/",
		MethodBase:  API_METHOD_BASE + "." + Method + ".",
		Store:       c,
		AccessToken: AccessToken,
	}
}

/**
 * @summary: 快手小店店铺类接口
 * @description:分页获取店铺授权品牌列表、校验店铺资质、获取店铺信息、通过图商poi获取快手poi详情
 * @Author: Wynters
 */
func (c *Store) Shop(AccessToken string) *Shop {
	var Method = getFuncName(2)
	return &Shop{
		ApiUrl:      API_BASE_URL + API_METHOD_BASE + "/" + Method + "/",
		MethodBase:  API_METHOD_BASE + "." + Method + ".",
		Store:       c,
		AccessToken: AccessToken,
	}
}

/**
 * @summary: 快手小店订单类接口
 * @description:商家同意买家修改地址申请、商家拒绝买家修改地址申请、商家改订单地址接口、获取卖家对应买家的订单列表、获取订单列表v2、批量解密、批量脱敏、订单详情v2、批量加密、批量获取索引串接口、获取订单列表旗帜、订单费率查询、关闭订单、获取订单费用详情、发货、物流更新、添加备注、修改订单规格
 * @Author: Wynters
 */
func (c *Store) Order(AccessToken string) *Order {
	var Method = getFuncName(2)
	return &Order{
		ApiUrl:      API_BASE_URL + API_METHOD_BASE + "/" + Method + "/",
		MethodBase:  API_METHOD_BASE + "." + Method + ".",
		Store:       c,
		AccessToken: AccessToken,
	}
}

/**
 * @summary: 快手小店快递类接口
 * @description:获取自定义区模板列表、电子面单追加子单、电子面单取消、电子面单取号、电子面单更新、获取自定义模板打印项列表、查询快递地址是否可达、获取所有标准电子面单模板、查询商家和物流商的订购关系以及面单使用情况
 * @Author: Wynters
 */
func (c *Store) Express(AccessToken string) *Express {
	var Method = getFuncName(2)
	return &Express{
		ApiUrl:      API_BASE_URL + API_METHOD_BASE + "/" + Method + "/",
		MethodBase:  API_METHOD_BASE + "." + Method + ".",
		Store:       c,
		AccessToken: AccessToken,
	}
}

/**
 * @summary: 服务市场API
 * @description:获取买家服务关系等
 * @Author: Wynters
 */
func (c *Store) Service(AccessToken string) *Service {
	var Method = getFuncName(2)
	return &Service{
		ApiUrl:      API_BASE_URL + API_METHOD_BASE + "/" + Method + "/",
		MethodBase:  API_METHOD_BASE + "." + Method + ".",
		Store:       c,
		AccessToken: AccessToken,
	}
}
