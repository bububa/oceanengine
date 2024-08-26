package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// AndroidBasicPackagePublish 发布安卓应用母包
// 发布安卓应用母包「支持所有账户体系」，发布审核通过的安卓应用母包。
// 注意： 单APP_ID接口调用限制为 3QPS
func AndroidBasicPackagePublish(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.AndroidBasicPackagePublishRequest) error {
	return clt.Post(ctx, "2/tools/app_management/android_basic_package/publish/", req, nil, accessToken)
}
