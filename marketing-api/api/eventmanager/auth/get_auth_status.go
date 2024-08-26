package auth

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager/auth"
)

// GetAuthStatus 查询鉴权开启状态
// 广告主查询鉴权开启状态。
func GetAuthStatus(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (bool, error) {
	req := auth.GetAuthStatusRequest{
		AdvertiserID: advertiserID,
	}
	var resp auth.GetAuthStatusResponse
	if err := clt.Get(ctx, "2/event_manager/auth/get_auth_status", req, &resp, accessToken); err != nil {
		return false, err
	}
	return resp.Data.Enabled, nil
}
