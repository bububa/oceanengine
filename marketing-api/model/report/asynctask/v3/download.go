package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/util"
)

// DownloadRequest 下载任务结果
type DownloadRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TaskID 任务 id
	TaskID uint64 `json:"task_id,omitempty"`
}

// Encode implement GetRequest interface
func (r DownloadRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("task_id", strconv.FormatUint(r.TaskID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
