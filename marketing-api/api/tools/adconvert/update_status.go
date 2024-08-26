package adconvert

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adconvert"
)

// UpdateStatus 更新转化目标操作状态
// 此接口用于针对转化目标的操作状态进行变更，将无用的转化目标删除放入回收站或将已删除至回收站的转化目标恢复
// 更改转化目标操作状态为删除后，正在使用该转化目标的广告计划不受影响；
// 转化目标删除后，对应已经在使用中的转化数据不受影响；
func UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *adconvert.UpdateStatusRequest) error {
	return clt.Post(ctx, "2/tools/ad_convert/update_status/", req, nil, accessToken)
}
