package ad

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
)

type GetRequest struct {
	AdvertiserID uint64        `json:"advertiser_id,omitempty"`
	Filtering    *GetFiltering `json:"filtering,omitempty"`
	Fields       []string      `json:"fields,omitempty"` // 查询字段集合, 如果指定, 则返回结果数组中, 每个元素是包含所查询字段的字典允许值: "id", "name", "budget", "budget_mode", "status", "opt_status","open_url", "modify_time", "start_time", "end_time", "bid","advertiser_id", "pricing", "flow_control_mode", "download_url", quick_app_url, "inventory_type", "schedule_type", "app_type", "cpa_bid","cpa_skip_first_phrase", "audience", "external_url", "package","campaign_id", "ad_modify_time", "ad_create_time","audit_reject_reason", "retargeting_type", "retargeting_tags","convert_id", "interest_tags", "hide_if_converted","external_actions", "device_type","auto_extend_enabled", "auto_extend_targets", "dpa_lbs", "dpa_city", "dpa_province", "dpa_recommend_type", "roi_goal","subscribe_url","form_id","form_index","app_desc","app_thumbnails"
	Page         int           `json:"page,omitempty"`
	PageSize     int           `json:"page_size,omitempty"`
}

func (r GetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

type GetFiltering struct {
	IDs          []uint64         `json:"ids,omitempty"`            // 按广告计划ID过滤，范围为1-100
	AdName       string           `json:"ad_name,omitempty"`        // 按广告计划name过滤，长度为1-30个字符
	PricingList  enum.PricingType `json:"pricing_list,omitempty"`   // 按出价方式过滤
	Status       enum.AdStatus    `json:"status,omitempty"`         // 按计划状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示
	AdCreateTime string           `json:"ad_create_time,omitempty"` // 广告计划创建时间，格式"yyyy-mm-dd"，表示过滤出当天创建的广告计划
	AdModifyTime string           `json:"ad_modify_time,omitempty"` // 广告计划更新时间，格式"yyyy-mm-dd"，表示过滤出当天更新的广告计划
}
