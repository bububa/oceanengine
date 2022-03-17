package adconvert

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateStatusRequest 更新转化目标操作状态
type UpdateStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ConvertID 转化id，其中较小数值convert_id为预定义转化
	ConvertID uint64 `json:"convert_id,omitempty"`
	// OptStatus 转化工具操作状态
	OptStatus enum.AdConvertOptStatus `json:"opt_status,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
