package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// MaterialBind 素材推送
// 通过此接口，用户可以进行同主体下不同广告主间的素材的推送。也就是说，将A广告主素材推送到，与A广告主主体（公司）相同的广告主。
// 推送后素材的名称不会改变，将使用推送的原素材名！
// 推送前需确认推送素材属于当前操作广告主，可通过【获取图片素材】和【获取视频素材】接口确认！
// 新上传素材存在同步延迟情况，建议等待2-3分钟再尝试操作推送！
// 当素材已存在待推送的广告主的素材库内时，不会重复推送，推送失败的结果会在推送失败列表展示！
func MaterialBind(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.MaterialBindRequest) ([]file.FailedMaterialBind, error) {
	var resp file.MaterialBindResponse
	err := clt.Post(ctx, "2/file/material/bind/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.FailList, nil
}
