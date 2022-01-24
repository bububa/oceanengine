package creative

import "github.com/bububa/oceanengine/marketing-api/enum"

// PlayableMaterial 基础试玩素材
type PlayableMaterial struct {
	// PlayableInfo 试玩信息
	PlayableInfo PlayableInfo `json:"playable_info,omitempty"`
}

// PlayableInfo 试玩信息
type PlayableInfo struct {
	// PlayableURL 试玩素材URL，可通过【获取试玩素材列表】进行获取。 只有穿山甲激励视频可以使用试玩素材，同时素材需要审核通过
	PlayableURL string `json:"playable_url,omitempty"`
	// PlayableUrlBasic 基础试玩素材url
	PlayableUrlBasic string `json:"playable_url_basic,omitempty"`
	// PlayableOrientation 基础试玩素材方向
	PlayableOrientation enum.PlayableOrientation `json:"playable_orientation,omitempty"`
	// PreviewUrl 基础试玩素材预览图
	PreviewUrl string `json:"preview_url,omitempty"`
}
