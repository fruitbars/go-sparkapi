# SparkAPI Go 客户端

一个用于与讯飞星火认知大模型SparkAPI 交互的 Go 客户端库。

## 功能特性
基于[spark-ai-go](https://github.com/iflytek/spark-ai-go)的一个封装，使用起来更加便捷
- 支持讯飞星火大模型 API 的基本调用
- 支持自定义 API 地址和域名参数
- 支持 WebSocket 协议的安全通信 (wss)
- 支持配置API地址和Domain

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
	go_sparkapi "github.com/fruitbars/go-sparkapi"
	"github.com/iflytek/spark-ai-go/sparkai/messages"
	"log"
)

func testSparkChatSimple(prompt string) {
	resp, err := go_sparkapi.SparkChatSimple(prompt)
	if err != nil {
		return
	}

	log.Println(resp.GetContent())
}

func testSparkChatSimpleWithCallback(prompt string, callback func(messages.ChatMessage) error) {
	resp, err := go_sparkapi.SparkChatSimpleWithCallback(prompt, callback)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("完整结果", resp.GetContent())
}

func main() {
	prompt := "天空是什么颜色的呢？"
	//testSparkChatSimple(prompt)

	testSparkChatSimpleWithCallback(prompt, func(msg messages.ChatMessage) error {
		log.Println(msg.GetType(), msg.GetContent())
		return nil
	})
}

```

记得将 `your_app_id`、`your_api_key` 和 `your_api_secret` 填写在.env文件中。

更多关于讯飞星火大模型 API 的信息，请参考 [星火认知大模型 Web API 文档](https://www.xfyun.cn/doc/spark/Web.html)。

## 贡献

欢迎贡献！如果您有建议或发现了错误，欢迎开启一个拉取请求或问题。

## 许可证

此库根据 MIT 许可证分发，请查看 LICENSE 获取更多信息。