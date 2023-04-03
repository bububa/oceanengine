package abtest

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/abtest"
)

// Update 更改实验
func Update(clt *core.SDKClient, accessToken string, req *abtest.UpdateRequest) (uint64, error) {
	var resp abtest.UpdateResponse
	if err := clt.Post("2/tools/ab_test/update/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.TestID, nil
}
