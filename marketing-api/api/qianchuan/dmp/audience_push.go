package dmp

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/dmp"
)

// AudiencePush 推送人群
// 用来实现人群推送功能
func AudiencePush(clt *core.SDKClient, accessToken string, req *dmp.AudiencePushRequest) (uint64, error) {
	var resp dmp.AudiencePushResponse
	if err := clt.Post("v1.0/qianchuan/audience/push/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.AudienceID, nil
}
