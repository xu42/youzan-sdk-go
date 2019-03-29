package api

import (
	"encoding/json"
	"github.com/xu42/youzan-sdk-go/util"
)

// Call 发起接口调用
func Call(request CallRequest) (response CallResponse, err error) {

	url := util.BuildURL(request.APIName, request.APIVersion, request.AccessToken)
	params := util.BuildPostParams(request.APIParams)
	body, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	response.Success = response.Result != nil

	return
}

// CallRequest 调用接口封装结构体
type CallRequest struct {
	AccessToken string
	APIName     string
	APIVersion  string
	APIParams   map[string]string
}

// CallResponse 接口响应封装结构体
type CallResponse struct {
	Success bool
	Result  map[string]interface{} `json:"response"`
	Error   struct {
		Code int
		Msg  string
	} `json:"error_response"`
}
