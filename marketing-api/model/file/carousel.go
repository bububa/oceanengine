package file

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
)

// Carousel 图集信息
type Carousel struct {
	// ID 图集id
	ID uint64 `json:"id,omitempty"`
	// CarouselID 图集id
	CarouselID uint64 `json:"carousel_id,omitempty"`
	// Uri 图集uri
	Uri string `json:"uri,omitempty"`
	// CarouselType 图集素材类型
	CarouselType enum.ImageMode `json:"carousel_type,omitempty"`
	// Images 图片信息
	Images []CarouselImage `json:"images,omitempty"`
	// Audio 音频信息
	Audio *Audio `json:"audio,omitempty"`
	// FileName 文件名
	FileName string `json:"file_name,omitempty"`
	// CreateTime 图集创建时间
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 图集更新时间
	UpdateTime string `json:"update_time,omitempty"`
	// Source  图文素材来源
	Source string `json:"source,omitempty"`
}

func (c *Carousel) UnmarshalJSON(data []byte) error {
	type C Carousel
	if err := json.Unmarshal(data, (*C)(c)); err != nil {
		return err
	}
	if c.ID == 0 {
		c.ID = c.CarouselID
	} else if c.CarouselID == 0 {
		c.CarouselID = c.ID
	}
	return nil
}

// AwemeCarousel 抖音图文素材
type AwemeCarousel struct {
	// ItemID 抖音图文id，在抖音号内上传图文素材时生成的id，是抖音端图文素材的唯一id
	ItemID uint64 `json:"item_id,omitempty"`
	//  MaterialID 抖音图文素材id，在查询素材维度的报表（会用该id作为维度，对应「素材管理」-「获取图文素材」等相关接口中的carousel_ids.id或carousel_id）
	MaterialID uint64 `json:"material_id,omitempty"`
	// CarouselTitle 抖音图文标题
	CarouselTitle string `json:"carousel_title,omitempty"`
	// CarouselPlayURL 抖音图文播放地址，格式为：https://douyin.com/video/XXXXXX
	CarouselPlayURL string `json:"carousel_play_url,omitempty"`
	// CarouselCoverURL 抖音图文素材内封面图片预览url，有效期10小时
	CarouselCoverURL string `json:"carousel_cover_url,omitempty"`
	// Images 抖音图文素材内的图片信息，最多35张图片
	// 按照图文本身的顺序返回，最前面的图片在最上面
	Images []CarouselImage `json:"images,omitempty"`
}

// CarouselImage 图集图片信息，包含图片id和图片主题
type CarouselImage struct {
	// ImageID 图片id，根据上传广告图片接口获取
	ImageID string `json:"image_id,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// Ratio 图片比例
	Ratio float64 `json:"ratio,omitempty"`
	// Size 图片大小
	Size uint64 `json:"size,omitempty"`
	// Width 图片宽度
	Width int `json:"width,omitempty"`
	// Height 图片高度
	Height int `json:"height,omitempty"`
	// URL 图片预览地址，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”
	// 链接仅做预览使用，预览链接有效期为1小时
	URL string `json:"url,omitempty"`
	// ImageSubject 图片主题
	ImageSubject *ImageSubject `json:"image_subject,omitempty"`
}

// ImageSubject 图片主题
type ImageSubject struct {
	// Description 图片详述
	Description string `json:"description,omitempty"`
	// Brand 品牌信息
	Brand string `json:"brand,omitempty"`
	// Text 图片文字
	Text string `json:"text,omitempty"`
	// Comment 备注
	Comment string `json:"comment,omitempty"`
}
