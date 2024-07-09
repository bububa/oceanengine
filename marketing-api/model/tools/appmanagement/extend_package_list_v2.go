package appmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ExtendPackageListV2Request 查询应用分包列表（支持所有账户体系） API Request
type ExtendPackageListV2Request struct {
	// AccountID 账户id，accout_type类型对应账户ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型
	// 允许值：BP 巨量纵横组织、AD 广告主账号、STAR 星图
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// PackageID 应用包ID，获取方法见接口文档
	PackageID string `json:"package_id,omitempty"`
	// Page 页数 默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	// 默认值：10，范围：1～1000
	PageSize int `json:"page_size,omitempty"`
	// Filtering 过滤条件
	Filtering *ExtendPackageListV2Filter `json:"filtering,omitempty"`
}

// ExtendPackageListV2Filter 过滤条件
type ExtendPackageListV2Filter struct {
	// Status 状态
	// 允许值："ALL"：全部
	// "NOT_UPDATE"： 未更新
	// "CREATING"：创建中
	// "UPDATING"：更新中
	// "PUBLISHED"：已发布
	// "CREATION_FAILED"：创建失败
	// "UPDATE_FAILED"：更新失败
	// 默认值："ALL"
	Status ExtendPackageStatus `json:"status,omitempty"`
	// ChannelID 渠道号，可通过渠道号筛选应用分包，单次支持传入渠道号个数<=50
	ChannelID []string `json:"channel_id,omitempty"`
	// UpdateTime 更新时间，可通过最近更新时间筛选应用分包，传入开始时间和结束时间查询该时间范围内更新过的应用分包
	UpdateTime *ExternalPackageListV2UpdateTime `json:"update_time,omitempty"`
}

type ExternalPackageListV2UpdateTime struct {
	// StartTime 开始时间，示例:"2024-07-01 00:00:00"
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间，示例:"2024-07-01 00:00:00"
	EndTime string `json:"end_time,omitempty"`
}

// Encode implement GetRequest interface
func (r ExtendPackageListV2Request) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	values.Set("package_id", r.PackageID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page != 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
