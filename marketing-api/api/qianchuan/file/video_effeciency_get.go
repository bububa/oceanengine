package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/file"
)

// VideoEffeciencyGet 获取低效素材
// 支持查询素材是否是低效素材，传入素材ID列表，返回低效素材列表。
// 1.什么是低效素材？
// 在一定历史投放周期内，素材累积消耗低于系统标准，即定义为低效素材。
// 目前低效素材仅针对视频素材进行识别。
// 2.低效素材有什么影响？
// 根据整体数据分析，无论是复用低效素材新建计划或将低效素材稍加修改后重新使用，低效素材起量的可能性较小，进而影响计划的跑量能力；
// 如计划下出现大量低效素材，则该计划预计跑量困难，占用在投计划配额、投放人力和投放预算，建议更换素材。
// 3. 如依然使用低效素材会产生哪些问题？
// 如果新建创意下素材全部为低效素材，那么创意将无法创建成功，建议更换其他素材。创建广告创意相关接口的拦截功能将于2023年1月初上线。请务必提前关注【获取低效素材】接口，进行低效素材识别和管理。
func VideoEffeciencyGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoEffeciencyGetRequest) ([]string, error) {
	var resp file.VideoEffeciencyGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/file/video/effeciency/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.InEffecientMaterialIDs, nil
}
