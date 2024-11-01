package security

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ScoreTotalGetRequest 查询账户累计积分 API Request
type ScoreTotalGetRequest struct {
	// Filter 过滤器
	Filter *ScoreTotalGetFilter `json:"filter,omitempty"`
	// BusinessLine 业务线，允许值：
	// AD 巨量广告
	BusinessLine string `json:"business_line,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页数，默认值：1，page范围为[1,99999]
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值：10，page_size范围为[1,100]
	PageSize int `json:"page_size,omitempty"`
}

type ScoreTotalGetFilter struct {
	// Year 积分年度，如：2000
	Year string `json:"year,omitempty"`
	// IllegalType 违规类型 可选值:
	// GENERAL 一般违规（AD）
	// SERIOUS 严重违规（AD）
	IllegalType enum.SecurityScoreIllegalType `json:"illegal_type,omitempty"`
}

func (r ScoreTotalGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("page", strconv.Itoa(r.Page))
	values.Set("page_size", strconv.Itoa(r.PageSize))
	if r.BusinessLine != "" {
		values.Set("business_line", r.BusinessLine)
	}
	if r.Filter != nil {
		values.Set("filter", string(util.JSONMarshal(r.Filter)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ScoreTotalGetResponse 查询账户累计积分 API Response
type ScoreTotalGetResponse struct {
	Data *ScoreTotalGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type ScoreTotalGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ScoreInfoList 积分详情
	ScoreInfoLIst []ScoreInfo `json:"score_info_list,omitempty"`
}

// ScoreInfo 积分详情
type ScoreInfo struct {
	// IllegalType 违规类型 可选值:
	// GENERAL 一般违规（AD）
	// SERIOUS 严重违规（AD）
	IllegalType enum.SecurityScoreIllegalType `json:"illegal_type,omitempty"`
	// Year 积分年度
	Year string `json:"year,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Score 积分
	Score int `json:"score,omitempty"`
}
