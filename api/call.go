package api

import (
	"encoding/json"
	"errors"

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

	if response.Result != nil {
		response.Success = true
		response.Data = response.Result
	}

	var errResponse CallErrorReponse
	if response.ErrorResponse != errResponse {
		response.Code = response.ErrorResponse.Code
		response.Message = response.ErrorResponse.Msg
		response.Success = false
	}

	if !response.Success {
		err = errors.New(response.GetErrorMessage())
	}

	return

}

// CallRequest 调用接口封装结构体
type CallRequest struct {
	AccessToken string
	APIName     string
	APIVersion  string
	APIParams   map[string]string
}

// CallErrorReponse 错误响应结构
type CallErrorReponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// CallDataReponse Data类型响应结构
type CallDataReponse struct {
	Code      int                    `json:"code"`
	Data      map[string]interface{} `json:"data"`
	Message   string                 `json:"message"`
	RequestID string                 `json:"request_id"`
	Success   bool
}

// CallResponse 接口响应封装结构体
type CallResponse struct {
	Result        map[string]interface{} `json:"response"`
	ErrorResponse CallErrorReponse       `json:"error_response"`
	CallDataReponse
}

// GetErrorMessage 获取错误信息
func (resp CallResponse) GetErrorMessage() string {

	if resp.Success {
		return ""
	}

	if resp.Code == 200 {
		return ""
	}
	return resp.Message
}
