package creative

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/creative"
)

// ProceduralCreativeUpdate 修改程序化创意（营销链路）
// 此接口用于全量更新广告程序化创意： 当计划下有创意时，需要通过此接口进行全量更新与创建创意（可以先通过创意详细信息接口获取全部信息在进行更新）
//
// 全量更新解释
//
// 全量更新即为最新一次更新将覆盖前一次的所有信息，例如：如果已有三个创意A、B、C，更新提交了创意B、C、D，则意味着覆盖删除创意A，新增了D，保留（或修改）了B、C
// 对于新建的创意，不需要上传创意ID，会生成新的创意ID；对于您希望保留的创意，不要忘记上传创意ID，否则将被覆盖
// 新版营销链路广告创意差异点
//
// 新版接口新增落地页，小程序等相关入参
// 下线推广卡片，附加创意等相关字段，替换为创意组件形式传入
// 投放位置与监测链接变更为计划接口传入
// 触发审核场景
//
// 在修改用户端可见的内容包括标题、图片/视频、来源、附加创意、落地页链接等时会触发审核；
// 对于广告位的修改有以下情况也会触发审核：①选择了网盟广告位然后增加头条广告位；②选择了头条广告位然后增加抖音广告位；这两种情况都会触发审核。其他情况比如选了头条广告位再增加火山、西瓜不会触发审核。
// 每个计划下程序化创意和自定义创意为二选一，且无法修改
// 每日最多创建500个创意（自定义创意+程序化创意）
// 其中视频的时长需要>=4s，否则会报错
// 对于不打算传的字段，不要传“”或者null，传了会校验
// 注意字段默认值，如果更新时候未传入带有默认值的字段，将以默认值为准
func ProceduralCreativeUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.ProceduralCreativeUpdateRequest) error {
	return clt.Post(ctx, "2/creative/procedural_creative/update/", req, nil, accessToken)
}
