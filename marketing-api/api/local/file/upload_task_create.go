package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/file"
)

// UploadTaskCreate 异步上传本地推视频
// 将视频文件通过连山云素材服务上传后获取到视频文件链接，再将获取到的连山云视频文件url作为入参的video_url通过素材库提供的视频上传接口进行文件上传
// 仅支持开发者购置连山云素材服务上传生成的tos链接上传，不支持其他三方链接地址
func UploadTaskCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.UploadTaskCreateRequest) (uint64, error) {
	var resp file.UploadTaskCreateResponse
	if err := clt.PostAPI(ctx, "v3.0/local/file/upload_task/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.TaskID, nil
}
