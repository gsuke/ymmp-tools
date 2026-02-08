package main

import (
	"bytes"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

// UTF-8 BOM
var utf8BOM = []byte{0xEF, 0xBB, 0xBF}

// removeBOM はデータからUTF-8 BOMを除去する
func removeBOM(data []byte) []byte {
	return bytes.TrimPrefix(data, utf8BOM)
}

// addBOM はデータの先頭にUTF-8 BOMを付与する
func addBOM(data []byte) []byte {
	return append(utf8BOM, data...)
}

// loadYmmp はファイルを読み込んでYmmp構造体に変換する
func loadYmmp(filePath string) (Ymmp, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Ymmp{}, fmt.Errorf("failed to read file: %w", err)
	}

	data = removeBOM(data)

	ymmp, err := UnmarshalYmmp(data)
	if err != nil {
		return Ymmp{}, fmt.Errorf("failed to parse YMMP: %w", err)
	}

	return ymmp, nil
}

// saveYmmp はYmmp構造体をファイルに出力する（BOM付き）
func saveYmmp(ymmp Ymmp, outputPath string) error {
	data, err := ymmp.Marshal()
	if err != nil {
		return fmt.Errorf("failed to marshal YMMP: %w", err)
	}

	data = addBOM(data)

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

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
	PrintSummary(ymmp.Timelines[1].Items[0])

	// 出力オプションが指定されている場合のみ出力
	if *outputPath != "" {
		if err := saveYmmp(ymmp, *outputPath); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Saved to: %s\n", *outputPath)
	}
}
