package main

import (
	"fmt"
	"github.com/guoqchen1001/youzan-sdk-go"
)

func main() {

	resp, err := youzan.GenSelfToken("bifrost-console", "bifrost-console", "160")
	fmt.Println(resp, err)

	params := map[string]string{
		"mobile":    "13800138000",
		"fans_type": "0",
		"fans_id":   "0",
	}

	result, err := youzan.Call(resp.Data.AccessToken, "youzan.scrm.customer.get", "3.0.0", params)
	fmt.Println(result.Success, result.Result, result.Error, err)
}
