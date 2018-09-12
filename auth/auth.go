package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"youzan/util"
)

const UrlOauthToken string = "https://open.youzan.com/oauth/token"

// 获取自用型AccessToken
func GenSelfToken(clientId, clientSecret string, kdtId uint64) (selfTokenResp SelfTokenResp, err error) {

	params := make(map[string]string)
	params["grant_type"] = "silent"
	params["client_id"] = clientId
	params["client_secret"] = clientSecret
	params["kdt_id"] = strconv.FormatUint(kdtId, 10)

	resp, err := http.DefaultClient.PostForm(UrlOauthToken, util.BuildPostParams(params))
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

// 自用型AccessToken结构体
type SelfTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
