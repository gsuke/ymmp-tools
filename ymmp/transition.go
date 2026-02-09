package ymmp

// ============================================================================
// ヘルパー関数: Bezierとアニメーション値の生成
// ============================================================================

// newDefaultBezierPoints は標準的なベジェ曲線のポイント配列を生成する。
// YMM4のデフォルト設定に準拠した2点のベジェ曲線を返す。
//
// 【構造】
//
//	Point[0]: 開始点 (0,0) + コントロールポイント
//	Point[1]: 終了点 (1,1) + コントロールポイント
func newDefaultBezierPoints() []Point {
	return []Point{
		{
			Point:         ControlPoint1{X: 0.0, Y: 0.0},
			ControlPoint1: ControlPoint1{X: -0.3, Y: -0.3},
			ControlPoint2: ControlPoint1{X: 0.3, Y: 0.3},
		},
		{
			Point:         ControlPoint1{X: 1.0, Y: 1.0},
			ControlPoint1: ControlPoint1{X: -0.3, Y: -0.3},
			ControlPoint2: ControlPoint1{X: 0.3, Y: 0.3},
		},
	}
}

// newDefaultBezier は標準的なベジェ曲線を生成する。
func newDefaultBezier() Bezier {
	return Bezier{
		Points:      newDefaultBezierPoints(),
		IsQuadratic: false,
	}
}

// newFontSize は指定した初期値でFontSize（アニメーション可能な数値）を生成する。
//
// 【パラメータ】
//
//	value: 初期値（X座標なら0.0、不透明度なら100.0など）
//
// 【戻り値】
//
//	FontSize: アニメーションなしの状態で初期化された構造体
func newFontSize(value float64) FontSize {
	return FontSize{
		Values: []Value{
			{Value: value},
		},
		Span:          0.0,
		AnimationType: "なし",
		Bezier:        newDefaultBezier(),
	}
}

// newFontSizePtr はnewFontSizeのポインタ版。
// TimelineItemのオプショナルフィールド用。
func newFontSizePtr(value float64) *FontSize {
	fs := newFontSize(value)
	return &fs
}

// ============================================================================
// TransitionItem生成関数
// ============================================================================

// NewReelTransition は巻き取りトランジション（Reelトランジション）のTimelineItemを生成する。
//
// 【パラメータ】
//
//	layer:  配置するレイヤー番号
//	frame:  開始フレーム番号
//	length: トランジションの長さ（フレーム数）
//
// 【戻り値】
//
//	TimelineItem: BenikazuraTransitionPackAのReelトランジションとして構成されたアイテム
//
// 【デフォルト値】
//
//	Direction:     "Left"（左方向への巻き取り）
//	RotationCount: 2.0（回転数）
//	Blur:          100.0
//	EasingType:    "Circ"
//
// 【注意】
//
//	BenikazuraTransitionPackAプラグインがYMM4にインストールされている必要がある
func NewReelTransition(layer, frame, length int64) TimelineItem {
	// 共通のデフォルト値
	fadeIn := 0.0
	fadeOut := 0.0
	blend := "Normal"
	isInverted := false
	isAlwaysOnTop := false
	isZOrderEnabled := false
	isClippingWithObjectAbove := false

	// トランジションタイプ（BenikazuraTransitionPackAのReel）
	transitionType := "BenikazuraTransitionPackA.Reel.ReelTransitionPlugin, BenikazuraTransitionPackA, Version=2.0.0.0, Culture=neutral, PublicKeyToken=null"

	// Reelトランジション固有のパラメータ
	direction := "Left"
	rotationCount := 2.0

	transitionParam := TransitionParameter{
		Type:          "BenikazuraTransitionPackA.Reel.ReelTransitionParameter, BenikazuraTransitionPackA",
		Direction:     &direction,
		RotationCount: &rotationCount,
		Blur:          newFontSizePtr(100.0),
		EasingType:    "Circ",
	}

	return TimelineItem{
		Type:                      "YukkuriMovieMaker.Project.Items.TransitionItem, YukkuriMovieMaker",
		X:                         newFontSizePtr(0.0),
		Y:                         newFontSizePtr(0.0),
		Z:                         newFontSizePtr(0.0),
		Opacity:                   newFontSizePtr(100.0),
		Zoom:                      newFontSizePtr(100.0),
		Rotation:                  newFontSizePtr(0.0),
		FadeIn:                    &fadeIn,
		FadeOut:                   &fadeOut,
		IsInverted:                &isInverted,
		Blend:                     &blend,
		IsAlwaysOnTop:             &isAlwaysOnTop,
		IsZOrderEnabled:           &isZOrderEnabled,
		IsClippingWithObjectAbove: &isClippingWithObjectAbove,
		VideoEffects:              []VideoEffect{},
		TransitionType:            &transitionType,
		TransitionParameter:       &transitionParam,
		BeforeVideoEffects:        []interface{}{},
		AfterVideoEffects:         []interface{}{},
		Group:                     0,
		Frame:                     frame,
		Layer:                     layer,
		KeyFrames: KeyFrames{
			Frames: []int64{},
			Count:  0,
		},
		Length:        length,
		PlaybackRate:  100.0,
		ContentOffset: "00:00:00",
		Remark:        "",
		IsLocked:      false,
		IsHidden:      false,
	}
}
