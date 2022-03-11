package advertiser

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// TransferableFundGetRequest 查询账户可转余额 API Request
type TransferableFundGetRequest struct {
	// AdvertiserID 广告主ID或代理商ID或组织ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r TransferableFundGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}

// TransferableFundGetResponse 查询账户可转余额 API Response
type TransferableFundGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *TransferableFund `json:"data,omitempty"`
}

// TransferableFund 账户可转余额
type TransferableFund struct {
	// GrantValid 可用赠款余额(单位元)
	GrantValid float64 `json:"grant_valid,omitempty"`
	// UnversalPrepayValid 通用预付可用余额(单位元)
	UniversalPrepayValid float64 `json:"universal_prepay_valid,omitempty"`
	// BrandPrepayValid 品牌预付可用余额(单位元)
	BrandPrepayValid float64 `json:"brand_prepay_valid,omitempty"`
	// BidPrepayValid 竞价预付可用余额(单位元)
	BidPrepayValid float64 `json:"bid_prepay_valid,omitempty"`
	// UniversalCreditValid 通用授信可用余额(单位元)
	UniversalCreditValid float64 `json:"universal_credit_valid,omitempty"`
	// BrandCreditValid 品牌授信可用余额(单位元)
	BrandCreditValid float64 `json:"brand_credit_valid,omitempty"`
	// BidCreditValid 竞价授信可用余额(单位元)
	BidCreditValid float64 `json:"bid_credit_valid,omitempty"`
}
