package creative

import "github.com/bububa/oceanengine/marketing-api/enum"

// VideoMaterial 视频素材信息
type VideoMaterial struct {
	// ImageInfo 视频封面图片，传1张
	ImageInfo *ImageInfo `json:"image_info,omitempty"`
	// VideoInfo 视频素材信息
	VideoInfo *VideoInfo `json:"video_info,omitempty"`
	// DpaVideoTemplateType 商品库视频生成类型，创建DPA创意时可传入，传入后该素材下image_info与video_info不生效
	DpaVideoTemplateType enum.DpaVideoTemplateType `json:"dpa_video_template_type,omitempty"`
	// DpaVideoTaskIDs 自定义商品库视频模板ID，创建DPA创意时可传入，传入后该素材下image_info与video_info不生效，长度限制1，从【获取 DPA 商品库视频模板】接口中获取
	DpaVideoTaskIDs []string `json:"dpa_video_task_ids,omitempty"`
}

// VideoInfo 视频素材信息
type VideoInfo struct {
	// VideoID 视频ID，image_mode为视频素材时填写 可通过【获取视频素材】接口获得
	VideoID string `json:"video_id,omitempty"`
}
