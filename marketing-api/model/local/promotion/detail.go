package promotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DetailRequest 获取广告详情 API Request
type DetailRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
}

// Encode implements GetRequest interface
func (r DetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	values.Set("promotion_id", strconv.FormatUint(r.PromotionID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DetailResponse 获取广告详情 API Response
type DetailResponse struct {
	Data *PromotionDetail `json:"data,omitempty"`
	model.BaseResponse
}
