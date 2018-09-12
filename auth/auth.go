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
func GenSelfToken(clientID, clientSecret string, kdtID uint64) (selfTokenResp SelfTokenResp, err error) {

	params := make(map[string]string)
	params["grant_type"] = "silent"
	params["client_id"] = clientID
	params["client_secret"] = clientSecret
	params["kdt_id"] = strconv.FormatUint(kdtID, 10)

	resp, err := http.DefaultClient.PostForm(URLOauthToken, util.BuildPostParams(params))
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &selfTokenResp)

	return
}

// SelfTokenResp 自用型AccessToken结构体
type SelfTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
