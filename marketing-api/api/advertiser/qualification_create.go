package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// QualificationCreate 批量上传投放资质
func QualificationCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.QualificationCreateRequest) error {
	return clt.Post(ctx, "2/advertiser/qualification/create_v2/", req, nil, accessToken)
}
