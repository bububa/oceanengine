package campaign

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取广告组 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件，若此字段不传，或传空则视为无限制条件
	Filtering *GetFiltering `json:"filtering,omitempty"`
	// Fields 查询字段集合, 如果指定, 则返回结果数组中, 每个元素是包含所查询字段的字典；允许值:"id", "name","budget", "budget_mode","landing_type","status","modify_time", "status","modify_time","campaign_modify_time","campaign_create_time"一定会返回
	Fields []string `json:"fields,omitempty"`
	// Page 当前页码: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小 默认值: 10， 取值范围：1-1000
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if r.Fields != nil {
		fields, _ := json.Marshal(r.Fields)
		values.Set("fields", string(fields))
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

// GetFiltering 过滤条件
type GetFiltering struct {
	// IDs 广告组ID过滤，数组，不超过100个
	IDs []uint64 `json:"ids,omitempty"`
	// CampaignName 广告组name过滤，长度为1-30个字符，其中1个中文字符算2位
	CampaignName string `json:"campaign_name,omitempty"`
	// LandingType 广告组推广目的过滤
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// Status 广告组状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示
	Status enum.CampaignStatus `json:"status,omitempty"`
	// CampaignCreateTime 广告组创建时间，格式yyyy-mm-dd，表示过滤出当天创建的广告组
	CampaignCreateTime string `json:"campaign_create_time,omitempty"`
}
