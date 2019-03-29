<h1 align="center"> Youzan-SDK-Go </h1>

<p align="center"> Youzan Yun client for Golang ...</p>

[![GoDoc](https://godoc.org/github.com/xu42/youzan-sdk-go?status.svg)](https://godoc.org/github.com/xu42/youzan-sdk-go)
[![Build Status](https://travis-ci.org/xu42/youzan-sdk-go.svg)](https://travis-ci.org/xu42/youzan-sdk-go)
[![release](https://img.shields.io/github/release/xu42/youzan-sdk-go.svg)](https://github.com/xu42/youzan-sdk-go/releases)
[![GitHub license](https://img.shields.io/badge/license-GNU-blue.svg)](https://raw.githubusercontent.com/xu42/youzan-sdk-go/master/LICENSE)

## NOTICE

有赞开放平台升级为有赞云，开发者需要进行迁移工作，此为兼容版本SDK，理论上可以无痛迁移(只需更改引入包版本)，但依然强烈建议你在生产环境迁移前进行充分的测试！

- [开放平台SDK 代号:carmen](../../tree/carmen)
- [兼容版本SDK 代号:compatible](../../tree/compatible)
- [有赞云版SDK 代号:bifrost](../../tree/bifrost)

## Install
```shell
go get github.com/xu42/youzan-sdk-go@compatible
```

## Usage
- Generate AccessToken, See [examples/auth/token.go](examples/auth/token.go)
- Call API, See [examples/api/call.go](examples/api/call.go)


## TODO
- [x] 支持自用型应用
- [x] 支持工具型应用
- [ ] 测试用例

## License
GNU General Public License v3.0
