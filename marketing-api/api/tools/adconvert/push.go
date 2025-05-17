package adconvert

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adconvert"
)

// Push 转化目标推送
// 通过当前接口，可以将转化目标ID推送给同公司主体下的其他账号，帮助您解决需要频繁创建转化目标ID的问题。
// 状态是已激活的转化目标ID才可进行推送，并且推送后的转化目标ID将自动变成已激活状态，帮助您免去繁琐的联调，但是推送后会有一段时间延迟，新的转化目标ID会在几分钟延迟后状态更新为已激活 推送后您可以通过 【查询转化详细信息】接口查询转化的状态
// 推送的广告主必须与转化目标ID所属广告主的公司主体一致；
// 一次可推送的广告主target_advertiser_ids上限为50
func Push(ctx context.Context, clt *core.SDKClient, accessToken string, req *adconvert.PushRequest) error {
	return clt.Post(ctx, "2/tools/ad_convert/push/", req, nil, accessToken)
}
