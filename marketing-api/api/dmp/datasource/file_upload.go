package datasource

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dmp/datasource"
)

// FileUpload 数据源文件上传
// 当用户需要上传本地数据到DMP数据平台上时，需要使用数据源文件上传功能。用户上传数据源文件后会返回文件路径file_path，用于调用【数据源创建】时创建相应的数据源。
// 文件内容：可接受的文件内容为设备id和用户id，需要在序列化时选择相应的匹配类型。 文件格式：数据源文件需要经过protocal buffer序列化,序列化具体操作方式参见【DMP相关附录-DMP上传数据格式】。数据源序列化后需要压缩成zip文件； 文件个数：不限制压缩包内文件个数。
// 压缩包大小不超过50M！
// 接口限制了10s超时，建议文件不要太大！
func FileUpload(ctx context.Context, clt *core.SDKClient, accessToken string, req *datasource.FileUploadRequest) (string, error) {
	var resp datasource.FileUploadResponse
	err := clt.Upload(ctx, "2/dmp/data_source/file/upload/", req, &resp, accessToken)
	if err != nil {
		return "", err
	}
	return resp.Data.FilePath, nil
}
