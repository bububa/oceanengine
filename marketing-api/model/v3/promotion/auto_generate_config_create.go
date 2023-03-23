package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AutoGenerateConfigCreateRequest 新建/修改白盒配置 API Request
type AutoGenerateConfigCreateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionID 广告ID，
	// 新建时无需传入，
	// 修改时必须传入，传入绑定的promotion_id
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// StrategyData 策略数据(列表中strategy_id需唯一, 即同一个策略（strategy）的策略配置项列表(strategy_state)，需放到同一个对象内，不可分开传递)
	StrategyData *StrategyData `json:"strategy_data,omitempty"`
}

// Encode implement PostRequest interface
func (r AutoGenerateConfigCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AutoGenerateConfigCreateResponse 新建/修改白盒配置 API Response
type AutoGenerateConfigCreateResponse struct {
	model.BaseResponse
	Data struct {
		// ConfigID 配置ID
		ConfigID uint64 `json:"config_id,omitempty"`
	} `json:"data,omitempty"`
}
