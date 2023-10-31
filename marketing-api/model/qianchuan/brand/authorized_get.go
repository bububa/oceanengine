package brand

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AuthorizedGetRequest 获取广告主绑定的品牌列表 API Request
type AuthorizedGetRequest struct {
	// AdvertiserID 千川广告主账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r AuthorizedGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AuthorizedGetResponse 获取广告主绑定的品牌列表 API Response
type AuthorizedGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// BrandInfos 授权品牌列表
		BrandInfos []Brand `json:"brand_infos,omitempty"`
	} `json:"data,omitempty"`
}
