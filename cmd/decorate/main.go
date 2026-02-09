package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"ymmp-tools/ymmp"

	flag "github.com/spf13/pflag"
)

func printHelp() {
	help := `Usage: decorate [options] <file.ymmp>

字幕ファイル（DSL形式）から装飾情報を読み取り、
指定されたLayerのTimelineItemに適用します。

【処理の流れ】
  1. .ymmpファイルと字幕ファイルを読み込む
  2. 指定タイムライン×レイヤーのアイテムをFrame順にソート
  3. 字幕ファイルの各行をパースしてDecoration配列を生成
  4. 対応するTimelineItem.Decorationsを置き換え

【DSL形式】
  <em>強調</>      → 強調色
  <pos>ポジティブ</> → ポジティブ色
  <neg>ネガティブ</> → ネガティブ色

Options:
%s
Example:
  decorate -t 1 -l 7 -i subtitles.txt -o output.ymmp input.ymmp
`
	fmt.Fprintf(os.Stderr, help, flag.CommandLine.FlagUsages())
}

func main() {
	// フラグの定義
	flag.SetInterspersed(true)
	outputPath := flag.StringP("output", "o", "", "出力先ファイルパス（必須）")
	timeline := flag.IntP("timeline", "t", -1, "Timeline番号（必須）")
	layer := flag.Int64P("layer", "l", -1, "Layer番号（必須）")
	subtitlePath := flag.StringP("input", "i", "", "字幕ファイルパス（必須）")
	settingsPath := flag.StringP("settings", "s", "settings.toml", "設定ファイルパス")
	flag.Parse()

	// 引数のチェック
	args := flag.Args()
	if len(args) < 1 || *timeline < 0 || *layer < 0 || *subtitlePath == "" || *outputPath == "" {
		printHelp()
		os.Exit(1)
	}

	inputPath := args[0]

	// YMMPファイルの読み込み
	project, err := ymmp.LoadYmmp(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading ymmp: %v\n", err)
		os.Exit(1)
	}

	// Timeline番号の範囲チェック
	if *timeline >= len(project.Timelines) {
		fmt.Fprintf(os.Stderr, "Error: Timeline番号 %d は範囲外です（0-%d）\n", *timeline, len(project.Timelines)-1)
		os.Exit(1)
	}

	// 字幕ファイルの読み込み
	subtitles, err := loadSubtitleFile(*subtitlePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading subtitle file: %v\n", err)
		os.Exit(1)
	}

	// Decorationを適用
	count, err := applyDecorations(
		project.Timelines[*timeline].Items,
		*layer,
		subtitles,
		*settingsPath,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error applying decorations: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Applied decorations to %d items\n", count)

	// 出力
	if err := ymmp.SaveYmmp(project, *outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving ymmp: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Saved to: %s\n", *outputPath)
}

// loadSubtitleFile は字幕ファイルを読み込み、各行を配列として返す。
// 空行は保持する（アイテムとの対応を維持するため）。
func loadSubtitleFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// \r\n のリテラル文字列を実際の改行に変換
		line := scanner.Text()
		line = strings.ReplaceAll(line, `\r`, "\r")
		line = strings.ReplaceAll(line, `\n`, "\n")
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// applyDecorations は字幕ファイルの内容をTimelineItemに適用する。
//
// 【処理の流れ】
//  1. 指定レイヤーのアイテムのインデックスを収集
//  2. Frame順にソート
//  3. 各アイテムに対して、対応する字幕行からDecorationを生成
//  4. TimelineItem.Decorationsを置き換え
//
// 【注意】
//   - 字幕ファイルの行数とアイテム数が一致しない場合はエラー
//   - 元のスライスを直接変更する（副作用あり）
func applyDecorations(
	items []ymmp.TimelineItem,
	layer int64,
	subtitles []string,
	settingsPath string,
) (int, error) {
	// ステップ1: 指定レイヤーのアイテムのインデックスを収集
	var indices []int
	for i, item := range items {
		if item.Layer == layer {
			indices = append(indices, i)
		}
	}

	// ステップ2: Frame順にソート
	sort.Slice(indices, func(a, b int) bool {
		return items[indices[a]].Frame < items[indices[b]].Frame
	})

	// ステップ3: アイテム数と字幕行数のチェック
	if len(indices) != len(subtitles) {
		return 0, fmt.Errorf(
			"アイテム数(%d)と字幕行数(%d)が一致しません",
			len(indices), len(subtitles),
		)
	}

	// ステップ4: 各アイテムにDecorationを適用
	for i, idx := range indices {
		subtitle := subtitles[i]

		// DSLをパースしてDecoration配列を生成
		decorations, _, err := ymmp.ParseDecorationDSL(subtitle, settingsPath)
		if err != nil {
			return 0, fmt.Errorf("行 %d のパースに失敗: %v", i+1, err)
		}

		// Decorationsを置き換え
		items[idx].Decorations = decorations
	}

	return len(indices), nil
}
