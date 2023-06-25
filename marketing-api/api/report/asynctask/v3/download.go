package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/report/asynctask/v3"
)

// Download 下载任务结果
func Download(clt *core.SDKClient, accessToken string, req *v3.DownloadRequest) ([]byte, error) {
	return clt.GetBytes("v3.0/report/custom/async_task/download/", req, accessToken)
}
