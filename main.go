package main

import (
	"log"
	"os"

	"github.com/fruitbars/go-sparkapi/sparkapi" // 替换为你的模块路径
)

func main() {
	// 创建一个日志记录器
	logger := log.New(os.Stdout, "SPARKAPI: ", log.LstdFlags)

	testURL := ""
	testDomain := ""
	// 创建一个 SparkClient 实例
	client := sparkapi.NewSparkClient("5ddb517c", "a82640492231a85e496887ccd0d4856f", "b13daf2b954ea8f524aab17d6fd595ce", logger, testURL, testDomain)

	// 调用 SparkAPI
	prompt := "天空是什么颜色的呢？"
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
