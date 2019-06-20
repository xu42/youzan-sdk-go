package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, err
}

// PostMultiPart 发起 Post MultiPart请求
func PostMultiPart(url string, params map[string]io.Reader) ([]byte, error) {

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	for key, v := range params {
		var fw io.Writer
		var err error
		// 文件类需要调用关闭接口
		if x, ok := v.(io.Closer); ok {
			defer x.Close()
		}
		// 根据文件创建
		if x, ok := v.(*os.File); ok {
			fw, err = writer.CreateFormFile(key, x.Name())
			if err != nil {
				return nil, err
			}

		} else {
			// 根据filed创建
			fw, err = writer.CreateFormField(key)
			if err != nil {
				return nil, err
			}
		}

		if _, err := io.Copy(fw, v); err != nil {
			return nil, err
		}

	}

	// 关闭closer
	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", "YZY-Open-Client bifrost - GO")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return respBody, err

}
