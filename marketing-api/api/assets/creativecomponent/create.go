package creativecomponent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/assets/creativecomponent"
)

// Create 创建组件 https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710672391183
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *creativecomponent.CreateRequest) (*creativecomponent.CreateResponseData, error) {
	var resp creativecomponent.CreateResponse

	if err := clt.Post(ctx, "2/assets/creative_component/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
