package appmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DownloadPackageGetRequest 查询解析应用包任务 API Request
type DownloadPackageGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// EventID 事件ID,由【提交解析应用包任务】获取
	EventID string `json:"event_id,omitempty"`
}

// Encode implement GetRequest interface
func (r DownloadPackageGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("event_id", r.EventID)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DownloadPackageGetResponse 查询解析应用包任务 API Response
type DownloadPackageGetResponse struct {
	model.BaseResponse
	Data struct {
		// PackageStatus 包解析状态，显示当前包解析状态
		PackageStatus DownloadPackageStatus `json:"package_status,omitempty"`
	} `json:"data,omitempty"`
}
