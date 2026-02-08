package main

import (
	"fmt"
	"os"
	"ymmp-tools/ymmp"

	flag "github.com/spf13/pflag"
)

func main() {
	// フラグの定義
	flag.SetInterspersed(true)
	outputPath := flag.StringP("output", "o", "", "出力先ファイルパス（必須）")
	timeline := flag.IntP("timeline", "t", -1, "Timeline番号（必須）")
	layers := flag.Int64SliceP("layer", "l", nil, "Layer番号（複数指定可、必須）")
	flag.Parse()

	// 引数のチェック
	args := flag.Args()
	if len(args) < 1 || *timeline < 0 || len(*layers) == 0 || *outputPath == "" {
		fmt.Println("Usage: extend [options] <file.ymmp>")
		fmt.Println()
		fmt.Println("指定されたLayerのアイテムのLengthを次のアイテムまで伸ばします。")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("  extend -t 1 -l 7 -l 8 -l 9 -o output.ymmp input.ymmp")
		fmt.Println("  extend -t 1 -l 7,8,9 -o output.ymmp input.ymmp")
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

	// Lengthを伸ばす処理
	ymmp.ExtendTimelineItems(project.Timelines[*timeline].Items, *layers)

	// 処理結果を表示
	for _, layer := range *layers {
		items := ymmp.TimelineItems(project.Timelines[*timeline].Items).
			FilterTimelineItems(layer, ymmp.ItemNone, -1).
			SortTimelineItemsByFrame()
		fmt.Printf("Layer %d: %d items processed\n", layer, len(items))
	}

	// 出力
	if err := ymmp.SaveYmmp(project, *outputPath); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Saved to: %s\n", *outputPath)
}
