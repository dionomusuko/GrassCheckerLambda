package slack

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

type payload struct {
	Text string `json:"text"`
}

func Send(str string) error {
	sl_api := os.Getenv("SLACK_API")
	//slack_apiにpost requestを投げる
	t, err := json.Marshal(payload{Text: str})
	if err != nil {
		log.Println(err)
	}
	resp, err := http.PostForm(sl_api, url.Values{"payload": {string(t)}})
	defer resp.Body.Close()
	log.Println("送信完了")

	return nil
}
