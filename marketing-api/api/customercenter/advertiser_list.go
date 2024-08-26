package customercenter

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/customercenter"
)

// AdvertiserList 获取纵横组织下资产账户列表（分页）
// 获取当前纵横组织下的资产账户列表（广告主，企业号），如果纵横组织下还有子级组织账号，会把cc_account_id下所有可操作的资产账户全量返回
// Token支持不同角色用户授权生成，某些无全量权限的角色（ 纵横组织(协作者) ）仅能看到当前纵横组织id下有权限操作的账户列表。操作权限受授权账号影响，所以可能会出现相同的纵横组织帐号在不同access_token请求下获取的列表不同的情况。
func AdvertiserList(ctx context.Context, clt *core.SDKClient, accessToken string, req *customercenter.AdvertiserListRequest) (*customercenter.AdvertiserListData, error) {
	var resp customercenter.AdvertiserListResponse
	if err := clt.Get(ctx, "2/customer_center/advertiser/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
