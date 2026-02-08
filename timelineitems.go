package main

// TimeLineItemsを条件で抽出する。
// itemType=""で指定なし。
func FilterTimeLineItems(
	items []TimelineItem,
	layer int64, // 0-ordered
	itemType string,
) []TimelineItem {
	result := make([]TimelineItem, 0)
	for _, item := range items {
		if item.Layer != layer {
			continue
		}
		if itemType != "" && item.Type != itemType {
			continue
		}
		result = append(result, item)
	}
	return result
}
