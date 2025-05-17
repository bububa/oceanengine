package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoAdGet 获取同主体下广告主视频素材
// 通过此接口，用户可以查询同主体下的广告主视频信息。
// 接口权限：查看视频的信息，获取视频md5、宽高、预览地址等内容，合理利用还可以搭建自己的素材库进行素材的管理。
// - 为保证接口使用的安全性避免调取他人的视频信息，该接口只可用于查询自己公司下的视频信息，即需查询的视频ID所属广告主账号的主体需与APPID对应开发者的主体保持一致，才可获取到图片的信息！
// - 如果您的开发者账号还未完成企业认证也将无法调用。请先完成企业认证！
func VideoAdGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoAdGetRequest) ([]file.Video, error) {
	var resp file.VideoGetResponse
	err := clt.Get(ctx, "2/file/video/ad/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.List, nil
}
