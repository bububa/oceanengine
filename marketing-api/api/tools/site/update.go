package site

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/site"
)

// Update 修改橙子建站站点
// 通过此接口，用户可以修改站点的基本信息。

// 基于应用下载安全要求，在创建和修改站点接口中，如果希望添加应用下载链接，有以下场景适用：
// ①DownloadEvent类型的按钮组件可以填入安卓/IOS下载链接
// 除以上情况外其他建站组件中的URL参数都不允许填入应用下载链接，如果填入将会创建失败。

// 目前bricks不支持部分更新，仅支持全量更新，更新bricks时，需要传递完整bricks信息，除需要更新的bricks信息以外，其余bricks中的信息需要和创建时保持不变。修改成功时，无业务参数返回！
// 图片大小需要在2M之内，否则会报错！
// 快应用链接仅支持以下三类
// http://hapjs.org/app//[path][?key=value]
// https://hapjs.org/app//[path][?key=value]
// hap://app//[paht][?key=value]
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *site.UpdateRequest) error {
	return clt.Post(ctx, "2/tools/site/update/", req, nil, accessToken)
}
