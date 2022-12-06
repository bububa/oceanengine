package form

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DetailRequest 获取表单详情 API Request
type DetailRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InstanceID 表单实例ID，可从【表单创建】【表单列表获取】接口获取到
	InstanceID uint64 `json:"instance_id,omitempty"`
}

// Encode implement GetRequest interface
func (r DetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("instance_id", strconv.FormatUint(r.InstanceID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DetailResponse 获取表单详情 API Response
type DetailResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// Form 表单详细
		Form *Form `json:"form,omitempty"`
	} `json:"data,omitempty"`
}
