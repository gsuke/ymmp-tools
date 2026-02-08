package main

// TimeLineItemsを条件で抽出する。
// layer=-1, itemType=ItemNone で指定なし。
// ソートされていないので注意。
// layerは0-basedなので念の為注意。
func FilterTimeLineItems(
	items []TimelineItem,
	layer int64, // 0-ordered
	itemType TimelineItemType,
) []TimelineItem {
	result := make([]TimelineItem, 0)
	for _, item := range items {
		if layer != -1 && item.Layer != layer {
			continue
		}
		if string(itemType) != "" && item.Type != string(itemType) {
			continue
		}
		result = append(result, item)
	}
	return result
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
