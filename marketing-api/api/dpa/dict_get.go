package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// DictGet 获取DPA词包
func DictGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.DictGetRequest) ([]dpa.Dict, error) {
	var resp dpa.DictGetResponse
	if err := clt.Get(ctx, "2/dpa/dict/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
