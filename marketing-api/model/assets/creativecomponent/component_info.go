package creativecomponent

import (
	"encoding/json"
	"fmt"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// ComponentInfo 组件信息
type ComponentInfo struct {
	// ComponentID 组件ID
	ComponentID model.Uint64 `json:"component_id,omitempty"`
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

type tmpComponentInfo struct {
	// ComponentID 组件ID
	ComponentID model.Uint64 `json:"component_id,omitempty"`
	// ComponentType 组件类型
	ComponentType enum.ComponentType `json:"component_type,omitempty"`
	// ComponentName 组件名称。长度小于等于20。一个中文长度为2
	ComponentName string `json:"component_name,omitempty"`
	// ComponentData 组件详细信息。不同的component_type对应的值不同，具体的结构见创建或更新接口定义
	ComponentData json.RawMessage `json:"component_data,omitempty"`
	// CreateTime 创建时间。格式"2020-12-25 15:12:08"
	CreateTime string `json:"create_time,omitempty"`
	// Status 组件审核状态。
	Status enum.ComponentStatus `json:"status,omitempty"`
}

// UnmarshalJSON implement json Unmarshal interface
func (i *ComponentInfo) UnmarshalJSON(b []byte) (err error) {
	var tmp tmpComponentInfo
	if err = json.Unmarshal(b, &tmp); err != nil {
		return
	}
	info := ComponentInfo{
		ComponentID:   tmp.ComponentID,
		ComponentType: tmp.ComponentType,
		ComponentName: tmp.ComponentName,
		CreateTime:    tmp.CreateTime,
		Status:        tmp.Status,
	}
	switch info.ComponentType {
	case enum.ComponentType_CHOICE_MAGNET:
		var data ChoiceMagnet
		if err = json.Unmarshal(tmp.ComponentData, &data); err != nil {
			fmt.Println(err)
			return
		}
		info.ComponentData = data
	case enum.ComponentType_VOTE_MAGNET:
		var data VoteMagnet
		if err = json.Unmarshal(tmp.ComponentData, &data); err != nil {
			fmt.Println(err)
			return
		}
		info.ComponentData = data
	case enum.ComponentType_IMAGE_MAGNET:
		var data ImageMagnet
		if err = json.Unmarshal(tmp.ComponentData, &data); err != nil {
			fmt.Println(err)
			return
		}
		info.ComponentData = data
	case enum.ComponentType_COMMERCE_MAGNET:
		var data CommerceMagnet
		if err = json.Unmarshal(tmp.ComponentData, &data); err != nil {
			fmt.Println(err)
			return
		}
		info.ComponentData = data
	case enum.ComponentType_COUPON_MAGNET:
		var data CouponMagnet
		if err = json.Unmarshal(tmp.ComponentData, &data); err != nil {
			fmt.Println(err)
			return
		}
		info.ComponentData = data
	case enum.ComponentType_PROMOTION_CARD:
		var data PromotionCard
		if err = json.Unmarshal(tmp.ComponentData, &data); err != nil {
			fmt.Println(err)
			return
		}
		info.ComponentData = data
	case enum.ComponentType_GAME_PACK:
		var data GamePack
		if err = json.Unmarshal(tmp.ComponentData, &data); err != nil {
			fmt.Println(err)
			return
		}
		info.ComponentData = data
	case enum.ComponentType_LIGHT_INTER_ACTIVE:
		var data LightInterActive
		if err = json.Unmarshal(tmp.ComponentData, &data); err != nil {
			fmt.Println(err)
			return
		}
		info.ComponentData = data
	case enum.ComponentType_RESERVATION_BUTTON:
		var data ReservationButton
		if err = json.Unmarshal(tmp.ComponentData, &data); err != nil {
			fmt.Println(err)
			return
		}
		info.ComponentData = data
	case enum.ComponentType_UNION_LIGHT_INTERACTIVE:
		var data UnionLightInteractive
		if err = json.Unmarshal(tmp.ComponentData, &data); err != nil {
			fmt.Println(err)
			return
		}
		info.ComponentData = data
	case enum.ComponentType_LUCKY_BOX:
		var data LuckyBox
		if err = json.Unmarshal(tmp.ComponentData, &data); err != nil {
			fmt.Println(err)
			return
		}
		info.ComponentData = data
	default:
		info.ComponentData = UnknownComponent(tmp.ComponentData)
	}
	*i = info
	return
}

var _ json.Unmarshaler = (*ComponentInfo)(nil)

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

// ChoiceMagnet 选择磁贴的component_data
type ChoiceMagnet struct {
	// Title 主标题。长度小于等于24。一个中文长度为2
	Title string `json:"title,omitempty"`
	// SubTitle 副标题。内容长度要求：44 ≤ x ≤60。一个中文长度为2
	SubTitle string `json:"sub_title,omitempty"`
	// ButtonList 按钮列表。长度必须为2
	ButtonList [2]Button `json:"button_list,omitempty"`
}

// Type implement ComponentData interface
func (m ChoiceMagnet) Type() enum.ComponentType {
	return enum.ComponentType_CHOICE_MAGNET
}

var _ ComponentData = (*ChoiceMagnet)(nil)

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

// Type implement ComponentData interface
func (m VoteMagnet) Type() enum.ComponentType {
	return enum.ComponentType_VOTE_MAGNET
}

var _ ComponentData = (*VoteMagnet)(nil)

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

// Type implement ComponentData interface
func (m ImageMagnet) Type() enum.ComponentType {
	return enum.ComponentType_IMAGE_MAGNET
}

var _ ComponentData = (*ImageMagnet)(nil)

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

// Type implement ComponentData interface
func (m CommerceMagnet) Type() enum.ComponentType {
	return enum.ComponentType_COMMERCE_MAGNET
}

var _ ComponentData = (*CommerceMagnet)(nil)

// CouponMagnet 优惠券磁贴的component_data
type CouponMagnet struct {
	// CommerceCards 电商磁铁信息，长度只能为1
	CommerceCards [1]CommerceCard `json:"commerce_cards,omitempty"`
}

// Type implement ComponentData interface
func (m CouponMagnet) Type() enum.ComponentType {
	return enum.ComponentType_COUPON_MAGNET
}

var _ ComponentData = (*CouponMagnet)(nil)

type ImageInfo struct {
	WebUri string `json:"web_uri,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
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
	EnablePersonalAction int         `json:"enable_personal_action,omitempty"`
	ImageInfo            []ImageInfo `json:"image_info,omitempty"`
}

// Type implement ComponentData interface
func (p PromotionCard) Type() enum.ComponentType {
	return enum.ComponentType_PROMOTION_CARD
}

var _ ComponentData = (*PromotionCard)(nil)

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

// Type implement ComponentData interface
func (g GamePack) Type() enum.ComponentType {
	return enum.ComponentType_GAME_PACK
}

var _ ComponentData = (*GamePack)(nil)

// CardImage 卡片图片
type CardImage struct {
	// ImageURL 卡片正面图片/盲盒素材
	ImageURL string `json:"image_url,omitempty"`
	// HiddenImageURL 卡片反面图片
	HiddenImageURL string `json:"hidden_image_url,omitempty"`
}

// CouponInfo 优惠券信息
type CouponInfo struct {
	// TItle 优惠标题
	// 字符数≤6
	Title string `json:"title,omitempty"`
	// CouponImage 优惠图片
	CouponImage string `json:"coupon_image,omitempty"`
	// DiscountTime 优惠时间
	DiscountTime string `json:"discount_time,omitempty"`
	// Amount 优惠金额
	// 字符数≤3
	Amount string `json:"amount,omitempty"`
	// Unit 金额单位
	Unit string `json:"unit,omitempty"`
}

// GiftPackageInfo 礼包码信息
type GiftPackageInfo struct {
	// Title 主标题。长度小于等于16，一个中文占2位。
	// 当style_type=14、15时有提示文案
	// 当style_type=16时，question回答正确有提示文案，否则报错
	Title string `json:"title,omitempty"`
	// Tipds 引导文案（如仅有3.4%用户获得）。长度小于等于30，一个中文占2位。当style_type=14时有提示文案
	// 当style_type=11翻卡领游戏礼包码(不带icon)时该字段必填
	Tips string `json:"tips,omitempty"`
	// Code 礼包兑换码（如VIP888）。长度小于等于16，一个中文占2位。
	Code string `json:"code,omitempty"`
	// GiftPackageImageInfo 礼包图片信息。数量限制为3，少于3或大于3报错。
	// 当style_type=10翻卡领游戏礼包码(带icon)时该字段必填
	// 当style_type=14、16时gift_package_image_id和gift_package_text为三组
	GiftPackageImageInfo []GiftPackageImageInfo `json:"gift_package_image_info,omitempty"`
	// CardPackageImageInfo 抽卡图片信息，数量限制为1。当style_type=15时该字段必填
	CardPackageImageInfo []CardPackageImageInfo `json:"card_package_image_info,omitempty"`
	// CtaText 按钮文案，当style_type=14、15、16时该字段必填
	// 枚举值：PREFER_SYSTEM_RECOMMAND-个性化文案、DOWNLOAD_AND_RECEIVE_NOW-立即下载领取、RECEIVE_NOW-立即领取、CUSTOM_CTA-自定义按钮文案
	CtaText string `json:"cta_text,omitempty"`
	// CtaTextCustom 当cta_text为自定义按钮文案时，该字段必填。长度小于等于12，一个中文占2位。
	CtaTextCustom string `json:"cta_text_custom,omitempty"`
	// CodeList 礼包码列表
	CodeList []GiftPackageCode `json:"code_list,omitempty"`
}

// GiftPackageImageInfo 礼包图片信息。数量限制为3，少于3或大于3报错。
type GiftPackageImageInfo struct {
	// GiftPackageImageID 礼包图片。要求图片大小在200k以内且尺寸比例为1：1，建议参考尺寸144*144px。
	GiftPackageImageID string `json:"gift_package_image_id,omitempty"`
	// GiftPackageText 礼包所属文案（如史诗表情X50）。长度小于等于16，一个中文占2位。
	GiftPackageText string `json:"gift_package_text,omitempty"`
}

// CardPackageImageInfo 抽卡图片信息，数量限制为1。当style_type=15时该字段必填
type CardPackageImageInfo struct {
	// CardFront 当style_type=15时该字段为默认url，支持自定义
	CardFront string `json:"card_front,omitempty"`
	// CardText 当style_type=15时该字段为卡牌说明，选填。若不填写，则不显示卡牌说明文案
	CardText string `json:"card_text,omitempty"`
}

// GiftPackageCode 礼包码
type GiftPackageCode struct {
	// Code 礼包码
	// 字符数≤12
	Code string `json:"code,omitempty"`
	// Probability 概率
	// 概率和=100
	Probability int `json:"probability,omitempty"`
}

// PairInfo 配对信息
type PairInfo struct {
	// BackgroundImage 背景
	BackgroundImage string `json:"background_image,omitempty"`
	// LeftAvatar 左头像
	LeftAvatar string `json:"left_avatar,omitempty"`
	// Title 激励文案
	// 字符数≤16
	Title string `json:"title,omitempty"`
}

// LightInterActive 轻互动组件
type LightInterActive struct {
	// StyleType 互动前玩法
	// 枚举值：1-翻卡、2-盲盒、3-翻页、4-擦除、5-刮刮卡
	StyleType int `json:"style_type,omitempty"`
	// ShowSeconds 出现时间
	// 互动组件相对视频的展现时间，建议游戏行业10s，非游戏行业4s。
	// 枚举值：4000-4s、10000-10s
	ShowSeconds int64 `json:"show_seconds,omitempty"`
	// Tips 提示标题
	// 字符数≤10
	Tips string `json:"tips,omitempty"`
	// ImageLIst 卡片图片列表/盲盒图片列表; 长度=1
	ImageList []CardImage `json:"image_list,omitempty"`
	// Title 卖点标题
	// 字符数≤7
	Title string `json:"title,omitempty"`
	// InteractiveType 互动后玩法
	// 枚举值：1-全局优惠、2-礼包码、3-配对、4-单品优惠、5-新人优惠
	InteractiveType int `json:"interactive_type,omitempty"`
	// RewardTips 领取后提示
	// 字符数≤16
	RewardTips string `json:"reward_tips,omitempty"`
	// CouponInfo 优惠券信息
	CouponInfo *CouponInfo `json:"coupon_info,omitempty"`
	// GiftPackageInfo 礼包码信息
	GiftPackageInfo *GiftPackageInfo `json:"gift_package_info,omitempty"`
	// PairInfo 配对信息
	PairInfo *PairInfo `json:"pair_info,omitempty"`
}

// Type Implement ComponentData interface
func (l LightInterActive) Type() enum.ComponentType {
	return enum.ComponentType_LIGHT_INTER_ACTIVE
}

var _ ComponentData = (*LightInterActive)(nil)

// ReservationButton 游戏预约按钮
type ReservationButton struct {
	// AppDesc 应用描述，长度是1-15个字（两个英文字符占1个字）
	AppDesc string `json:"app_desc,omitempty"`
	// AppIntroduction 应用介绍，长度是1-20个字（两个英文字符占1个字）
	AppIntroduction string `json:"app_introduction,omitempty"`
	// AppThumbnails 应用图片集，图片id，从【素材管理-获取图片素材】接口获取;
	// 如果组件类型component_type是RESERVATION_BUTTON（游戏预约按钮）时必须三个，建议宽高比 9:16
	AppThumbnails []string `json:"app_thumbnails,omitempty"`
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
	// AppLogo 应用logo，可通过【获取图片素材】接口获取
	AppLogo string `json:"app_logo,omitempty"`
	// PackageName 应用包名称
	PackageName string `json:"package_name,omitempty"`
	// DownloadURL 应用下载链接
	// IOS下载链接：需要为iTunes官方地址
	// Android下载链接：需要为「应用管理中心」提供下载链接。
	DownloadURL string `json:"download_url,omitempty"`
}

// Type Implement ComponentData interface
func (r ReservationButton) Type() enum.ComponentType {
	return enum.ComponentType_RESERVATION_BUTTON
}

var _ ComponentData = (*ReservationButton)(nil)

// RewardInfo 红包信息。
type RewardInfo struct {
	// Title 主标题。长度小于等于16，一个中文占2位。
	Title string `json:"title,omitempty"`
	// RewardTips 互动后玩法引导文案（如恭喜获得XX元红包 ）。长度小于等于30，一个中文占2位。
	RewardTips string `json:"reward_tips,omitempty"`
	// BonusAmount 奖励物。取值为0.1～1000.0之间的数字，至多保留一位小数。
	BonusAmount float64 `json:"bonus_amount,omitempty"`
	// AddonText 特殊说明，即红包描述文案（如实际金额以APP内为准）。长度小于等于30，一个中文占2位。
	AddonText string `json:"addon_text,omitempty"`
	// CtaText 按钮文案。枚举值：PREFER_SYSTEM_RECOMMAND-个性化文案、DOWNLOAD_AND_RECEIVE_NOW-立即下载领取、RECEIVE_NOW-立即领取
	CtaText string `json:"cta_text,omitempty"`
}

// ProductCouponInfo 商品优惠信息。
type ProductCouponInfo struct {
	// RewardTips 互动后玩法引导文案（如恭喜抽中商品，下载即可领取）。长度小于等于30，一个中文占2位。
	RewardTips string `json:"reward_tips,omitempty"`
	// ProductImageID 商品图片。要求图片大小在200k以内且尺寸比例为1：1，建议参考尺寸144*144px。
	ProductImageID string `json:"product_image_id,omitempty"`
	// ProductName 商品文案。长度小于等于30，一个中文占2位。
	ProductName string `json:"product_name,omitempty"`
	// AddonText 特殊说明。长度小于等于30，一个中文占2位。
	AddonText string `json:"addon_text,omitempty"`
	// CtaText 按钮文案。枚举值：PREFER_SYSTEM_RECOMMAND-个性化文案、DOWNLOAD_AND_RECEIVE_NOW-立即下载领取、RECEIVE_NOW-立即领取
	CtaText string `json:"cta_text,omitempty"`
	// CtaTextCustom 当cta_text为自定义按钮文案时，该字段必填。长度小于等于12，一个中文占2位。
	CtaTextCustom string `json:"cta_text_custom,omitempty"`
	// ProductImageInfo 商品图信息。数量上限为3。
	// 当style_type=5-翻卡-领商品1和8翻卡-领商品(单图)时该字段必填,需传1组商品图片和商品文案，否则报错。
	// 当style_type=9翻卡-领商品(多图)时该字段必填,需传3组商品图片和商品文案，否则报错。
	ProductImageInfo []ProductImageInfo `json:"product_image_info,omitempty"`
}

// ProductImageInfo 商品图信息。数量上限为3。
type ProductImageInfo struct {
	// ProductImageID 商品图片。要求图片大小在200k以内且尺寸比例为1：1，建议参考尺寸280*280px。
	ProductImageID string `json:"product_image_id,omitempty"`
	// ProductName 商品文案。长度小于等于30，一个中文占2位
	ProductName string `json:"product_name,omitempty"`
}

// OffsetTimeInfo 抵时长信息。
type OffsetTimeInfo struct {
	// RewardReduceDuration 抵扣时长。可输入5～25（包含）之间的整数。选填，若不填默认为5秒。
	RewardReduceDuration int64 `json:"reward_reduce_duration,omitempty"`
}

// UnionLightInteractive 穿山甲轻互动组件的component_data
type UnionLightInteractive struct {
	// StyleType 互动前玩法
	// 枚举值：1-翻卡-领红包1、2-领红包样式2、3-领红包样式3、4-红包雨、5-翻卡-领商品1、6-答题、7-滑动、12-宝箱雨、13-答题领红包、14-宝箱领礼包码、15-抽卡领礼包码、16-答题领礼包码场景
	StyleType int `json:"style_type,omitempty"`
	// ShowSeconds 出现时间
	// 互动组件相对视频的展现时间，可输入3.0～25.0（包含）之间的数字，至多保留1位小数。
	// 当style_type为2、3、4时，只允许输入5.0～25.0（包含）之间的数字
	ShowSeconds int64 `json:"show_seconds,omitempty"`
	// Tips 引导文案（如点击翻卡，赢取无门槛红包提现）。长度小于等于30，一个中文占2位。
	Tips string `json:"tips,omitempty"`
	// CompSkin 组件样式。允许值：FLIPCARD-卡片，TREASUREBOX-宝箱，REWARD-激励、NATIVE-原生。
	// 当style_type为1、5时，允许值为FLIPCARD，TREASUREBOX；当style_type为6、7时，允许值为REWARD、NATIVE。
	CompSkin string `json:"comp_skin,omitempty"`
	// Question 问题。长度小于等于30，一个中文占2位。
	// 当style_type为6-答题时，必填。
	Question string `json:"question,omitempty"`
	// Option1 选项一。长度小于等于16，一个中文占2位。
	// 当style_type为6-答题时，必填。
	Option1 string `json:"option_1,omitempty"`
	// Option2 选项二。长度小于等于16，一个中文占2位。
	// 当style_type为6-答题时，必填。
	Option2 string `json:"option_2,omitempty"`
	// AnswerOption 正确选项。允许值：1-选项一，2-选项二
	// 当style_type为6-答题时，必填。
	AnswerOption int `json:"answer_option,omitempty"`
	// GiftPackageInfo 礼包码信息
	GiftPackageInfo *GiftPackageInfo `json:"gift_package_info,omitempty"`
	// Title 互动前场景主标题，自定义上传文案。
	// 当style_type=2、3时，有默认值url:「翻卡有奖励」、当style_type=13时，有默认值url:「答题有奖励」。长度小于等于16，一个字符占两位
	// 当style_type=8、9时，有默认值url:「翻卡有奖励」，长度小于等于12，一个字符占两位
	// 当style_type=10、11时有默认值url：「翻卡有福利」
	// 当style_type=14时有默认值url：「领取宝箱好礼」
	// 当style_type=15时有默认值url：「翻卡得奖励」
	Title string `json:"title,omitempty"`
	// TreasureIcon 宝箱（或红包）icon。当style_type=1、2、3、4、8、9、12、14时，有默认icon URL，支持用户自定义
	TreasureIcon string `json:"treasure_icon,omitempty"`
	// CardBack 卡牌背面，当style_type=15时必填。有默认url，支持替换
	CardBack string `json:"card_back,omitempty"`
	// InteractiveType 互动后玩法。枚举值：1-红包、2-商品优惠、3-抵扣时长
	InteractiveType int `json:"interactive_type,omitempty"`
	// RewardInfo 红包信息。
	RewardInfo *RewardInfo `json:"reward_info,omitempty"`
	// ProductCouponInfo 商品优惠信息。
	ProductCouponInfo *ProductCouponInfo `json:"product_coupon_info,omitempty"`
	// OffsetTimeInfo 抵时长信息。
	OffsetTimeInfo *OffsetTimeInfo `json:"offset_time_info,omitempty"`
}

// Type Implement ComponentData interface
func (u UnionLightInteractive) Type() enum.ComponentType {
	return enum.ComponentType_UNION_LIGHT_INTERACTIVE
}

var _ ComponentData = (*UnionLightInteractive)(nil)

// LuckyBox 幸运盒子
type LuckyBox struct {
	// BoxUrlLogo 品牌logo图，url格式。尺寸：200px*200px，格式：jpg或png（二选一），大小：300kb
	BoxUrlLogo string `json:"box_url_logo,omitempty"`
	// BoxColor 提供RGB色值，如“#161823”
	BoxColor string `json:"box_color,omitempty"`
	// TitleFirst 标题第一段文案，默认为“预备！”。客户可以自定义文案，字符个数限制：<=3个中文字数（6个字符）
	TitleFirst string `json:"title_first,omitempty"`
	// Duration 倒计时时间，默认“5”。选择盒子倒计时时间 ，5-12秒；为了保证素材核心信息不被感染，建议5s之后出现
	Duration int `json:"duration,omitempty"`
	// TitleSecond 标题第二段文案，默认为“秒后有大量盒盒掉落”，秒固定不可配置，“后有大量盒盒掉落”文案可配置。可以自定义文案，字符个数限制：<=9个中文字数（18个字符）
	TitleSecond string `json:"title_second,omitempty"`
	// SubTitle 副标题文案，可以自定义文案，字符个数限制：<=14个中文字数（28个字符）
	SubTitle string `json:"sub_title,omitempty"`
	// SpecialBoxType 特效盒子类型，默认1金色盒子（当前只有一种类型，后续会开发其他类型）
	SpecialBoxType int `json:"special_box_type,omitempty"`
	// HintText 互动提示文案，用户提供自定义文案，字符个数限制：<=14个中文字数（28个字符）
	HintText string `json:"hint_text,omitempty"`
}

// Type Implement ComponentData interface
func (u LuckyBox) Type() enum.ComponentType {
	return enum.ComponentType_LUCKY_BOX
}

var _ ComponentData = (*LuckyBox)(nil)

// UnknownComponent 未知组件
type UnknownComponent []byte

// Type Implement ComponentData interface
func (u UnknownComponent) Type() enum.ComponentType {
	return enum.ComponentType_UNKNOWN
}

var _ ComponentData = (*UnknownComponent)(nil)
