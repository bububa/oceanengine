package majordomo

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/majordomo"
)

// 广告主列表（管家）
// 获取客户中心管家下的广告主ID列表。注意：如果发现没有查询到一些广告主，需要您在客户中心确认是否这些广告主是否和此管家解绑关系。
func MajordomoSelect(clt *core.SDKClient, accessToken string, req *majordomo.AdvertiserSelectRequest) (*majordomo.AdvertiserSelectResponse, error) {
	var resp majordomo.AdvertiserSelectResponse
	err := clt.Get("2/majordomo/advertiser/select", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
