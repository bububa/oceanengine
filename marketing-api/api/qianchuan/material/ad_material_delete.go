package material

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/material"
)

// AdMaterialDelete 删除广告创意
func AdMaterialDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *material.AdMaterialDeleteRequest) error {
	return clt.Post(ctx, "v1.0/qianchuan/ad/material/delete/", req, nil, accessToken)
}
