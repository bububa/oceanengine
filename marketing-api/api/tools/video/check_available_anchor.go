package video

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/video"
)

// CheckAvailableAnchor 查询视频是否挂载下载类锚点
// 建议优先选择带游戏/网服/平台电商的行业锚点视频投放，更有可能获得平台优质流量补贴
func CheckAvailableAnchor(ctx context.Context, clt *core.SDKClient, accessToken string, req *video.CheckAvailableAnchorRequest) ([]video.AvailableAnchor, error) {
	var resp video.CheckAvailableAnchorResponse
	if err := clt.Get(ctx, "2/tools/video/check_available_anchor/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
