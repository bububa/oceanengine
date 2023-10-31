package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CompensateStatusGetRequest 获取计划成本保障状态 API Request
type CompensateStatusGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 计划id列表，每次最多传入50个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r CompensateStatusGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_ids", string(util.JSONMarshal(r.AdIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CompensateStatusGetResponse 获取计划成本保障状态 API Response
type CompensateStatusGetResponse struct {
	model.BaseResponse
	Data struct {
		List []CompensateStatus `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// CompensateStatus 计划成本保障状态
type CompensateStatus struct {
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Status 当前请求是否成功, 枚举值:
	// SUCCESS: 查询成功
	// FAILED: 查询失败，请重试
	Status qianchuan.CompensateRequestStatus `json:"status,omitempty"`
	// CompensateStatus 计划成本保障状态，允许值
	// IN_EFFECT: 成本保障生效中
	// INVALID: 成本保障已失效
	// CONFIRMING: 成本保障确认中
	// PAID: 成本保障已到账
	// ENDED: 成本保障已结束
	// DEFAULT：无成本保障状态
	CompensateStatus qianchuan.CompensateStatus `json:"compensate_status,omitempty"`
	// Reason 成本保障失效/结束原因
	Reason string `json:"reason,omitempty"`
	// URL 赔付规则链接
	URL string `json:"url,omitempty"`
}
