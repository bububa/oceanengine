package audiencepackage

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/audiencepackage"
)

// AdBind 计划绑定定向包
func AdBind(clt *core.SDKClient, accessToken string, req *audiencepackage.AdBindRequest) (uint64, error) {
	var resp audiencepackage.AdBindResponse
	err := clt.Post("2/audience_package/ad/bind/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.AudiencePackageID, nil
}
