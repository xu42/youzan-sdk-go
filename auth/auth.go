package auth

import (
	"encoding/json"

	"github.com/xu42/youzan-sdk-go/util"
)

// URLOauthToken 认证Token
const URLOauthToken string = "https://open.youzanyun.com/auth/token"

// GenSelfToken 获取自用型AccessToken
func GenSelfToken(request GenSelfTokenRequest) (response GenSelfTokenResponse, err error) {
	body, err := get(request.toMap())
	err = json.Unmarshal(body, &response)
	return
}

// GenToolToken 生成工具型Token
func GenToolToken(request GenToolTokenRequest) (response GenToolTokenResponse, err error) {
	body, err := get(request.toMap())
	err = json.Unmarshal(body, &response)
	return
}

// get get
func get(data map[string]string) (body []byte, err error) {

	params := util.BuildPostParams(data)
	body, err = util.PostJSON(URLOauthToken, params)
	return
}

// GenTokenBaseRequest 获取AccessToken基本请求参数结构体
type GenTokenBaseRequest struct {
	AuthorizeType string `json:"authorize_type"`
	ClientID      string `json:"client_id"`
	ClientSecret  string `json:"client_secret"`
}

// GenSelfTokenRequest 获取自用型AccessToken请求参数结构体grant_id
type GenSelfTokenRequest struct {
	GenTokenBaseRequest
	GrantID string `json:"grant_id"`
}

// GenToolTokenRequest  获取工具型AccessToken请求参数结构体
type GenToolTokenRequest struct {
	GenTokenBaseRequest
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

// GenTokenBaseResponse 获取AccessToken基本响应参数结构体
type GenTokenBaseResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GenSelfTokenResponse 获取自用型AccessToken响应参数结构体
type GenSelfTokenResponse struct {
	GenTokenBaseResponse
	Data struct {
		AccessToken string `json:"access_token"`
		Expires     int64  `json:"expires"`
		Scope       string `json:"scope"`
	} `json:"data"`
}

// GenToolTokenResponse 获取工具型AccessToken响应参数结构体
type GenToolTokenResponse struct {
	GenTokenBaseResponse
	Data struct {
		AccessToken  string `json:"access_token"`
		Expires      int64  `json:"expires"`
		Scope        string `json:"scope"`
		TokenType    string `json:"token_type"`
		RefreshToken string `json:"refresh_token"`
	} `json:"data"`
}

func (req *GenTokenBaseRequest) toMap(grantType string) (m map[string]string) {
	m = make(map[string]string)
	m["client_secret"] = req.ClientSecret
	m["client_id"] = req.ClientID
	m["authorize_type"] = grantType
	return
}

func (req *GenSelfTokenRequest) toMap() (m map[string]string) {
	m = make(map[string]string)
	m = req.GenTokenBaseRequest.toMap("silent")
	m["grant_id"] = req.GrantID
	return
}

func (req *GenToolTokenRequest) toMap() (m map[string]string) {
	m = make(map[string]string)
	m = req.GenTokenBaseRequest.toMap("authorization_code")
	m["code"] = req.Code
	m["redirect_uri"] = req.RedirectURI
	return
}
