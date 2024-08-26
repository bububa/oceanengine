package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// ExtendPackageListV2 查询应用分包列表 （支持所有账户体系）
// 查询分包列表（支持所有账户体系），可查询该应用包相关信息和包含的分包信息。
// 查询该应用包下所有的应用分包，不受分包新建方式（系统自动、手动上传渠道号分包）影响，都可以获取到
func ExtendPackageListV2(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.ExtendPackageListV2Request) (*appmanagement.ExtendPackageListData, error) {
	var resp appmanagement.ExtendPackageListResponse
	if err := clt.Get(ctx, "2/tools/app_management/extend_package/list_v2/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
