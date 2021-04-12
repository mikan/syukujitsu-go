# syukujitsu-go

内閣府の祝日 CSV を以下の `syukujitsu.Entry` 構造体のスライスにパースします。

```go
type Entry struct {
	Year  int
	Month int
	Day   int
	Name  string
}
```

## 使い方

```bash
go get github.com/mikan/syukujitsu-go
```

```go
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
```

ここにソースがあります: [demo.go](cmd/demo/demo.go)

## 予めダウンロードしたファイルを使いたい場合

```bash
go get github.com/mikan/syukujitsu-go
curl -O https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv
```

```go
package main

import (
	"fmt"
	"time"

	"github.com/mikan/syukujitsu-go"
)

func main() {
	entries, err := syukujitsu.LoadAndParse("syukujitsu.csv")
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
```

ここにソースがあります: [demo.go](cmd/load/demo.go)

## なぜ shukujitsu ではなく syukujitsu なのですか？

ファイル名に合わせました。でも URL を良く見るとディレクトリ名は shukujitsu なんですよね。

```
https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv
```

## ライセンス

[BSD 3-clause](LICENSE)

## 作った人

[@mikan](https://github.com/mikan)
