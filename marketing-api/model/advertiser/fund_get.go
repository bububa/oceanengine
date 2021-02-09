package advertiser

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type FundGetRequest struct {
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

func (r FundGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}

type FundGetResponse struct {
	model.BaseResponse
	Data *FundGetResponseData `json:"data,omitempty"`
}

type FundGetResponseData struct {
	AdvertiserID        uint64  `json:"advertiser_id,omitempty"`          // 广告主ID或代理商ID
	Name                string  `json:"name,omitempty"`                   // 账户名
	Email               string  `json:"email,omitempty"`                  // 联系邮箱
	Balance             float64 `json:"balance,omitempty"`                // 账户总余额(单位元)
	ValidBalance        float64 `json:"valid_balance,omitempty"`          // 账户可用总余额(单位元)
	Cash                float64 `json:"cash,omitempty"`                   // 现金余额(单位元)
	ValidCash           float64 `json:"valid_cash,omitempty"`             // 现金可用余额(单位元)
	Grant               float64 `json:"grant,omitempty"`                  // 赠款余额(单位元)
	ValidGrant          float64 `json:"valid_grant,omitempty"`            // 赠款可用余额(单位元)
	ReturnGoodsAbs      float64 `json:"return_goods_abs,omitempty"`       // 返货余额(单位元)，仅支持部分广告主
	ValidReturnGoodsAbs float64 `json:"valid_return_goods_abs,omitempty"` // 返货可用余额(单位元)，仅支持部分广告主
	ReturnGoodsCost     float64 `json:"return_goods_cost,omitempty"`      // 返货支出(单位元)，仅支持部分广告主
}
