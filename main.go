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
	temperature := 0.7
	topk := 5
	maxtokens := 100
	version := "v1"
	system := "test system"

	response, err := client.CallSpark(prompt, temperature, topk, maxtokens, version, system)
	if err != nil {
		logger.Fatalf("Failed to call SparkAPI: %v", err)
	}

	// 输出响应
	logger.Printf("Response from SparkAPI: %s", response)
}
