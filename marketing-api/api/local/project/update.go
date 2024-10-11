package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/project"
)

// Update 更新项目
// 接口的更新逻辑为增量更新
// 定向参数如需设置对应受众为不限：string类型受众设置无限统一传入不限枚举，list类型的受众设置无限统一传入空数组，即[]
// 对于不打算更新的字段，不要传“”或者null，传了会校验
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.UpdateRequest) error {
	return clt.PostAPI(ctx, "v3.0/local/project/update/", req, nil, accessToken)
}
