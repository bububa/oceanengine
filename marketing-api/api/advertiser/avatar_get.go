package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// AvatarGet 获取广告主头像信息
func AvatarGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.AvatarGetResponseData, error) {
	req := &advertiser.AvatarGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.AvatarGetResponse
	if err := clt.Get(ctx, "2/advertiser/avatar/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
