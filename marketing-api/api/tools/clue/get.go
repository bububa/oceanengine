package clue

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/clue"
)

// Get 获取线索列表
// 该接口用于获取广告主的飞鱼线索列表。
// 为保证接口使用的安全性避免调取他人的线索信息，该接口只可用于查询自己广告主下的线索信息，即需查询的广告主账号的主体需与APPID对应开发者的主体保持一致，才可获取到线索的信息，否则会报错！广告者主体与开发者主体不一致，又要使用该接口时，解决方式请参考【Qauth2.0授权-其他】
// 其中应答参数是按creative_time倒序返回的。如果您有数据处理的需求，可以参考此排序逻辑。
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *clue.GetRequest) (*clue.GetResponseData, error) {
	var resp clue.GetResponse
	if err := clt.Get(ctx, "2/tools/clue/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
