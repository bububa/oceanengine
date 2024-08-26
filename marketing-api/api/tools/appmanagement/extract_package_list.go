package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// ExtendPackageList 查询应用分包列表
// 通过应用包id，查询该应用包相关信息和包含的分包信息。
// 信息包括应用包id，分包的列表，状态,渠道号，下载链接，分包失败原因等；
// 查询该应用包下所有的应用分包，不受分包新建方式（系统自动、手动上传渠道号分包）影响，都可以获取到。
// 查询应用分包时，可以通过状态进行过滤，状态status枚举值（全部、未更新、创建中、更新中、已发布、更新失败、创建失败），默认返回全部状态的分包列表
func ExtendPackageList(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.ExtendPackageListRequest) (*appmanagement.ExtendPackageListData, error) {
	var resp appmanagement.ExtendPackageListResponse
	if err := clt.Get(ctx, "2/tools/app_management/extend_package/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
