package appmanagement

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DownloadPackageParseRequest 提交解析应用包任务 API Request
type DownloadPackageParseRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DownloadURL 下载链接
	DownloadURL string `json:"download_url,omitempty"`
}

// Encode implement PostRequest interface
func (r DownloadPackageParseRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DownloadPackageParseResponse 提交解析应用包任务 API Response
type DownloadPackageParseResponse struct {
	model.BaseResponse
	Data struct {
		// EventID 事件ID，用于查询解析包状态接口，有效期为24小时
		EventID string `json:"event_id,omitempty"`
	} `json:"data,omitempty"`
}
