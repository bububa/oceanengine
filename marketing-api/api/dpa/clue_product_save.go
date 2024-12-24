package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// ClueProductSave 创建/编辑升级版商品
func ClueProductSave(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ClueProductSaveRequest) (*dpa.ClueProductSaveResult, error) {
	var resp dpa.ClueProductSaveResponse
	if err := clt.PostAPI(ctx, "2/dpa/clue_product/save/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
