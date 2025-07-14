package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// PlayableCreate 上传试玩/直玩素材
func PlayableCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.PlayableCreateRequest) (*file.PlayableMaterial, error) {
	var resp file.PlayableCreateResponse
	if err := clt.PostAPI(ctx, "v3.0/file/playable/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
