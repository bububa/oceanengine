package comment

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// Get 获取评论列表
func Get(clt *core.SDKClient, accessToken string, req *comment.GetRequest) (*comment.GetResponseData, error) {
	var resp comment.GetResponse
	err := clt.Get("2/tools/comment/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
