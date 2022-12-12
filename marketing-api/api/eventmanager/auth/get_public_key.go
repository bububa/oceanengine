package auth

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager/auth"
)

// GetPublicKey 查询公钥
// 广告主查询公钥。
func GetPublicKey(clt *core.SDKClient, accessToken string, req *auth.GetPublicKeyRequest) (*auth.PublicKey, error) {
	var resp auth.GetPublicKeyResponse
	if err := clt.Get("2/event_manager/auth/get_public_key", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.PubKey, nil
}
