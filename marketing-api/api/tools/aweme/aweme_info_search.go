package aweme

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/aweme"
)

// AwemeInfoSearch 查询抖音帐号和类目信息
func AwemeInfoSearch(clt *core.SDKClient, accessToken string, req *aweme.AwemeInfoSearchRequest) ([]aweme.AwemeInfoSearchResult, error) {
	var resp aweme.AwemeInfoSearchResponse
	err := clt.Get("2/tools/aweme_info_search/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
