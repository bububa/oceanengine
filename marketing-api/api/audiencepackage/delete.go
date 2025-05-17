package audiencepackage

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/audiencepackage"
)

// Delete 删除定向包
func Delete(ctx context.Context, clt *core.SDKClient, accessToken string, req *audiencepackage.DeleteRequest) (uint64, error) {
	var resp audiencepackage.DeleteResponse
	err := clt.Post(ctx, "2/audience_package/delete/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.AudiencePackageID, nil
}
