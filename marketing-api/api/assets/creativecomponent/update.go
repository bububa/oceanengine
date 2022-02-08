package creativecomponent

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/assets/creativecomponent"
)

// Update 更新组件 https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710673083407
func Update(clt *core.SDKClient, accessToken string, req *creativecomponent.UpdateRequest) (*creativecomponent.PostResponseData, error) {
	var resp creativecomponent.PostResponse

	err := clt.Post("2/assets/creative_component/update/", req, &resp, accessToken)

	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}
