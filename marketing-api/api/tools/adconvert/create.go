package adconvert

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adconvert"
)

// Create 创建转化目标
// 通过工具能力，创建一个能被计划使用的转化跟踪工具。可以通过设定转化目标（比如多少人打开应用）跟踪转化状态，衡量推广转化效果。按照转化跟踪类型目的区分目前支持两种：跟踪线索以及跟踪应用。
// 如果广告主期望获得销售线索（表单提交、销售线索等）可以使用追踪线索相关，线索追踪中支持JS、Xpath、API（线索）
// 如果拥有自己的应用，希望获取用户指定行为（下载、安装、激活等）可以使用应用推广相关，应用追踪支持应用API、应用SDK
// 您可以通过转化ID列表或转化详细信息接口的status参数来查询转化状态
// 创建后状态为未激活状态，需要前往巨量引擎投放平台激活，再去创建广告计划 转化类型和激活规则对应关系如下：
// 线索-XPATH：创建后状态默认为激活；
// 线索-JS：创建后状态为未激活，需要下载代码在自有落地页添加，完成检测转化上报；
// 线索-API：创建后为状态为未激活，需要完成检测转化上报（过程中需技术人员接收广告点击数据并将相关转化事件回传至回调地址）
// 应用下载SDK：首次创建成功后状态为未激活，需要检测转化上报，完成后激活；
// 应用下载API：创建成功后状态为未激活，需要检测转化上报，完成后激活；
// 其他非应用下载类的转化每次创建都要求激活联调。但是如果你使用转化推送接口，推送已激活的转化，那推送后的转化目标是免联调直接进入激活状态（请注意状态变更有几分钟的延迟）
// 「convert_data_type」转化统计方式中EVERY_ONE（每一次）为白名单功能，可联系对接销售或运营咨询
// 创建应用下载类转化目标时，Android类应用必须要填写package_name的应用包名，并且要在平台侧的应用管理中完成上传解析才可以使用。 平台侧操作：投放平台-「资产」-「移动应用」-「新建应用」-「上传」
// 使用时注意区分应用中文名app_name与包名package_name app_name：应用名称，例如：“计算器”、“记事本” package_name：应用包包名，例如“com.xx.xx”
// 当前应用下载类转化目标中，仅AD_CONVERT_TYPE_ACTIVE激活、AD_CONVERT_TYPE_PAY付费，支持设置对应深度转化目标
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *adconvert.CreateRequest) (*adconvert.Convert, error) {
	var resp adconvert.CreateResponse
	if err := clt.Post(ctx, "2/tools/ad_convert/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
