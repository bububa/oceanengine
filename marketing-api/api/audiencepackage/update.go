package audiencepackage

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/audiencepackage"
)

// Update 更新定向包
func Update(clt *core.SDKClient, accessToken string, req *audiencepackage.UpdateRequest) (uint64, error) {
	var resp audiencepackage.UpdateResponse
	err := clt.Post("2/audience_package/update/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.AudiencePackageID, nil
}
