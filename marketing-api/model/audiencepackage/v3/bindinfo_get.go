package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BindInfoGetRequest 定向包查询关联项目信息 API Request
type BindInfoGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AudiencePackageID 定向包ID
	AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	// Page 分页
	Page int `json:"page,omitempty"`
	// Size  分页尺寸
	Size int `json:"size,omitempty"`
}

// Encode implements GetRequest interface
func (r BindInfoGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("audience_package_id", strconv.FormatUint(r.AudiencePackageID, 10))
	values.Set("page", strconv.Itoa(r.Page))
	values.Set("size", strconv.Itoa(r.Size))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// BindInfoGetResponse 定向包查询关联项目信息 API Response
type BindInfoGetResponse struct {
	Data *BindInfoGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type BindInfoGetResult struct {
	// PageInfo 翻页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 项目信息列表
	List []BindProject `json:"list,omitempty"`
}

// BindProject 项目信息
type BindProject struct {
	// ProjectName 项目名称
	ProjectName string `json:"project_name,omitempty"`
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
}
