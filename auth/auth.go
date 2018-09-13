package auth

import (
	"encoding/json"
	"github.com/xu42/youzan-sdk-go/util"
	"io/ioutil"
	"net/http"
	"strconv"
)

// URLOauthToken 认证Token
const URLOauthToken string = "https://open.youzan.com/oauth/token"

// GenSelfToken 获取自用型AccessToken
func GenSelfToken(request GenSelfTokenRequest) (response GenSelfTokenResponse, err error) {

	params := make(map[string]string)
	params["grant_type"] = "silent"
	params["client_id"] = request.ClientID
	params["client_secret"] = request.ClientSecret
	params["kdt_id"] = strconv.FormatUint(request.KdtID, 10)

	resp, err := http.DefaultClient.PostForm(URLOauthToken, util.BuildPostParams(params))
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)

	return
}

// GenSelfTokenRequest 获取自用型AccessToken请求参数结构体
type GenSelfTokenRequest struct {
	ClientID     string
	ClientSecret string
	KdtID        uint64
}

// GenSelfTokenResponse 获取自用型AccessToken响应参数结构体
type GenSelfTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
