package keywordsbidratio

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/keywordsbidratio"
)

// Create 设置优词提量系数和生效维度
func Create(clt *core.SDKClient, accessToken string, req *keywordsbidratio.CreateRequest) error {
	return clt.PostAPI("v3.0/tools/keywords_bid_ratio/create/", req, nil, accessToken)
}
