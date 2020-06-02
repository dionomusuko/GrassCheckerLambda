package encoder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Commit struct {
	PushedAt string `json:"pushed_at"`
}

func JsonEncoder(url string) string {
	//github.apiからリポジトリ全体のjsonをとってくる
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//jsonからpushed_atのみ抜き出す
	var commits []Commit
	var tmp string

	if err := json.Unmarshal(body, &commits); err != nil {
		log.Fatal(err)
	}
	t := time.Now().Format("2006-01-02")

	//for文を使って今日の日付と一致しているか確認
	for _, commit := range commits {
		tmp = commit.PushedAt
		if strings.Index(tmp, t) == 0 {
			return "お疲れ様でした"
			break
		}
	}
	return "まだコミットされてません\n" +
		"プライベートリポジトリにコミット済みの場合はこちらのメッセージを無視してください"
}
