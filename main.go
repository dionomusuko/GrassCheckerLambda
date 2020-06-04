package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dionomusuko/GrassChecker/encoder"
	"github.com/dionomusuko/GrassChecker/slack"
	"os"
)

func HandleLambdaEvent() error {
	github_api := os.Getenv("GITHUB_API")
	str := encoder.JsonEncoder(github_api)
	slack.Send(str)
	return nil
}

func main() {
	lambda.Start(HandleLambdaEvent())
}
