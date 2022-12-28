package ad

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

type IListFilter interface {
	GetIDs() []uint64
	GetName() string
	GetCreateTime() string
	GetModifyTime() string
}

// GetRequest 获取广告计划 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *GetFiltering `json:"filtering,omitempty"`
	// Fields 查询字段集合, 如果指定, 则返回结果数组中, 每个元素是包含所查询字段的字典允许值: "id", "name", "budget", "budget_mode", "status", "opt_status","open_url", "modify_time", "start_time", "end_time", "bid","advertiser_id", "pricing", "flow_control_mode", "download_url", quick_app_url, "inventory_type", "schedule_type", "app_type", "cpa_bid","cpa_skip_first_phrase", "audience", "external_url", "package","campaign_id", "ad_modify_time", "ad_create_time","audit_reject_reason", "retargeting_type", "retargeting_tags","convert_id", "interest_tags", "hide_if_converted","external_actions", "device_type","auto_extend_enabled", "auto_extend_targets", "dpa_lbs", "dpa_city", "dpa_province", "dpa_recommend_type", "roi_goal","subscribe_url","form_id","form_index","app_desc","app_thumbnails"
	Fields []string `json:"fields,omitempty"`
	// Page 页数默认值: 1，page必须大于0
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值:10，page_size范围为1-1000
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
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetFiltering 过滤条件
type GetFiltering struct {
	// IDs 按广告计划ID过滤，范围为1-100
	IDs []uint64 `json:"ids,omitempty"`
	// AdName 按广告计划name过滤，长度为1-30个字符
	AdName string `json:"ad_name,omitempty"`
	// PricingList 按出价方式过滤
	PricingList enum.PricingType `json:"pricing_list,omitempty"`
	// Status 按计划状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示
	Status enum.AdStatus `json:"status,omitempty"`
	// AdCreateTime 广告计划创建时间，格式"yyyy-mm-dd"，表示过滤出当天创建的广告计划
	AdCreateTime string `json:"ad_create_time,omitempty"`
	// AdModifyTime 广告计划更新时间，格式"yyyy-mm-dd"，表示过滤出当天更新的广告计划
	AdModifyTime string `json:"ad_modify_time,omitempty"`
}

func (f GetFiltering) GetIDs() []uint64 {
	return f.IDs
}
func (f GetFiltering) GetName() string {
	return f.AdName
}
func (f GetFiltering) GetCreateTime() string {
	return f.AdCreateTime
}
func (f GetFiltering) GetModifyTime() string {
	return f.AdModifyTime
}

// GetResponse 获取广告计划 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// List 广告数组
	List []Ad `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
