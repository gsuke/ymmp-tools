package ymmp

import (
	"os"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
)

// TODO: settings.tomlに対応していないかもしれない

// ============================================================================
// 設定ファイル関連
// ============================================================================

// DecorationSettings は装飾タグの色設定を保持する。
type DecorationSettings struct {
	Decoration struct {
		Colors map[string]string `toml:"colors"`
	} `toml:"decoration"`
}

// デフォルトの色設定
var defaultColors = map[string]string{
	"em":  "#FF94F9FF", // 強調
	"pos": "#FFFFAAAA", // ポジティブ
	"neg": "#FFCDA8EC", // ネガティブ
}

// LoadDecorationSettings は設定ファイルから装飾色を読み込む。
// ファイルが見つからない場合はデフォルト値を返す。
func LoadDecorationSettings(path string) (map[string]string, error) {
	// ファイルが存在しない場合はデフォルト値を返す
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return defaultColors, nil
	}

	var settings DecorationSettings
	if _, err := toml.DecodeFile(path, &settings); err != nil {
		return nil, err
	}

	// デフォルト値とマージ（設定ファイルの値を優先）
	colors := make(map[string]string)
	for k, v := range defaultColors {
		colors[k] = v
	}
	for k, v := range settings.Decoration.Colors {
		colors[k] = v
	}

	return colors, nil
}

// ============================================================================
// 型定義
// ============================================================================

// TextSpan はパース後のテキスト区間を表す。
// DSLタグを解析した結果、テキストの各区間がどの装飾を持つかを記録する。
type TextSpan struct {
	Text  string  // タグを除去した実際のテキスト
	Tag   string  // 装飾タグ名（"em", "pos", "neg", または "" で装飾なし）
	Color *string // 装飾色（nil の場合は装飾なし）
}

// ============================================================================
// メイン関数: ParseDecorationDSL
// ============================================================================

// ParseDecorationDSL はDSL形式の文字列をパースしてDecoration配列を生成する。
//
// 【処理の流れ】
//  1. 設定ファイルから装飾色を読み込む
//  2. DSL文字列をパースしてTextSpan配列を生成
//  3. TextSpan配列をDecoration配列に変換（改行で分割）
//
// 【入力例】
//
//	"<pos>魅力</>を<em>隅々まで\r\n語っていく</>つもりなので"
//
// 【出力】
//
//	[]Decoration: 各区間の装飾情報（改行をまたぐ場合は分割される）
//
// 【注意】
//   - \r\n は2文字としてカウント（CR, LFそれぞれ1文字）
//   - 改行をまたぐ装飾は分割される
func ParseDecorationDSL(input string, settingsPath string) ([]Decoration, string, error) {
	// ステップ1: 設定ファイルから色を読み込む
	colors, err := LoadDecorationSettings(settingsPath)
	if err != nil {
		return nil, "", err
	}

	// ステップ2: DSLをパースしてTextSpan配列を生成
	spans := parseDSLToSpans(input, colors)

	// ステップ3: TextSpan配列からプレーンテキストを生成
	plainText := extractPlainText(spans)

	// ステップ4: TextSpan配列をDecoration配列に変換
	decorations := spansToDecorations(spans)

	return decorations, plainText, nil
}

// ============================================================================
// DSLパーサー
// ============================================================================

// parseDSLToSpans はDSL文字列をパースしてTextSpan配列を生成する。
//
// 【処理内容】
// 正規表現でタグをマッチし、タグ内外のテキストをTextSpanとして収集する。
//
// 【DSL形式】
//   - 開始タグ: <em>, <pos>, <neg>
//   - 閉じタグ: </> （一律この形式）
//   - 入れ子はサポートしない
//
// 【例】
//
//	入力: "<pos>魅力</>を<em>語る</>"
//	出力: [
//	  {Text: "魅力", Tag: "pos", Color: "#FFFFAAAA"},
//	  {Text: "を", Tag: "", Color: nil},
//	  {Text: "語る", Tag: "em", Color: "#FF94F9FF"},
//	]
func parseDSLToSpans(input string, colors map[string]string) []TextSpan {
	// タグにマッチする正規表現
	// <tag>content</> の形式をマッチ
	// 注意: 入れ子のタグはサポートしない
	// 注意: 閉じタグは一律 </> とする
	// 注意: (?s) フラグで . が改行(\r\n)にもマッチするようにする
	tagPattern := regexp.MustCompile(`(?s)<(em|pos|neg)>(.*?)</>`)

	var spans []TextSpan
	lastEnd := 0

	// 全てのタグをマッチして処理
	matches := tagPattern.FindAllStringSubmatchIndex(input, -1)

	for _, match := range matches {
		// match[0]:match[1] = 全体のマッチ範囲
		// match[2]:match[3] = タグ名の範囲
		// match[4]:match[5] = コンテンツの範囲
		fullStart, fullEnd := match[0], match[1]
		tagStart, tagEnd := match[2], match[3]
		contentStart, contentEnd := match[4], match[5]

		// タグの前にテキストがあれば追加（装飾なし）
		if fullStart > lastEnd {
			beforeText := input[lastEnd:fullStart]
			spans = append(spans, TextSpan{
				Text:  beforeText,
				Tag:   "",
				Color: nil,
			})
		}

		// タグ内のテキストを追加（装飾あり）
		tagName := input[tagStart:tagEnd]
		content := input[contentStart:contentEnd]
		color := colors[tagName]
		spans = append(spans, TextSpan{
			Text:  content,
			Tag:   tagName,
			Color: &color,
		})

		lastEnd = fullEnd
	}

	// 最後のタグ以降のテキストがあれば追加
	if lastEnd < len(input) {
		spans = append(spans, TextSpan{
			Text:  input[lastEnd:],
			Tag:   "",
			Color: nil,
		})
	}

	return spans
}

// extractPlainText はTextSpan配列からプレーンテキスト（タグなし）を抽出する。
func extractPlainText(spans []TextSpan) string {
	var sb strings.Builder
	for _, span := range spans {
		sb.WriteString(span.Text)
	}
	return sb.String()
}

// ============================================================================
// Decoration配列への変換
// ============================================================================

// spansToDecorations はTextSpan配列をDecoration配列に変換する。
//
// 【処理内容】
// 各TextSpanを走査し、改行(\r\n)で装飾を分割しながらDecoration配列を生成する。
//
// 【改行の扱い】
//   - \r\n は2文字としてカウント（CR=1文字, LF=1文字）
//   - 改行部分は IsLineBreak=true の別Decorationとして生成
//   - 装飾が改行をまたぐ場合、改行の前後で別々のDecorationに分割
//
// 【例】
//
//	入力spans: [{Text: "あ\r\nい", Color: "#FF0000"}]
//	出力: [
//	  {Start: 0, Length: 1, Foreground: "#FF0000", IsLineBreak: false},
//	  {Start: 1, Length: 2, Foreground: nil, IsLineBreak: true},
//	  {Start: 3, Length: 1, Foreground: "#FF0000", IsLineBreak: false},
//	]
func spansToDecorations(spans []TextSpan) []Decoration {
	var decorations []Decoration
	currentPos := int64(0)

	for _, span := range spans {
		// 改行で分割
		parts := splitByLineBreak(span.Text)

		for i, part := range parts {
			if part.IsLineBreak {
				// 改行部分: IsLineBreak=true, 装飾なし
				decorations = append(decorations, Decoration{
					Start:         currentPos,
					Length:        int64(len(part.Text)), // \r\n = 2文字
					IsBold:        false,
					IsItalic:      false,
					Scale:         1.0,
					Font:          nil,
					Foreground:    nil,
					IsLineBreak:   true,
					HasDecoration: false,
				})
			} else if len(part.Text) > 0 {
				// 通常テキスト部分
				decorations = append(decorations, Decoration{
					Start:         currentPos,
					Length:        int64(len([]rune(part.Text))), // rune単位でカウント
					IsBold:        false,
					IsItalic:      false,
					Scale:         1.0,
					Font:          nil,
					Foreground:    span.Color,
					IsLineBreak:   false,
					HasDecoration: span.Color != nil,
				})
			}

			// 位置を進める
			if part.IsLineBreak {
				currentPos += int64(len(part.Text)) // \r\n は2文字
			} else {
				currentPos += int64(len([]rune(part.Text)))
			}

			_ = i // unused warning防止
		}
	}

	return decorations
}

// ============================================================================
// 改行分割
// ============================================================================

// TextPart は改行で分割されたテキストの一部を表す。
type TextPart struct {
	Text        string // テキスト内容（改行の場合は "\r\n"）
	IsLineBreak bool   // この部分が改行かどうか
}

// splitByLineBreak はテキストを改行(\r\n)で分割する。
//
// 【例】
//
//	入力: "あいう\r\nえお"
//	出力: [
//	  {Text: "あいう", IsLineBreak: false},
//	  {Text: "\r\n", IsLineBreak: true},
//	  {Text: "えお", IsLineBreak: false},
//	]
func splitByLineBreak(text string) []TextPart {
	var parts []TextPart
	remaining := text

	for {
		// \r\n を探す
		idx := strings.Index(remaining, "\r\n")
		if idx == -1 {
			// 改行なし: 残りのテキストを追加して終了
			if len(remaining) > 0 {
				parts = append(parts, TextPart{
					Text:        remaining,
					IsLineBreak: false,
				})
			}
			break
		}

		// 改行の前のテキスト
		if idx > 0 {
			parts = append(parts, TextPart{
				Text:        remaining[:idx],
				IsLineBreak: false,
			})
		}

		// 改行自体
		parts = append(parts, TextPart{
			Text:        "\r\n",
			IsLineBreak: true,
		})

		// 次の部分へ
		remaining = remaining[idx+2:]
	}

	return parts
}
