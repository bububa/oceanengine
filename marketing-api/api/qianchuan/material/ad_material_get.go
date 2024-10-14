package material

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/material"
)

// AdMaterialGet 获取计划下素材列表
// 获取计划下的素材信息，包括素材的元信息，审核状态，关联创意ID，派生信息，是否删除等。
func AdMaterialGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *material.AdMaterialGetRequest) (*material.AdMaterialGetResult, error) {
	var resp material.AdMaterialGetResponse
	if err := clt.GetAPI(ctx, "v1.0/qianchuan/ad/material/get/", req, resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
