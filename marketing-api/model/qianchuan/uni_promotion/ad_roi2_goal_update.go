package unipromotion

import "github.com/bububa/oceanengine/marketing-api/util"

// AdRoi2GoalUpdateRequest 更新全域推广控成本计划支付ROI目标
type AdRoi2GoalUpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UpdateRoi2Infos 需要更新的roi信息
	// 注意：一次最多支持10条计划。
	UpdateRoi2Infos []UpdateRoi2Info `json:"update_roi2_infos,omitempty"`
}

// UpdateRoi2Info 需要更新的roi信息
type UpdateRoi2Info struct {
	// AdID 商品全域计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Roi2Goal roi目标
	// 单位，最多支持两位小数
	Roi2Goal float64 `json:"roi2_goal,omitempty"`
}

// Encode implements PostRequest inteface
func (r AdRoi2GoalUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
