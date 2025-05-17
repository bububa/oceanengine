package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// AndroidBasicPackageGet 查询安卓应用母包
// 查询安卓应用母包信息
func AndroidBasicPackageGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.AndroidBasicPackageGetRequest) (*appmanagement.AndroidBasicPackageGetResult, error) {
	var resp appmanagement.AndroidBasicPackageGetResponse
	if err := clt.Get(ctx, "2/tools/app_management/android_basic_package/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
