package promotion

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CostProtectStatusGetRequest 批量获取计划成本保障状态 API Request
type CostProtectStatusGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionIDs 广告id，每次最多传入50个
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r CostProtectStatusGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.PromotionIDs)
	values.Set("promotion_ids", string(ids))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CostProtectStatusGetResponse 批量获取计划成本保障状态 API Response
type CostProtectStatusGetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data struct {
		// List .
		List []CostProtectStatus `json:"compensate_status_info_list,omitempty"`
	} `json:"data,omitempty"`
}

// CostProtectStatus guan广告成本保护状态
type CostProtectStatus struct {
	// QueryID 计划id
	QueryID uint64 `json:"query_id,omitempty"`
	// Status 计划成本保障状态
	Status string `json:"status,omitempty"`
	// CompensationStatus 成本保障状态 枚举值:
	// DEFAULT_STATUS: 当前计划未产生首次send或者不在成本保障范围内
	// IN_EFFECT: 成本保障生效中
	// INVALID: 成本保障已失效
	// CONFIRMING: 成本保障确认中
	// PAID: 成本保障已到账
	// ENDED: 成本保障已结束
	CompensationStatus string `json:"compensate_status,omitempty"`
	// InvalidReasons 成本保障失效原因 枚举值:
	// AUD_CHANGES单日修改定向超过2次,
	// BID_CHANGES单日修改出价超过2次
	// ROI_CHANGES单日修改roi_goal超过2次
	// ANTI_SPAM命中反作弊
	// BID_TYPE_EXPIRED选择的出价产品暂不支持成本保障
	// MANUAL_JUDGE_SPAM有异常的作弊行为
	// AUD_BID_CHANGES单日修改定向/出价超过2次
	// AUD_ROI_CHANGES单日修改定向/ROI目标超过2次
	InvalidReasons map[string]string `json:"invalid_reasons,omitempty"`
	// EndReasons 成本保障结束原因
	// UN_OBERCOST 超成本比例没有达到1.2倍
	// ROI_REAL_EXPECTED实际roi大于目标roi的80%
	// CONVERT_UNDER_THRESHOLD转化数没有达到门槛
	// CURRENCY_PRECISION赔付金额小于0.01元
	EndReasons map[string]string `json:"end_reasons,omitempty"`
	// CompensationAmount 赔付金额
	CompensationAmount float64 `json:"compensation_amount,omitempty"`
	// URL 赔付规则链接
	URL string `json:"url,omitempty"`
}
