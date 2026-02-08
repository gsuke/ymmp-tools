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
