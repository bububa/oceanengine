package advertiser

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QualificationSelectRequest 获取投放资质信息（新版） API Request
type QualificationSelectRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r QualificationSelectRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QualificationSelectResponse 获取投放资质信息（新版）
type QualificationSelectResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List 资质列表
		List []AdQualification `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
