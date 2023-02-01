package creative

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
)

// ImageMaterial 创意图片素材
type ImageMaterial struct {
	// ImageMode 素材类型，必填，注意：程序化创意不支持组图 CREATIVE_IMAGE_MODE_GROUP，其他类型图片都支持，如横版/竖版大图、小图。详见【附录-素材类型】
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// ImageInfo 图片素材信息
	ImageInfo *ImageInfoWrapper `json:"image_info,omitempty"`
	// TemplateImage 图片模版信息，创建DPA创意时可传入,选择模板后image_info传入内容无效
	TemplateImage *TemplateImage `json:"template_image,omitempty"`
}

// ImageInfoWrapper image_info wrapper
// image_info 可能为slice也可能为object
type ImageInfoWrapper struct {
	Object *ImageInfo
	List   []ImageInfo
}

func (i *ImageInfoWrapper) IsObject() bool {
	return i.Object != nil
}

// UnmarshalJSON implement json Unmarshal interface
func (i *ImageInfoWrapper) UnmarshalJSON(b []byte) (err error) {
	var ret ImageInfoWrapper
	if b[0] == '[' && b[len(b)-1] == ']' {
		if err := json.Unmarshal(b, &ret.List); err != nil {
			return err
		}
	} else if b[0] == '{' && b[len(b)-1] == '}' {
		ret.Object = new(ImageInfo)
		if err := json.Unmarshal(b, ret.Object); err != nil {
			return err
		}
	}
	if ret.Object == nil && len(ret.List) == 0 {
		i = nil
	} else {
		*i = ret
	}
	return nil
}

// MmarshalJSON implement json Marshal interface
func (i *ImageInfoWrapper) MarshalJSON() ([]byte, error) {
	if i.Object != nil {
		return json.Marshal(i.Object)
	} else if len(i.List) == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(i.List)
}

var _ json.Unmarshaler = (*ImageInfoWrapper)(nil)
var _ json.Marshaler = (*ImageInfoWrapper)(nil)

// ImageInfo 图片素材信息
type ImageInfo struct {
	// ImageMode 素材类型，必填，注意：程序化创意不支持组图 CREATIVE_IMAGE_MODE_GROUP，其他类型图片都支持，如横版/竖版大图、小图。详见【附录-素材类型】
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// ImageID 图片ID，image_mode为图片素材时填写 可通过【获取图片素材】接口获得
	ImageID string `json:"image_id,omitempty"`
	// VideoID 视频ID，视频素材时填写。可通过【获取视频素材】接口获得
	VideoID string `json:"video_id,omitempty"`
	// ImageIDs 图片ID列表，非视频素材时填写。图片ID和视频ID可通过【获取图片素材】接口获得。组图类型传3张图，其他图片类型传1张，否则会报错。图片大小不能超过1.5M
	ImageIDs []string `json:"image_ids,omitempty"`
	// TemplateIDs 模版ID列表
	TemplateIDs []uint64 `json:"template_ids,omitempty"`
	// TemplateID DPA模板ID，针对DPA广告，且对应的素材类型是大图、小图、组图。可通过【获取DPA模板】接口查询模版ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// TemplateDataList 模版自定义参数
	TemplateDataList []TemplateData `json:"template_data_list,omitempty"`
	// Name 图片描述，搜索创意素材类型为橱窗素材时可传入，长度限制4-6字，两个英文字符占1位。
	Name string `json:"name,omitempty"`
}

// TemplateImage 图片模版信息
type TemplateImage struct {
	// TemplateID 图片素材类型-DPA模板ID，针对DPA广告，且对应的素材类型是大图、小图、组图。可通过【获取DPA模板】接口查询模版ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// TemplateDataList 模版自定义参数
	TemplateDataList []TemplateData `json:"template_data_list,omitempty"`
}

// TemplateData 模版自定义参数
type TemplateData struct {
	// BackgroundImageID 自定义背景图片ID，图片尺寸必须与模版背景图尺寸一致。图片ID可通过【获取图片素材】接口获得
	BackgroundImageID string `json:"background_image_id,omitempty"`
}
