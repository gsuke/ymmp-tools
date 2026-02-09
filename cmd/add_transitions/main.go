package main

import (
	"fmt"
	"os"
	"ymmp-tools/ymmp"

	flag "github.com/spf13/pflag"
)

var length int64 = 6 // トランジションの長さ(フレーム数)

func main() {
	// フラグの定義
	flag.SetInterspersed(true)
	outputPath := flag.StringP("output", "o", "", "出力先ファイルパス（必須）")
	timeline := flag.IntP("timeline", "t", -1, "Timeline番号（必須）")
	srcLayer := flag.Int64P("src-layer", "s", -1, "元アイテムがあるレイヤー番号（必須）")
	dstLayer := flag.Int64P("dst-layer", "d", -1, "出力先のレイヤー番号（必須）")
	flag.Parse()

	// 引数のチェック
	args := flag.Args()
	if len(args) < 1 || *timeline < 0 || *srcLayer < 0 || *dstLayer < 0 || *outputPath == "" {
		fmt.Println("Usage: add_transitions [options] <file.ymmp>")
		fmt.Println()
		fmt.Println("指定されたレイヤーのアイテムの位置にトランジションアイテムを挿入します。")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("  add_transitions -t 0 -s 7 -d 8 -o output.ymmp input.ymmp")
		os.Exit(1)
	}

	inputPath := args[0]

	// YMMPファイルの読み込み
	project, err := ymmp.LoadYmmp(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Timeline番号の範囲チェック
	if *timeline >= len(project.Timelines) {
		fmt.Printf("Error: Timeline番号 %d は範囲外です（0-%d）\n", *timeline, len(project.Timelines)-1)
		os.Exit(1)
	}

	// 元レイヤーのアイテムを取得（Frameでソート）
	srcItems := ymmp.TimelineItems(project.Timelines[*timeline].Items).
		FilterTimelineItems(*srcLayer, ymmp.ItemNone, -1).
		SortTimelineItemsByFrame()

	if len(srcItems) == 0 {
		fmt.Printf("Error: レイヤー %d にアイテムがありません\n", *srcLayer)
		os.Exit(1)
	}

	// 各アイテムの位置にトランジションを作成
	transitionCount := 0
	for _, item := range srcItems {
		transition := ymmp.NewReelTransition(*dstLayer, item.Frame, length)
		project.Timelines[*timeline].Items = append(project.Timelines[*timeline].Items, transition)
		transitionCount++
	}

	fmt.Printf("Layer %d: %d items found\n", *srcLayer, len(srcItems))
	fmt.Printf("Layer %d: %d transitions inserted\n", *dstLayer, transitionCount)

	// 出力
	if err := ymmp.SaveYmmp(project, *outputPath); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Saved to: %s\n", *outputPath)
}
