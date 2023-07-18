package dmp

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OrientationPackageGetRequest 获取定向包列表 API Request
type OrientationPackageGetRequest struct {
	// AdvertiserID 千川广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *OrientationPackageGetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认为20，取值范围：1-100
	PageSize int `json:"page_size,omitempty"`
}

// OrientationPackageGetFilter 过滤条件
type OrientationPackageGetFilter struct {
	// Name 定向包名称，长度限制20个汉字
	Name string `json:"name,omitempty"`
	// ID 定向包ID，最多支持100个
	ID []uint64 `json:"id,omitempty"`
}

// Encode implement GetRequest interface
func (r OrientationPackageGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// OrientationPackageGetResponse 获取定向包列表 API Response
type OrientationPackageGetResponse struct {
	model.BaseResponse
	Data *OrientationPackageGetResult `json:"data,omitempty"`
}

type OrientationPackageGetResult struct {
	// List 定向包列表
	List []Orientation `json:"list,omitempty"`
	// PageInfo 页码信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
