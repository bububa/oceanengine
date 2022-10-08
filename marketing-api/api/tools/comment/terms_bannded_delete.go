package comment

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// TermsBanndedDelete 删除屏蔽词
// 本接口用于删除评论屏蔽词。屏蔽词会针对该广告主下的所有评论生效，命中屏蔽词的抖音评论将直接隐藏，不对外进行展示，删除后后续评论可正常展示。屏蔽词管理模块目前只针对“抖音”广告位生效。屏蔽词数量最多500个。
func TermsBanndedDelete(clt *core.SDKClient, accessToken string, req *comment.TermsBanndedDeleteRequest) error {
	return clt.Post("2/tools/comment/terms_bannded/delete/", req, nil, accessToken)
}
