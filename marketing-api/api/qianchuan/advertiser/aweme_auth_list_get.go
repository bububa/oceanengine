package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/advertiser"
)

// AwemeAuthListGet 获取千川账户下抖音号授权列表
// 获取千川账户下抖音号授权列表，对齐千川PC抖音号授权管理模块
func AwemeAuthListGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.AwemeAuthListGetRequest) (*advertiser.AwemeAuthListGetResult, error) {
	var resp advertiser.AwemeAuthListGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/aweme_auth_list/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
