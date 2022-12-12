package auth

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager/auth"
)

// GetAllPublicKey 查询全部公钥
// 广告主查询公钥。
func GetAllPublicKeys(clt *core.SDKClient, accessToken string, req *auth.GetAllPublicKeyRequest) ([]auth.PublicKey, error) {
	var resp auth.GetAllPublicKeyResponse
	if err := clt.Get("2/event_manager/auth/get_all_public_keys", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.PubKeys, nil
}
