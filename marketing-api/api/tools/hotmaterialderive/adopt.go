package hotmaterialderive

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/hotmaterialderive"
)

// Adopt 采纳爆款裂变视频
func Adopt(ctx context.Context, clt *core.SDKClient, accessToken string, req *hotmaterialderive.AdoptRequest) ([]hotmaterialderive.AdoptResult, error) {
	var resp hotmaterialderive.AdoptResponse
	if err := clt.PostAPI(ctx, "v3.0/tools/hot_material_derive/adopt/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Results, nil
}
