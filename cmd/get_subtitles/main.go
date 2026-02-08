package main

import (
	"fmt"
	"os"
	"strings"
	"ymmp-tools/ymmp"

	flag "github.com/spf13/pflag"
)

func main() {
	// フラグの定義
	flag.SetInterspersed(true)
	timeline := flag.IntP("timeline", "t", -1, "Timeline番号（必須）")
	layer := flag.Int64P("layer", "l", -1, "Layer番号（必須）")
	flag.Parse()

	// 引数のチェック
	args := flag.Args()
	if len(args) < 1 || *timeline < 0 || *layer < 0 {
		fmt.Fprintln(os.Stderr, "Usage: get_subtitles [options] <file.ymmp>")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "指定されたLayerの字幕（Serif）をフレーム順に出力します。")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Example:")
		fmt.Fprintln(os.Stderr, "  get_subtitles -t 1 -l 7 input.ymmp")
		os.Exit(1)
	}

	inputPath := args[0]

	// YMMPファイルの読み込み
	project, err := ymmp.LoadYmmp(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Timeline番号の範囲チェック
	if *timeline >= len(project.Timelines) {
		fmt.Fprintf(os.Stderr, "Error: Timeline番号 %d は範囲外です（0-%d）\n", *timeline, len(project.Timelines)-1)
		os.Exit(1)
	}

	// 指定レイヤーのアイテムを取得してフレーム順にソート
	items := ymmp.TimelineItems(project.Timelines[*timeline].Items).
		FilterTimelineItems(*layer, ymmp.ItemNone, -1).
		SortTimelineItemsByFrame()

	// 字幕を出力
	for _, item := range items {
		if item.Serif != nil {
			// \r\n をリテラル文字列として出力
			escaped := strings.ReplaceAll(*item.Serif, "\r", `\r`)
			escaped = strings.ReplaceAll(escaped, "\n", `\n`)
			fmt.Println(escaped)
		}
	}
}
