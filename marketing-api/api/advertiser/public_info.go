package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// PublicInfo 广告主公开信息
func PublicInfo(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.PublicInfoRequest) ([]advertiser.PublicInfo, error) {
	var resp advertiser.PublicInfoResponse
	err := clt.Get(ctx, "2/advertiser/public_info/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
