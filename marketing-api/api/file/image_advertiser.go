package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// ImageAdvertiser 上传广告主图片
// 通过此接口，用户可以按照一定方式上传符合格式的广告主资质相关图片，例如营业执照等，接口会返回"code_0"和"message_OK"，代表上传成功
// 图片格式：jpg、jpeg、png、bmp、gif，大小1.5M内
// 此接口上传的是广告主资质图片，如需上传广告素材图片请调用【上传广告图片】接口！
func ImageAdvertiser(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.ImageAdvertiserRequest) (*file.Image, error) {
	var resp file.ImageAdvertiserResponse
	err := clt.Upload(ctx, "2/file/image/advertiser/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
