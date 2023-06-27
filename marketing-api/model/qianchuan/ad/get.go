package ad

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取账户下计划列表（不含创意） API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// RequestAwemeInfo 是否包含抖音号信息，允许值：0：不包含；1：包含；默认不返回
	RequestAwemeInfo int `json:"request_aweme_info,omitempty"`
	// Filtering 过滤条件，若此字段不传，或传空则视为无限制条件
	Filtering *GetFiltering `json:"filtering,omitempty"`
	// Page 当前页码: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小 默认值: 10， 取值范围：1-1000
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.RequestAwemeInfo == 1 {
		values.Set("request_aweme_info", "1")
	}
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
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetFiltering 过滤条件
type GetFiltering struct {
	// IDs 按计划ID过滤，list长度限制 1-100
	IDs []uint64 `json:"ids,omitempty"`
	// AdName 按计划名称过滤，长度为1-30个字符
	AdName string `json:"ad_name,omitempty"`
	// Status 按计划状态过滤，不传入即默认返回“所有不包含已删除”，其他规则详见【附录-广告计划查询状态】
	Status qianchuan.AdStatusForSearch `json:"status,omitempty"`
	// CampaignScene 按营销场景过滤，允许值：DAILY_SALE日常销售（默认）
	CampaignScene qianchuan.CampaignScene `json:"campaign_scene,omitempty"`
	// MarketingScene 按广告类型过滤，允许值：ALL 全部，FEED 通投广告，SEARCH 搜索广告，默认为FEED
	MarketingScene qianchuan.MarketingScene `json:"marketing_scene,omitempty"`
	// PromotionWay 按推广方式过滤，允许值：STANDARD专业推广、SIMPLE极速推广
	PromotionWay enum.PromotionWay `json:"promotion_way,omitempty"`
	// MarketingGoal 广告组营销目标，允许值：VIDEO_PROM_GOODS：短视频带货、LIVE_PROM_GOODS：直播带货
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// CampaignID 按广告组ID过滤
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// AdCreateStartDate 计划创建开始时间，格式："yyyy-mm-dd"
	AdCreateStartDate string `json:"ad_create_start_date,omitempty"`
	// AdCreateEndDate 计划创建结束时间，与ad_create_start_date搭配使用，格式："yyyy-mm-dd"，时间跨度不能超过180天
	AdCreateEndDate string `json:"ad_create_end_date,omitempty"`
	// AdModifyTime 计划修改时间，精确到小时，格式："yyyy-mm-dd HH"
	AdModifyTime string `json:"ad_modify_time,omitempty"`
	// AwemeID 根据抖音号过滤
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// AutoManageFilter 按是否为托管计划过滤，允许值：ALL ：不限，AUTO_MANAGE ：托管计划，NORMAL ：非托管计划，默认为ALL
	AutoManageFilter string `json:"auto_manage_filter,omitempty"`
}

// GetResponse 获取账户下计划列表（不含创意） API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// PageInfo 分页信息
	PageInfo model.PageInfo `json:"page_info,omitempty"`
	// FailList 获取失败的计划id列表
	FailList []uint64 `json:"fail_list,omitempty"`
	// List 计划列表
	List []Ad `json:"list,omitempty"`
}
