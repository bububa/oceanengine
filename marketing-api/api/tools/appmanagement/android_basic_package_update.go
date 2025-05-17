package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// AndroidBasicPackageUpdate 更新安卓应用母包
// 更新安卓应用母包对应的信息
// 注意： 单APP_ID接口调用限制为 3QPS
func AndroidBasicPackageUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.AndroidBasicPackageUpdateRequest) error {
	return clt.Post(ctx, "2/tools/app_management/android_basic_package/update/", req, nil, accessToken)
}
