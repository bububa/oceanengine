package star

import "github.com/bububa/oceanengine/marketing-api/enum"

// Order 星图客户任务订单
type Order struct {
	// ID 订单id
	ID uint64 `json:"order_id,omitempty"`
	// Title 作品名称
	Title string `json:"title,omitempty"`
	// UniversalOrderStatus 订单状态
	UniversalOrderStatus enum.UniversalOrderStatus `json:"universal_order_status,omitempty"`
	// AuthorID 达人id
	AuthorID uint64 `json:"author_id,omitempty"`
	// AuthorName 达人名称
	AuthorName string `json:"author_name,omitempty"`
	// AvatarURL 达人头像
	AvatarURL string `json:"avatar_url,omitempty"`
	// CreateTime 订单创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
	// DemandID 任务id
	DemandID uint64 `json:"demand_id,omitempty"`
	// HeadImageURL 封面图
	HeadImageURL string `json:"head_image_url,omitempty"`
	// VideoID 视频id，每个视频唯一（建议使用item_id）
	VideoID string `json:"video_id,omitempty"`
	// ItemID 视频id，与星图平台前端video_url中展现的视频id一致，每个视频唯一
	ItemID uint64 `json:"item_id,omitempty"`
	// VideoURL 视频链接
	VideoURL string `json:"video_url,omitempty"`
}
