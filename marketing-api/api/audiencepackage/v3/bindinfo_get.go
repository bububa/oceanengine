package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/audiencepackage/v3"
)

// BindInfoGet 定向包查询关联项目信息
// 定向包查询关联项目信息
// 可通过【获取定向包】接口获取定向包ID，根据定向包ID查询该定向包关联了哪些项目
func BindInfoGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.BindInfoGetRequest) (*v3.BindInfoGetResult, error) {
	var resp v3.BindInfoGetResponse
	if err := clt.GetAPI(ctx, "v3.0/audience_package_bindinfo/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
