package wechat

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GameListRequest 获取微信小游戏列表 API Request
type GameListRequest struct {
	// AccountID 广告主ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType  账户类型
	// 允许值：BP 巨量纵横组织账号、AD 广告主账号
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// Filtering 过滤条件
	Filtering *AppletListFilter `json:"filtering,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值为10，最大值100
	PageSize int `json:"page_size,omitempty"`
}

// GameListFilter 过滤条件
type GameListFilter struct {
	// Name 小程序名称或备注的模糊匹配
	Name string `json:"name,omitempty"`
	// AssetStatus  资产状态 允许值：升级版UPGRADED、 原版ORIGINAL
	AssetStatus enum.WechatAssetStatus `json:"asset_status,omitempty"`
	// AuditStatus 审核状态 允许值: AUDIT_ACCEPTED 审核通过、AUDITING 审核中、AUDIT_REJECTED 审核不通过。不传表示全部状态
	AuditStatus enum.WechatAuditStatus `json:"audit_status,omitempty"`
	// SearchType  搜索类型
	// 允许值： CREATE_ONLY 只查询该账户创建的应用（默认值）、SHARE_ONLY 只查询被共享的应用
	SearchType enum.WechatSearchType `json:"search_type,omitempty"`
	// CreateTime 按创建时间查询的时间范围
	CreateTime *ListTimeRange `json:"create_time,omitempty"`
}

// Encode implement GetRequest interface
func (r GameListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// GameListResponse 获取微信小游戏列表 API Response
type GameListResponse struct {
	model.BaseResponse
	Data *GameListResult `json:"data,omitempty"`
}

type GameListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 微信小游戏列表
	List []WechatGame `json:"list,omitempty"`
}
