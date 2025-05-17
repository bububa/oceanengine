package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// TermsBannedDelete 删除屏蔽词
// 本接口用于删除评论屏蔽词。屏蔽词会针对该广告主下的所有评论生效，命中屏蔽词的抖音评论将直接隐藏，不对外进行展示，删除后后续评论可正常展示。屏蔽词管理模块目前只针对“抖音”广告位生效。屏蔽词数量最多500个。
func TermsBannedDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.TermsBannedDeleteRequest) error {
	return clt.PostAPI(ctx, "v3.0/tools/comment/terms_banned/delete/", req, nil, accessToken)
}
