package comment

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// ReplyGet 获取评论回复列表
// 获取评论回复（用于获取二级评论）
func ReplyGet(clt *core.SDKClient, accessToken string, req *comment.ReplyGetRequest) (*comment.ReplyGetResponseData, error) {
	var resp comment.ReplyGetResponse
	if err := clt.Get("v3.0/tools/comment_reply/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
