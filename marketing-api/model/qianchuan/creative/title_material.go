package creative

import "github.com/bububa/oceanengine/marketing-api/enum/qianchuan"

// TitleMaterial 标题素材
type TitleMaterial struct {
	// ID 素材唯一标识
	ID uint64 `json:"id,omitempty"`
	// Title 创意标题
	Title string `json:"title,omitempty"`
	// TitleType 素材类型，可选值
	// CUSTOM自定义标题
	// COMMODITY_CARD商品卡标题
	// 注意：商品卡标题限制条件如下：
	// 广告主类型为商家
	// 抖音号选择的是「官方」或「自运营」
	TitleType qianchuan.TitleType `json:"title_type,omitempty"`
	// AwemeCarouselID 抖音主页图文id
	AwemeCarouselID uint64 `json:"aweme_carousel_id,omitempty"`
	// DynamicWords 动态词包对象列表
	DynamicWords []DynamicWord `json:"dynamic_words,omitempty"`
}
