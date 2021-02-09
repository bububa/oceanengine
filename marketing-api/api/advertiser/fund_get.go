package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

func Fund(clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.FundGetResponse, error) {
	req := &advertiser.FundGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.FundGetResponse
	err := clt.Get("2/advertiser/fund/get", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
