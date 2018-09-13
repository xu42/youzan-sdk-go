package youzan

import (
	"github.com/xu42/youzan-sdk-go/api"
	"github.com/xu42/youzan-sdk-go/auth"
)

// GenSelfToken 生成自用型Token
func GenSelfToken(clientID, clientSecret string, kdtID uint64) (resp auth.GenSelfTokenResponse, err error) {
	return auth.GenSelfToken(auth.GenSelfTokenRequest{ClientID: clientID, ClientSecret: clientSecret, KdtID: kdtID})
}

// Call 发起接口调用
func Call(accessToken, apiName, apiVersion string, params map[string]string) (resp api.CallResponse, err error) {
	return api.Call(api.CallRequest{AccessToken: accessToken, APIName: apiName, APIVersion: apiVersion, APIParams: params})
}
