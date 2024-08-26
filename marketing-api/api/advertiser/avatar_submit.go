package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// AvatarSubmit 更新广告主账户头像
func AvatarSubmit(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.AvatarSubmitRequest) error {
	// var resp advertiser.AvatarSubmitResponse
	return clt.Post(ctx, "2/advertiser/avatar/submit/", req, nil, accessToken)
}
