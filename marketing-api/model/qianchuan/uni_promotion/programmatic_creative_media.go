package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
)

// ProgrammaticCreativeMedia 程序化创意信息
type ProgrammaticCreativeMedia struct {
	// VideoMaterial 视频信息
	VideoMaterial []VideoMaterial `json:"video_material,omitempty"`
	// TitleMaterial 标题信息
	// 注意：如果视频全部为抖音主页视频，不支持设置title
	TitleMaterial []TitleMaterial `json:"title_material,omitempty"`
	// BlockVideoMaterial 排除抖音主页视频列表
	BlockVideoMaterial []VideoMaterial `json:"block_video_material,omitempty"`
}

// VideoMaterial 视频信息
type VideoMaterial struct {
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// VideoDuration 视频时长
	VideoDuration float64 `json:"video_duration,omitempty"`
	// Title 视频标题
	Title string `json:"title,omitempty"`
	// URL 素材库视频预览链接
	// 注意：对于抖音主页视频不会返回预览链接，可自行通过如下规则拼接：https://www.douyin.com/video/{aweme_item_id}
	URL string `json:"url,omitempty"`
	// CoverImage 视频封面
	CoverImage *Image `json:"cover_image,omitempty"`
	// ImageMode 素材类型，支持视频
	// VIDEO_LARGE 横版视频
	// VIDEO_VERTICAL 竖版视频
	// 注意：当视频素材选择抖音号下视频时，image_mode必须传竖版视频
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// VideoID 视频ID，可通过【获取视频素材】接口获得视频素材id
	VideoID string `json:"video_id,omitempty"`
	// VideoCoverID 视频封面ID，可通过【获取图片素材】接口获得图片素材id
	VideoCoverID string `json:"video_cover_id,omitempty"`
	// AwemeItemID 抖音视频ID，可通过【获取抖音号下视频素材】接口获得视频id，使用抖音视频的时候默认忽略video_id
	AwemeItemID uint64 `json:"aweme_item_id,omitempty"`
	// AwemeCarouselID 抖音图文id，可通过【获取全域可排除抖音视频/图文列表】接口获得视频id
	AwemeCarouselID uint64 `json:"aweme_carousel_id,omitempty"`
	// StarTraffic 是否星图达人视频，可选值:
	// 1是
	StarTraffic int `json:"star_traffic,omitempty"`
}

// Image 图片
type Image struct {
	// Height 宽
	Height int `json:"height,omitempty"`
	// Width 高
	Width int `json:"width,omitempty"`
	// ID 图片id
	ID string `json:"id,omitempty"`
	// ImageURL 图片链接
	ImageURL string `json:"image_url,omitempty"`
}

// TitleMaterial 标题信息
type TitleMaterial struct {
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// Title 标题
	Title string `json:"title,omitempty"`
	// TitleType  素材类型，可选值:
	// COMMODITY_CARD 商品卡标题
	// CUSTOM 自定义标题
	// 注意：如果添加了商品卡图片，必须添加一个商品卡标题
	TitleType qianchuan.TitleType `json:"title_type,omitempty"`
	// DynamicWords 动态词包对象列表，最多支持两个词包。可使用【查询动态创意词包】获得，结合标题中的词包格式您需要填写相同个数与顺序的词包ID，如果实际ID顺序与标题中词包名顺序不一致我们将以词包ID顺序为准
	DynamicWords []DynamicWord `json:"dynamic_words,omitempty"`
}

// DynamicWord 动态词包对象
type DynamicWord struct {
	// WordID 动态词包ID
	WordID string `json:"word_id,omitempty"`
	// DictName 创意词包名称
	DictName string `json:"dict_name,omitempty"`
	// DefaultWord 创意词包默认词
	DefaultWord string `json:"default_word,omitempty"`
}
