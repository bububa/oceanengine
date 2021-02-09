package report

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
)

type GetRequest struct {
	AgentID         uint64                   `json:"agent_id,omitempty"`         // 代理商id。
	AdvertiserID    uint64                   `json:"advertiser_id,omitempty"`    // 广告主ID
	StartDate       time.Time                `json:"start_date,omitempty"`       // 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	EndDate         time.Time                `json:"end_date,omitempty"`         // 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	GroupBy         []enum.StatGroupBy       `json:"group_by,omitempty"`         // 分组条件默认为 STAT_GROUP_BY_FIELD_STAT_TIME
	TimeGranularity enum.StatTimeGranularity `json:"time_granularity,omitempty"` // 时间粒度, 默认值: STAT_TIME_GRANULARITY_DAILY
	OrderField      string                   `json:"order_field,omitempty"`      // 排序字段，所有的统计指标均可参与排序
	OrderType       enum.OrderType           `json:"order_type,omitempty"`       // 排序方式；默认值: DESC；允许值: ASC, DESC
	Page            int                      `json:"page,omitempty"`             // 页码；默认值: 1
	PageSize        int                      `json:"page_size,omitempty"`        // 页面大小，即每页展示的数据量；默认值: 20；取值范围: 1-1000
	Filtering       *StatFiltering           `json:"filtering,omitempty"`        // 过滤字段，json格式，支持字段如下
}

type StatFiltering struct {
	AdvertiserIDs         []uint64                 `json:"advertiser_ids,omitempty"`          // 广告主 id 列表，若选填该字段，则最少应上传1个广告主id，最多支持同时查询100个广告主。
	Active                string                   `json:"active,omitempty"`                  // 用户活跃状态。可选枚举值: ACTIVE表示活跃用户,SILENT表示沉默用户,ALL表示所有用户。不填写该字段则默认值为ALL。
	FirstIndustry         string                   `json:"first_industry,omitempty"`          // 一级行业名称
	SecondIndustry        string                   `json:"secondy_industry,omitempty"`        // 二级行业名称
	AccountSource         enum.AccountSource       `json:"account_source,omitempty"`          // 账户类型。可选枚举值: LUBAN_ACCOUNT,NORMAL_ADVERTISER
	AccountStatus         enum.AccountStatus       `json:"account_status,omitempty"`          // 账户状态
	StartAuditPassTime    string                   `json:"start_audit_pass_time,omitempty"`   // 过审时间范围的开始时间。闭区间。格式：YYYY-MM-DD。表示过滤出账户资质审核通过时间在给定的过审时间范围内的广告主数据。
	EndAuditPassTime      string                   `json:"end_audit_pass_time,omitempty"`     // 过审时间范围的结束时间。闭区间。格式：YYYY-MM-DD。表示过滤出账户资质审核通过时间在给定的过审时间范围内的广告主数据。
	CampaignIDs           []uint64                 `json:"campaign_ids,omitempty"`            // 广告组id列表：按照campaign_id过滤，最多支持100个
	CampaignName          string                   `json:"campaign_name,omitempty"`           // 广告组名称：按照campaign_name过滤，字符串长度限制1-30
	AdIDs                 []uint64                 `json:"ad_ids,omitempty"`                  // 广告计划id列表：按照 ad_id 过滤，最多支持100个
	AdName                string                   `json:"ad_name,omitempty"`                 // 广告计划名称：按照ad_name过滤，字符串长度限制1-30
	CreativeIDs           []uint64                 `json:"creative_ids,omitempty"`            // 广告创意id列表：按照 creative_id 过滤，最多支持100个
	InventoryTypes        []enum.StatInventoryType `json:"inventory_types,omitempty"`         // 广告位置列表：按照广告位置过滤
	Pricings              []enum.PricingType       `json:"pricings,omitempty"`                // 出价方式列表：按照出价方式过滤
	ImageModes            []enum.ImageMode         `json:"image_modes,omitempty"`             // 素材类型列表：按照类型过滤
	CreativeMaterialModes []string                 `json:"creative_material_modes,omitempty"` // 创意类型列表：按照创意类型过滤，STATIC_ASSEMBLE 表示程序化创意，INTERVENE表示自定义创意
	LandingTypes          []enum.LandingType       `json:"landing_types,omitempty"`           // 推广目的列表：按照广告组推广目的过滤, 允许值: "LINK","APP","DPA","GOODS","STORE","SHOP","AWEME"
	Status                enum.CreativeStatus      `json:"status,omitempty"`                  // 广告创意状态：按照广告创意状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示
}

func (r GetRequest) Encode() string {
	values := &url.Values{}
	values.Set("start_date", r.StartDate.Format("2006-01-02"))
	values.Set("end_date", r.EndDate.Format("2006-01-02"))
	if r.AgentID > 0 {
		values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	}
	if r.AdvertiserID > 0 {
		values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	if r.GroupBy != nil {
		groupBy, _ := json.Marshal(r.GroupBy)
		values.Set("group_by", string(groupBy))
	}
	if r.TimeGranularity != "" {
		values.Set("time_granularity", string(r.TimeGranularity))
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
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	return values.Encode()
}
