package tools

import "github.com/bububa/oceanengine/marketing-api/util"

// ShopAuthRequest 店铺新客定向授权 API Request
type ShopAuthRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ShopID 要授权的店铺id
	ShopID uint64 `json:"shop_id,omitempty"`
	// ShopAuthTimeType 授权时间类型，默认不限，允许值
	// NONE 不限
	// CUSTOM 自定义
	ShopAuthTimeType string `json:"shop_auth_time_type,omitempty"`
	// EndTime 授权结束时间，当shop_auth_time_type=CUSTOM时必填
	// 格式为yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty"`
}

// Encode implement PostRequest interface
func (r ShopAuthRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
