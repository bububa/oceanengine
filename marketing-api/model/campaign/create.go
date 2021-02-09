package campaign

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

type CreateRequest struct {
	AdvertiserID       uint64           `json:"advertiser_id,omitempty"`        // 广告主ID
	CampaignName       string           `json:"campaign_name,omitempty"`        // 广告组名称，长度为1-100个字符，其中1个中文字符算2位
	Operation          string           `json:"operation,omitempty"`            // 广告组状态, 允许值: "enable","disable"默认值：enable开启状态
	BudgetMode         enum.BudgetMode  `json:"budget_mode,omitempty"`          // 广告组预算类型
	Budget             float64          `json:"budget,omitempty"`               // 广告组预算，取值范围: ≥ 0; 当budget_mode为"BUDGET_MODE_DAY"时,必填,且日预算不少于300	元
	LandingType        enum.LandingType `json:"landing_type,omitempty"`         // 广告组推广目的
	UniqueFk           string           `json:"unique_fk,omitempty"`            // 第三方唯一键，传该值时保证接口重试的幂等性，请注意，带有相同unique_fk的请求服务端会视为同一个广告处理。仅在创建接口传入且无法修改，如果创建时传入了已存在的唯一键值，那么会返回该唯一键值所对应的广告组ID。该值可用于内部系统会生成的唯一ID与头条ID做关联的场景，避免超时重试实际上一次创建请求又成功导致的重复创建问题，通过unique_fk可与内部系统ID实现关联并避免重复创建，可结合实际场景选择使用，广告组中的unique_fk要求不重复，与计划中的unique_fk无相关。
	DeliveryRelatedNum enum.CampaignDPA `json:"delivery_related_num,omitempty"` // 广告组商品类型
}

func (r CreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type CreateResponse struct {
	model.BaseResponse
	Data *CreateResponseData `json:"data,omitempty"`
}

type CreateResponseData struct {
	CampaignID uint64 `json:"campaign_id,omitempty"` // 广告组id
}
