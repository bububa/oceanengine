package auth

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager/auth"
)

// DelPublicKey 删除公钥
// 广告主删除公钥。
// 鉴权关闭状态下，可直接删除；鉴权开启状态下，平台会检查最近一段时间（24小时）是否有待删除密钥的使用记录，如果没有则可删除，否则删除失败。
func DelPublicKey(ctx context.Context, clt *core.SDKClient, accessToken string, req *auth.DelPublicKeyRequest) error {
	return clt.Post(ctx, "2/event_manager/auth/del_public_key", req, nil, accessToken)
}
