package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/tools"
)

// AwemeAuth 广告主添加抖音号
// 支持抖音号授权，即将抖音号授权给广告主做广告投放
// 线下联系合作达人
// 获取达人合作码
// 添加可授权抖音号
// 注意：如果广告主未在千川PC添加过抖音号，需要先在PC完成过一次添加抖音号操作（签署《巨量千川商业合作授权协议》），否则调用接口会报错。
func AwemeAuth(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.AwemeAuthRequest) (bool, error) {
	var resp tools.AwemeAuthResponse
	if err := clt.PostAPI(ctx, "v1.0/qianchuan/tools/aweme_auth/", req, &resp, accessToken); err != nil {
		return false, err
	}
	return resp.Data.AuthSuccess, nil
}
