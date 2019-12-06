package main

import (
	"fmt"
	"os"
	"strings"

	aw "github.com/deanishe/awgo"
)

var (
	// このプロジェクト全体で参照するaw.Workflow
	wf *aw.Workflow
)

func init() {
	// 初期化時にインスタンスを生成する
	wf = aw.New()
}

func run() {
	licenses, err := getLicenses("./credits.json")
	if err != nil {
		// wf.Fatal~~は，エラーメッセージを出力して終了させる
		wf.FatalError(err)
	}
	for _, c := range licenses {
		// ループを回して一つずつItemを作っていく
		// Itemの持つ各メソッドはItem自身を返すのでメソッドチェーンで書ける
		item := wf.NewItem(c.Name).
			// Argは選択されたときに次のコンポーネントへ渡す引数となる文字列
			Arg(c.Content).
			Subtitle(strings.Split(c.Content, "\n")[0]).
			// Valid(true)としないと選択できないので注意
			Valid(true)
		// そのアイテムがCommand+Enterで選択されたときの動作を変更する
		item.Cmd().
			Arg(c.URL).
			Subtitle(fmt.Sprintf("Open %s in browser", c.URL)).
			Valid(true)
	}
	// 与えられた文字列でアイテムをフィルタリングする
	args := os.Args
	if len(args) > 1 {
		// https://godoc.org/github.com/deanishe/awgo/fuzzy
		wf.Filter(args[1])
	}
	// 最終的に表示すべきアイテムが無かったときに表示するエラー文
	wf.WarnEmpty("No credits were found.", "Try different query.")
	// 標準出力へ最終的なJSONをプリントする
	wf.SendFeedback()
}

func main() {
	// 内部でpanic等をうまくハンドリングしてくれる
	wf.Run(run)
}
