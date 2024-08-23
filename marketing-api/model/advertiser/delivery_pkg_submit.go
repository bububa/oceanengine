package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeliveryPkgSubmitRequest 提交推广产品资质 API Request
type DeliveryPkgSubmitRequest struct {
	// DeliveryPkg 推广产品资质信息
	DeliveryPkg *DeliveryPkg `json:"delivery_pkg,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement PostRequest interface
func (r DeliveryPkgSubmitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeliverPkgSubmitResponse 提交推广产品资质 API Response
type DeliveryPkgSubmitResponse struct {
	Data *DeliveryPkgSubmitResult `json:"data,omitempty"`
	model.BaseResponse
}

type DeliveryPkgSubmitResult struct {
	// QualificationIDs 系统生成的资质id，每份资质对应一个id
	QualificationIDs []uint64 `json:"qualification_ids,omitempty"`
	// PkgID 推广产品组id，可用于后续的查询或编辑
	PkgID uint64 `json:"pkg_id,omitempty"`
}
