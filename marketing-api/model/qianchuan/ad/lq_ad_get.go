package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// LqAdGetRequest 获取低效计划列表 API Request
type LqAdGetRequest struct {
	// AdvertiserID 千川广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *LqAdGetFilter `json:"filtering,omitempty"`
}

type LqAdGetFilter struct {
	// MarketingScene 按广告类型过滤，允许值：
	// ALL 全部
	// FEED 通投广告，默认为FEED
	// SEARCH 搜索广告
	// SHOPPING_MALL商城广告
	MarketingScene qianchuan.MarketingScene `json:"marketing_scene,omitempty"`
}

// Encode implement GetRequest interface
func (r LqAdGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// LqAdGetResponse 获取低效计划列表 API Response
type LqAdGetResponse struct {
	model.BaseResponse
	Data struct {
		// AdIDs 低效计划ID列表
		AdIDs []uint64 `json:"ad_ids,omitempty"`
	} `json:"data,omitempty"`
}
