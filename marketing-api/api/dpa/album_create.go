package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// AlbumCreate 上传短剧剧目
// 特别注意：
// 创建短剧为异步动作，提交创建任务→上传短剧完成耗时较长，创建完后返回的album_id未必处于正常状态，建议您定时轮询「查询短剧可投状态」，获取短剧创建结果和短剧是否可投
// 请求本接口时需要使用APP Access Token，通过【获取APP Access Token】接口获取，同时传入token对应的app_id请求
func AlbumCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.AlbumCreateRequest) (string, error) {
	var resp dpa.AlbumCreateResponse
	if err := clt.PostAPI(ctx, "v3.0/dpa/album/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.AlbumID, nil
}
