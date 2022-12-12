package creative

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取创意列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *GetFiltering `json:"filtering,omitempty"`
	// Fields 查询字段集合, 如果指定, 则返回结果数组中, 每个元素是包含所查询字段的字典，默认全部指定
	// 允许值: "creative_id", "ad_id", "advertiser_id", "status","opt_status", "image_mode", "title", "creative_word_ids","third_party_id", "image_ids", "image_id", "video_id","materials"
	Fields []string `json:"fields,omitempty"`
	// Page 页数默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值: 10，大小上限1000
	PageSize int `json:"page_size,omitempty"`
	// Cursor 页码游标值，第一次拉取，传入0
	// 同时传入时，cursor优先级大于page
	// 注：page+page_size与cursor+count为两种分页方式
	// cursor+count适用于获取数据记录数≥10000的场景
	Cursor int `json:"cursor,omitempt"`
	// Count 页面数据量
	// 注：page+page_size与cursor+count为两种分页方式
	// cursor+count适用于获取数据记录数≥10000的场景
	Count int `json:"count,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if len(r.Fields) > 0 {
		fields, _ := json.Marshal(r.Fields)
		values.Set("fields", string(fields))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Cursor > 0 || r.Count > 0 {
		values.Set("cursor", strconv.Itoa(r.Cursor))
		values.Set("count", strconv.Itoa(r.Count))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetFiltering 过滤条件
type GetFiltering struct {
	// CampaignID 按照campaign_id过滤
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// AdID 按照ad_id过滤
	AdID uint64 `json:"ad_id,omitempty"`
	// CreativeIDs 按照creative_id过滤，最多传100个。创意ID需属于当前广告主，否则会报错
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// LandingType 按照广告组推广目的过滤
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// Pricing 按照广告计划出价方式过滤
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// Status 按照创意状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示
	Status enum.CreativeStatus `json:"status,omitempty"`
	// ImageMode 按照创意素材类型过滤
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// CreativeCreateTime 广告创意创建时间，格式yyyy-MM-dd，表示过滤出当天创建的广告创意
	CreativeCreateTime string `json:"creative_create_time,omitempty"`
	// CreativeModifyTime 广告创意更新时间，格式yyyy-MM-dd，表示过滤出当天更新的广告创意
	CreativeModifyTime string `json:"creative_modify_time,omitempty"`
}
