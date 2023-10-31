package creative

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StrategyListRequest 获取模板（白盒策略）列表 API Request
type StrategyListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StrategyTypes 策略类型(支持多选)
	StrategyTypes []enum.StrategyType `json:"strategy_type,omitempty"`
	// Page 页码，从1开始, 默认1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小, 取值[1-100], 默认10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r StrategyListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("strategy_types", string(util.JSONMarshal(r.StrategyTypes)))
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

// StrategyListResponse 获取模板（白盒策略）列表 API Response
type StrategyListResponse struct {
	model.BaseResponse
	Data *StrategyListData `json:"data,omitempty"`
}

type StrategyListData struct {
	// StrategyModels 策略列表
	StrategyModels []promotion.StrategyData `json:"strategy_models,omitempty"`
	// PageInfo 翻页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
