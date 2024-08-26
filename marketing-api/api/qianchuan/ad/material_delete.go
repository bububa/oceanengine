package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// MaterialDelete 删除广告创意
func MaterialDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.MaterialDeleteRequest) error {
	return clt.Post(ctx, "v1.0/qianchuan/ad/material/delete/", req, nil, accessToken)
}
