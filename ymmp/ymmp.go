// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    ymmp, err := UnmarshalYmmp(bytes)
//    bytes, err = ymmp.Marshal()

package ymmp

import (
	"bytes"
	"encoding/json"
	"errors"
)

func UnmarshalYmmp(data []byte) (Ymmp, error) {
	var r Ymmp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Ymmp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Ymmp struct {
	FilePath              string               `json:"FilePath"`
	SelectedTimelineIndex int64                `json:"SelectedTimelineIndex"`
	Timelines             []Timeline           `json:"Timelines"`
	Characters            []Character          `json:"Characters"`
	CollapsedGroups       []interface{}        `json:"CollapsedGroups"`
	LayoutXML             string               `json:"LayoutXml"`
	ToolStates            map[string]ToolState `json:"ToolStates"`
}

type Character struct {
	Name                             string                       `json:"Name"`
	GroupName                        string                       `json:"GroupName"`
	Color                            string                       `json:"Color"`
	Layer                            int64                        `json:"Layer"`
	KeyGesture                       KeyGesture                   `json:"KeyGesture"`
	Voice                            Voice                        `json:"Voice"`
	Volume                           FontSize                     `json:"Volume"`
	Pan                              FontSize                     `json:"Pan"`
	PlaybackRate                     float64                      `json:"PlaybackRate"`
	VoiceParameter                   VoiceParameter               `json:"VoiceParameter"`
	ContentOffset                    string                       `json:"ContentOffset"`
	AdditionalTime                   float64                      `json:"AdditionalTime"`
	VoiceFadeIn                      float64                      `json:"VoiceFadeIn"`
	VoiceFadeOut                     float64                      `json:"VoiceFadeOut"`
	EchoIsEnabled                    bool                         `json:"EchoIsEnabled"`
	EchoInterval                     float64                      `json:"EchoInterval"`
	EchoAttenuation                  float64                      `json:"EchoAttenuation"`
	CustomVoiceIsEnabled             bool                         `json:"CustomVoiceIsEnabled"`
	CustomVoiceIncludeSubdirectories bool                         `json:"CustomVoiceIncludeSubdirectories"`
	CustomVoiceDirectory             string                       `json:"CustomVoiceDirectory"`
	CustomVoiceFileName              string                       `json:"CustomVoiceFileName"`
	AudioEffects                     []CharacterAudioEffect       `json:"AudioEffects"`
	IsJimakuVisible                  bool                         `json:"IsJimakuVisible"`
	IsJimakuLocked                   bool                         `json:"IsJimakuLocked"`
	X                                FontSize                     `json:"X"`
	Y                                FontSize                     `json:"Y"`
	Z                                FontSize                     `json:"Z"`
	Opacity                          FontSize                     `json:"Opacity"`
	Zoom                             FontSize                     `json:"Zoom"`
	Rotation                         FontSize                     `json:"Rotation"`
	JimakuFadeIn                     float64                      `json:"JimakuFadeIn"`
	JimakuFadeOut                    float64                      `json:"JimakuFadeOut"`
	Blend                            string                       `json:"Blend"`
	IsInverted                       bool                         `json:"IsInverted"`
	IsAlwaysOnTop                    bool                         `json:"IsAlwaysOnTop"`
	IsZOrderEnabled                  bool                         `json:"IsZOrderEnabled"`
	IsClippingWithObjectAbove        bool                         `json:"IsClippingWithObjectAbove"`
	Font                             string                       `json:"Font"`
	FontSize                         FontSize                     `json:"FontSize"`
	LineHeight2                      FontSize                     `json:"LineHeight2"`
	LetterSpacing2                   FontSize                     `json:"LetterSpacing2"`
	WordWrap                         string                       `json:"WordWrap"`
	MaxWidth                         FontSize                     `json:"MaxWidth"`
	BasePoint                        string                       `json:"BasePoint"`
	FontColor                        string                       `json:"FontColor"`
	Style                            string                       `json:"Style"`
	StyleColor                       string                       `json:"StyleColor"`
	Bold                             bool                         `json:"Bold"`
	Italic                           bool                         `json:"Italic"`
	IsTrimEndSpace                   bool                         `json:"IsTrimEndSpace"`
	IsDevidedPerCharacter            bool                         `json:"IsDevidedPerCharacter"`
	DisplayInterval                  float64                      `json:"DisplayInterval"`
	DisplayDirection                 string                       `json:"DisplayDirection"`
	HideInterval                     float64                      `json:"HideInterval"`
	HideDirection                    string                       `json:"HideDirection"`
	JimakuVideoEffects               []CharacterJimakuVideoEffect `json:"JimakuVideoEffects"`
	TachieType                       string                       `json:"TachieType"`
	TachieCharacterParameter         TachieCharacterParameter     `json:"TachieCharacterParameter"`
	MouseSmooth                      int64                        `json:"MouseSmooth"`
	IsTachieLocked                   bool                         `json:"IsTachieLocked"`
	TachieX                          FontSize                     `json:"TachieX"`
	TachieY                          FontSize                     `json:"TachieY"`
	TachieZ                          FontSize                     `json:"TachieZ"`
	TachieOpacity                    FontSize                     `json:"TachieOpacity"`
	TachieZoom                       FontSize                     `json:"TachieZoom"`
	TachieRotation                   FontSize                     `json:"TachieRotation"`
	TachieFadeIn                     float64                      `json:"TachieFadeIn"`
	TachieFadeOut                    float64                      `json:"TachieFadeOut"`
	TachieBlend                      string                       `json:"TachieBlend"`
	TachieIsInverted                 bool                         `json:"TachieIsInverted"`
	TachieIsAlwaysOnTop              bool                         `json:"TachieIsAlwaysOnTop"`
	TachieIsZOrderEnabled            bool                         `json:"TachieIsZOrderEnabled"`
	TachieIsClippingWithObjectAbove  bool                         `json:"TachieIsClippingWithObjectAbove"`
	TachieDefaultItemParameter       TachieItemParameter          `json:"TachieDefaultItemParameter"`
	TachieItemVideoEffects           []TachieItemVideoEffect      `json:"TachieItemVideoEffects"`
	TachieDefaultFaceParameter       TachieFaceParameter          `json:"TachieDefaultFaceParameter"`
	TachieDefaultFaceEffects         []interface{}                `json:"TachieDefaultFaceEffects"`
	AdditionalForegroundTemplateName interface{}                  `json:"AdditionalForegroundTemplateName"`
	AdditionalBackgroundTemplateName interface{}                  `json:"AdditionalBackgroundTemplateName"`
	VoiceItemLength                  int64                        `json:"VoiceItemLength"`
	TachieItemLength                 int64                        `json:"TachieItemLength"`
	VoiceItemKeyFrames               KeyFrames                    `json:"VoiceItemKeyFrames"`
	TachieItemKeyFrames              KeyFrames                    `json:"TachieItemKeyFrames"`
}

type CharacterAudioEffect struct {
	Type           string   `json:"$type"`
	Label          string   `json:"Label"`
	Ratio          FontSize `json:"Ratio"`
	Threshold      FontSize `json:"Threshold"`
	MakeupGain     FontSize `json:"MakeupGain"`
	Attack         FontSize `json:"Attack"`
	Release        FontSize `json:"Release"`
	MinAmplitudeDB float64  `json:"MinAmplitudeDb"`
	IsEnabled      bool     `json:"IsEnabled"`
	Remark         string   `json:"Remark"`
}

type FontSize struct {
	Values        []Value `json:"Values"`
	Span          float64 `json:"Span"`
	AnimationType string  `json:"AnimationType"`
	Bezier        Bezier  `json:"Bezier"`
}

type Bezier struct {
	Points      []Point `json:"Points"`
	IsQuadratic bool    `json:"IsQuadratic"`
}

type Point struct {
	Point         ControlPoint1 `json:"Point"`
	ControlPoint1 ControlPoint1 `json:"ControlPoint1"`
	ControlPoint2 ControlPoint1 `json:"ControlPoint2"`
}

type ControlPoint1 struct {
	X float64 `json:"X"`
	Y float64 `json:"Y"`
}

type Value struct {
	Value float64 `json:"Value"`
}

type CharacterJimakuVideoEffect struct {
	Type             string    `json:"$type"`
	StrokeThickness  *FontSize `json:"StrokeThickness,omitempty"`
	Blur             FontSize  `json:"Blur"`
	X                FontSize  `json:"X"`
	Y                FontSize  `json:"Y"`
	Opacity          FontSize  `json:"Opacity"`
	Zoom             FontSize  `json:"Zoom"`
	Rotation         *FontSize `json:"Rotation,omitempty"`
	StrokeBrush      *Brush    `json:"StrokeBrush,omitempty"`
	IsOutlineOnly    *bool     `json:"IsOutlineOnly,omitempty"`
	IsAngular        *bool     `json:"IsAngular,omitempty"`
	IsEnabled        bool      `json:"IsEnabled"`
	Remark           string    `json:"Remark"`
	Angle            *FontSize `json:"Angle,omitempty"`
	IsRotateAtCenter *bool     `json:"IsRotateAtCenter,omitempty"`
	Brush            *Brush    `json:"Brush,omitempty"`
}

type Brush struct {
	Type      string               `json:"Type"`
	Parameter StrokeBrushParameter `json:"Parameter"`
}

type StrokeBrushParameter struct {
	Type  string `json:"$type"`
	Color string `json:"Color"`
}

type KeyGesture struct {
	Key       int64 `json:"Key"`
	Modifiers int64 `json:"Modifiers"`
}

type TachieCharacterParameter struct {
	Type             string  `json:"$type"`
	FilePath         string  `json:"FilePath"`
	MouthSensitivity float64 `json:"MouthSensitivity"`
}

type TachieFaceParameter struct {
	Type           string   `json:"$type"`
	IsEnabled      bool     `json:"IsEnabled"`
	FilePath       *string  `json:"FilePath"`
	EnableLayers   []string `json:"EnableLayers"`
	EyeAnimation   string   `json:"EyeAnimation"`
	MouthAnimation string   `json:"MouthAnimation"`
}

type TachieItemParameter struct {
	Type                 string   `json:"$type"`
	FilePath             string   `json:"FilePath"`
	EnableLayers         []string `json:"EnableLayers"`
	IsHiddenWhenNoSpeech bool     `json:"IsHiddenWhenNoSpeech"`
}

type KeyFrames struct {
	Frames []int64 `json:"Frames"`
	Count  int64   `json:"Count"`
}

type TachieItemVideoEffect struct {
	Type               string    `json:"$type"`
	StrokeThickness    *FontSize `json:"StrokeThickness,omitempty"`
	Blur               *FontSize `json:"Blur,omitempty"`
	X                  *FontSize `json:"X,omitempty"`
	Y                  *FontSize `json:"Y,omitempty"`
	Opacity            *FontSize `json:"Opacity,omitempty"`
	Zoom               *FontSize `json:"Zoom,omitempty"`
	Rotation           *FontSize `json:"Rotation,omitempty"`
	StrokeBrush        *Brush    `json:"StrokeBrush,omitempty"`
	IsOutlineOnly      *bool     `json:"IsOutlineOnly,omitempty"`
	IsAngular          *bool     `json:"IsAngular,omitempty"`
	IsEnabled          bool      `json:"IsEnabled"`
	Remark             string    `json:"Remark"`
	Angle              *FontSize `json:"Angle,omitempty"`
	IsRotateAtCenter   *bool     `json:"IsRotateAtCenter,omitempty"`
	Brush              *Brush    `json:"Brush,omitempty"`
	MotionType         *string   `json:"MotionType,omitempty"`
	StartNaturally     *bool     `json:"StartNaturally,omitempty"`
	EndNaturally       *bool     `json:"EndNaturally,omitempty"`
	Loop               *bool     `json:"Loop,omitempty"`
	Interval           *float64  `json:"Interval,omitempty"`
	Invert             *bool     `json:"Invert,omitempty"`
	Speed              *float64  `json:"Speed,omitempty"`
	PositionCorrection *float64  `json:"PositionCorrection,omitempty"`
	ZoomCorrection     *float64  `json:"ZoomCorrection,omitempty"`
	RotationCorrection *float64  `json:"RotationCorrection,omitempty"`
}

type Voice struct {
	API string `json:"API"`
	Arg string `json:"Arg"`
}

type VoiceParameter struct {
	Type              string  `json:"$type"`
	StyleID           int64   `json:"StyleID"`
	MorphingStyleID   int64   `json:"MorphingStyleID"`
	MorphingRate      float64 `json:"MorphingRate"`
	Speed             int64   `json:"Speed"`
	Pitch             float64 `json:"Pitch"`
	Intonation        float64 `json:"Intonation"`
	PrePhonemeLength  float64 `json:"PrePhonemeLength"`
	PostPhonemeLength float64 `json:"PostPhonemeLength"`
	PauseLengthScale  float64 `json:"PauseLengthScale"`
}

type Timeline struct {
	ID            string         `json:"ID"`
	Name          string         `json:"Name"`
	VideoInfo     VideoInfo      `json:"VideoInfo"`
	VerticalLine  VerticalLine   `json:"VerticalLine"`
	Items         []TimelineItem `json:"Items"`
	LayerSettings LayerSettings  `json:"LayerSettings"`
	CurrentFrame  int64          `json:"CurrentFrame"`
	Length        int64          `json:"Length"`
	MaxLayer      int64          `json:"MaxLayer"`
}

type TimelineItem struct {
	Type                      string                  `json:"$type"`
	FilePath                  *string                 `json:"FilePath,omitempty"`
	X                         *FontSize               `json:"X,omitempty"`
	Y                         *FontSize               `json:"Y,omitempty"`
	Z                         *FontSize               `json:"Z,omitempty"`
	Opacity                   *FontSize               `json:"Opacity,omitempty"`
	Zoom                      *FontSize               `json:"Zoom,omitempty"`
	Rotation                  *FontSize               `json:"Rotation,omitempty"`
	FadeIn                    *float64                `json:"FadeIn,omitempty"`
	FadeOut                   *float64                `json:"FadeOut,omitempty"`
	Blend                     *string                 `json:"Blend,omitempty"`
	IsInverted                *bool                   `json:"IsInverted,omitempty"`
	IsClippingWithObjectAbove *bool                   `json:"IsClippingWithObjectAbove,omitempty"`
	IsAlwaysOnTop             *bool                   `json:"IsAlwaysOnTop,omitempty"`
	IsZOrderEnabled           *bool                   `json:"IsZOrderEnabled,omitempty"`
	VideoEffects              []VideoEffect           `json:"VideoEffects,omitempty"`
	Group                     int64                   `json:"Group"`
	Frame                     int64                   `json:"Frame"`
	Layer                     int64                   `json:"Layer"`
	KeyFrames                 KeyFrames               `json:"KeyFrames"`
	Length                    int64                   `json:"Length"`
	PlaybackRate              float64                 `json:"PlaybackRate"`
	ContentOffset             string                  `json:"ContentOffset"`
	Remark                    string                  `json:"Remark"`
	IsLocked                  bool                    `json:"IsLocked"`
	IsHidden                  bool                    `json:"IsHidden"`
	Text                      *string                 `json:"Text,omitempty"`
	Decorations               []Decoration            `json:"Decorations,omitempty"`
	Font                      *string                 `json:"Font,omitempty"`
	FontSize                  *FontSize               `json:"FontSize,omitempty"`
	LineHeight2               *FontSize               `json:"LineHeight2,omitempty"`
	LetterSpacing2            *FontSize               `json:"LetterSpacing2,omitempty"`
	WordWrap                  *string                 `json:"WordWrap,omitempty"`
	MaxWidth                  *FontSize               `json:"MaxWidth,omitempty"`
	BasePoint                 *string                 `json:"BasePoint,omitempty"`
	FontColor                 *string                 `json:"FontColor,omitempty"`
	Style                     *string                 `json:"Style,omitempty"`
	StyleColor                *string                 `json:"StyleColor,omitempty"`
	Bold                      *bool                   `json:"Bold,omitempty"`
	Italic                    *bool                   `json:"Italic,omitempty"`
	IsTrimEndSpace            *bool                   `json:"IsTrimEndSpace,omitempty"`
	IsDevidedPerCharacter     *bool                   `json:"IsDevidedPerCharacter,omitempty"`
	DisplayInterval           *float64                `json:"DisplayInterval,omitempty"`
	DisplayDirection          *string                 `json:"DisplayDirection,omitempty"`
	HideInterval              *float64                `json:"HideInterval,omitempty"`
	HideDirection             *string                 `json:"HideDirection,omitempty"`
	ShapeType2                *string                 `json:"ShapeType2,omitempty"`
	ShapeParameter            *ShapeParameter         `json:"ShapeParameter,omitempty"`
	GroupRange                *int64                  `json:"GroupRange,omitempty"`
	IsGroupOnly               *bool                   `json:"IsGroupOnly,omitempty"`
	IsCompressFrame           *bool                   `json:"IsCompressFrame,omitempty"`
	CompositeCenter           *string                 `json:"CompositeCenter,omitempty"`
	CharacterName             *string                 `json:"CharacterName,omitempty"`
	TachieItemParameter       *TachieItemParameter    `json:"TachieItemParameter,omitempty"`
	IsWaveformEnabled         *bool                   `json:"IsWaveformEnabled,omitempty"`
	Serif                     *string                 `json:"Serif,omitempty"`
	Hatsuon                   *string                 `json:"Hatsuon,omitempty"`
	Pronounce                 *Pronounce              `json:"Pronounce,omitempty"`
	LipSyncFrames             []LipSyncFrame          `json:"LipSyncFrames,omitempty"`
	VoiceLength               *string                 `json:"VoiceLength,omitempty"`
	VoiceCache                *string                 `json:"VoiceCache,omitempty"`
	Volume                    *FontSize               `json:"Volume,omitempty"`
	Pan                       *FontSize               `json:"Pan,omitempty"`
	VoiceParameter            *VoiceParameter         `json:"VoiceParameter,omitempty"`
	VoiceFadeIn               *float64                `json:"VoiceFadeIn,omitempty"`
	VoiceFadeOut              *float64                `json:"VoiceFadeOut,omitempty"`
	EchoIsEnabled             *bool                   `json:"EchoIsEnabled,omitempty"`
	EchoInterval              *float64                `json:"EchoInterval,omitempty"`
	EchoAttenuation           *float64                `json:"EchoAttenuation,omitempty"`
	AudioEffects              []ItemAudioEffect       `json:"AudioEffects,omitempty"`
	JimakuVisibility          *string                 `json:"JimakuVisibility,omitempty"`
	JimakuFadeIn              *float64                `json:"JimakuFadeIn,omitempty"`
	JimakuFadeOut             *float64                `json:"JimakuFadeOut,omitempty"`
	JimakuVideoEffects        []ItemJimakuVideoEffect `json:"JimakuVideoEffects,omitempty"`
	TachieFaceParameter       *TachieFaceParameter    `json:"TachieFaceParameter,omitempty"`
	TachieFaceEffects         []TachieFaceEffect      `json:"TachieFaceEffects,omitempty"`
	AudioTrackIndex           *int64                  `json:"AudioTrackIndex,omitempty"`
	IsLooped                  *bool                   `json:"IsLooped,omitempty"`
	TransitionType            *string                 `json:"TransitionType,omitempty"`
	TransitionParameter       *TransitionParameter    `json:"TransitionParameter,omitempty"`
	BeforeVideoEffects        []interface{}           `json:"BeforeVideoEffects,omitempty"`
	AfterVideoEffects         []interface{}           `json:"AfterVideoEffects,omitempty"`
	Blur                      *FontSize               `json:"Blur,omitempty"`
	InvertMask                *bool                   `json:"InvertMask,omitempty"`
}

type ItemAudioEffect struct {
	Type              string      `json:"$type"`
	Label             string      `json:"Label"`
	Ratio             *FontSize   `json:"Ratio,omitempty"`
	Threshold         *FontSize   `json:"Threshold,omitempty"`
	MakeupGain        *FontSize   `json:"MakeupGain,omitempty"`
	Attack            *FontSize   `json:"Attack,omitempty"`
	Release           *FontSize   `json:"Release,omitempty"`
	MinAmplitudeDB    *float64    `json:"MinAmplitudeDb,omitempty"`
	IsEnabled         bool        `json:"IsEnabled"`
	Remark            string      `json:"Remark"`
	Algorithm         *string     `json:"Algorithm,omitempty"`
	PitchType         *string     `json:"PitchType,omitempty"`
	PitchValue        *PitchValue `json:"PitchValue,omitempty"`
	Strength          *FontSize   `json:"Strength,omitempty"`
	FftSize           *int64      `json:"FftSize,omitempty"`
	HopSize           *float64    `json:"HopSize,omitempty"`
	LeftDelay         *FontSize   `json:"LeftDelay,omitempty"`
	RightDelay        *FontSize   `json:"RightDelay,omitempty"`
	Volume            *float64    `json:"Volume,omitempty"`
	IsInEffect        *bool       `json:"IsInEffect,omitempty"`
	IsOutEffect       *bool       `json:"IsOutEffect,omitempty"`
	EffectTimeSeconds *float64    `json:"EffectTimeSeconds,omitempty"`
	EasingType        *string     `json:"EasingType,omitempty"`
	EasingMode        *string     `json:"EasingMode,omitempty"`
}

type PitchValue struct {
	Type    string  `json:"$type"`
	Percent float64 `json:"Percent"`
	Rate    float64 `json:"Rate"`
}

type Decoration struct {
	Start         int64       `json:"Start"`
	Length        int64       `json:"Length"`
	IsBold        bool        `json:"IsBold"`
	IsItalic      bool        `json:"IsItalic"`
	Scale         float64     `json:"Scale"`
	Font          interface{} `json:"Font"`
	Foreground    *string     `json:"Foreground"`
	IsLineBreak   bool        `json:"IsLineBreak"`
	HasDecoration bool        `json:"HasDecoration"`
}

type ItemJimakuVideoEffect struct {
	Type             string    `json:"$type"`
	StrokeThickness  *FontSize `json:"StrokeThickness,omitempty"`
	Blur             *FontSize `json:"Blur,omitempty"`
	X                *FontSize `json:"X,omitempty"`
	Y                *FontSize `json:"Y,omitempty"`
	Opacity          FontSize  `json:"Opacity"`
	Zoom             *FontSize `json:"Zoom,omitempty"`
	Rotation         *FontSize `json:"Rotation,omitempty"`
	StrokeBrush      *Brush    `json:"StrokeBrush,omitempty"`
	IsOutlineOnly    *bool     `json:"IsOutlineOnly,omitempty"`
	IsAngular        *bool     `json:"IsAngular,omitempty"`
	IsEnabled        bool      `json:"IsEnabled"`
	Remark           string    `json:"Remark"`
	Round            *FontSize `json:"Round,omitempty"`
	BlendMode        *string   `json:"BlendMode,omitempty"`
	IsBackgroundOnly *bool     `json:"IsBackgroundOnly,omitempty"`
	Brush            *Brush    `json:"Brush,omitempty"`
	Top              *FontSize `json:"Top,omitempty"`
	Bottom           *FontSize `json:"Bottom,omitempty"`
	Left             *FontSize `json:"Left,omitempty"`
	Right            *FontSize `json:"Right,omitempty"`
	Angle            *FontSize `json:"Angle,omitempty"`
	IsRotateAtCenter *bool     `json:"IsRotateAtCenter,omitempty"`
	GradientType     *string   `json:"GradientType,omitempty"`
	Stops            []Stop    `json:"Stops,omitempty"`
	Size             *FontSize `json:"Size,omitempty"`
	ExtendMode       *string   `json:"ExtendMode,omitempty"`
	Blend            *string   `json:"Blend,omitempty"`
}

type Stop struct {
	Offset float64 `json:"Offset"`
	Color  string  `json:"Color"`
}

type LipSyncFrame struct {
	Time  string `json:"Time"`
	Shape string `json:"Shape"`
}

type Pronounce struct {
	Type          string      `json:"$type"`
	AudioQuery    AudioQuery  `json:"AudioQuery"`
	LipSyncFrames interface{} `json:"LipSyncFrames"`
}

type AudioQuery struct {
	Type               string         `json:"$type"`
	AccentPhrases      []AccentPhrase `json:"accent_phrases"`
	SpeedScale         float64        `json:"speedScale"`
	PitchScale         float64        `json:"pitchScale"`
	IntonationScale    float64        `json:"intonationScale"`
	VolumeScale        float64        `json:"volumeScale"`
	PrePhonemeLength   float64        `json:"prePhonemeLength"`
	PostPhonemeLength  float64        `json:"postPhonemeLength"`
	OutputSamplingRate int64          `json:"outputSamplingRate"`
	OutputStereo       bool           `json:"outputStereo"`
	Kana               string         `json:"kana"`
	PauseLength        interface{}    `json:"pauseLength"`
	PauseLengthScale   float64        `json:"pauseLengthScale"`
}

type AccentPhrase struct {
	Moras           []Mora `json:"moras"`
	Accent          int64  `json:"accent"`
	PauseMora       *Mora  `json:"pause_mora"`
	IsInterrogative bool   `json:"is_interrogative"`
}

type Mora struct {
	Text            string   `json:"text"`
	Consonant       *string  `json:"consonant"`
	ConsonantLength *float64 `json:"consonant_length"`
	Vowel           string   `json:"vowel"`
	VowelLength     float64  `json:"vowel_length"`
	Pitch           float64  `json:"pitch"`
}

type ShapeParameter struct {
	Type            string               `json:"$type"`
	Dent            *FontSize            `json:"Dent,omitempty"`
	BarLength       *FontSize            `json:"BarLength,omitempty"`
	BarThickness    *FontSize            `json:"BarThickness,omitempty"`
	BarShapeType    *string              `json:"BarShapeType,omitempty"`
	SizeMode        *string              `json:"SizeMode,omitempty"`
	Size            FontSize             `json:"Size"`
	AspectRate      *FontSize            `json:"AspectRate,omitempty"`
	Width           *FontSize            `json:"Width,omitempty"`
	Height          *FontSize            `json:"Height,omitempty"`
	StrokeThickness *FontSize            `json:"StrokeThickness,omitempty"`
	Brush           *ShapeParameterBrush `json:"Brush,omitempty"`
	Density         *FontSize            `json:"Density,omitempty"`
	Thickness       *FontSize            `json:"Thickness,omitempty"`
	Length          *FontSize            `json:"Length,omitempty"`
	CenterWidth     *FontSize            `json:"CenterWidth,omitempty"`
	Speed           *FontSize            `json:"Speed,omitempty"`
	Stroke          *string              `json:"Stroke,omitempty"`
	Round           *FontSize            `json:"Round,omitempty"`
}

type ShapeParameterBrush struct {
	Type      string          `json:"Type"`
	Parameter PurpleParameter `json:"Parameter"`
}

type PurpleParameter struct {
	Type            string    `json:"$type"`
	Color           *string   `json:"Color,omitempty"`
	Stops           []Stop    `json:"Stops,omitempty"`
	Size            *FontSize `json:"Size,omitempty"`
	Offset          *FontSize `json:"Offset,omitempty"`
	Angle           *FontSize `json:"Angle,omitempty"`
	ExtendMode      *string   `json:"ExtendMode,omitempty"`
	StrokeColor     *string   `json:"StrokeColor,omitempty"`
	BackgroundColor *string   `json:"BackgroundColor,omitempty"`
	Thickness       *FontSize `json:"Thickness,omitempty"`
	Width           *FontSize `json:"Width,omitempty"`
	Height          *FontSize `json:"Height,omitempty"`
	Zoom            *FontSize `json:"Zoom,omitempty"`
	X               *FontSize `json:"X,omitempty"`
	Y               *FontSize `json:"Y,omitempty"`
	Aspect          *FontSize `json:"Aspect,omitempty"`
	IsInverted      *bool     `json:"IsInverted,omitempty"`
}

type TachieFaceEffect struct {
	Type               string  `json:"$type"`
	MotionType         string  `json:"MotionType"`
	StartNaturally     bool    `json:"StartNaturally"`
	EndNaturally       bool    `json:"EndNaturally"`
	Loop               bool    `json:"Loop"`
	Interval           float64 `json:"Interval"`
	Invert             bool    `json:"Invert"`
	Speed              float64 `json:"Speed"`
	PositionCorrection float64 `json:"PositionCorrection"`
	ZoomCorrection     float64 `json:"ZoomCorrection"`
	RotationCorrection float64 `json:"RotationCorrection"`
	IsEnabled          bool    `json:"IsEnabled"`
	Remark             string  `json:"Remark"`
}

type TransitionParameter struct {
	Type                  string    `json:"$type"`
	File                  *string   `json:"File,omitempty"`
	Tolerance             *FontSize `json:"Tolerance,omitempty"`
	Angle                 *FontSize `json:"Angle,omitempty"`
	EasingType            string    `json:"EasingType"`
	EasingMode            *string   `json:"EasingMode,omitempty"`
	KeepAspect            *bool     `json:"KeepAspect,omitempty"`
	IsFixedCoveringScale  *bool     `json:"IsFixedCoveringScale,omitempty"`
	IsReversed            *bool     `json:"IsReversed,omitempty"`
	Target                *string   `json:"Target,omitempty"`
	Direction             *string   `json:"Direction,omitempty"`
	ZoomIn                *FontSize `json:"ZoomIn,omitempty"`
	ZoomOut               *FontSize `json:"ZoomOut,omitempty"`
	Rotation              *FontSize `json:"Rotation,omitempty"`
	IsFlip                *bool     `json:"IsFlip,omitempty"`
	Strength              *FontSize `json:"Strength,omitempty"`
	Threshold             *FontSize `json:"Threshold,omitempty"`
	Blur                  *FontSize `json:"Blur,omitempty"`
	IsFixedSizeEnabled    *bool     `json:"IsFixedSizeEnabled,omitempty"`
	IsColorizationEnabled *bool     `json:"IsColorizationEnabled,omitempty"`
	Color                 *string   `json:"Color,omitempty"`
	RotationCount         *float64  `json:"RotationCount,omitempty"`
	Angle1                *FontSize `json:"Angle1,omitempty"`
	Angle2                *FontSize `json:"Angle2,omitempty"`
	Amplitude             *FontSize `json:"Amplitude,omitempty"`
	WaveLength            *FontSize `json:"WaveLength,omitempty"`
	Period                *FontSize `json:"Period,omitempty"`
}

type VideoEffect struct {
	Type                 string    `json:"$type"`
	SizeMode             *string   `json:"SizeMode,omitempty"`
	Size                 *Size     `json:"Size,omitempty"`
	Dot                  *bool     `json:"Dot,omitempty"`
	IsEnabled            bool      `json:"IsEnabled"`
	Remark               string    `json:"Remark"`
	StrokeThickness      *FontSize `json:"StrokeThickness,omitempty"`
	Blur                 *FontSize `json:"Blur,omitempty"`
	X                    *FontSize `json:"X,omitempty"`
	Y                    *FontSize `json:"Y,omitempty"`
	Opacity              *FontSize `json:"Opacity,omitempty"`
	Zoom                 *FontSize `json:"Zoom,omitempty"`
	Rotation             *FontSize `json:"Rotation,omitempty"`
	StrokeBrush          *Brush    `json:"StrokeBrush,omitempty"`
	IsOutlineOnly        *bool     `json:"IsOutlineOnly,omitempty"`
	IsAngular            *bool     `json:"IsAngular,omitempty"`
	ZoomX                *FontSize `json:"ZoomX,omitempty"`
	ZoomY                *FontSize `json:"ZoomY,omitempty"`
	IsNearestNeighbor    *bool     `json:"IsNearestNeighbor,omitempty"`
	GradientType         *string   `json:"GradientType,omitempty"`
	Stops                []Stop    `json:"Stops,omitempty"`
	ExtendMode           *string   `json:"ExtendMode,omitempty"`
	Blend                *string   `json:"Blend,omitempty"`
	Angle                *FontSize `json:"Angle,omitempty"`
	Width                *FontSize `json:"Width,omitempty"`
	Horizontal           *string   `json:"Horizontal,omitempty"`
	Vertical             *string   `json:"Vertical,omitempty"`
	IsKeepPosition       *bool     `json:"IsKeepPosition,omitempty"`
	IsHardBorderMode     *bool     `json:"IsHardBorderMode,omitempty"`
	Length               *FontSize `json:"Length,omitempty"`
	IsRotateAtCenter     *bool     `json:"IsRotateAtCenter,omitempty"`
	Brush                *Brush    `json:"Brush,omitempty"`
	MotionType           *string   `json:"MotionType,omitempty"`
	StartNaturally       *bool     `json:"StartNaturally,omitempty"`
	EndNaturally         *bool     `json:"EndNaturally,omitempty"`
	Loop                 *bool     `json:"Loop,omitempty"`
	Interval             *Interval `json:"Interval"`
	Invert               *bool     `json:"Invert,omitempty"`
	Speed                *float64  `json:"Speed,omitempty"`
	PositionCorrection   *float64  `json:"PositionCorrection,omitempty"`
	ZoomCorrection       *float64  `json:"ZoomCorrection,omitempty"`
	RotationCorrection   *float64  `json:"RotationCorrection,omitempty"`
	Z                    *FontSize `json:"Z,omitempty"`
	Span                 *FontSize `json:"Span,omitempty"`
	DistanceMode         *string   `json:"DistanceMode,omitempty"`
	KeyColor             *string   `json:"KeyColor,omitempty"`
	Threshold2           *FontSize `json:"Threshold2,omitempty"`
	Smoothness           *FontSize `json:"Smoothness,omitempty"`
	Despill              *FontSize `json:"Despill,omitempty"`
	IsInvert             *bool     `json:"IsInvert,omitempty"`
	EasingType           *string   `json:"EasingType,omitempty"`
	EasingMode           *string   `json:"EasingMode,omitempty"`
	AngleX               *float64  `json:"AngleX,omitempty"`
	AngleY               *float64  `json:"AngleY,omitempty"`
	CenterPoint          *string   `json:"CenterPoint,omitempty"`
	CenterX              *float64  `json:"CenterX,omitempty"`
	CenterY              *float64  `json:"CenterY,omitempty"`
	IsInEffect           *bool     `json:"IsInEffect,omitempty"`
	IsOutEffect          *bool     `json:"IsOutEffect,omitempty"`
	EffectTimeSeconds    *float64  `json:"EffectTimeSeconds,omitempty"`
	Yaw                  *FontSize `json:"Yaw,omitempty"`
	Pitch                *FontSize `json:"Pitch,omitempty"`
	Roll                 *FontSize `json:"Roll,omitempty"`
	IsCentering          *bool     `json:"IsCentering,omitempty"`
	JumpHeight           *FontSize `json:"JumpHeight,omitempty"`
	Stretch              *FontSize `json:"Stretch,omitempty"`
	Period               *FontSize `json:"Period,omitempty"`
	Distortion           *FontSize `json:"Distortion,omitempty"`
	File                 *string   `json:"File,omitempty"`
	Tolerance            *FontSize `json:"Tolerance,omitempty"`
	KeepAspect           *bool     `json:"KeepAspect,omitempty"`
	IsFixedCoveringScale *bool     `json:"IsFixedCoveringScale,omitempty"`
	IsReversedInEffect   *bool     `json:"IsReversedInEffect,omitempty"`
	IsReversedOutEffect  *bool     `json:"IsReversedOutEffect,omitempty"`
}

type Size struct {
	Type          *string   `json:"$type,omitempty"`
	Width         *FontSize `json:"Width,omitempty"`
	Height        *FontSize `json:"Height,omitempty"`
	Label         *string   `json:"Label,omitempty"`
	Values        []Value   `json:"Values,omitempty"`
	Span          *float64  `json:"Span,omitempty"`
	AnimationType *string   `json:"AnimationType,omitempty"`
	Bezier        *Bezier   `json:"Bezier,omitempty"`
}

type LayerSettings struct {
	Items []LayerSettingsItem `json:"Items"`
}

type LayerSettingsItem struct {
	Layer    int64   `json:"Layer"`
	Label    string  `json:"Label"`
	Color    string  `json:"Color"`
	IsHidden bool    `json:"IsHidden"`
	Volume   float64 `json:"Volume"`
}

type VerticalLine struct {
	IsEnabled  bool   `json:"IsEnabled"`
	StartFrame int64  `json:"StartFrame"`
	LineType   string `json:"LineType"`
	Line       Line   `json:"Line"`
	Group      int64  `json:"Group"`
}

type Line struct {
	Type string  `json:"$type"`
	BPM  float64 `json:"BPM"`
}

type VideoInfo struct {
	FPS    int64 `json:"FPS"`
	Hz     int64 `json:"Hz"`
	Width  int64 `json:"Width"`
	Height int64 `json:"Height"`
}

type ToolState struct {
	Title      string  `json:"Title"`
	SavedState *string `json:"SavedState"`
}

type Interval struct {
	Double   *float64
	FontSize *FontSize
}

func (x *Interval) UnmarshalJSON(data []byte) error {
	x.FontSize = nil
	var c FontSize
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FontSize = &c
	}
	return nil
}

func (x *Interval) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, false, nil, x.FontSize != nil, x.FontSize, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")
}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
