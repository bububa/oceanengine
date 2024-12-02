package advertiser

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// FundGetRequest 查询账号余额 API Request
type FundGetRequest struct {
	// AdvertiserID 广告主ID或代理商ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// GrantTypeSplit 是否拆分赠款类型，允许值：
	// 开启 ON ，关闭 OFF（默认）
	GrantTypeSplit enum.OnOff `json:"grant_type_split,omitempty"`
}

// Encode implement GetRequest interface
func (r FundGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.GrantTypeSplit != "" {
		values.Set("grant_type_split", string(r.GrantTypeSplit))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// FundGetResponse 查询账号余额 API Response
type FundGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *FundGetResult `json:"data,omitempty"`
}

// FundGetResponseData 账号余额
type FundGetResult struct {
	// AdvertiserID 广告主ID或代理商ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 账户名
	Name string `json:"name,omitempty"`
	// Email 联系邮箱
	Email string `json:"email,omitempty"`
	FundInfo
}

type FundInfo struct {
	// Balance 账户总余额(单位元)
	Balance model.Float64 `json:"balance,omitempty"`
	// ValidBalance 账户可用总余额(单位元)
	ValidBalance model.Float64 `json:"valid_balance,omitempty"`
	// Cash 现金余额(单位元)
	Cash model.Float64 `json:"cash,omitempty"`
	// ValidCash 现金可用余额(单位元)
	ValidCash model.Float64 `json:"valid_cash,omitempty"`
	// Grant 赠款余额(单位元)
	Grant model.Float64 `json:"grant,omitempty"`
	// DefaultGrant 默认广告位可用赠款余额(单位元)
	DefaultGrant model.Float64 `json:"default_grant,omitempty"`
	// CommonGrant 通用广告位可用赠款余额(单位元)
	CommonGrant model.Float64 `json:"common_grant,omitempty"`
	// SearchGrant 搜索广告位可用赠款余额(单位元)
	SearchGrant model.Float64 `json:"search_grant,omitempty"`
	// UnionGrant 穿山甲广告位可用赠款余额(单位元)
	UnionGrant model.Float64 `json:"union_grant,omitempty"`
	// ValidGrant 赠款可用余额(单位元)
	ValidGrant model.Float64 `json:"valid_grant,omitempty"`
	// ReturnGoodsAbs 返货余额(单位元)，仅支持部分广告主
	ReturnGoodsAbs model.Float64 `json:"return_goods_abs,omitempty"`
	// ValidReturnGoodsAbs 返货可用余额(单位元)，仅支持部分广告主
	ValidReturnGoodsAbs model.Float64 `json:"valid_return_goods_abs,omitempty"`
	// ReturnGoodsCost 返货支出(单位元)，仅支持部分广告主
	ReturnGoodsCost model.Float64 `json:"return_goods_cost,omitempty"`
	// ReturnGoodsGrant 返货赠款
	ReturnGoodsGrant model.Float64 `json:"return_goods_grant,omitempty"`
	// CompensationGrant 赔付赠款
	CompensationGrant model.Float64 `json:"compensation_grant,omitempty"`
	// ReturnGoodsValidGrant 返货可用赠款
	ReturnGoodsValidGrant model.Float64 `json:"return_goods_valid_grant,omitempty"`
	// CompensationValidGrant 赔付可用赠款
	CompensationValidGrant model.Float64 `json:"compensation_valid_grant,omitempty"`
	// WalletID 钱包id（广告主账户绑定的共享子钱包id）
	WalletID model.Uint64 `json:"wallet_id,omitempty"`
	// WalletName 钱包名称（广告主账户绑定的共享子钱包名称）
	WalletName string `json:"wallet_name,omitempty"`
	// WalletTotalBalanceValid 账户绑定的子钱包的可用共享余额（单位元）
	WalletTotalBalanceValid model.Float64 `json:"wallet_total_balance_valid,omitempty"`
}
