package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// AlbumStatusGet 查询短剧可投状态
// 通过短剧alblumID查询短剧可投状态
// 特别注意：请求本接口时需要使用APP Access Token，通过【获取APP Access Token】接口获取
func AlbumStatusGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.AlbumStatusGetRequest) (*dpa.AlbumStatusGetResult, error) {
	var resp dpa.AlbumStatusGetResponse
	if err := clt.GetAPI(ctx, "v3.0/dpa/album_status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
