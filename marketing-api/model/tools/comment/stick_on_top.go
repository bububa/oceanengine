package comment

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StickOnTopRequest 置顶评论 API Request
type StickOnTopRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CommentID 评论id
	CommentID uint64 `json:"comment_id,omitempty"`
	// 操作类型，可选值:
	// CANCEL_STICK: 取消置顶
	// STICK_ON_TOP: 置顶
	StickType enum.CommentStickType `json:"stick_type,omitempty"`
}

// Encode implement PostRequest interface
func (r StickOnTopRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
