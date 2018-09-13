package main

import (
	"fmt"
	"github.com/xu42/youzan-sdk-go"
)

func main() {
	// 获取自用型AccessToken
	result, err := youzan.GenSelfToken("CLIENT_ID", "CLIENT_SECRET", 110)
	fmt.Println(result, result.AccessToken, err)
}
