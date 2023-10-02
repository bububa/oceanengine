package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AssetLinkListRequest 获取字节小程序/小游戏详情内容 API Request
type AssetLinkListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *AssetLinkListFilter `json:"filtering,omitempty"`
	// Page 页码,默认值:1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小,默认值:10，最大值100
	PageSize int `json:"page_size,omitempty"`
}

type AssetLinkListFilter struct {
	// InstanceID 资产id
	InstanceID string `json:"instance_id,omitempty"`
	// CreateTime 按创建时间查询的时间范围
	CreateTime *model.DateRange `json:"create_time,omitempty"`
}

// Encode implement GetRequest interface
func (r AssetLinkListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
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

// AssetLinkListResponse 获取字节小程序/小游戏详情内容 API Response
type AssetLinkListResponse struct {
	model.BaseResponse
	Data *AssetLinkListResult `json:"data,omitempty"`
}

type AssetLinkListResult struct {
	// PageInfo 页面信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 字节小游戏/小程序列表
	List []AssetLink `json:"list,omitempty"`
}

// AssetLink 字节小游戏/小程序
type AssetLink struct {
	// LinkID 链接id
	LinkID uint64 `json:"link_id,omitempty"`
	// InstanceID 资产id
	InstanceID uint64 `json:"instance_id,omitempty"`
	// AppID app id
	AppID string `json:"app_id,omitempty"`
	// AdvertiserID 所属广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Link 启动链接
	Link string `json:"link,omitempty"`
	// LinkRemark 链接备注
	LinkRemark string `json:"link_remark,omitempty"`
	// StartPage 启动页面
	StartPage string `json:"start_page,omitempty"`
	// StartParam 启动参数
	StartParam string `json:"start_param,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ModifyTime 更新时间
	ModifyTime string `json:"modify_time,omitempty"`
}
