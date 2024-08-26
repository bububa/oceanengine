package dmp

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/dmp"
)

// OrientationPackageGet 获取定向包列表
// 用于获取千川广告账户下的定向包列表
func OrientationPackageGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.OrientationPackageGetRequest) (*dmp.OrientationPackageGetResult, error) {
	var resp dmp.OrientationPackageGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/orientation_package/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
