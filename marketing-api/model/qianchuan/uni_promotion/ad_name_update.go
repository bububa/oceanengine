package unipromotion

import "github.com/bububa/oceanengine/marketing-api/util"

// AdNameUpdateRequest 更新商品全域推广计划名称 API Request
type AdNameUpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 商品全域计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Name 商品计划名称，长度为1-100个字符，其中1个汉字算2位字符
	Name string `json:"name,omitempty"`
}

// Encode implements PostRequest inteface
func (r AdNameUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
