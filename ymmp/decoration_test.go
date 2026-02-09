package ymmp

import (
	"testing"
)

// ============================================================================
// ヘルパー関数: Decoration作成を簡略化
// ============================================================================

// dec は装飾ありのDecorationを作成する（簡略化用）
func dec(start, length int64, color string) Decoration {
	return Decoration{
		Start:         start,
		Length:        length,
		IsBold:        false,
		IsItalic:      false,
		Scale:         1.0,
		Font:          nil,
		Foreground:    &color,
		IsLineBreak:   false,
		HasDecoration: true,
	}
}

// plain は装飾なしのDecorationを作成する（簡略化用）
func plain(start, length int64) Decoration {
	return Decoration{
		Start:         start,
		Length:        length,
		IsBold:        false,
		IsItalic:      false,
		Scale:         1.0,
		Font:          nil,
		Foreground:    nil,
		IsLineBreak:   false,
		HasDecoration: false,
	}
}

// lineBreak は改行のDecorationを作成する（簡略化用）
func lineBreak(start int64) Decoration {
	return Decoration{
		Start:         start,
		Length:        2, // \r\n = 2文字
		IsBold:        false,
		IsItalic:      false,
		Scale:         1.0,
		Font:          nil,
		Foreground:    nil,
		IsLineBreak:   true,
		HasDecoration: false,
	}
}

// ============================================================================
// テスト: splitByLineBreak
// ============================================================================

func TestSplitByLineBreak(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []TextPart
	}{
		{
			name:  "改行なし",
			input: "あいうえお",
			expected: []TextPart{
				{Text: "あいうえお", IsLineBreak: false},
			},
		},
		{
			name:  "改行1つ",
			input: "あいう\r\nえお",
			expected: []TextPart{
				{Text: "あいう", IsLineBreak: false},
				{Text: "\r\n", IsLineBreak: true},
				{Text: "えお", IsLineBreak: false},
			},
		},
		{
			name:  "先頭に改行",
			input: "\r\nあいう",
			expected: []TextPart{
				{Text: "\r\n", IsLineBreak: true},
				{Text: "あいう", IsLineBreak: false},
			},
		},
		{
			name:  "末尾に改行",
			input: "あいう\r\n",
			expected: []TextPart{
				{Text: "あいう", IsLineBreak: false},
				{Text: "\r\n", IsLineBreak: true},
			},
		},
		{
			name:  "連続する改行",
			input: "あ\r\n\r\nい",
			expected: []TextPart{
				{Text: "あ", IsLineBreak: false},
				{Text: "\r\n", IsLineBreak: true},
				{Text: "\r\n", IsLineBreak: true},
				{Text: "い", IsLineBreak: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitByLineBreak(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("長さが異なる: got %d, want %d", len(result), len(tt.expected))
				return
			}
			for i, part := range result {
				if part.Text != tt.expected[i].Text || part.IsLineBreak != tt.expected[i].IsLineBreak {
					t.Errorf("parts[%d]: got {%q, %v}, want {%q, %v}",
						i, part.Text, part.IsLineBreak,
						tt.expected[i].Text, tt.expected[i].IsLineBreak)
				}
			}
		})
	}
}

// ============================================================================
// テスト: parseDSLToSpans
// ============================================================================

func TestParseDSLToSpans(t *testing.T) {
	colors := map[string]string{
		"em":  "#FF94F9FF",
		"pos": "#FFFFAAAA",
		"neg": "#FFCDA8EC",
	}

	tests := []struct {
		name     string
		input    string
		expected []TextSpan
	}{
		{
			name:  "タグなし",
			input: "plain text",
			expected: []TextSpan{
				{Text: "plain text", Tag: "", Color: nil},
			},
		},
		{
			name:  "emタグのみ",
			input: "<em>強調</>",
			expected: []TextSpan{
				{Text: "強調", Tag: "em", Color: ptr("#FF94F9FF")},
			},
		},
		{
			name:  "posタグのみ",
			input: "<pos>ポジティブ</>",
			expected: []TextSpan{
				{Text: "ポジティブ", Tag: "pos", Color: ptr("#FFFFAAAA")},
			},
		},
		{
			name:  "negタグのみ",
			input: "<neg>ネガティブ</>",
			expected: []TextSpan{
				{Text: "ネガティブ", Tag: "neg", Color: ptr("#FFCDA8EC")},
			},
		},
		{
			name:  "複数タグ",
			input: "<pos>魅力</>を<em>語る</>",
			expected: []TextSpan{
				{Text: "魅力", Tag: "pos", Color: ptr("#FFFFAAAA")},
				{Text: "を", Tag: "", Color: nil},
				{Text: "語る", Tag: "em", Color: ptr("#FF94F9FF")},
			},
		},
		{
			name:  "タグ内に改行",
			input: "<em>隅々まで\r\n語っていく</>",
			expected: []TextSpan{
				{Text: "隅々まで\r\n語っていく", Tag: "em", Color: ptr("#FF94F9FF")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseDSLToSpans(tt.input, colors)
			if len(result) != len(tt.expected) {
				t.Errorf("長さが異なる: got %d, want %d", len(result), len(tt.expected))
				return
			}
			for i, span := range result {
				exp := tt.expected[i]
				if span.Text != exp.Text || span.Tag != exp.Tag {
					t.Errorf("spans[%d]: got {Text:%q, Tag:%q}, want {Text:%q, Tag:%q}",
						i, span.Text, span.Tag, exp.Text, exp.Tag)
				}
			}
		})
	}
}

func ptr(s string) *string {
	return &s
}

// ============================================================================
// テスト: spansToDecorations（改行分割の確認）
// ============================================================================

func TestSpansToDecorations_LineBreakSplit(t *testing.T) {
	// 【仕様】改行をまたぐ装飾は分割される
	//
	// 入力: "<em>あいう\r\nえお</>"
	// 期待: 3つのDecorationに分割
	//   1. "あいう" (装飾あり, 3文字)
	//   2. "\r\n"   (改行, 2文字)
	//   3. "えお"   (装飾あり, 2文字)

	color := "#FF94F9FF"
	spans := []TextSpan{
		{Text: "あいう\r\nえお", Tag: "em", Color: &color},
	}

	result := spansToDecorations(spans)

	expected := []Decoration{
		dec(0, 3, color),  // "あいう" (3文字)
		lineBreak(3),      // "\r\n" (2文字, Start=3)
		dec(5, 2, color),  // "えお" (2文字, Start=5)
	}

	assertDecorations(t, result, expected)
}

func TestSpansToDecorations_MixedContent(t *testing.T) {
	// 【仕様】装飾あり/なしが混在するケース
	//
	// 入力: "<pos>魅力</>を<em>語る</>"
	// 期待: 3つのDecoration
	//   1. "魅力" (pos装飾, 2文字)
	//   2. "を"   (装飾なし, 1文字)
	//   3. "語る" (em装飾, 2文字)

	posColor := "#FFFFAAAA"
	emColor := "#FF94F9FF"
	spans := []TextSpan{
		{Text: "魅力", Tag: "pos", Color: &posColor},
		{Text: "を", Tag: "", Color: nil},
		{Text: "語る", Tag: "em", Color: &emColor},
	}

	result := spansToDecorations(spans)

	expected := []Decoration{
		dec(0, 2, posColor), // "魅力"
		plain(2, 1),         // "を"
		dec(3, 2, emColor),  // "語る"
	}

	assertDecorations(t, result, expected)
}

// ============================================================================
// テスト: 全角半角混在
// ============================================================================

func TestSpansToDecorations_MixedWidthCharacters(t *testing.T) {
	// 【仕様】全角・半角が混在しても、rune単位でカウントする
	//
	// 入力: "ABC漢字123"
	// 期待: 1つのDecoration (9文字 = 3 + 2 + 3 + 1 rune)
	//
	// 注意: Lengthはrune単位（文字数）であり、バイト数でも表示幅でもない

	spans := []TextSpan{
		{Text: "ABC漢字123", Tag: "", Color: nil},
	}

	result := spansToDecorations(spans)

	// "ABC漢字123" = 9 runes (A,B,C,漢,字,1,2,3 = 8文字... あれ、9文字?)
	// A, B, C, 漢, 字, 1, 2, 3 = 8文字
	expected := []Decoration{
		plain(0, 8), // 8 runes
	}

	assertDecorations(t, result, expected)
}

func TestSpansToDecorations_HalfWidthOnly(t *testing.T) {
	// 【仕様】半角のみのケース
	//
	// 入力: "Hello"
	// 期待: 1つのDecoration (5文字)

	color := "#FF0000FF"
	spans := []TextSpan{
		{Text: "Hello", Tag: "em", Color: &color},
	}

	result := spansToDecorations(spans)

	expected := []Decoration{
		dec(0, 5, color), // "Hello" = 5 runes
	}

	assertDecorations(t, result, expected)
}

func TestSpansToDecorations_FullWidthOnly(t *testing.T) {
	// 【仕様】全角のみのケース
	//
	// 入力: "日本語テスト"
	// 期待: 1つのDecoration (6文字)

	color := "#00FF00FF"
	spans := []TextSpan{
		{Text: "日本語テスト", Tag: "pos", Color: &color},
	}

	result := spansToDecorations(spans)

	expected := []Decoration{
		dec(0, 6, color), // "日本語テスト" = 6 runes
	}

	assertDecorations(t, result, expected)
}

// ============================================================================
// テスト: 複雑なケース（README.mdの例に近い）
// ============================================================================

func TestSpansToDecorations_ComplexCase(t *testing.T) {
	// 【仕様】README.mdの例に近い複雑なケース
	//
	// 入力: "<pos>あいうえおあいうえ</>\r\n<pos>あいうあいう</>あいうえお!!!"
	//
	// 期待される分割:
	//   1. "あいうえおあいうえ" (pos装飾, 9文字)
	//   2. "\r\n" (改行, 2文字)
	//   3. "あいうあいう" (pos装飾, 6文字)
	//   4. "あいうえお!!!" (装飾なし, 8文字)

	posColor := "#FFFFAAAA"
	spans := []TextSpan{
		{Text: "あいうえおあいうえ", Tag: "pos", Color: &posColor},
		{Text: "\r\n", Tag: "", Color: nil},
		{Text: "あいうあいう", Tag: "pos", Color: &posColor},
		{Text: "あいうえお!!!", Tag: "", Color: nil},
	}

	result := spansToDecorations(spans)

	expected := []Decoration{
		dec(0, 9, posColor),  // "あいうえおあいうえ" = 9文字
		lineBreak(9),         // "\r\n" = 2文字
		dec(11, 6, posColor), // "あいうあいう" = 6文字
		plain(17, 8),         // "あいうえお!!!" = 8文字
	}

	assertDecorations(t, result, expected)
}

// ============================================================================
// テスト: ParseDecorationDSL（統合テスト）
// ============================================================================

func TestParseDecorationDSL_Integration(t *testing.T) {
	// 【仕様】DSL文字列から直接Decoration配列を生成
	//
	// 【DSL形式】
	//   - 開始タグ: <em>, <pos>, <neg>
	//   - 閉じタグ: </> （一律この形式、入れ子なし）
	//
	// 入力: "<pos>魅力</>を<em>隅々まで\r\n語っていく</>つもりなので"
	//
	// 期待されるplainText: "魅力を隅々まで\r\n語っていくつもりなので"
	//
	// 期待されるDecoration:
	//   1. "魅力" (pos, 2文字)
	//   2. "を" (なし, 1文字)
	//   3. "隅々まで" (em, 4文字)
	//   4. "\r\n" (改行, 2文字)
	//   5. "語っていく" (em, 5文字)
	//   6. "つもりなので" (なし, 6文字)

	input := "<pos>魅力</>を<em>隅々まで\r\n語っていく</>つもりなので"

	decorations, plainText, err := ParseDecorationDSL(input, "nonexistent.toml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// プレーンテキストの確認
	expectedPlainText := "魅力を隅々まで\r\n語っていくつもりなので"
	if plainText != expectedPlainText {
		t.Errorf("plainText: got %q, want %q", plainText, expectedPlainText)
	}

	// Decorationの確認
	posColor := "#FFFFAAAA"
	emColor := "#FF94F9FF"
	expected := []Decoration{
		dec(0, 2, posColor),  // "魅力"
		plain(2, 1),          // "を"
		dec(3, 4, emColor),   // "隅々まで"
		lineBreak(7),         // "\r\n"
		dec(9, 5, emColor),   // "語っていく"
		plain(14, 6),         // "つもりなので"
	}

	assertDecorations(t, decorations, expected)
}

// ============================================================================
// ヘルパー: Decoration配列の比較
// ============================================================================

func assertDecorations(t *testing.T, got, want []Decoration) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("Decoration数が異なる: got %d, want %d", len(got), len(want))
		t.Logf("got: %+v", got)
		t.Logf("want: %+v", want)
		return
	}

	for i := range got {
		g, w := got[i], want[i]

		if g.Start != w.Start {
			t.Errorf("[%d] Start: got %d, want %d", i, g.Start, w.Start)
		}
		if g.Length != w.Length {
			t.Errorf("[%d] Length: got %d, want %d", i, g.Length, w.Length)
		}
		if g.IsLineBreak != w.IsLineBreak {
			t.Errorf("[%d] IsLineBreak: got %v, want %v", i, g.IsLineBreak, w.IsLineBreak)
		}
		if g.HasDecoration != w.HasDecoration {
			t.Errorf("[%d] HasDecoration: got %v, want %v", i, g.HasDecoration, w.HasDecoration)
		}

		// Foreground の比較（nil考慮）
		if (g.Foreground == nil) != (w.Foreground == nil) {
			t.Errorf("[%d] Foreground: got %v, want %v", i, g.Foreground, w.Foreground)
		} else if g.Foreground != nil && *g.Foreground != *w.Foreground {
			t.Errorf("[%d] Foreground: got %q, want %q", i, *g.Foreground, *w.Foreground)
		}
	}
}
