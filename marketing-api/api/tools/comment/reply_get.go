package comment

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// ReplyGet 获取评论回复列表
// 说明：评论回复会经过审核，一般有10s的审核时间，如果审核未通过的评论将无法展现和查询出来，并且无拒审理由，意味着10S后回复列表没有查出来那回复内容大概率就是被审核拒绝了。
func ReplyGet(clt *core.SDKClient, accessToken string, req *comment.ReplyGetRequest) (*comment.ReplyGetResponseData, error) {
	var resp comment.ReplyGetResponse
	err := clt.Get("2/tools/comment_reply/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
