package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建广告 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// Name 广告名称，长度是1-50个字（两个英文字符占1个字）。名称不可重复
	Name string `json:"name,omitempty"`
	// Operation 广告状态， 允许值: ENABLE开启(默认值）、DISABLE关闭
	Operation enum.OptStatus `json:"operation,omitempty"`
	// PromotionMaterials 广告素材组合
	PromotionMaterials []PromotionMaterial `json:"promotion_materials,omitempty"`
	// Source 广告来源，字数限制：[1-10]
	// 当landing_type = LINK、MICRO_GAME时必填
	Source string `json:"source,omitempty"`
	// Budget 预算(出价方式为OCPM时，不少于300元；24小时内修改预算操作，不能超过20次，24小时是指自然天的24小时；单次修改预算幅度不能低于100元（增加或者减少）；修改后预算金额，不能低于当前已消费金额的105%，以整百单位向上取整；取值范围: ≥ 0
	Budget float64 `json:"budget,omitempty"`
	// CpaBid 目标转化出价/预期成本， 取值范围：0.1-10000元
	// 注意：当 bid_type = NO_BID时，不填写该字段，否则会报错
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// DeepCpaBid 深度优化出价，deep_bid_type = DEEP_BID_MIN 时必填。取值范围：0.1~10000元
	// 注意：当 bid_type = NO_BID时，不填写该字段，否则会报错
	DeepCpaBid float64 `json:"deep_cpabid,omitempty"`
	// RoiGoal 深度转化ROI系数, 范围(0,5]，精度：保留小数点后四位，当 deep_bid_type = ROI_COEFFICIENT 时必填
	// 注意：当 bid_type=NO_BID时，不填写该字段，否则会报错
	RoiGoal float64 `json:"roi_goal,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建广告 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// PromotionID 广告ID
		PromotionID uint64 `json:"promotion_id,omitempty"`
	} `json:"data,omitempty"`
}
