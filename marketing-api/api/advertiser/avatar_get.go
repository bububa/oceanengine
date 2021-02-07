package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// 获取广告主头像信息
func Avatar(clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.AvatarGetResponse, error) {
	req := &advertiser.AvatarGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.AvatarGetResponse
	err := clt.Get("2/advertiser/avatar/get", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
