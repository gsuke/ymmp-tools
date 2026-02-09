# YMMP Tools

ゆっくりMovieMaker4のプロジェクトファイル `.ymmp` を直接書き換えることで、動画編集作業を効率化するCLIツール集。

## 機能

* `extend`: タイムラインアイテムを次のアイテムのところまで自動で引き伸ばす
* `get_subtitles`: 字幕を取得する
* `decorate`: 字幕の一部分に色をつける(DSLを使用)
* `add_transitions`: 背景などの切り替わりのタイミングでトランジションエフェクトを追加する

## 使い方

以下のワークフローを推奨。

```shell
# 1. 公式の「台本編集」より、ボイスを追加し、発音や間を調整する。

# 2. extend
extend input.ymmp -t 1 -l 7 -l 8 -l 9 -o output.ymmp

# 3. get_subtitles
get_subtitles input.ymmp -t 1 -l 7 --mark-long > subtitles.txt

# 4. subtitles.txt に改行コード `\r\n` や 字幕色を追加する。

# 5. decorate
decorate input.ymmp -i subtitles.txt -t 1 -l 7  -o output.ymmp

# 6. add_transitions
add_transitions input.ymmp -t 0 -s 7 -d 8 -o output.ymmp
```

## 開発

### データ構造の調査

quicktypeを使用して、プロジェクトファイルの構造を自動で分析できます。

```
quicktype <input>.ymmp -o ymmp.go --no-enums
```

また、細かいデータ構造は実物を見ましょう。

### テスト

```
go test -v ./cmd/...
```

### TODO

`TODO` でgrepしてください。
