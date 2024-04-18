package file

import (
	"encoding/json"
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"strconv"
)

// VideoAdRequest 上传广告视频 API Request
type VideoAdRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UploadType 图片上传方式; 默认值：UPLOAD_BY_FILE; 允许值：UPLOAD_BY_FILE、UPLOAD_BY_URL
	UploadType enum.UploadType `json:"upload_type,omitempty"`
	// VideoSignature 视频的md5值(用于服务端校验)
	VideoSignature string `json:"video_signature,omitempty"`
	// VideoFile 视频文件; 允许格式：mp4、mpeg、3gp、avi（10s超时限制）
	VideoFile *model.UploadField `json:"video_file,omitempty"`
	// Filename 素材的文件名，可自定义素材名，不传则默认取文件名，最长255个字符; 注：若同一素材已进行上传，重新上传不会改名
	Filename string `json:"filename,omitempty"`
	// Label 标签
	Label []string `json:"label,omitempty"`
	// VideoUrl 视频 url地址，如http://xxx.xxx，upload_type为UPLOAD_BY_URL必填
	VideoUrl string `json:"video_url,omitempty"`
	// IsAigc 图片素材是否为AIGC生成
	IsAigc bool `json:"is_aigc,omitempty"`
	// isGuideVideo 上传视频场景是否是引导视频（此功能仅游戏行业白名单客户可用，如需使用请联系运营）
	IsGuideVideo bool `json:"is_guide_video,omitempty"`
}

// Encode implement UploadRequest interface
func (r VideoAdRequest) Encode() []model.UploadField {
	fields := []model.UploadField{
		{
			Key:   "advertiser_id",
			Value: strconv.FormatUint(r.AdvertiserID, 10),
		},
	}
	if r.UploadType == enum.UPLOAD_BY_URL {
		fields = append(fields, model.UploadField{
			Key:   "upload_type",
			Value: string(r.UploadType),
		}, model.UploadField{
			Key:   "video_url",
			Value: r.VideoUrl,
		})
	} else if r.VideoFile != nil {
		filename := r.VideoFile.Value
		if filename == "" {
			filename = "file"
		}
		fields = append(fields, model.UploadField{
			Key:   "video_signature",
			Value: r.VideoSignature,
		}, model.UploadField{
			Key:    "video_file",
			Value:  filename,
			Reader: r.VideoFile.Reader,
		})
	}
	if r.Filename != "" {
		fields = append(fields, model.UploadField{
			Key:   "filename",
			Value: r.Filename,
		})
	}
	if r.IsAigc {
		fields = append(fields, model.UploadField{
			Key:   "is_aigc",
			Value: "true",
		})
	}
	if r.IsGuideVideo {
		fields = append(fields, model.UploadField{
			Key:   "is_guide_video",
			Value: "true",
		})
	}
	if len(r.Label) > 0 {
		value, _ := json.Marshal(r.Label)
		fields = append(fields, model.UploadField{
			Key:   "label",
			Value: string(value),
		})
	}
	return fields
}

// VideoAdResponse 上传广告视频 API Response
type VideoAdResponse struct {
	model.BaseResponse
	Data *Video `json:"data,omitempty"`
}
