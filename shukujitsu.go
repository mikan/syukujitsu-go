package syukujitsu

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Entry は祝日1日分の情報を保持する構造体です。
type Entry struct {
	Year  int
	Month int
	Day   int
	Name  string
}

const csvURL = "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"

// FetchAndParse は内閣府ウェブサイトから祝日 CSV を取得して Entry スライスに変換します。
func FetchAndParse(ctx context.Context) ([]Entry, error) {
	data, err := fetch(ctx, csvURL)
	if err != nil {
		return nil, err
	}
	return Parse(data)
}

// LoadAndParse は予めダウンロードしておいた祝日 CSV を Entry スライスに変換します。
func LoadAndParse(name string) ([]Entry, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("ファイル %s を読み込めませんでした: %w", name, err)
	}
	return Parse(data)
}

// Find は Entry スライス内に time.Time と一致する祝日があればその祝日名と true を、一致しなければ空文字と false を返却します。
func Find(entries []Entry, time time.Time) (name string, found bool) {
	for _, entry := range entries {
		if time.Year() == entry.Year && int(time.Month()) == entry.Month && time.Day() == entry.Day {
			return entry.Name, true
		}
	}
	return "", false
}

// Parse はバイトスライスにロードされた CSV データを Entry スライスに変換します。
func Parse(data []byte) ([]Entry, error) {
	records, err := csv.NewReader(transform.NewReader(bytes.NewReader(data), japanese.ShiftJIS.NewDecoder())).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("データの解析に失敗しました: %w", err)
	}
	var entries []Entry
	for i, row := range records {
		if i == 0 {
			continue
		}
		if len(row) != 2 {
			return nil, fmt.Errorf("想定外のデータに遭遇しました: 行 %d = %v", i+1, row)
		}
		ymd := strings.Split(row[0], "/")
		if len(ymd) != 3 {
			return nil, fmt.Errorf("年月日を認識できません: 行 %d = %v", i+1, row)
		}
		y, err := strconv.Atoi(ymd[0])
		if err != nil {
			return nil, fmt.Errorf("年月日 (年) を認識できません: 行 %d = %v", i+1, row)
		}
		m, err := strconv.Atoi(ymd[1])
		if err != nil {
			return nil, fmt.Errorf("年月日 (月) を認識できません: 行 %d = %v", i+1, row)
		}
		d, err := strconv.Atoi(ymd[2])
		if err != nil {
			return nil, fmt.Errorf("年月日 (日) を認識できません: 行 %d = %v", i+1, row)
		}
		entries = append(entries, Entry{Year: y, Month: m, Day: d, Name: row[1]})
	}
	return entries, nil
}

func fetch(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("要求の構築に失敗しました: %w", err)
	}
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, fmt.Errorf("接続に失敗しました: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("通信のクローズに失敗しました: %e", err)
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("データの取得に失敗しました: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("データの取得に失敗しました: HTTP %s %s", resp.Status, string(body))
	}
	return body, nil
}
