/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-07 18:13:48
 * @LastEditTime: 2022-11-08 13:52:29
 * @FilePath: \GoKwai\addons\open_kwai\core\result\oauth2.go
 */
package result

type Oauth2_CodeToAccessToken struct {
	Result                int      `json:"result"`
	AccessToken           string   `json:"access_token"`
	OpenID                string   `json:"open_id"`
	ExpiresIn             int      `json:"expires_in"`
	TokenType             string   `json:"token_type"`
	RefreshToken          string   `json:"refresh_token"`
	RefreshTokenExpiresIn int      `json:"refresh_token_expires_in"`
	Scopes                []string `json:"scopes"`
}
