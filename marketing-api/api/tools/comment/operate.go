package comment

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// Operate 操作广告计划的评论
// 为保证接口使用的安全性，避免操作他人的评论信息，该接口只可用于操作广告主自身计划下的评论信息，即需查询的广告主账号的主体需与APPID对应开发者的主体保持一致，才可操作评论的信息，否则会报错！ 说明：评论回复会经过审核，一般有10s的审核时间，如果审核未通过的评论将无法展现和查询出来，并且无拒审理由，意味着10S后回复列表没有查出来那回复内容大概率就是被审核拒绝了。
func Operate(clt *core.SDKClient, accessToken string, req *comment.OperateRequest) (*comment.OperateResponseData, error) {
	var resp comment.OperateResponse
	err := clt.Post("2/tools/comment/operate/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
