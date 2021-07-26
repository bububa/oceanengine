package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// PublicInfo 广告主公开信息
func PublicInfo(clt *core.SDKClient, accessToken string, req *advertiser.PublicInfoRequest) (*advertiser.PublicInfoResponse, error) {
	var resp advertiser.PublicInfoResponse
	err := clt.Get("2/advertiser/public_info/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
