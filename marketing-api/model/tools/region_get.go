package tools

import (
	"net/url"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// RegionGetRequest 获取地域列表 API Request
type RegionGetRequest struct {
	// RegionType 地域类型，目前只支持：BUSINESS_DISTRICT(商圈);允许值:"BUSINESS_DISTRICT"
	RegionType string `json:"region_type,omitempty"`
	// RegionLevel 只获取某层级数据，详见【附录-地域层级】
	RegionLevel enum.RegionLevel `json:"region_level,omitempty"`
}

// Encode implement GetRequest interface
func (r RegionGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("region_type", r.RegionType)
	values.Set("region_level", string(r.RegionLevel))
	return values.Encode()
}

// RegionGetResponse  获取地域列表 API Response
type RegionGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		List []Region `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// Region 地域
type Region struct {
	// ID
	ID uint64 `json:"id,omitempty"`
	// Name 名称
	Name string `json:"name,omitempty"`
	// ParentID 父级id
	ParentID uint64 `json:"parent_id,omitempty"`
	// RegionLevel 地域所在层级
	RegionLevel string `json:"region_level,omitempty"`
}
