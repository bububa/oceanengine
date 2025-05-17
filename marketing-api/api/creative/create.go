package creative

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/creative"
)

// Create 创建广告创意
// 此接口用于创建广告创意，对于搜索广告的创建可参照【搜索广告投放】
//
// 创意包括了首选投放位置、基础创意、创意分类、监测链接等信息，对于其中的概念解释可以参考：【广告创意】
// 创意创建涉及了多个资产管理：素材管理、落地页管理等，开发者需要提前根据开放接口构建这些资产功能，以免创建计划卡住 不同首选投放位置支持素材类型不同，具体可参考：【广告位与素材类型支持关系】
// 如果创建创意遇到问题，可参考 常见问题
//
// * 每个计划下程序化创意和自定义创意为二选一，且无法修改
// * 程序化创意: 最多10个标题、12个图片素材和10个视频素材;如果创建的是程序化创意（程序化创意实际会按照传入的title_list和image_list进行组合，对于效果不好的组合无法通过审核，获取到的都是审核通过的创意），只有在审核之后才会获取到创意数据与创意id
// * 自定义创意限制计划不能超过10个创意
// * 每日最多创建500个创意（自定义创意+程序化创意）
// * 素材类型：不同广告位要求素材类型不同,其中每一种【素材类型】都有自己的规格，请传入符合要求的素材，否则会报错
// * 其中视频的时长需要>=4s，否则会报错
// * 监测链接：当在计划纬度设置了转化id，如果在创建创意时不传监测链接，会自动获取转化id里监测链接；如果在创建（更新）创意时传入对应的监测链接，会取传入的监测链接，但是对于应用下载推广，即便主动传入点击监测链接，也会取转化id监测链接
// * 对于不打算传的字段，不要传“”或者null，传了会校验
// * 如果计划ID下已有创意信息，需要使用update_v2接口进行修改或者新增创意素材，否则会报错
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.CreateRequest) (*creative.CreateResponseData, error) {
	var resp creative.CreateResponse
	err := clt.Post(ctx, "2/creative/create_v2/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
