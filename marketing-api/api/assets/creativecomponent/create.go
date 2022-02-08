package creativecomponent

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/assets/creativecomponent"
)

// Post 创建组件 https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710672391183
func Post(clt *core.SDKClient, accessToken string, req *creativecomponent.PostRequest) (*creativecomponent.PostResponseData, error) {
	var resp creativecomponent.PostResponse

	err := clt.Post("2/assets/creative_component/create/", req, &resp, accessToken)

	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}
