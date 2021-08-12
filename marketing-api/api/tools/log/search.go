package log

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/log"
)

// Search 日志查询
// 可查询ad后台操作日志，默认查询最近7天的数据，最多查询跨度为一个月。
func Search(clt *core.SDKClient, accessToken string, req *log.SearchRequest) (*log.SearchResponseData, error) {
	var resp log.SearchResponse
	err := clt.Get("2/tools/log_search/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
