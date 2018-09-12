package youzan

import (
	"encoding/json"
	"github.com/xu42/youzan-sdk-go/util"
	"io/ioutil"
	"net/http"
)

// Call 发起接口调用
func Call(accessToken, apiName, apiVersion string, params map[string]string) (response Response, err error) {

	params["access_token"] = accessToken

	resp, err := http.DefaultClient.PostForm(util.BuildURL(apiName, apiVersion), util.BuildPostParams(params))

	body, err := ioutil.ReadAll(resp.Body)
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

// Response 接口响应封装结构体
type Response struct {
	Success bool
	Result  map[string]interface{} `json:"response"`
	Error   struct {
		Code int
		Msg  string
	} `json:"error_response"`
}
