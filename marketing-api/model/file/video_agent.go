package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// VideoAgentRequest 【代理商】上传自产首发素材至方舟（搬运治理） API Request
type VideoAgentRequest struct {
	// AgentID 代理商id
	AgentID uint64 `json:"agent_id,omitempty"`
	// UploadType 视频上传方式，暂仅支持文件上传，可选值:
	// UPLOAD_BY_FILE 文件上传（默认值）
	UploadType enum.UploadType `json:"upload_type,omitempty"`
	// Filename 视频上传名称
	Filename string `json:"filename,omitempty"`
	// VideoFile 视频文件数据，upload_type为UPLOAD_BY_FILE必填
	VideoFile *model.UploadField `json:"video_file,omitempty"`
	// VideoUrl 图片url地址，如http://xxx.xxx，upload_type为UPLOAD_BY_URL必填
	VideoUrl string `json:"video_url,omitempty"`
	// IsAigc 标注是否为AIGC生成
	IsAigc bool `json:"is_aigc,omitempty"`
	// IsNeedAuth 标注是否允许生效授权保护（前提素材需要是自产！！），目前只支持true 默认值: True
	IsNeedAuth bool `json:"is_need_auth,omitempty"`
}

// Encode implement UploadRequest interface
func (r VideoAgentRequest) Encode() []model.UploadField {
	fields := []model.UploadField{
		{
			Key:   "agent_id",
			Value: strconv.FormatUint(r.AgentID, 10),
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
	} else {
		fields = append(fields, model.UploadField{
			Key:   "is_aigc",
			Value: "failse",
		})
	}
	if r.IsNeedAuth {
		fields = append(fields, model.UploadField{
			Key:   "is_need_auth",
			Value: "true",
		})
	} else {
		fields = append(fields, model.UploadField{
			Key:   "is_need_auth",
			Value: "failse",
		})
	}
	return fields
}

// VideoAgentResponse 【代理商】上传自产首发素材至方舟（搬运治理） API Response
type VideoAgentResponse struct {
	model.BaseResponse
	Data struct {
		VideoInfo *Video `json:"video_info,omitempty"`
	} `json:"data,omitempty"`
}
