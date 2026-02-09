package main

import (
	"fmt"
	"os"
	"strings"
	"ymmp-tools/ymmp"

	"github.com/mattn/go-runewidth"
	flag "github.com/spf13/pflag"
)

// runewidthの条件（全角=2, 半角=1）
var runeCondition = func() *runewidth.Condition {
	c := runewidth.NewCondition()
	c.EastAsianWidth = true
	return c
}()

func printHelp() {
	help := `Usage: get_subtitles [options] <file.ymmp>

指定されたLayerの字幕（Serif）をフレーム順に出力します。

Options:
%s
Example:
  get_subtitles -t 1 -l 7 input.ymmp
  get_subtitles -t 1 -l 7 --mark-long input.ymmp
`
	fmt.Fprintf(os.Stderr, help, flag.CommandLine.FlagUsages())
}

func main() {
	// フラグの定義
	flag.SetInterspersed(true)
	timeline := flag.IntP("timeline", "t", -1, "Timeline番号（必須）")
	layer := flag.Int64P("layer", "l", -1, "Layer番号（必須）")
	markLong := flag.Bool("mark-long", false, "幅32を超える行に [!] マークを付ける")
	flag.Parse()

	// 引数のチェック
	args := flag.Args()
	if len(args) < 1 || *timeline < 0 || *layer < 0 {
		printHelp()
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
			// \r\n を除去
			escaped := strings.ReplaceAll(*item.Serif, "\r", "")
			escaped = strings.ReplaceAll(escaped, "\n", "")

			// 文字幅が32を超えているならば、マークを付ける（オプション有効時のみ）
			prefix := ""
			if *markLong {
				width := runeCondition.StringWidth(escaped)
				if width > 32 {
					prefix = "[!]"
				}
			}

			fmt.Printf("%s%s\n", prefix, escaped)
		}
	}
}
