package adconvert

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adconvert"
)

// Read 查询转化目标详细信息
// 查询转化ID的详细信息
// 整体返回信息依赖于转化来源convert_source_type，不同转化来源返回的详情参数有差异
// 如果是需要广告计划可用的转化，请使用接口【查询可用转化ID】
func Read(ctx context.Context, clt *core.SDKClient, accessToken string, req *adconvert.ReadRequest) (*adconvert.Convert, error) {
	var resp adconvert.ReadResponse
	err := clt.Get(ctx, "2/tools/ad_convert/read/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
