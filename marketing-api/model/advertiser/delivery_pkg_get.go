package advertiser

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeliveryPkgGetRequest 查询推广产品资质 API Request
type DeliveryPkgGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PkgID 推广产品组id，是推广产品的组标识
	PkgID uint64 `json:"pkg_id,omitempty"`
}

// Encode implement GetRequest interface
func (r DeliveryPkgGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("pkg_id", strconv.FormatUint(r.PkgID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DeliveryPkgGetResponse 查询推广产品资质 API Response
type DeliveryPkgGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// DeliveryPkg 推广产品资质信息
		DeliveryPkg *DeliveryPkg `json:"delivery_pkg,omitempty"`
	} `json:"data,omitempty"`
}
