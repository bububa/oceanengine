package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// QualificationSubmit 提交广告主资质（新版）
// 提交广告主资质信息为全量接口。更新的时候需要全量先获取所有资质，然后更新相应资质。
func QualificationSubmit(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.QualificationSubmitRequest) error {
	return clt.PostAPI(ctx, "v3.0/advertiser/qualification/submit/", req, nil, accessToken)
}
