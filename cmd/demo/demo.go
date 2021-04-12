package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mikan/syukujitsu-go"
)

func main() {
	entries, err := syukujitsu.FetchAndParse(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d 件の祝日を読み込みました\n", len(entries))
	if name, found := syukujitsu.Search(entries, time.Now()); found {
		fmt.Printf("今日は%sです！\n", name)
	} else {
		fmt.Printf("今日は祝日ではありません\n")
	}
}
