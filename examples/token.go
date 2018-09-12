package main

import (
	"fmt"
	"youzan/auth"
)

func main() {
	// 获取自用型AccessToken
	result, err := auth.GenSelfToken("CLIENT_ID", "CLIENT_SECRET", 110)
	fmt.Println(result, result.AccessToken, err)
}
