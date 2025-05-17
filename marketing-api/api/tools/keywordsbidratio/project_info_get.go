package keywordsbidratio

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/keywordsbidratio"
)

// ProjectInfoGet 查询优词绑定的项目信息
func ProjectInfoGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *keywordsbidratio.ProjectInfoGetRequest) (*keywordsbidratio.ProjectInfo, error) {
	var resp keywordsbidratio.ProjectInfoGetResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/keywords_project_info/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ProjectInfo, nil
}
