package site

import (
	"context"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/site"
)

// Create 创建橙子建站站点
// 通过此接口，用户可以创建站点（用于存放落地页），之后才能创建落地页。创建站点接口会返回"code_0"，代表站点创建成功。

// 基于应用下载安全要求，在创建和修改站点接口中，如果希望添加应用下载链接，有以下场景适用：
// ①DownloadEvent类型的按钮组件可以填入安卓/IOS下载链接
// 除以上情况外其他建站组件中的URL参数都不允许填入应用下载链接，如果填入将会创建失败。

// Create 此接口为创建橙子建站站点，如需进行发布请调用【更改橙子建站站点状态】接口将站点及其落地页发布到线上，未发布则保存在本地，无法使用！
// 图片大小需要在2M之内，否则会报错！
// 创建完成后会返回建站id，建站地址可由如下格式拼装得到：https://www.chengzijianzhan.com/tetris/page/XXXXXXXXXXXXX/（其中XX是建站ID，拼装后就可获得在广告投放流程中使用的投放的URL，以及预览URL）！
// 快应用链接仅支持以下三类
// http://hapjs.org/app//[path][?key=value]
// https://hapjs.org/app//[path][?key=value]
// hap://app//[paht][?key=value]
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *site.CreateRequest) (uint64, error) {
	var resp site.CreateResponse
	if err := clt.Post(ctx, "2/tools/site/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return strconv.ParseUint(resp.Data.SiteID, 10, 64)
}
