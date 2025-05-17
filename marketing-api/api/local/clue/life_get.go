package clue

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/clue"
)

// LifeGet 获取本地推线索列表
// 该接口用于获取广告主的本地推线索列表。
// 线索列表中，不包含来客账户下自然流量线索。
// 为保证接口使用的安全性避免调取他人的线索信息：
// 该接口仅允许巨量引擎工作台（组织）账户授权后调用，不支持代理商账户授权。
// 该接口只可用于查询自己本地推账户下的线索信息，即需查询的本地推账号的主体需与APPID对应开发者的主体保持一致，才可获取到线索的信息，否则会报错！本地推主体与开发者主体不一致，又要使用该接口时，请在授权时勾选「敏感物料授权」
// 您要获取的线索信息中包含大量用户个人信息。除您另行获得用户的同意外，您仅可将相关的用户个人信息用于用户授权范围内的必要用途；涉嫌向他人非法提供、售卖用户个人信息的，可能构成刑事犯罪，须依法承担相应法律责任。
// 其中应答参数是按creative_time倒序返回的。如果您有数据处理的需求，可以参考此排序逻辑。
func LifeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *clue.LifeGetRequest) (*clue.LifeGetResult, error) {
	var resp clue.LifeGetResponse
	if err := clt.GetAPI(ctx, "2/tools/clue/life/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
