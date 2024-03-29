package finance

// Wallet 账户钱包信息
type Wallet struct {
	// TotalBalanceAbs 账户总余额
	TotalBalanceAbs float64 `json:"total_balance_abs,omitempty"`
	// GrantBalance 赠款余额
	GrantBalance float64 `json:"grant_balance,omitempty"`
	// UnionValidGrantBalance 赠款余额-穿山甲-可用
	UnionValidGrantBalance float64 `json:"union_valid_grant_balance,omitempty"`
	// SearchValidGrantBalance 赠款余额-巨量搜索广告-可用
	SearchValidGrantBalance float64 `json:"search_valid_grant_balance,omitempty"`
	// CommonValidGrantBalance 赠款余额-巨量信息流广告-可用
	CommonValidGrantBalance float64 `json:"common_valid_grant_balance,omitempty"`
	// DefaultValidGrantBalance 赠款余额-通用-可用
	DefaultValidGrantBalance float64 `json:"default_valid_grant_balance,omitempty"`
	// GeneralTotalBalance 通用余额
	GeneralTotalBalance float64 `json:"general_total_balance,omitempty"`
	// GeneralBalanceValid 通用余额-可用余额
	GeneralBalanceValid float64 `json:"general_balance_valid,omitempty"`
	// GeneralBalanceValidNoGrant 通用余额-可用余额-非赠款
	GeneralBalanceValidNoGrant float64 `json:"general_balance_valid_no_grant,omitempty"`
	// GeneralBalanceValidGrantUnion 通用余额-可用余额-赠款-穿山甲
	GeneralBalanceValidGrantUnion float64 `json:"general_balance_valid_grant_union,omitempty"`
	// GeneralBalanceValidGrantSearch 通用余额-可用余额-赠款-巨量搜索广告
	GeneralBalanceValidGrantSearch float64 `json:"general_balance_valid_grant_search,omitempty"`
	// GeneralBalanceValidGrantCommon 通用余额-可用余额-赠款-巨量信息流广告
	GeneralBalanceValidGrantCommon float64 `json:"general_balance_valid_grant_common,omitempty"`
	// GeneralBalanceValidGrantDefault 通用余额-可用余额-赠款-通用
	GeneralBalanceValidGrantDefault float64 `json:"general_balance_valid_grant_default,omitempty"`
	// GeneralBalanceInvalid 通用余额-不可用余额
	GeneralBalanceInvalid float64 `json:"general_balance_invalid,omitempty"`
	// GeneralBalanceInvalidOrder 通用余额-不可用余额-随心推已下单
	GeneralBalanceInvalidOrder float64 `json:"general_balance_invalid_order,omitempty"`
	// GeneralBalanceInvalidFrozen 通用余额-不可用余额-冻结
	GeneralBalanceInvalidFrozen float64 `json:"general_balance_invalid_frozen,omitempty"`
	// BrandBalance 品牌余额
	BrandBalance float64 `json:"brand_balance,omitempty"`
	// BrandBalanceValid 品牌余额-可用
	BrandBalanceValid float64 `json:"brand_balance_valid,omitempty"`
	// BrandBalanceValidNoGrant 品牌余额-可用-非赠款
	BrandBalanceValidNoGrant float64 `json:"brand_balance_valid_no_grant,omitempty"`
	// BrandBalanceValidGrant 品牌余额-可用余额-赠款
	BrandBalanceValidGrant float64 `json:"brand_balance_valid_grant,omitempty"`
	// BrandBalanceInvalid 品牌余额-不可用余额
	BrandBalanceInvalid float64 `json:"brand_balance_invalid,omitempty"`
	// BrandBalanceInvalidFrozen 品牌余额-不可用余额-冻结
	BrandBalanceInvalidFrozen float64 `json:"brand_balance_invalid_frozen,omitempty"`
	// DeductionCouponBalance 消返红包余额
	DeductionCouponBalance float64 `json:"deduction_coupon_balance,omitempty"`
	// DeductionCouponBalanceAll 消返红包余额（通用）
	DeductionCouponBalanceAll float64 `json:"deduction_coupon_balance_all,omitempty"`
	// DeductionCouponBalanceOther 消返红包余额（代投）
	DeductionCouponBalanceOther float64 `json:"deduction_coupon_balance_other,omitempty"`
	// DeductionCouponBalanceSelf 消返红包余额（自投）
	DeductionCouponBalanceSelf float64 `json:"deduction_coupon_balance_self,omitempty"`
	// GrantExpiring 15天内赠款到期金额
	GrantExpiring float64 `json:"grant_expiring,omitempty"`
	// ShareBalance 共享赠款余额
	ShareBalance float64 `json:"share_balance,omitempty"`
	// ShareBalanceValidGrantUnion 共享赠款余额-可用余额-赠款-穿山甲
	ShareBalanceValidGrantUnion float64 `json:"share_balance_valid_grant_union,omitempty"`
	// ShareBalanceValidGrantSearch 共享赠款余额-可用余额-赠款-巨量搜索广告
	ShareBalanceValidGrantSearch float64 `json:"share_balance_valid_grant_search,omitempty"`
	// ShareBalanceValidGrantCommon 共享赠款余额-可用余额-赠款-巨量信息流广告
	ShareBalanceValidGrantCommon float64 `json:"share_balance_valid_grant_common,omitempty"`
	// ShareBalanceValidGrantDefault 共享赠款余额-可用余额-赠款-通用
	ShareBalanceValidGrantDefault float64 `json:"share_balance_valid_grant_default,omitempty"`
	// ShareBalanceValid 共享赠款余额-可用余额
	ShareBalanceValid float64 `json:"share_balance_valid,omitempty"`
	// ShareBalanceExpiring 共享赠款余额-30天内到期余额
	ShareBalanceExpiring float64 `json:"share_balance_expiring,omitempty"`
	// ShareExpiringDetailList 共享赠款余额到期详情
	ShareExpiringDetailList *ShareExpiringDetail `json:"share_expiring_detail_list,omitempty"`
}

// ShareExpiringDetail 共享赠款余额到期详情
type ShareExpiringDetail struct {
	// Category 类别，允许值：
	// CONFIRM 站内信息流及其他
	// DEFAULT 通用
	// SEARCH 站内搜索
	// UNION 网盟穿山甲
	Category string `json:"category,omitempty"`
	// Amount 金额
	Amount float64 `json:"amount,omitempty"`
	// ExpireTime 到期时间
	ExpireTime int64 `json:"expire_time,omitempty"`
}
