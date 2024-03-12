# SparkAPI Go 客户端

一个用于与讯飞星火认知大模型SparkAPI 交互的 Go 客户端库。

## 功能特性

- 支持讯飞星火大模型 API 的基本调用
- 支持自定义 API 地址和域名参数
- 支持 WebSocket 协议的安全通信 (wss)
- 支持讯飞星火大模型的 v1、v2、v3 和v3.5版本

## 安装

要使用 SparkAPI 客户端库，请首先使用 `go get` 命令进行安装：

```bash
go get github.com/fruitbars/go-sparkapi
```

## 使用方法

以下是使用此库的一个简单例子：

```go
package main

import (
	"log"
	"os"

	"github.com/fruitbars/go-sparkapi/sparkapi" // 替换为你的模块路径
)

func main() {
	// 创建一个日志记录器
	logger := log.New(os.Stdout, "SPARKAPI: ", log.LstdFlags)

	// 创建一个 SparkClient 实例
	client := sparkapi.NewSparkClient("yourAppID", "yourAPIKey", "yourAPISecret", logger, "", "")

	// 调用 SparkAPI
	prompt := "Hello, world!"
	temperature := 0.5
	topk := 4
	maxtokens := 8192
	version := "v3.5"
	system := ""

	response, err := client.CallSpark(prompt, temperature, topk, maxtokens, version, system)
	if err != nil {
		logger.Fatalf("Failed to call SparkAPI: %v", err)
	}

	// 输出响应
	logger.Printf("Response from SparkAPI: %s", response)
}
```

记得将 `your_app_id`、`your_api_key` 和 `your_api_secret` 替换为您实际的 SparkAPI 凭证。

更多关于讯飞星火大模型 API 的信息，请参考 [星火认知大模型 Web API 文档](https://www.xfyun.cn/doc/spark/Web.html)。

## 贡献

欢迎贡献！如果您有建议或发现了错误，欢迎开启一个拉取请求或问题。

## 许可证

此库根据 MIT 许可证分发，请查看 LICENSE 获取更多信息。