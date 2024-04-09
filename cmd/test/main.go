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
