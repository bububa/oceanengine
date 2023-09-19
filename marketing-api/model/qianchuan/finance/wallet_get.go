package finance

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// WalletGetRequest 获取账户钱包信息 API Request
type WalletGetRequest struct {
	// AdvertiserID 千川广告主/代理商账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r WalletGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// WalletGetResponse 获取账户钱包信息 API Response
type WalletGetResponse struct {
	model.BaseResponse
	// Data 返回数据，单位为千分之一分，即174938000.00=1749.38000元
	Data *Wallet `json:"data,omitempty"`
}
