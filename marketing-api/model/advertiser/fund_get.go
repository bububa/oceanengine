package advertiser

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// FundGetRequest 查询账号余额 API Request
type FundGetRequest struct {
	// AdvertiserID 广告主ID或代理商ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r FundGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}

// FundGetResponse 查询账号余额 API Response
type FundGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *FundGetResponseData `json:"data,omitempty"`
}

// FundGetResponseData 账号余额
type FundGetResponseData struct {
	// AdvertiserID 广告主ID或代理商ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 账户名
	Name string `json:"name,omitempty"`
	// Email 联系邮箱
	Email string `json:"email,omitempty"`
	// Balance 账户总余额(单位元)
	Balance float64 `json:"balance,omitempty"`
	// ValidBalance 账户可用总余额(单位元)
	ValidBalance float64 `json:"valid_balance,omitempty"`
	// Cash 现金余额(单位元)
	Cash float64 `json:"cash,omitempty"`
	// ValidCash 现金可用余额(单位元)
	ValidCash float64 `json:"valid_cash,omitempty"`
	// Grant 赠款余额(单位元)
	Grant float64 `json:"grant,omitempty"`
	// ValidGrant 赠款可用余额(单位元)
	ValidGrant float64 `json:"valid_grant,omitempty"`
	// ReturnGoodsAbs 返货余额(单位元)，仅支持部分广告主
	ReturnGoodsAbs float64 `json:"return_goods_abs,omitempty"`
	// ValidReturnGoodsAbs 返货可用余额(单位元)，仅支持部分广告主
	ValidReturnGoodsAbs float64 `json:"valid_return_goods_abs,omitempty"`
	// ReturnGoodsCost 返货支出(单位元)，仅支持部分广告主
	ReturnGoodsCost float64 `json:"return_goods_cost,omitempty"`
}
