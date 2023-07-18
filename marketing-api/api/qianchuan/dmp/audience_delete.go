package dmp

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/dmp"
)

// AudienceDelete 删除人群
// 用来实现删除人群功能
func AudienceDelete(clt *core.SDKClient, accessToken string, req *dmp.AudienceDeleteRequest) (uint64, error) {
	var resp dmp.AudienceDeleteResponse
	if err := clt.Post("v1.0/qianchuan/audience/delete/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.AudienceID, nil
}
