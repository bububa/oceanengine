package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adraise/v3"
)

// VersionGet “获取起量版本信息”用于获取计划在多次起量过程中产生的起量版本号及对应的起止时间。
func VersionGet(clt *core.SDKClient, accessToken string, req *v3.VersionGetRequest) (*v3.VersionGetResponseData, error) {
	var resp v3.VersionGetResponse
	if err := clt.Get("v3.0/tools/promotion_raise_version/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
