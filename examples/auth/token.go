package main

import (
	"fmt"
	"github.com/xu42/youzan-sdk-go"
)

func main() {

	// 获取自用型AccessToken
	resp, err := youzan.GenSelfToken("CLIENT_ID", "CLIENT_SECRET", "110")
	fmt.Println(resp, resp.AccessToken, err)

	// 获取自用型AccessToken
	resp1, err1 := youzan.GenToolToken("CLIENT_ID", "CLIENT_SECRET", "CODE", "URI")
	fmt.Println(resp1, resp1.AccessToken, err1)
}
