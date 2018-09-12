package util

import (
	"fmt"
	"net/url"
	"strings"
)

const UrlApiBase string = "https://open.youzan.com/api/oauthentry/%s/%s/%s"

func BuildPostParams(data map[string]string) url.Values {

	params := make(url.Values)

	for key, value := range data {
		params.Set(key, value)
	}

	return params
}

func BuildUrl(apiName, apiVersion string) (url string) {

	sl := strings.Split(apiName, ".")
	url = fmt.Sprintf(UrlApiBase, strings.Join(sl[0:len(sl)-1], "."), apiVersion, sl[len(sl)-1])

	return
}
