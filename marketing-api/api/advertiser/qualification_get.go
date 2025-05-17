package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// QualificationGet 获取广告主资质（新版）
// 获取广告主资质信息为全量接口，会返回广告主所有资质，注意如果广告主没有任何资质，这个接口的data将会是空。
func QualificationGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.Qualification, error) {
	req := &advertiser.QualificationGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.QualificationGetResponse
	if err := clt.GetAPI(ctx, "v3.0/advertiser/qualification/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
