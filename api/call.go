package api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/xu42/youzan-sdk-go/util"
)

// Call 发起接口调用
func Call(request CallRequest) (response CallResponse, err error) {

	url := util.BuildURL(request.APIName, request.APIVersion, request.AccessToken)
	params := util.BuildPostParams(request.APIParams)

	fmt.Print(string(params[:]))

	body, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	err = response.Analysis()

	return

}

// Upload 调用接口，文件上传类
func Upload(request UploadRequest) (response CallResponse, err error) {

	url := util.BuildURL(request.APIName, request.APIVersion, request.AccessToken)

	body, err := util.PostMultiPart(url, request.APIParams)

	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	err = response.Analysis()

	return

}

// CallRequest 调用接口封装结构体
type CallRequest struct {
	AccessToken string
	APIName     string
	APIVersion  string
	APIParams   map[string]interface{}
}

// UploadRequest 调用上传文件型接口结构体
type UploadRequest struct {
	AccessToken string
	APIName     string
	APIVersion  string
	APIParams   map[string]io.Reader
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

// CallGWErrorResponse 通用网关错误
type CallGWErrorResponse struct {
	Code int    `json:"err_code"`
	Msg  string `json:"err_msg"`
}

// CallResponse 接口响应封装结构体
type CallResponse struct {
	Result          map[string]interface{} `json:"response"`
	ErrorResponse   CallErrorReponse       `json:"error_response"`
	GWErrorResponse CallGWErrorResponse    `json:"gw_err_resp"`
	CallDataReponse
}

// Analysis 解析处理响应内容
func (response *CallResponse) Analysis() error {

	// gw_error_response 网关错误
	var gwErrResp CallGWErrorResponse
	if response.GWErrorResponse != gwErrResp {
		return fmt.Errorf("[%d]%s", response.GWErrorResponse.Code, response.GWErrorResponse.Msg)
	}

	// err_response 响应错误
	var errResponse CallErrorReponse
	if response.ErrorResponse != errResponse {
		return fmt.Errorf("[%d]%s", response.ErrorResponse.Code, response.ErrorResponse.Msg)
	}

	// response结构，有此内容时默认调用成功，此时赋值给reponse.Data，方便统一解析
	if response.Result != nil {
		response.Data = response.Result
		response.Success = true
	}

	// code-messgae-data 结构，根据success判断成功与否
	if !response.Success {
		return fmt.Errorf("[%d]%s", response.Code, response.Message)
	}

	return nil
}
