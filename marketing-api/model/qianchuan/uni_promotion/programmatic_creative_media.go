package unipromotion

import "github.com/bububa/oceanengine/marketing-api/enum"

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
}

// TitleMaterial 标题信息
type TitleMaterial struct {
	// Title 标题
	Title string `json:"title,omitempty"`
}
