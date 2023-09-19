package comment

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/enterprise"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListGetRequest 获取评论列表 API Request
type ListGetRequest struct {
	// CcAccountID 纵横组织ID，纵横组织管理员或协作者授权后，通过【获取已授权账户】接口，查询到“账号角色为CUSTOMER_ADMIN-管理员授权的纵横组织、或CUSTOMER_OPERATOR-协作者授权的纵横组织”的advertiser_id，即为纵横组织ID
	CcAccountID uint64 `json:"cc_account_id,omitempty"`
	// EDouyinID 企业号账户ID，获取到授权的纵横组织ID后，再通过【获取纵横组织下资产账户列表（分页）】接口，查询到e_douyin_id，即为企业号账户ID，需确保传入的企业号账户ID与纵横组织ID已建立绑定关系，且绑定关系未解除
	EDouyinID string `json:"e_douyin_id,omitempty"`
	// StartTime 开始时间，默认30天前 只支持天级的。e.g.2022-01-01
	// 开始时间必须大于2021-01-01
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间，默认当天 只支持天级的。e.g.2022-01-02
	// 结束时间必须晚于开始时间，查询时间跨度不能大于3个月
	EndTime string `json:"end_time,omitempty"`
	// OrderField 排序字段，支持排序字段：create_time评论创建时间
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序方式，允许值: ASC、DESC默认值：DESC
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页数，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，支持范围1-100 之间，默认值：20
	PageSize int `json:"page_size,omitempty"`
	// Filter 筛选字段
	Filter *ListGetFilter `json:"filter,omitempty"`
}

// ListGetFilter 筛选字段
type ListGetFilter struct {
	// Content 关键词筛选，默认按照匹配度排序
	Content string `json:"content,omitempty"`
	// ReplyStatusByEDouyin 回复状态(是否被企业号回复)
	// 允许值：NO_REPLY未回复、REPLY已回复
	ReplyStatusByEDouyin enterprise.CommentReplyStatus `json:"reply_status_by_e_douyin,omitempty"`
	// Level 评论层级
	// 允许值：LEVEL_ONE一级评论、LEVEL_TWO二级评论，即回复
	Level enterprise.CommentLevel `json:"level,omitempty"`
	// Source 流量来源
	// 允许值：FROM_NATURAL自然流量、FROM_DOUPLUSDou+、FROM_PERFORM竞价广告、FROM_BRAND品牌广告、FROM_OTHER其他流量
	Source enterprise.CommentSource `json:"source,omitempty"`
	// AdvertiserID 流量来源的广告账户id列表
	AdvertiserID []uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 流量来源的广告组id列表
	CampaignID []uint64 `json:"campaign_id,omitempty"`
	// AdID 流量来源的广告计划id列表
	AdID []uint64 `json:"ad_id,omitempty"`
	// CreativeID 流量来源的广告创意id列表
	CreativeID []uint64 `json:"creative_id,omitempty"`
	// ItemType 评论所属的视频类型 允许值：ITEM_AD广告视频、ITEM_CONTENT非广告视频(抖音视频)
	ItemType enterprise.ItemType `json:"item_type,omitempty"`
	// MaterialID 评论所属的广告视频素材id，通过【获取视频素材】接口获取
	MaterialID uint64 `json:"material_id,omitempty"`
	// ItemID 评论所属的抖音视频id，通过【获取企业号视频列表】接口获取
	ItemID []uint64 `json:"item_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ListGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("cc_account_id", strconv.FormatUint(r.CcAccountID, 10))
	values.Set("e_douyin_id", r.EDouyinID)
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
	}
	if r.Filter != nil {
		values.Set("filter", string(util.JSONMarshal(r.Filter)))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
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

// ListGetResponse 获取评论列表 API Response
type ListGetResponse struct {
	model.BaseResponse
	Data *ListGetResult `json:"data,omitempty"`
}

type ListGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// CommentList 评论列表
	CommentList []Comment `json:"comment_list,omitempty"`
}
