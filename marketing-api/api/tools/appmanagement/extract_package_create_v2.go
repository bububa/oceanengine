package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// ExtendPackageCreateV2 创建应用分包 （支持所有账户体系）
// 创建应用分包，支持所有账户体系下创建应用分包
// 母包package创建成功后，可以通过当前接口创建分包
func ExtendPackageCreateV2(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.ExtendPackageCreateV2Request) error {
	return clt.Post(ctx, "2/tools/app_management/extend_package/create_v2/", req, nil, accessToken)
}
