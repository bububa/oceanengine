package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// ExtendPackageUpdate 更新应用子包版本
// 通过广告主id和应用包id，更新全部或部分应用子包版本。
func ExtendPackageUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.ExtendPackageUpdateRequest) error {
	return clt.Post(ctx, "2/tools/app_management/extend_package/update/", req, nil, accessToken)
}
