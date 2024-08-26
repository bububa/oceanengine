package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// AvatarUpload 获取广告主账户头像ID
// 本接口用户获取广告主账户头像的image_id，您可使用该id调用「更新广告主账户头像」接口完成账户头像更新
// 【注意】本接口的功能仅用于获取image_id，上传成功 ≠ 更新头像，更新头像的接口是「更新广告主账户头像」接口
func AvatarUpload(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.AvatarUploadRequest) (string, error) {
	var resp advertiser.AvatarUploadResponse
	if err := clt.Upload(ctx, "2/advertiser/avatar/upload/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.ImageID, nil
}
