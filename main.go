package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {
	// フラグの定義
	flag.SetInterspersed(true)
	outputPath := flag.StringP("output", "o", "", "出力先ファイルパス")
	flag.Parse()

	// 引数のチェック
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: ymmp-tools [options] <file.ymmp>")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputPath := args[0]

	// YMMPファイルの読み込み
	ymmp, err := loadYmmp(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// サンプルコード、そのまま残しておく
	PrintSummary(
		FilterTimeLineItems(ymmp.Timelines[1].Items, 1, "")[0],
	)

	// 出力オプションが指定されている場合のみ出力
	if *outputPath != "" {
		if err := saveYmmp(ymmp, *outputPath); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Saved to: %s\n", *outputPath)
	}
}
