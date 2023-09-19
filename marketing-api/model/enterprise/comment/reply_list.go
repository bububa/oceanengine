package comment

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ReplyListRequest 获取评论的回复 API Request
type ReplyListRequest struct {
	// CcAccountID 纵横组织ID，纵横组织管理员或协作者授权后，通过【获取已授权账户】接口，查询到“账号角色为CUSTOMER_ADMIN-管理员授权的纵横组织、或CUSTOMER_OPERATOR-协作者授权的纵横组织”的advertiser_id，即为纵横组织ID
	CcAccountID uint64 `json:"cc_account_id,omitempty"`
	// EDouyinID 企业号账户ID，获取到授权的纵横组织ID后，再通过【获取纵横组织下资产账户列表（分页）】接口，查询到e_douyin_id，即为企业号账户ID，需确保传入的企业号账户ID与纵横组织ID已建立绑定关系，且绑定关系未解除
	EDouyinID string `json:"e_douyin_id,omitempty"`
	// CommentID 评论id
	CommentID uint64 `json:"comment_id,omitempty"`
	// Page 页数，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，支持范围1-100 之间，默认值：20
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements the GetRequest interface
func (r ReplyListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("cc_account_id", strconv.FormatUint(r.CcAccountID, 10))
	values.Set("e_douyin_id", r.EDouyinID)
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

// ReplyListResponse 获取评论的回复 API Response
type ReplyListResponse struct {
	model.BaseResponse
	Data *ReplyListResult `json:"data,omitempty"`
}

type ReplyListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ReplyList 回复列表
	ReplyList []CommentReply `json:"reply_list,omitempty"`
}
