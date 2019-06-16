package main

import (
	"fmt"
	"github.com/guoqchen1001/youzan-sdk-go"
)

func main() {

	// 获取自用型AccessToken
	resp, err := youzan.GenSelfToken("CLIENT_ID", "CLIENT_SECRET", "110")
	fmt.Println(resp, resp.Data.AccessToken, err)

	// 获取工具型型AccessToken
	resp1, err1 := youzan.GenToolToken("CLIENT_ID", "CLIENT_SECRET", "CODE", "URI")
	fmt.Println(resp1, resp1.Data.AccessToken, err1)
}
