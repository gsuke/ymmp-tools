package main

import (
	"bytes"
	"fmt"
	"os"
)

// UTF-8 BOM
var utf8BOM = []byte{0xEF, 0xBB, 0xBF}

// データからUTF-8 BOMを除去する
func removeBOM(data []byte) []byte {
	return bytes.TrimPrefix(data, utf8BOM)
}

// データの先頭にUTF-8 BOMを付与する
func addBOM(data []byte) []byte {
	return append(utf8BOM, data...)
}

// ファイルを読み込んでYmmp構造体に変換する
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

// Ymmp構造体をファイルに出力する（BOM付き）
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
