package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// TermsBannedUpdate 更新屏蔽词
// 本接口用于对单一评论屏蔽词进行更新，可以选择指定的屏蔽词进行更新，效果等同于删除老屏蔽词后再添加新屏蔽词。该屏蔽词针对该广告主下的所有评论生效，命中屏蔽词的抖音评论将直接隐藏，不对外进行展示。屏蔽词管理模块目前只针对“抖音”广告位生效。屏蔽词数量最多500个。
// 如果new_terms已存在，则等同于删除origin_terms；如果origin_terms不存在，则等同于新增new_terms
func TermsBannedUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.TermsBannedUpdateRequest) error {
	return clt.PostAPI(ctx, "v3.0/tools/comment/terms_banned/update/", req, nil, accessToken)
}
