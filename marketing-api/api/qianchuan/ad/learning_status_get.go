package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// LearningStatusGet 获取计划学习期状态
func LearningStatusGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.LearningStatusGetRequest) ([]ad.LearningStatus, error) {
	var resp ad.LearningStatusGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/ad/learning_status/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
