package coupon

// CodeType 券码类型
type CodeType string

const (
	// CodeType_API 合作方api实时传入的券码
	CodeType_API CodeType = "API"
	// CodeType_COMMON 通用码，固定值，不支持核销
	CodeType_COMMON CodeType = "COMMON"
	// CodeType_MERCHANT 商家上传券码，需调用上传券码接口上传券码文件
	CodeType_MERCHANT CodeType = "MERCHANT"
	// CodeType_MERCHANT 平台产生券码
	CodeType_PLATFORM CodeType = "PLATFORM"
)

// UseType 核销类型
type UseType string

const (
	// UseType_OFFLINE 用户线下核销
	UseType_OFFLINE UseType = "OFFLINE"
	// UseType_LINK 用户使用时直接跳转新链接
	UseType_LINK UseType = "LINK"
)

// IndustryType 行业列表
type IndustryType string

const (
	// IndustryType_ENTERTAINMENT 生活娱乐
	IndustryType_ENTERTAINMENT IndustryType = "ENTERTAINMENT"
	// IndustryType_FINANCIAL 互联网金融
	IndustryType_FINANCIAL IndustryType = "FINANCIAL"
	// IndustryType_FOOD 美食
	IndustryType_FOOD IndustryType = "FOOD"
	// IndustryType_GAME 游戏
	IndustryType_GAME IndustryType = "GAME"
	// IndustryType_TICKET 休闲票务
	IndustryType_TICKET IndustryType = "TICKET"
	// IndustryType_OTHER 其他
	IndustryType_OTHER IndustryType = "OTHER"
)

// ResourceType 优惠券类型
type ResourceType string

const (
	// ResourceType_COMMON 通用券
	ResourceType_COMMON ResourceType = "COMMON"
	// ResourceType_DISCOUNT 折扣券
	ResourceType_DISCOUNT ResourceType = "DISCOUNT"
	// ResourceType_FULL 满减券
	ResourceType_FULL ResourceType = "FULL"
	// ResourceType_GAME 游戏礼包
	ResourceType_GAME ResourceType = "GAME"
	// ResourceType_PHYSICAL 实物券
	ResourceType_PHYSICAL ResourceType = "PHYSICAL"
)

// Resource 优惠券信息
type Resource struct {
	// ResourceID 优惠券ID
	ResourceID uint64 `json:"resource_id,omitempty"`
	// Title 优惠券标题，不超过15字
	Title string `json:"title,omitempty"`
	// MerchantName 商户名称，不超过15字
	MerchantName string `json:"merchant_name,omitempty"`
	// ValidDays 有效时间段，从领取日开始计算，最大值不超过365
	// 设置此值后，valid_start和valid_end会被忽略，0表示不设置
	ValidDays int `json:"valid_days,omitempty"`
	// ValidStart 有效期开始时间 格式：%Y-%m-%d
	ValidStart string `json:"valid_start,omitempty"`
	// ValidEnd 有效期结束时间 格式：%Y-%m-%d
	ValidEnd string `json:"valid_end,omitempty"`
	// CodeType 券码类型，可选值:
	// API：合作方api实时传入的券码
	// COMMON：通用码，固定值，不支持核销
	// MERCHANT：商家上传券码，需调用上传券码接口上传券码文件
	// PLATFORM：平台产生券码
	CodeType CodeType `json:"code_type,omitempty"`
	// CodeReferURL 商家上传券码类型时，有效券码文件的url
	CodeReferURL string `json:"code_refer_url,omitempty"`
	// Code 通用券码类型时设置的券码，4-30位，只支持字母、数字
	Code string `json:"code,omitempty"`
	// Condition 使用须知，长度不超过50字
	Condition string `json:"condition,omitempty"`
	// Notification 领取须知，长度不超过500
	Notification string `json:"notification,omitempty"`
	// HeadImageURL 卡券头图地址，即优惠券背景图
	HeadImageURL string `json:"head_image_url,omitempty"`
	// LogoImageURL 商家logo地址
	LogoImageURL string `json:"logo_image_url,omitempty"`
	// ServiceNum 客服电话
	ServiceNum string `json:"service_num,omitempty"`
	// Stock 优惠券的库存，0表示不限制
	Stock int `json:"stock,omitempty"`
	// UseType 核销类型
	// OFFLINE：用户线下核销
	// LINK：用户使用时直接跳转新链接
	UseType UseType `json:"use_type,omitempty"`
	// Link use_type为LINK时必填，跳转链接
	Link string `json:"link,omitempty"`
	// IndustryType 行业列表，可选值:
	// ENTERTAINMENT：生活娱乐
	// FINANCIAL：互联网金融
	// FOOD：美食
	// GAME：游戏
	// TICKET：休闲票务
	// OTHER：其他
	IndustryType IndustryType `json:"industry_type,omitempty"`
	// StoreIDs 可用门店ID列表
	// 通过 【查询门店接口】获取
	// 建议不超过500
	StoreIDs []uint64 `json:"store_ids,omitempty"`
	// MinAmount 最低消费金额，单位为分，如：500 表示500分，即5元
	MinAmount int64 `json:"min_amount,omitempty"`
	// ResourceType 优惠券类型。详见【附录-枚举值】，可选值:
	// COMMON：通用券
	// DISCOUNT：折扣券
	// FULL：满减券
	// GAME：游戏礼包
	// PHYSICAL：实物券
	ResourceType ResourceType `json:"resource_type,omitempty"`
	// GiftList 游戏礼包列表
	GiftList []Gift `json:"gift_list,omitempty"`
	// IosLink ios端下载链接
	IosLink string `json:"ios_link,omitempty"`
	// AndroidLink 安卓端下载链接
	AndroidLink string `json:"android_link,omitempty"`
	// Discount 折扣券的折扣，值为1到99，如：1代表0.1折，90代表9折。
	Discount int `json:"discount,omitempty"`
	// MaxAmount 优惠券类型为折扣券时的最大减免金额，单位为分，如：500 表示500分，即5元
	MaxAmount int64 `json:"max_amount,omitempty"`
	// ReliefAmount 满减券的减免金额，单位为分，如：500 表示500分，即5元
	ReliefAmount int64 `json:"relief_amount,omitempty"`
	// AuditMsg 审核信息，成功时通常为空，失败时描述原因
	AuditMsg string `json:"audit_msg,omitempty"`
}
