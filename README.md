<h1 align="center"> Youzan-SDK-Go </h1>

<p align="center"> Youzan Yun client for Golang ...</p>

[![GoDoc](https://godoc.org/github.com/xu42/youzan-sdk-go?status.svg)](https://godoc.org/github.com/xu42/youzan-sdk-go)
[![Build Status](https://travis-ci.org/xu42/youzan-sdk-go.svg)](https://travis-ci.org/xu42/youzan-sdk-go)
[![release](https://img.shields.io/github/release/xu42/youzan-sdk-go.svg)](https://github.com/xu42/youzan-sdk-go/releases)
[![GitHub license](https://img.shields.io/badge/license-GNU-blue.svg)](https://raw.githubusercontent.com/xu42/youzan-sdk-go/master/LICENSE)


## NOTICE

有赞开放平台升级为有赞云，开发者需要进行迁移工作，此为有赞云版本SDK。如果你之前未接入有赞云，建议直接使用本版本进行业务开发。如果已接入开放平台需迁移，可以使用兼容版本也可以直接使用本版本进行开发。

- [开放平台SDK 代号:carmen](../../tree/carmen)
- [兼容版本SDK 代号:compatible](../../tree/compatible)
- [有赞云版SDK 代号:bifrost](../../tree/bifrost)


## Install
```shell
go get github.com/xu42/youzan-sdk-go
or
go get github.com/xu42/youzan-sdk-go@bifrost
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
