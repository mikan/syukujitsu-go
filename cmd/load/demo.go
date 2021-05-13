package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mikan/syukujitsu-go"
)

func main() {
	path := flag.String("f", "syukujitsu.csv", "syukujitsu.csv へのパス")
	flag.Parse()
	entries, err := syukujitsu.LoadAndParse(*path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d 件の祝日を読み込みました\n", len(entries))
	if name, found := syukujitsu.Find(entries, time.Now()); found {
		fmt.Printf("今日は%sです！\n", name)
	} else {
		fmt.Printf("今日は祝日ではありません\n")
	}
}
