package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// ExtendPackageCreate 创建应用分包
// 通过广告主id和应用包id，为应用包创建对应的分包信息。
// 目前开放平台不支持母包package的创建，创建母包需要前往 投放平台-资产-移动应用 ；
// 母包package创建成功后，可以通过当前接口创建分包。
// 单个package目前支持分包上限3000个；
// 单次创建分包数量上限是100。
func ExtendPackageCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.ExtendPackageCreateRequest) error {
	return clt.Post(ctx, "2/tools/app_management/extend_package/create/", req, nil, accessToken)
}
