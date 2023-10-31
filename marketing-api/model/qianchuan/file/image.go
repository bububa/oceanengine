package file

import "github.com/bububa/oceanengine/marketing-api/enum"

// Image 图片
type Image struct {
	// ID 图片ID
	ID string `json:"id,omitempty"`
	// MaterialID 素材id，即多合一报表中的素材id，一个素材唯一对应一个素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// Size 图片大小
	Size uint64 `json:"size,omitempty"`
	// Width 图片宽度
	Width int `json:"width,omitempty"`
	// Height 图片高度
	Height int `json:"height,omitempty"`
	// URL 图片预览地址，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”
	// 链接仅做预览使用，预览链接有效期为1小时
	URL string `json:"url,omitempty"`
	// Format 图片格式
	Format string `json:"format,omitempty"`
	// Signature 图片md5
	Signature string `json:"signature,omitempty"`
	// CreateTime 素材的上传时间，格式："yyyy-mm-dd HH:MM:SS"
	CreateTime string `json:"create_time,omitempty"`
	// Filename 素材的文件名
	Filename string `json:"filename,omitempty"`
	// ImageMode 素材类型，允许值:
	// SQUARE 方图
	// LARGE 大图
	// SMALL 小图
	// LARGE_VERTICAL 大图竖图
	// UNION_SPLASH 穿山甲开屏图片
	// VIDEO_LARGE 横版视频
	// VIDEO_VERTICAL 竖版视频
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// Tag 素材标签
	Tag string `json:"tag,omitempty"`
	// 是否AI生成
	// false：不是AI生成
	// true：AI生成
	IsAiCreate bool `json:"is_ai_create,omitempty"`
}
