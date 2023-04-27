package promotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AutoGenerateConfigGetRequest 查询配置详情 API Request
type AutoGenerateConfigGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ConfigID 配置id
	ConfigID uint64 `json:"config_id,omitempty"`
}

// Encode implement GetRequest interface
func (r AutoGenerateConfigGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("config_id", strconv.FormatUint(r.ConfigID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AutoGenerateConfigGetResponse 查询配置详情 API Response
type AutoGenerateConfigGetResponse struct {
	model.BaseResponse
	Data *AutoGenerateConfig `json:"data,omitempty"`
}
