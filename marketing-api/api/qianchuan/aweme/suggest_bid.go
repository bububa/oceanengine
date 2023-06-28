package aweme

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// SuggestBid 获取随心推短视频建议出价
// 小店获取短视频建议出价
func SuggestBid(clt *core.SDKClient, accessToken string, req *aweme.SuggestBidRequest) (float64, error) {
	var resp aweme.SuggestBidResponse
	if err := clt.Get("v1.0/qianchuan/aweme/suggest_bid/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.SuggestBid, nil
}
