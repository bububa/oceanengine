package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// DownloadPackageGet 查询包解析状态
// 查询包解析状态，用于创建包含游戏礼包码的广告计划时，【提交解析应用包任务】后，查询应用包的解析状态。 当返回参数package_status包解析状态为“SUCCESS”时表示包解析已经完成。此时，您可以提交创建广告计划。
// 如果在没有解析完成时提交广告计划可能导致游戏礼包码的原生转化卡信息展示不全。
func DownloadPackageGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.DownloadPackageGetRequest) (appmanagement.DownloadPackageStatus, error) {
	var resp appmanagement.DownloadPackageGetResponse
	if err := clt.Get(ctx, "2/tools/app_management/download/package/get/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.PackageStatus, nil
}
