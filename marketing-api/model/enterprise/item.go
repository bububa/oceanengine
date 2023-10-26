package enterprise

import (
	"github.com/bububa/oceanengine/marketing-api/enum/enterprise"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// Item 视频
type Item struct {
	// ItemType 视频类型 允许值：ITEM_AD广告视频、ITEM_CONTENT非广告视频(抖音视频)
	ItemType enterprise.ItemType `json:"item_type,omitempty"`
	// MaterialID 广告视频素材id
	MaterialID int `json:"material_id,omitempty"`
	// ItemID 抖音视频id
	ItemID model.Uint64 `json:"item_id,omitempty"`
	// ItemTitle 抖音视频标题
	ItemTitle string `json:"item_title,omitempty"`
	// ItemCoverURL 抖音视频封面url
	ItemCoverURL string `json:"item_cover_url,omitempty"`
	// ItemDuration 抖音视频时长 单位是秒
	ItemDuration int64 `json:"item_duration,omitempty"`
	// ItemCreateTime 抖音视频发布时间
	ItemCreateTime string `json:"item_create_time,omitempty"`
	// CommentCount 评论数量
	CommentCount model.Int64 `json:"comment_count,omitempty"`
}
