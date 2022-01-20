package star

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/star"
)

// ClueList 获取星图订单投后线索
// 此接口用于根据星图id和订单id，获取星图客户的投后线索；
// 若星图客户所发布的抖音传播任务，为线索收集类任务，则星图客户或代理商，可通过此接口，获取达人接单视频中所收集的线索列表。
func ClueList(clt *core.SDKClient, accessToken string, req *star.ClueListRequest) (*star.ClueListResponseData, error) {
	var resp star.ClueListResponse
	err := clt.Get("2/star/clue/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
