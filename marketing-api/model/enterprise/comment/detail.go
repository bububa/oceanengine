package comment

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DetailRequest 获取评论从属信息 API Request
type DetailRequest struct {
	// CcAccountID 纵横组织ID，纵横组织管理员或协作者授权后，通过【获取已授权账户】接口，查询到“账号角色为CUSTOMER_ADMIN-管理员授权的纵横组织、或CUSTOMER_OPERATOR-协作者授权的纵横组织”的advertiser_id，即为纵横组织ID
	CcAccountID uint64 `json:"cc_account_id,omitempty"`
	// EDouyinID 企业号账户ID，获取到授权的纵横组织ID后，再通过【获取纵横组织下资产账户列表（分页）】接口，查询到e_douyin_id，即为企业号账户ID，需确保传入的企业号账户ID与纵横组织ID已建立绑定关系，且绑定关系未解除
	EDouyinID string `json:"e_douyin_id,omitempty"`
	// CommentID 评论id
	CommentID uint64 `json:"comment_id,omitempty"`
}

// Encode implement GetRequest interface
func (r DetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("cc_account_id", strconv.FormatUint(r.CcAccountID, 10))
	values.Set("e_douyin_id", r.EDouyinID)
	values.Set("comment_id", strconv.FormatUint(r.CommentID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DetailResponse 获取评论从属信息 API Response
type DetailResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data *DetailResult `json:"data,omitempty"`
}

// DetailResult 评论从属信息
type DetailResult struct {
	// ItemID 评论所属的抖音视频id，若评论属于广告视频，则无以下信息
	ItemID uint64 `json:"item_id,omitempty"`
	// ItemTitle 评论所属的抖音视频标题
	ItemTitle string `json:"item_title,omitempty"`
	// ItemCoverURL 抖音视频封面url
	ItemCoverURL string `json:"item_cover_url,omitempty"`
	// ItemDuration 抖音视频时长 单位是秒
	ItemDuration int64 `json:"item_duration,omitempty"`
	// ItemCreateTime 抖音视频发布时间
	ItemCreateTime string `json:"item_create_time,omitempty"`
	// ItemDiggCount 评论所属的抖音视频的点赞数（2022年6月后，该字段不再支持）
	ItemDiggCount int `json:"item_digg_count,omitempty"`
	// ReplyCommentID 被回复的评论id，即当前评论的上级评论id
	ReplyCommentID uint64 `json:"reply_comment_id,omitempty"`
	// ReplyCommentText 上级评论内容
	ReplyCommentText string `json:"reply_comment_text,omitempty"`
	// ReplyCommentImage 上级评论图片url
	ReplyCommentImage string `json:"reply_comment_image,omitempty"`
	// ReplyCommentCreateTime 上级评论发布时间
	ReplyCommentCreateTime string `json:"reply_comment_create_time,omitempty"`
	// ReplyOpenID 上级评论用户open_id
	ReplyOpenID string `json:"reply_open_id,omitempty"`
}
