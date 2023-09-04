package shop

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取店铺账户信息 API Request
type GetRequest struct {
	// ShopIDs 店铺id
	ShopIDs []uint64 `json:"shop_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	ids, _ := json.Marshal(r.ShopIDs)
	values.Set("shop_ids", string(ids))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 获取店铺账户信息 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List shop账号列表
		List []Shop `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
