package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoAgent 【代理商】上传自产首发素材至方舟（搬运治理）
// 代理商可以通过此接口将「首发素材」上传视频素材至巨量方舟，上传后「首发素材」即可自动完成保护（前置需要先完成整体保护授权，参考详细文档介绍「搬运治理-首发保护」说明手册（可对外））。保护后系统将根据代理授权范围识别搬运素材生效打压，避免其他方抢夺代理的流量。
// 注意1：建议代理商务必将自己制作的「首发素材」先通过本接口上传到方舟平台（素材才可能按照代理授权生效，否则可能会被客户先做授权保护），避免素材被其他方抢先保护。
// 注意2：非首发素材，即使进行上传也不会生效授权保护（点此查看「首发素材」相关说明）
// 注意3：调用接口上传的素材可在“方舟平台-优化-搬运治理页面”查看
func VideoAgent(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoAgentRequest) (*file.Video, error) {
	var resp file.VideoAgentResponse
	err := clt.Upload(ctx, "2/file/video/agent/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.VideoInfo, nil
}
