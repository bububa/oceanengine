package creativecomponent

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
)

// ComponentInfo 组件信息
type ComponentInfo struct {
	// ComponentType 组件类型
	ComponentType enum.ComponentType `json:"component_type,omitempty"`
	// ComponentName 组件名称。长度小于等于20。一个中文长度为2
	ComponentName string `json:"component_name,omitempty"`
	// ComponentData 组件详细信息。不同的component_type对应的值不同，具体的结构见创建或更新接口定义
	ComponentData ComponentData `json:"component_data,omitempty"`
	// CreateTime 创建时间。格式"2020-12-25 15:12:08"
	CreateTime string `json:"create_time,omitempty"`
	// Status 组件审核状态。
	Status enum.ComponentStatus `json:"status,omitempty"`
}

// ComponentData 组件详细信息 interface
type ComponentData interface {
	Type() enum.ComponentType
}

// Button 组件按钮
type Button struct {
	// ButtonType 按钮类型。 0跳转落地页 2显示图片
	ButtonType int `json:"button_type"`
	// ButtonText 按钮文案。长度小于等于24一个中文长度为2
	ButtonText string `json:"button_text,omitempty"`
	// ExternalURL 落地页链接。 可使用橙子建站的链接或第三方的链接 button_type为跳转落地页时候必传
	ExternalURL string `json:"external_url,omitempty"`
	// ResultText 结果文案。 button_type为显示图片时候必传; 长度小于等于24。一个中文长度为2
	ResultText string `json:"result_text,omitempty"`
	// ShowTips 按钮点击后显示文案。button_type为显示图片时候必传。长度小于等于16一个中文长度为2
	ShowTips string `json:"show_tips,omitempty"`
	// ImageID 图片id。尺寸<=472*164px 大小<1M button_type为显示图片时候必传
	ImageID string `json:"image_id,omitempty"`
}

// ChoiceMagne 选择磁贴的component_data
type ChoiceMagnet struct {
	// Title 主标题。长度小于等于24。一个中文长度为2
	Title string `json:"title,omitempty"`
	// SubTitle 副标题。内容长度要求：44 ≤ x ≤60。一个中文长度为2
	SubTitle string `json:"sub_title,omitempty"`
	// ButtonList 按钮列表。长度必须为2
	ButtonList [2]Button `json:"button_list,omitempty"`
}

func (m ChoiceMagnet) Type() enum.ComponentType {
	return enum.ComponentType_CHOICE_MAGNET
}

// VoteMagnet 投票磁贴的component_data
type VoteMagnet struct {
	// Title 主标题。长度小于等于24。一个中文长度为2
	Title string `json:"title,omitempty"`
	// SubTitle 副标题。内容长度要求 44 ≤ x ≤60。一个中文长度为2
	SubTitle string `json:"sub_title,omitempty"`
	// VoteCardWebURL 落地页连接。可使用橙子建站的链接或第三方的链接
	VoteCardWebURL string `json:"vote_card_web_url,omitempty"`
	// ClickText 按钮点击后显示文案。长度小于等于16。一个中文长度为2
	ClickText string `json:"click_text,omitempty"`
	// ButtonList 按钮列表。长度必须等于2
	ButtonList [2]Button `json:"button_list,omitempty"`
}

func (m VoteMagnet) Type() enum.ComponentType {
	return enum.ComponentType_VOTE_MAGNET
}

// ImageMagnet 图片磁贴的component_data
type ImageMagnet struct {
	// StartTime 投放开始时间。格式: "2020-12-01"
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间。格式: "2020-12-02" 投放结束时间不能早于开始时间
	EndTime string `json:"end_time,omitempty"`
	// ImageID 图片id。尺寸<=540*276px ，大小<2M
	ImageID string `json:"image_id,omitempty"`
	// ShowTime 出现时间类型。1 系统默认，2自定义。不传代表系统默认
	ShowTime int `json:"show_time,omitempty"`
	// Duration 出现时间，单位s。当show_time为自定义时必传，内容长度要求：10 ≤ x ≤59。当show_time为系统默认时，可不传，默认为10
	Duration int64 `json:"duration,omitempty"`
}

func (m ImageMagnet) Type() enum.ComponentType {
	return enum.ComponentType_IMAGE_MAGNET
}

// CommerceCard 电商磁铁/优惠券磁贴信息
type CommerceCard struct {
	// StartTime 投放开始时间。格式: "2020-12-01"
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间。格式: "2020-12-02" 投放结束时间不能早于开始时间
	EndTime string `json:"end_time,omitempty"`
	// MagnetTitle 磁贴标题。长度最大为10。一个中文长度为1
	MagnetTitle string `json:"magnet_title,omitempty"`
	// Title 文案内容。长度最大为15。一个中文长度为1
	Title string `json:"title,omitempty"`
	// ButtonText 按钮文案。只支持立即下单和立即购买
	ButtonText string `json:"button_text,omitempty"`
	// ImageID 图片id。尺寸<=174:174，大小<1M
	ImageID string `json:"image_id,omitempty"`
	// Name 活动名称。长度最大为15。一个中文长度为2
	Name string `json:"name,omitempty"`
	// CouponAmount 优惠金额（整数的字符串格式）
	CouponAmount string `json:"coupon_amount,omitempty"`
	// CouponCondition 优惠条件。长度最大为16。一个中文长度为2
	CouponCondition string `json:"coupon_condition,omitempty"`
	// EffectiveStartDate 优惠开始时间。格式: "2020-12-01"
	EffectiveStartDate string `json:"effective_start_date,omitempty"`
	// EffectiveEndDate 优惠结束时间。格式: "2020-12-01"
	EffectiveEndDate string `json:"effective_end_date,omitempty"`
	// ShowTime 出现时间类型。1 系统默认，2自定义。不传代表系统默认
	ShowTime int `json:"show_time,omitempty"`
	// Duration 出现时间，单位s。当show_time为自定义时必传，内容长度要求：10 ≤ x ≤59。当show_time为系统默认时，可不传，默认为10
	Duration int64 `json:"duration,omitempty"`
}

// CommerceMagnet 电商磁贴的component_data
type CommerceMagnet struct {
	// CommerceCards 电商磁铁信息，长度只能为1
	CommerceCards [1]CommerceCard `json:"commerce_cards,omitempty"`
}

// CouponMagnet 优惠券磁贴的component_data
type CouponMagnet struct {
	// CommerceCards 电商磁铁信息，长度只能为1
	CommerceCards [1]CommerceCard `json:"commerce_cards,omitempty"`
}

func (m CouponMagnet) Type() enum.ComponentType {
	return enum.ComponentType_COUPON_MAGNET
}

// PromotionCard 推广卡片的component_data
type PromotionCard struct {
	// ImageID 图片id。建议尺寸：108*108，大小为1M
	ImageID string `json:"image_id,omitempty"`
	// Title 卡片标题。长度最大为20。一个中文长度为1
	Title string `json:"title,omitempty"`
	// ButtonText 行动号召。内容长度要求：2 ≤ x ≤6。一个中文长度为1
	ButtonText string `json:"button_text,omitempty"`
	// ProductSellingPoints 推广卖点。最多选择10个推广卖点 ，内容长度要求：6 ≤ x ≤9 。一个中文长度为1
	ProductSellingPoints []string `json:"product_selling_points,omitempty"`
	// EnablePersonalAction 是否开启智能优选，1-开启、0-不开启
	EnablePersonalAction int `json:"enable_personal_action,omitempty"`
}

func (p PromotionCard) Type() enum.ComponentType {
	return enum.ComponentType_PROMOTION_CARD
}

// GamePack 游戏礼包码的component_data
type GamePack struct {
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
	// PackageName 应用包名
	PackageName string `json:"package_name,omitempty"`
	// AppDownloadURL 应用下载链接
	AppDownloadURL string `json:"app_download_url,omitempty"`
	// AppLogo 应用logo
	AppLogo string `json:"app_logo,omitempty"`
	// AppDescription 应用描述。长度最大为30。一个中文长度为2
	AppDescription string `json:"app_description,omitempty"`
	// BatchID 礼包券活动id
	BatchID uint64 `json:"batch_id,omitempty"`
	// AppThumbnails 应用图片ID，可通过【获取图片素材】获取。最多2张图片。建议尺寸：1280 * 720 ≤ 尺寸 ≤ 2560 * 1440
	AppThumbnails []string `json:"app_thumbnails,omitempty"`
}

func (g GamePack) Type() enum.ComponentType {
	return enum.ComponentType_GAME_PACK
}
