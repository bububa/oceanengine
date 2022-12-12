package asynctask

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
	// RangeFrom 分片下载开始位置（从 0 开始）。闭区间。缺省表示从第一个字节开始下载。
	RangeFrom int64 `json:"range_from,omitempty"`
	// RangeTo 分片下载结束位置。闭区间。缺省表示下载到最后一个字节。特别的， -1 表示到最后一个字节。
	RangeTo int64 `json:"range_to,omitempty"`
}

// Encode implement GetRequest interface
func (r DownloadRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("task_id", strconv.FormatUint(r.TaskID, 10))
	if r.RangeFrom > 0 {
		values.Set("range_from", strconv.FormatInt(r.RangeFrom, 10))
	}
	if r.RangeFrom != 0 {
		values.Set("range_to", strconv.FormatInt(r.RangeTo, 10))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
