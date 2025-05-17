package auth

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager/auth"
)

// AddPublicKey 新增公钥
// 广告主将公钥保存到平台。
// 为支持切换密钥的场景，广告主最多可使用两对公/私钥对，两对的label约定为“primary”、“backup”。
// 为确保广告主使用了正确的公钥格式和签名算法，在新增公钥时有一步验签的动作，广告主使用私钥对指定文本“BYTEDANCE”签名。
func AddPublicKey(ctx context.Context, clt *core.SDKClient, accessToken string, req *auth.AddPublicKeyRequest) (*auth.PublicKey, error) {
	var resp auth.AddPublicKeyResponse
	if err := clt.Post(ctx, "2/event_manager/auth/add_public_key", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.PubKey, nil
}
