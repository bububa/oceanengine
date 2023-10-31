package comment

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ReplyRequest 回复评论 API Request
type ReplyRequest struct {
	// CcAccountID 纵横组织ID，纵横组织管理员或协作者授权后，通过【获取已授权账户】接口，查询到“账号角色为CUSTOMER_ADMIN-管理员授权的纵横组织、或CUSTOMER_OPERATOR-协作者授权的纵横组织”的advertiser_id，即为纵横组织ID
	CcAccountID uint64 `json:"cc_account_id,omitempty"`
	// EDouyinID 企业号账户ID，获取到授权的纵横组织ID后，再通过【获取纵横组织下资产账户列表（分页）】接口，查询到e_douyin_id，即为企业号账户ID，需确保传入的企业号账户ID与纵横组织ID已建立绑定关系，且绑定关系未解除
	EDouyinID string `json:"e_douyin_id,omitempty"`
	// CommentID 评论id列表
	CommentIDs []uint64 `json:"comment_ids,omitempty"`
	// ReplyText 回复内容 最多100字 汉字中英文都算1个字
	ReplyText string `json:"reply_text,omitempty"`
}

// Encode implement PostRequest interface
func (r ReplyRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ReplyResponse 回复评论 API Response
type ReplyResponse struct {
	model.BaseResponse
	Data *ReplyResult `json:"data,omitempty"`
}

type ReplyResult struct {
	// Success 回复成功的回复
	Success []CommentReply `json:"success,omitempty"`
	// Errors 回复失败的回复
	Errors []ReplyError `json:"errors,omitempty"`
}

// ReplyError 回复失败的回复
type ReplyError struct {
	// ReplyID 回复失败的回复id
	ReplyID uint64 `json:"reply_id,omitempty"`
	// Message 失败原因
	Message string `json:"message,omitempty"`
}

// Error implement errors interface
func (r ReplyError) Error() string {
	return util.StringsJoin(r.Message, ", 回复ID: ", strconv.FormatUint(r.ReplyID, 10))
}
