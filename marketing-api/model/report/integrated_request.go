package report

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

type IntegratedRequest struct {
	AdvertiserID  uint64               `json:"advertiser_id,omitempty"`  // 广告主ID
	StartDate     time.Time            `json:"start_date,omitempty"`     // 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	EndDate       time.Time            `json:"end_date,omitempty"`       // 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	GroupBy       []enum.StatGroupBy   `json:"group_by,omitempty"`       // 分组条件默认为 STAT_GROUP_BY_FIELD_STAT_TIME
	Fields        []string             `json:"fields,omitempty"`         // 指定需要的指标名称。如果选择了该字段，请至少提供一个指标；注意：1.如果没有指定，那么只返回支持的默认指标名称（参见文档下方说明）；2.对于不同的分组条件，支持不同的指标（参见文档下方说明）
	OrderField    string               `json:"order_field,omitempty"`    // 排序字段，所有的统计指标均可参与排序
	OrderType     enum.OrderType       `json:"order_type,omitempty"`     // 排序方式；默认值: DESC；允许值: ASC, DESC
	Page          int                  `json:"page,omitempty"`           // 页码；默认值: 1
	PageSize      int                  `json:"page_size,omitempty"`      // 页面大小，即每页展示的数据量；默认值: 20；取值范围: 1-1000
	Filtering     *IntegratedFiltering `json:"filtering,omitempty"`      // 过滤字段，json格式，如果campaign_ids不填默认按照广告主过滤，支持字段如下
	PostFiltering *PostFiltering       `json:"post_filtering,omitempty"` // 后置过滤条件，即聚合之后进行过滤。类似于 sql 语法中的 having
}

type IntegratedFiltering struct {
	CampaignIDs           []uint64                   `json:"campaign_id,omitempty"`            // 广告组id列表：按照campaign_id过滤，最多支持100个
	AdIDs                 []uint64                   `json:"ad_id,omitempty"`                  // 广告计划id列表：按照 ad_id 过滤，最多支持100个
	CreativeIDs           []uint64                   `json:"creative_id,omitempty"`            // 广告创意id列表：按照 creative_id 过滤，最多支持100个
	InventoryTypes        []enum.StatInventoryType   `json:"inventory_type,omitempty"`         // 广告位置列表：按照广告位置过滤
	Pricings              []enum.PricingType         `json:"pricing,omitempty"`                // 出价方式列表：按照出价方式过滤
	ImageModes            []enum.ImageMode           `json:"image_mode,omitempty"`             // 素材类型列表：按照类型过滤
	CreativeMaterialModes []string                   `json:"creative_material_mode,omitempty"` // 创意类型列表：按照创意类型过滤，STATIC_ASSEMBLE 表示程序化创意，INTERVENE表示自定义创意
	LandingTypes          []enum.LandingType         `json:"landing_type,omitempty"`           // 推广目的列表：按照广告组推广目的过滤, 允许值: "LINK","APP","DPA","GOODS","STORE","SHOP","AWEME"
	Bidwords              []string                   `json:"bidword,omitempty"`                // 按关键词过滤, 传入关键词，不可传空字符串，最多支持100个。只支持关键词/搜索词报表，关键词可通过【搜索广告-获取关键词】接口进行获取
	MaterialIDs           []uint64                   `json:"material_id,omitempty"`            // 按照素材id过滤，id值为为大于1的整数，最多支持100个，material_id可通过【素材管理】接口获取
	PlayableIDs           []uint64                   `json:"playable_id,omitempty"`            // 按照试玩素材id过滤，与playable_url、playable_name取交集；长度范围1-100，id值为为大于1且不超过2^63-1的整数，playable_id可通过【工具-试玩素材管理】获取
	PlayableUrls          []string                   `json:"playable_url,omitempty"`           // 按照试玩素材链接过滤，与playable_id、playable_name取交集；长度范围1-100
	PlayableName          string                     `json:"playable_name,omitempty"`          // 按照试玩素材名字过滤，支持模糊匹配，长度不超过100;
	PlayableOrientations  []enum.PlayableOrientation `json:"playable_orientation,omitempty"`   // 试玩素材展示方向，重复无效，详见【附录-试玩素材方向】，允许值: "BOTH", "PORTRAIT", "LANDSCAPE"。其中"BOTH"指横竖屏均适配的试玩素材，而非"PORTRAIT", "LANDSCAPE"的集合
}

type PostFiltering struct {
	Cost        *model.FloatRange `json:"cost,omitempty"`         // 按总消耗过滤，左闭右开区间，正负均可，min必须小于max
	Convert     *model.FloatRange `json:"convert,omitempty"`      // 按转化数过滤，左闭右开区间，正负均可，min必须小于max
	ConvertCost *model.FloatRange `json:"convert_cost,omitempty"` // 按转化成本过滤，左闭右开区间，正负均可，min必须小于max
}

func (r IntegratedRequest) Encode() string {
	values := &url.Values{}
	values.Set("start_date", r.StartDate.Format("2006-01-02"))
	values.Set("end_date", r.EndDate.Format("2006-01-02"))
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.GroupBy != nil {
		groupBy, _ := json.Marshal(r.GroupBy)
		values.Set("group_by", string(groupBy))
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
	if r.PostFiltering != nil {
		filtering, _ := json.Marshal(r.PostFiltering)
		values.Set("post_filtering", string(filtering))
	}
	return values.Encode()
}
