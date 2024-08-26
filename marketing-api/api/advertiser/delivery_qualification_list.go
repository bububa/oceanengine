package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// DeliveryQualificationList 投放资质查询
// 用于查询账户投放资质
func DeliveryQualificationList(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.DeliveryQualificationListRequest) (*advertiser.DeliveryQualificationListData, error) {
	var resp advertiser.DeliveryQualificationListResponse
	if err := clt.GetAPI(ctx, "v3.0/advertiser/delivery_qualification/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
