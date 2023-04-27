package abtest

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/abtest"
)

// Create 创建实验
// 该接口用于创建AB实验。
func Create(clt *core.SDKClient, accessToken string, req *abtest.CreateRequest) (uint64, error) {
	var resp abtest.CreateResponse
	if err := clt.Post("2/tools/ab_test/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.TestID, nil
}
