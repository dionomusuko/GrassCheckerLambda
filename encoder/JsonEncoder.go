package encoder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Commit struct {
	PushedAt string `json:"pushed_at"`
}

func JsonEncoder(url string) []Commit {

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
	var commits []Commit
	if err := json.Unmarshal(body, &commits); err != nil {
		log.Fatal(err)
	}
	return commits
}
