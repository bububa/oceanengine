package dmp

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/dmp"
)

// AudienceFilePartUpload 大文件分片上传
// 支持用户上传本地大文件分片到DMP数据平台上
// 注意：
// 文件内容为一行一个id，以txt/csv（一个或多个文件）格式zip压缩，非合法id数据行不影响其他合法id行的解析
// 分片上传是指对压缩后需要上传的zip文件按一定大小进行分片，上传的是同一个zip文件拆分后的多个部分，并非是对原始txt/csv文件拆分后压缩成多个zip文件上传
// 文件上传后生成的人群包会自动去重，但为减轻计算压力，业务方应尽量自行去重
// 文件大于25M，需使用该接口分片上传；单个分片大小，需大于5M且小于25M，建议20M；总分片个数不能超过10000
func AudienceFilePartUpload(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.AudienceFilePartUploadRequest) (string, error) {
	var resp dmp.AudienceFilePartUploadResponse
	if err := clt.Upload(ctx, "v1.0/qianchuan/audience_file/part_upload/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.FileKey, nil
}
