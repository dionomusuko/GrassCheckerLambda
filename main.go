package main

import (
	"fmt"
	"strings"
	"time"
)

func GrassChecker(commit string) string {
	t := time.Now().Format("2006-01-02")
	ans := strings.Index(commit, t)
	if ans == 0 {
		return "今日も1日お疲れ様でした"
	} else {
		return "まだコミットされてません\n" +
			"プライベートリポジトリにコミット済みの場合はこちらのメッセージを無視してください"
	}
}

func main() {
	fmt.Println(GrassChecker("2020-05-31"))
}
