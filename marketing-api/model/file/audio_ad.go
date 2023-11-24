package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// AudioAdRequest 上传图文内的音频素材 API Request
type AudioAdRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UploadType 音频上传方式, 允许值:
	// UPLOAD_BY_FILE 文件上传（默认值）
	// UPLOAD_BY_URL 网址上传
	UploadType enum.UploadType `json:"upload_type,omitempty"`
	// AudioSignature 音频的md5值(用于服务端校验)
	// upload_type为UPLOAD_BY_FILE时必填
	AudioSignature string `json:"audio_signature,omitempty"`
	// AudioFile 音频文件，upload_type为UPLOAD_BY_FILE时必填
	// 允许格式：m4a、mp3
	// 文件大小：≤20M
	AudioFile *model.UploadField `json:"audio_file,omitempty"`
	// AudioURL 音频url地址，如http://xxx.xxx，upload_type为UPLOAD_BY_URL必填
	AudioURL string `json:"audio_url,omitempty"`
}

// Encode implement UploadRequest interface
func (r AudioAdRequest) Encode() []model.UploadField {
	fields := []model.UploadField{
		{
			Key:   "advertiser_id",
			Value: strconv.FormatUint(r.AdvertiserID, 10),
		},
	}
	if r.UploadType != "" {
		fields = append(fields, model.UploadField{
			Key:   "upload_type",
			Value: string(r.UploadType),
		})
	}
	if r.AudioFile != nil {
		filename := r.AudioFile.Value
		if filename == "" {
			filename = "file"
		}
		fields = append(fields, model.UploadField{
			Key:   "audio_signature",
			Value: r.AudioSignature,
		}, model.UploadField{
			Key:    "audio_file",
			Value:  filename,
			Reader: r.AudioFile.Reader,
		})
	}
	if r.AudioURL != "" {
		fields = append(fields, model.UploadField{
			Key:   "audio_url",
			Value: r.AudioURL,
		})
	}
	return fields
}

// AudioAdResponse 上传图文内的音频素材 API Response
type AudioAdResponse struct {
	model.BaseResponse
	Data struct {
		// AudioInfo 音频信息
		AudioInfo *Audio `json:"audio_info,omitempty"`
	} `json:"data,omitempty"`
}
