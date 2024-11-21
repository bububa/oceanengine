package task

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	task "github.com/bububa/oceanengine/marketing-api/model/star/star-ad-unite-task"
)

// ItemList 获取星广联投(星图版)视频维度数据
// 此接口可根据星图客户id和星广联投(星图版)任务id，获取某段时间下该任务分天的视频数据
// 其中预估实时消耗、转化量可拉取当天实时数据。
// 累计结算消耗、播放量仅支持T+1拉取，第二天10点后可获取稳定数据
func ItemList(ctx context.Context, clt *core.SDKClient, accessToken string, req *task.ItemListRequest) ([]task.ItemStatInfo, error) {
	var resp task.ItemListResponse
	err := clt.GetAPI(ctx, "2/star/star_ad_unite_task_item/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.StatInfo, nil
}
