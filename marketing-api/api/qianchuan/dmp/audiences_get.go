package dmp

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/dmp"
)

// AudiencesGet 查询人群包列表
// 该接口用于查询广告主下可用的人群包列表和信息，支持两类：不限营销目标的平台精选人群包、自定义人群包。
// 直播带货专属的平台精选人群包的使用，可在创编计划时直接从live_platform_tags字段传入对应枚举，获取计划信息时同理。
func AudiencesGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.AudiencesGetRequest) (*dmp.AudiencesGetResponseData, error) {
	var resp dmp.AudiencesGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/dmp/audiences/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
