package ymmp

import "sort"

type TimelineItems []TimelineItem

// TimelineItemsを条件で抽出する。
// layer=-1, itemType=ItemNone, frame=-1 で指定なし。
// ソートしたい場合は、SortTimelineItemsByFrame()を繋げること。
// layerは0-basedなので念の為注意。
func (items TimelineItems) FilterTimelineItems(
	layer int64,
	itemType TimelineItemType,
	frame int64,
) TimelineItems {
	result := make(TimelineItems, 0)
	for _, item := range items {
		if layer != -1 && item.Layer != layer {
			continue
		}
		if string(itemType) != "" && item.Type != string(itemType) {
			continue
		}
		if frame != -1 && item.Frame != frame {
			continue
		}
		result = append(result, item)
	}
	return result
}

// Frameの昇順でソートする。
func (items TimelineItems) SortTimelineItemsByFrame() TimelineItems {
	sorted := make(TimelineItems, len(items))
	copy(sorted, items)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Frame < sorted[j].Frame
	})
	return sorted
}

// ExtendTimelineItems は指定されたレイヤーのアイテムのLengthを次のアイテムまで伸ばす。
// 【副作用あり】この関数は items スライスの要素を直接変更する。
func ExtendTimelineItems(items []TimelineItem, layers []int64) {
	// 処理の流れ:
	//  1. 指定された各レイヤーについて以下を実行
	//  2. そのレイヤーに属するアイテムのインデックスを収集
	//  3. Frame（開始位置）の昇順でインデックスをソート
	//  4. 各アイテムのLengthを「次のアイテムのFrame - 自分のFrame」に更新
	//  5. 右端（最後）のアイテムは次がないので変更しない
	//
	// 例: Layer 7 に Frame=0, Frame=100, Frame=250 の3つのアイテムがある場合
	//
	//	変更前: [Frame=0, Length=50], [Frame=100, Length=80], [Frame=250, Length=100]
	//	変更後: [Frame=0, Length=100], [Frame=100, Length=150], [Frame=250, Length=100] ← 右端は変更なし
	//
	// なぜインデックスを使うのか:
	//   - items スライスには複数レイヤーのアイテムが混在している
	//   - 直接フィルタすると元のスライスへの参照が失われ、変更が反映されない
	//   - インデックスを保持することで、元の items[i].Length を直接更新できる

	for _, layer := range layers {
		// ステップ1: このレイヤーに属するアイテムの「元スライス内でのインデックス」を収集
		// 例: items = [Layer0のアイテム, Layer7のアイテム, Layer0のアイテム, Layer7のアイテム, ...]
		//     → Layer7 の場合、indices = [1, 3, ...] のようになる
		var indices []int
		for i, item := range items {
			if item.Layer == layer {
				indices = append(indices, i)
			}
		}

		// ステップ2: 収集したインデックスをFrame順（時系列順）にソート
		// indices自体をソートするが、比較時は items[indices[a]].Frame を参照する
		// これにより「元スライス内のインデックス」を「Frame順」に並べ替える
		sort.Slice(indices, func(a, b int) bool {
			return items[indices[a]].Frame < items[indices[b]].Frame
		})

		// ステップ3: 最後のアイテム以外のLengthを更新
		// len(indices)-1 までループすることで、最後のアイテムをスキップ
		for i := 0; i < len(indices)-1; i++ {
			currentIdx := indices[i] // 現在のアイテムの元スライス内インデックス
			nextIdx := indices[i+1]  // 次のアイテムの元スライス内インデックス
			// 新しいLength = 次のアイテムの開始位置 - 自分の開始位置
			items[currentIdx].Length = items[nextIdx].Frame - items[currentIdx].Frame
		}
	}
}

type TimelineItemType string

const (
	ItemNone       TimelineItemType = "" // 条件指定用
	ItemShape      TimelineItemType = "YukkuriMovieMaker.Project.Items.ShapeItem, YukkuriMovieMaker"
	ItemTachie     TimelineItemType = "YukkuriMovieMaker.Project.Items.TachieItem, YukkuriMovieMaker"
	ItemText       TimelineItemType = "YukkuriMovieMaker.Project.Items.TextItem, YukkuriMovieMaker"
	ItemVoice      TimelineItemType = "YukkuriMovieMaker.Project.Items.VoiceItem, YukkuriMovieMaker"
	ItemImage      TimelineItemType = "YukkuriMovieMaker.Project.Items.ImageItem, YukkuriMovieMaker"
	ItemAudio      TimelineItemType = "YukkuriMovieMaker.Project.Items.AudioItem, YukkuriMovieMaker"
	ItemVideo      TimelineItemType = "YukkuriMovieMaker.Project.Items.VideoItem, YukkuriMovieMaker"
	ItemGroup      TimelineItemType = "YukkuriMovieMaker.Project.Items.GroupItem, YukkuriMovieMaker"
	ItemTachieFace TimelineItemType = "YukkuriMovieMaker.Project.Items.TachieFaceItem, YukkuriMovieMaker"
	ItemTransition TimelineItemType = "YukkuriMovieMaker.Project.Items.TransitionItem, YukkuriMovieMaker"
	ItemEffect     TimelineItemType = "YukkuriMovieMaker.Project.Items.EffectItem, YukkuriMovieMaker"
)
