package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dionomusuko/GrassChecker/encoder"
	"github.com/dionomusuko/GrassChecker/slack"
	//"github.com/joho/godotenv"
	"os"
)

func HandleLambdaEvent() error {
	//環境変数の処理
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	github_api := os.Getenv("GITHUB_API")
	str := encoder.JsonEncoder(github_api)
	slack.Send(str)
	return nil
}

func main() {
	lambda.Start(HandleLambdaEvent())
}
