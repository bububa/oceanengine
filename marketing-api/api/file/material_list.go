package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// MaterialList 获取素材标签列表
// 根据 adv 检索素材列表的接口，返回结果包含素材属性，支持素材属性标签筛选，目前仅支持视频素材
func MaterialList(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.MaterialListRequest) (*file.MaterialListData, error) {
	var resp file.MaterialListResponse
	if err := clt.Get(ctx, "2/file/material/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
