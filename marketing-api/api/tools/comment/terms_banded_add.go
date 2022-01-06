package comment

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// TermsBandedAdd 添加屏蔽词
// 本接口用于添加评论屏蔽词。该屏蔽词针对该广告主下的所有评论生效，命中屏蔽词的抖音评论将直接隐藏，不对外进行展示。屏蔽词管理模块目前只针对“抖音”广告位生效。屏蔽词数量最多500个。
func TermsBandedAdd(clt *core.SDKClient, accessToken string, req *comment.TermsBandedAddRequest) error {
	var resp model.BaseResponse
	return clt.Post("2/tools/comment/terms_banded/add/", req, &resp, accessToken)
}
