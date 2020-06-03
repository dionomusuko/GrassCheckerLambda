package main

import (
	"encoding/json"
	"github.com/dionomusuko/GrassChecker/encoder"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/url"
	"os"
)

type payload struct {
	Text string `json:"text"`
}

func main() {
	//環境変数の処理
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	github_api := os.Getenv("GITHUB_API")
	sl_api := os.Getenv("SLACK_API")

	str := encoder.JsonEncoder(github_api)
	t, err := json.Marshal(payload{Text: str})
	if err != nil {
		log.Println(err)
	}
	resp, err := http.PostForm(sl_api, url.Values{"payload": {string(t)}})
	defer resp.Body.Close()
	log.Println("送信完了")
}
