package campaign

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

type UpdateRequest struct {
	AdvertiserID uint64          `json:"advertiser_id,omitempty"` // 广告主ID
	CampaignID   uint64          `json:"campaign_id,omitempty"`   // 广告组ID，广告组ID需要属于广告主ID，否则会报错！
	ModifyTime   string          `json:"modify_time,omitempty"`   // 时间戳（从campaign/get/接口得到,用于判断是否基于最新信息修改，正确获取与填写，以免报错！)
	CampaignName string          `json:"campaign_name,omitempty"` // 广告组名称，长度为1-100个字符，其中1个中文字符算2位
	BudgetMode   enum.BudgetMode `json:"budget_mode,omitempty"`   // 广告组预算类型
	Budget       float64         `json:"budget,omitempty"`        // 广告组预算，取值范围: ≥ 0; 当budget_mode为"BUDGET_MODE_DAY"时,必填,且日预算不少于300	元
}

func (r UpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type UpdateResponse struct {
	model.BaseResponse
	Data *UpdateResponseData `json:"data,omitempty"`
}

type UpdateResponseData struct {
	CampaignID uint64 `json:"campaign_id,omitempty"` // 广告组id
}
