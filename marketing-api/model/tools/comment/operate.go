package comment

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// OperateRequest 评论操作 API Request
type OperateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CommentIDs 评论id列表; operate_type类型为"REPLY_TO_REPLY"、"STICK"、"CANCEL_STICK"时只允许传入1个评论ID；operate_type类型为"REPLY"时，只允许传入小于等于20个评论ID;operate_type类型为"HIDE"时，只允许传入小于等于50个评论ID
	CommentIDs []uint64 `json:"comment_ids,omitempty"`
	// InventoryType 广告位 允许值："INVENTORY_AWEME_FEED"（抖音）
	InventoryType string `json:"inventory_type,omitempty"`
	// OperateType 操作类型，详情见【附录-评论操作类型】;允许值："REPLY"、"REPLY_TO_REPLY"、"STICK_ON_TOP"、"CANCEL_STICK"、"HIDE"、"BLOCK_USERS"（屏蔽用户目前只支持抖音）
	OperateType enum.CommentOperateType `json:"operate_type,omitempty"`
	// ReplyText 回复内容，当operate_type为"REPLY"，"REPLY_TO_REPLY"时必填
	ReplyText string `json:"reply_text,omitempty"`
	// ReplyID 回复的id，回复二级评论时传入, 用户可从reply_infos中获取，当operate_type为"REPLY_TO_REPLY"时必填
	ReplyID uint64 `json:"reply_id,omitempty"`
}

// Encode implement PostRequest interface
func (r OperateRequest) Encode() []byte {
	js, _ := json.Marshal(r)
	return js
}

// OperateResponse 评论操作 API Response
type OperateResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *OperateResponseData `json:"data,omitempty"`
}

// OperateResponseData json 返回值
type OperateResponseData struct {
	// ReplyInfos 评论列表
	ReplyInfos []OperateReply `json:"reply_infos,omitempty"`
	// SuccessCommentIDs 回复成功的comment_id
	SuccessCommentIDs []uint64 `json:"success_comment_ids,omitempty"`
}

// OperateReply 评论回复
type OperateReply struct {
	// ReplyToCommentID 回复的一级评论ID
	ReplyToCommentID json.Number `json:"reply_to_comment_id,omitempty"`
	// ID 回复ID
	ID json.Number `json:"id,omitempty"`
	// ReplyToReplyID 二级回复所属的上级回复ID，只有二级回复才有。二级回复：表示回复“评论中其他回复内容”
	ReplyToReplyID json.Number `json:"reply_to_reply_id,omitempty"`
	// UserID 回复用户ID
	UserID uint64 `json:"user_id,omitempty"`
	// ReplyToUserID 回复的那条回复的用户ID，只有二级回复才有
	ReplyToUserID uint64 `json:"reply_to_user_id,omitempty"`
	// Text 回复内容
	Text string `json:"text,omitempty"`
	// CreateTime 回复创建时间
	CreateTime string `json:"create_time,omitempty"`
}
