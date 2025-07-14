package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// PlayableList 获取试玩/直玩素材
func PlayableList(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.PlayableListRequest) (*file.PlayableListResult, error) {
	var resp file.PlayableListResponse
	if err := clt.GetAPI(ctx, "v3.0/file/playable/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
