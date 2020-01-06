package youzan

import (
	"io"

	"github.com/xu42/youzan-sdk-go/api"
	"github.com/xu42/youzan-sdk-go/auth"
)

// GenSelfToken 获取自用型AccessToken
func GenSelfToken(clientID, clientSecret, grantID string) (resp auth.GenSelfTokenResponse, err error) {
	return auth.GenSelfToken(auth.GenSelfTokenRequest{
		GenTokenBaseRequest: auth.GenTokenBaseRequest{ClientID: clientID, ClientSecret: clientSecret, AuthorizeType: "silent"},
		GrantID:             grantID})
}

// GenSelfTokenWithRefresh 获取自用型AccessToken同时获取RefreshToken
func GenSelfTokenWithRefresh(clientID, clientSecret, grantID string) (resp auth.GenSelfTokenResponse, err error) {
	return auth.GenSelfToken(auth.GenSelfTokenRequest{
		GenTokenBaseRequest: auth.GenTokenBaseRequest{ClientID: clientID, ClientSecret: clientSecret, AuthorizeType: "silent"},
		GrantID:             grantID,
		Refresh:             true})
}

// GenToolToken 获取工具型Token
func GenToolToken(clientID, clientSecret, code, redirectURL string) (resp auth.GenToolTokenResponse, err error) {
	return auth.GenToolToken(auth.GenToolTokenRequest{
		GenTokenBaseRequest: auth.GenTokenBaseRequest{ClientID: clientID, ClientSecret: clientSecret, AuthorizeType: "authorization_code"},
		Code:                code, RedirectURI: redirectURL})
}

// Call 发起接口调用
func Call(accessToken, apiName, apiVersion string, params map[string]interface{}) (resp api.CallResponse, err error) {
	return api.Call(api.CallRequest{AccessToken: accessToken, APIName: apiName, APIVersion: apiVersion, APIParams: params})
}

// Upload 发起接口调用-文件上传类接口
func Upload(accessToken, apiName, apiVersion string, params map[string]io.Reader) (resp api.CallResponse, err error) {

	var request api.UploadRequest
	request.APIName = apiName
	request.APIParams = params
	request.APIVersion = apiVersion
	request.AccessToken = accessToken
	return api.Upload(request)

}
