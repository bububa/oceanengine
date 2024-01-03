package enum

// AnchorShareType 锚点共享关系
type AnchorShareType string

const (
	// AnchorShareType_OWN_ANCHOR 自有锚点
	AnchorShareType_OWN_ANCHOR AnchorShareType = "OWN_ANCHOR"
	// AnchorShareType_SHARE_ANCHOR 共享锚点
	AnchorShareType_SHARE_ANCHOR AnchorShareType = "SHARE_ANCHOR"
)
