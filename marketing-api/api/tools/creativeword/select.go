package creativeword

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/creativeword"
)

// Select 查询动态创意词包
// 创建动态创意词包：创建的词包你可以使用在创意标题中做固定词位的替换，例如创建地理位置的词包，针对用户的位置属性显示你词包中相应的地理位置词。如果相应查看可使用的词包请参考查询创意词包接口。
func Select(ctx context.Context, clt *core.SDKClient, accessToken string, req *creativeword.SelectRequest) ([]creativeword.CreativeWord, error) {
	var resp creativeword.SelectResponse
	err := clt.Get(ctx, "2/tools/creative_word/select/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
