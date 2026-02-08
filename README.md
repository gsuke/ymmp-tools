# .ymmp Tools

```
quicktype <input>.ymmp -o ymmp.go --no-enums
```

# セリフ(TimeLineItem.Serif)形式

TimeLineItem.Serif の改行は `\r\n`

## デコレーションの例

サンプル文字: `あいうえおあいうえ\r\nあいうあいうあいうえお!!!`

* 改行をまたいでデコレーションする場合は、分割する。
* 改行コード `\r\n` は2文字としてカウントする。

```json
"Decorations": [
  {
    "Start": 0,
    "Length": 9,
    "IsBold": false,
    "IsItalic": false,
    "Scale": 1.0,
    "Font": null,
    "Foreground": "#FFFFAAAA",
    "IsLineBreak": false,
    "HasDecoration": true
  },
  {
    "Start": 9,
    "Length": 2,
    "IsBold": false,
    "IsItalic": false,
    "Scale": 1.0,
    "Font": null,
    "Foreground": null,
    "IsLineBreak": true,
    "HasDecoration": false
  },
  {
    "Start": 11,
    "Length": 6,
    "IsBold": false,
    "IsItalic": false,
    "Scale": 1.0,
    "Font": null,
    "Foreground": "#FFFFAAAA",
    "IsLineBreak": false,
    "HasDecoration": true
  },
  {
    "Start": 17,
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
