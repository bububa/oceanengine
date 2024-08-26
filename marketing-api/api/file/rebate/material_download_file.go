package rebate

import (
	"context"
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file/rebate"
)

// MaterialDownloadFile 下载任务结果
// 通过指定的task_id,获取对应的数据明细文件
func MaterialDownloadFile(ctx context.Context, clt *core.SDKClient, accessToken string, req *rebate.MaterialDownloadFileRequest) (json.RawMessage, error) {
	var resp rebate.MaterialDownloadFileResponse
	if err := clt.Get(ctx, "2/file/rebate/material_download/download_file/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
