package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// URLAPIBase API URL
const URLAPIBase string = "https://open.youzanyun.com/api/%s/%s?access_token=%s"

// BuildPostParams 组装HTTP POST参数
func BuildPostParams(data map[string]string) []byte {
	jsonStr, _ := json.Marshal(data)
	return jsonStr
}

// BuildURL 组装接口URL
func BuildURL(apiName, apiVersion, accessToken string) (url string) {
	return fmt.Sprintf(URLAPIBase, apiName, apiVersion, accessToken)
}

// PostJSON 发起Post Json请求
func PostJSON(url string, params []byte) ([]byte, error) {

	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(params))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "YZY-Open-Client bifrost - GO")

	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, err
}
