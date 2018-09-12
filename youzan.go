package youzan

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"youzan/util"
)

func Call(accessToken, apiName, apiVersion string, params map[string]string) (response Response, err error) {

	params["access_token"] = accessToken

	resp, err := http.DefaultClient.PostForm(util.BuildUrl(apiName, apiVersion), util.BuildPostParams(params))

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

type Response struct {
	Success bool
	Result  map[string]interface{} `json:"response"`
	Error   struct {
		Code int
		Msg  string
	} `json:"error_response"`
}
