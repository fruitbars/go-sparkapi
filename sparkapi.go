package go_sparkapi

import (
	"context"
	"fmt"
	"github.com/iflytek/spark-ai-go/sparkai/llms/spark"
	"github.com/iflytek/spark-ai-go/sparkai/llms/spark/client/sparkclient"
	"github.com/iflytek/spark-ai-go/sparkai/messages"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	AppIdEnvVarName        = "SPARKAI_APP_ID"     //nolint:gosec
	ApiKeyEnvVarName       = "SPARKAI_API_KEY"    //nolint:gosec
	ApiSecretEnvVarName    = "SPARKAI_API_SECRET" //nolint:gosec
	SparkDomainEnvVarName  = "SPARKAI_DOMAIN"
	sparkVersionEnvVarName = "SPARKAI_API_VERSION" //nolint:gosec
	BaseURLEnvVarName      = "SPARKAI_URL"         //nolint:gosec
	organizationEnvVarName = "SPARK_ORGANIZATION"  //nolint:gosec
)

var (
	SPARK_API_KEY    string
	SPARK_API_SECRET string
	SPARK_API_BASE   string
	SPARK_APP_ID     string
	SPARK_DOMAIN     string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SPARK_API_KEY = os.Getenv(ApiKeyEnvVarName)
	SPARK_API_SECRET = os.Getenv(ApiSecretEnvVarName)
	SPARK_API_BASE = os.Getenv(BaseURLEnvVarName)
	SPARK_APP_ID = os.Getenv(AppIdEnvVarName)
	SPARK_DOMAIN = os.Getenv(SparkDomainEnvVarName)
}

type SparkChatRequest struct {
	prompt       string
	temperature  float64
	topk         int64
	maxtokens    int64
	system       string
	questionType string
}

func SparkChatWithCallback(req SparkChatRequest, callback func(messages.ChatMessage) error) (messages.ChatMessage, error) {
	_, client, err := spark.NewClient(spark.WithBaseURL(SPARK_API_BASE), spark.WithApiKey(SPARK_API_KEY), spark.WithApiSecret(SPARK_API_SECRET), spark.WithAppId(SPARK_APP_ID), spark.WithAPIDomain(SPARK_DOMAIN))
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	ctx := context.Background()
	r := &sparkclient.ChatRequest{
		Domain: &SPARK_DOMAIN,
		Messages: []messages.ChatMessage{
			&messages.GenericChatMessage{
				Role:    "user",
				Content: req.prompt,
			},
		},
		TopK:        &req.topk,
		Temperature: &req.temperature,
		MaxTokens:   &req.maxtokens,
	}
	return client.CreateChatWithCallBack(ctx, r, callback)
}

func SparkChat(req SparkChatRequest) (messages.ChatMessage, error) {
	return SparkChatWithCallback(req, nil)
}

func SparkChatSimple(prompt string) (messages.ChatMessage, error) {
	req := SparkChatRequest{
		prompt: prompt,
	}
	return SparkChatWithCallback(req, nil)
}

func SparkChatSimpleWithCallback(prompt string, callback func(messages.ChatMessage) error) (messages.ChatMessage, error) {
	req := SparkChatRequest{
		prompt: prompt,
	}
	return SparkChatWithCallback(req, callback)
}
