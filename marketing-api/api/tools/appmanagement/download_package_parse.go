package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// DownloadPackageParse 提交解析应用包任务
// 当您需要创建一个附加创意类型为游戏礼包码的广告计划时，如果该游戏礼包的download_url为未在字节的广告系统中上传或使用过的新download_url，则需要使用该接口提交包解析任务。
// 使用提交解析应用包任务时，您只需要传入参数advertiser_id和该游戏礼包的download_url，即可解析出APP相关信息（包名、appname、icon等），这些信息将在礼包码原生转化卡中展示。
// 提交解析应用包任务后，您可以使用【查询包解析状态】接口对当前包解析状态进行查询。查询到返回参数package_status为"SUCCESS"后，方可提交创建广告计划。
func DownloadPackageParse(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.DownloadPackageParseRequest) (string, error) {
	var resp appmanagement.DownloadPackageParseResponse
	if err := clt.Post(ctx, "2/tools/app_management/download/package/parse/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.EventID, nil
}
