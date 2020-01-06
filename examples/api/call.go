package main

import (
	"fmt"

	"github.com/xu42/youzan-sdk-go"
)

func main() {

	resp, err := youzan.GenSelfToken("CLIENT_ID", "CLIENT_SECRET", "110")
	fmt.Println(resp, err)

	params := make(map[string]interface{})
	params["mobile"] = "13800138000"
	params["fans_type"] = 0
	params["fans_id"] = 0

	result, err := youzan.Call(resp.Data.AccessToken, "youzan.scrm.customer.get", "3.0.0", params)
	fmt.Println(result, err)
}
