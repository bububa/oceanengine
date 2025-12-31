package unipromotion

import "github.com/bububa/oceanengine/marketing-api/enum"

// MultiProductCreative 商品全域创意素材信息
type MultiProductCreative struct {
	// ProductID 商品id
	ProductID uint64 `json:"product_id,omitempty"`
	// CreativeType 创意类型，可选值:
	// PROGRAMMATIC_CREATIVE 程序化创意
	CreativeType enum.CreativeMaterialMode `json:"creative_type,omitempty"`
	// HideInAweme 抖音主页可见性设置，和抖音号关系类型相关，返回值参考【附录-抖音号授权类型】
	// 仅单次展示可见 true
	// 主页始终可见 false
	// 官方+自运营（bind_type为OFFICIAL或SELF）
	// 1、全是抖音号主页视频，无需传，传了亦无效
	// 2、存在非抖音号主页原生视频，支持设置
	// 达人（bind_type不为OFFICIAL或SELF）
	// 1、不支持设置，传了亦无效
	HideInAweme bool `json:"hide_in_aweme,omitempty"`
	// VideoMaterial 视频素材
	VideoMaterial []VideoMaterial `json:"video_material,omitempty"`
	// ImageMaterial 图片素材
	ImageMaterial []ImageMaterial `json:"image_material,omitempty"`
	// TitleMaterial 支持0-30个标题
	// 注意
	// 1、当选择非抖音主页视频or图片时，请至少添加一个标题，否则素材不生效
	// 2、如果素材全都是抖音主页视频，不支持添加标题
	TitleMaterial []TitleMaterial `json:"title_material,omitempty"`
	// CreativeCard 投放卡片
	// 注意：在商品全域下，如果是无号商家，无需入参；如果是有号商家，必填
	CreativeCard *CreativeCard `json:"creative_card,omitempty"`
	// BlockVideoMaterial 排除抖音主页视频/图文列表
	BlockVideoMaterial []VideoMaterial `json:"block_video_material,omitempty"`
	// CarouselMaterial 图文素材
	CarouselMaterial []CarouselMaterial `json:"carousel_material,omitempty"`
}

// ImageMaterial 图片素材
type ImageMaterial struct {
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// ImageID 图片id
	ImageID string `json:"image_id,omitempty"`
	// ImageURL 图片链接
	ImageURL string `json:"image_url,omitempty"`
	// ImageMode 素材类型，仅支持商品卡方图SQUARE
	ImageMode string `json:"image_mode,omitempty"`
	// ImageIDs 图片ID列表，当 image_mode 为 图片素材 时必填，图片ID可通过【素材管理】中的接口获得。目前仅支持上传一张图片，图片大小不能超过1.5M
	ImageIDs []string `json:"image_ids,omitempty"`
}

// CreativeCard 投放卡片
type CreativeCard struct {
	// PromotionCardTitle 默认值: 视频同款商品
	PromotionCardTitle string `json:"promotion_card_title,omitempty"`
	// PromotionCardSellingPoints 投放卡片卖点列表，卖点文字长度要求为11~18个字符，汉字算2个字符
	PromotionCardSellingPoints []string `json:"promotion_card_selling_points,omitempty"`
	// PromotionCardImageID 投放卡片配图，可通过【获取图片素材】接口获得图片素材id
	PromotionCardImageID string `json:"promotion_card_image_id,omitempty"`
	// PromotionCardActionButton  投放卡片行动号召按钮文案，可选值:
	// 专属优惠 专属优惠
	// 了解更多 了解更多
	// 去逛逛 去逛逛
	// 更多秒杀 更多秒杀
	// 更多精彩 更多精彩
	// 查看详情 查看详情
	// 立即抢购 立即抢购
	// 领取福利 领取福利
	//  默认值: 查看详情
	PromotionCardActionButton string `json:"promotion_card_action_button,omitempty"`
	// PromotionCardButtonStartOptimization 是否对行动号召按钮文案启用智能优选，默认false关闭
	PromotionCardButtonStartOptimization bool `json:"promotion_card_button_smart_optimization,omitempty"`
}

// CarouselMaterial 图文素材
type CarouselMaterial struct {
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// Title 图文标题
	Title string `json:"title,omitempty"`
	// Description 图文描述
	Description string `json:"description,omitempty"`
	// Image 图片列表
	Image []Image `json:"image,omitempty"`
	// MusicURL 音频链接
	MusicURL string `json:"music_url,omitempty"`
	// AwemeCarouselID 抖音图文id
	// 注意：使用aweme_carousel_id时会忽略carousel_id
	AwemeCarouselID uint64 `json:"aweme_carousel_id,omitempty"`
	// CarouselID 图文id
	CarouselID uint64 `json:"carousel_id,omitempty"`
}
