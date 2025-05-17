package dmp

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/dmp"
)

// AudienceCreateByFile 上传人群
// 基于已上传的数据文件，实现人群上传功能
func AudienceCreateByFile(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.AudienceCreateByFileRequest) (uint64, error) {
	var resp dmp.AudienceCreateByFileResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/audience/create_by_file/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.AudienceID, nil
}
