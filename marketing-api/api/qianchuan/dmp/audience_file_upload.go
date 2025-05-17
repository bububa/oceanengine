package dmp

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/dmp"
)

// AudienceFileUpload 小文件直接上传
// 支持用户上传本地小文件到DMP数据平台上
// 注意：
// 1、压缩包大小不得超过25M！
// 2、文件内容为一行一个id，以txt/csv（一个或多个文件）格式zip压缩，非合法id数据行不影响其他合法id行的解析
func AudienceFileUpload(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.AudienceFileUploadRequest) (*dmp.AudienceFileUploadResult, error) {
	var resp dmp.AudienceFileUploadResponse
	if err := clt.Upload(ctx, "v1.0/qianchuan/audience_file/upload/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
