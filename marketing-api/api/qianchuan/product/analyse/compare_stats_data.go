package analyse

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/product/analyse"
)

// CompareStatsData 商品竞争分析详情-效果对比
// 获取商品竞争分析详情-效果对比
func CompareStatsData(clt *core.SDKClient, accessToken string, req *analyse.CompareStatsDataRequest) (*analyse.CompareStatsData, error) {
	var resp analyse.CompareStatsDataResponse
	if err := clt.Get("v1.0/qianchuan/product/analyse/compare_stats_data/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
