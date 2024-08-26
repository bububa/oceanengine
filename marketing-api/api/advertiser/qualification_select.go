package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// QualificationSelect 获取投放资质信息（新版）
// 获取广告主资质信息，资质分为描述，营业执照，开户资质，投放资质。不同类型的资质有不同的字段，具体字段见下表。
func QualificationSelect(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) ([]advertiser.AdQualification, error) {
	req := &advertiser.QualificationSelectRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.QualificationSelectResponse
	if err := clt.Get(ctx, "2/advertiser/qualification/select_v2/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
