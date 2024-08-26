package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise/comment"
)

// ListGet 获取评论列表
// 根据企业号账户id，查询企业号下的评论列表，列表内容包含评论id、评论内容、评论用户open_id、评论所属视频、所属视频类型、评论时间、评论层级、流量来源(自然流量、dou+、竞价、品牌）等。
//
// 不支持查询评论用户头像、评论用户抖音号id；
// 评论用户昵称已模糊化处理；
// 用户open_id，用来识别用户唯一性和调用回复等操作接口，每个抖音用户对每个开发者的open_id是唯一的，对于不同开发者，同一用户的open_id不同。
func ListGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ListGetRequest) (*comment.ListGetResult, error) {
	var resp comment.ListGetResponse
	if err := clt.Get(ctx, "v1.0/enterprise/comment/list/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
