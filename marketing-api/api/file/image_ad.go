package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// ImageAd 上传广告图片
// 通过此接口，用户可以上传和广告相关的素材图片，例如创意素材。
// 图片格式：jpg、jpeg、png、bmp、gif，大小1.5M内
// 上传的图片一定要符合格式，才会在投放平台素材库显示！
// 若同一素材已进行上传，重新上传不会改名！
func ImageAd(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.ImageAdRequest) (*file.Image, error) {
	var resp file.ImageAdResponse
	err := clt.Upload(ctx, "2/file/image/ad/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
