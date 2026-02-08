# .ymmp Tools

```
quicktype <input>.ymmp -o ymmp.go --no-enums
```

# セリフ(TimeLineItem.Serif)形式

TimeLineItem.Serif の改行は `\r\n`

デコレーションの例 (12文字テキストで、先頭4文字だけ色つき)

```json
"Decorations": [
  {
    "Start": 0,
    "Length": 4,
    "IsBold": false,
    "IsItalic": false,
    "Scale": 1.0,
    "Font": null,
    "Foreground": "#FF94F9FF",
    "IsLineBreak": false,
    "HasDecoration": true
  },
  {
    "Start": 4,
    "Length": 8,
    "IsBold": false,
    "IsItalic": false,
    "Scale": 1.0,
    "Font": null,
    "Foreground": null,
    "IsLineBreak": false,
    "HasDecoration": false
  }
]
```
