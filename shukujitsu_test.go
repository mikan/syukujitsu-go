package syukujitsu

import (
	"context"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	entries, err := FetchAndParse(context.Background())
	if err != nil {
		t.Fatalf("FetchAndParse() の実行に失敗しました: %e", err)
	}

	// 2021年の元日
	name, found := Search(entries, time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))
	if !found || name != "元日" {
		t.Errorf("2021/1/1 (元旦) が見つかりません: 結果=(%s,%v)", name, found)
	}

	// 2021年の元旦の次の日
	name, found = Search(entries, time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local))
	if found || len(name) > 0 {
		t.Errorf("2021/1/2 が祝日と判定されました: 結果=(%s,%v)", name, found)
	}
}
