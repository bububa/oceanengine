package comment

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OperateRequest 评论操作 API Request
type OperateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CommentIDs 评论id列表，只允许传入小于等于20个评论id
	CommentIDs []uint64 `json:"comment_ids,omitempty"`
	// ReplyText 回复内容，最多支持100字
	ReplyText string `json:"reply_text,omitempty"`
}

// Encode implement PostRequest interface
func (r OperateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// OperateResponse 评论操作 API Response
type OperateResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data struct {
		// SuccessCommentIDs 回复成功的评论id列表
		SuccessCommentIDs []uint64 `json:"success_comment_ids,omitempty"`
	} `json:"data,omitempty"`
}
