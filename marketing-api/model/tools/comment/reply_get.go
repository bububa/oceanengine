package comment

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ReplyGetRequest 获取评论回复列表 API Request
type ReplyGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CommentID 父评论id
	CommentID uint64 `json:"comment_id,omitempty"`
	// Page 页数 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r ReplyGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("comment_id", strconv.FormatUint(r.CommentID, 10))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ReplyGetResponse 获取评论回复列表 API Response
type ReplyGetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *ReplyGetResponseData `json:"data,omitempty"`
}

// ReplyGetResponseData json 返回值
type ReplyGetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ReplyList 评论列表
	ReplyList []Reply `json:"reply_list,omitempty"`
}
