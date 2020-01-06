package main

import (
	"fmt"
	"github.com/xu42/youzan-sdk-go"
)

func main() {

	// 获取自用型AccessToken
	resp, err := youzan.GenSelfToken("CLIENT_ID", "CLIENT_SECRET", "110")
	fmt.Println(resp, err)

	// 获取自用型Token同时获取RefreshToken
	resp1, err1 := youzan.GenSelfTokenWithRefresh("CLIENT_ID", "CLIENT_SECRET", "110")
	fmt.Println(resp1, err1)

	// 获取工具型型AccessToken
	resp2, err2 := youzan.GenToolToken("CLIENT_ID", "CLIENT_SECRET", "CODE", "URI")
	fmt.Println(resp2, err2)
}
