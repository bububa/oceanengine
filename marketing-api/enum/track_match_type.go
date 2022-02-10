package enum

// TrackMatchType 归因方式
type TrackMatchType = int

const (
	// Track_CLICK 点击归因
	Track_CLICK TrackMatchType = 0
	// Track_IMPRESSION 展现归因
	Track_IMPRESSION TrackMatchType = 1
	// Track_VALID_IMPRESSION 有效展现归因
	Track_VALID_IMPRESSION TrackMatchType = 2
)
