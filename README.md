# SparkAPI Go 客户端

一个用于与讯飞星火认知大模型SparkAPI 交互的 Go 客户端库。

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
    "github.com/fruitbars/go-sparkapi"
    "log"
)

func main() {
    // 设置您的凭证和您要发送的消息
    appID := "your_app_id"
    apiKey := "your_api_key"
    apiSecret := "your_api_secret"
    prompt := "你好，世界！"

    // 初始化客户端
    client := sparkapi.NewClient(appID, apiKey, apiSecret)

    // 调用 API
    response, err := client.CallSpark(prompt, 0.5, 5, 150, "v1", "")
    if err != nil {
        log.Fatalf("CallSpark 调用失败: %v", err)
    }

    log.Printf("响应: %s", response)
}
```

记得将 `your_app_id`、`your_api_key` 和 `your_api_secret` 替换为您实际的 SparkAPI 凭证。

## 文档

SparkAPI 客户端库中可用的函数包括：

- `NewClient(appID, apiKey, apiSecret string) *SparkAPIClient`：创建一个新的 SparkAPIClient 实例的构造函数。
- `(*SparkAPIClient) CallSpark(prompt string, temperature float64, topk int, maxtokens int, version string, system string) (string, error)`：使用指定的参数调用 SparkAPI 并返回响应。

## 贡献

欢迎贡献！如果您有建议或发现了错误，欢迎开启一个拉取请求或问题。

## 许可证

此库根据 MIT 许可证分发，请查看 LICENSE 获取更多信息。