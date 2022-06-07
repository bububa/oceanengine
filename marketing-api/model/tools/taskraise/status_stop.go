package taskraise

import (
	"encoding/json"
)

// StatusStopRequest 关闭优选起量任务
type StatusStopRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ReportID 任务id
	ReportID uint64 `json:"report_id,omitempty"`
}

// Encode implement PostRequest interface
func (r StatusStopRequest) Encode() []byte {
	data, _ := json.Marshal(r)
	return data
}
