package customaudience

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dmp/customaudience"
)

// Read 人群包详细信息
// 用户可以通过调用此接口，查询广告主下的指定人群包信息。支持查询已删除的人群包信息，具体包含的信息内容请查看应答参数。
// 人群包创建时间大于6个月，且近6个月没绑定计划以及产生消耗，该人群包会过期。（对应该接口status返回26）对于过期的人群包，只能重新建包，再次推送。
func Read(ctx context.Context, clt *core.SDKClient, accessToken string, req *customaudience.ReadRequest) ([]customaudience.CustomAudience, error) {
	var resp customaudience.ReadResponse
	err := clt.Get(ctx, "2/dmp/custom_audience/read/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.CustomAudienceList, nil
}
