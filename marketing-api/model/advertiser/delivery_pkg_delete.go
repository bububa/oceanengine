package advertiser

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeliveryPkgDeleteRequest 批量删除推广产品资质 API Request
type DeliveryPkgDeleteRequest struct {
	// PkgIDs 推广产品pkg_id列表，支持传入多个批量删除，list长度1~50
	PkgIDs []uint64 `json:"pkg_ids,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest interface
func (r DeliveryPkgDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

type DeliveryPkgDeleteResponse struct {
	Data *DeliveryPkgDeleteResult `json:"data,omitempty"`
	model.BaseResponse
}

type DeliveryPkgDeleteResult struct {
	// PkgIDs 删除成功的推广产品id
	PkgIDs []uint64 `json:"pkg_ids,omitempty"`
	// Errors 删除失败的行业产品列表
}

// DeliveryPkgDeleteError 删除失败的行业产品
type DeliveryPkgDeleteError struct {
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message,omitempty"`
	// PkgID 行业产品id
	PkgID uint64 `json:"pkg_id,omitempty"`
}

// Error implements error interface
func (e DeliveryPkgDeleteError) Error() string {
	return util.StringsJoin(e.ErrorMessage, "(", strconv.FormatUint(e.PkgID, 10), ")")
}
