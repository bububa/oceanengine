package union

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// FlowPackageGetRequest 获取穿山甲流量包 API Request
type FlowPackageGetRequest struct {
	// AdvertiserID 获取穿山甲流量包
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤字段
	Filtering *FlowPackageGetFilter `json:"filtering,omitempty"`
	// Page 页数
	// 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	// 默认值: 10，最大值100
	PageSize int `json:"page_size,omitempty"`
}

// FlowPackageGetFilter 过滤字段
type FlowPackageGetFilter struct {
	// FlowPackageIDs 按照流量包ID过滤，最大数量限制：100
	FlowPackageIDs []uint64 `json:"flow_package_ids,omitempty"`
	// FlowPackageType 按照流量包类型过滤，允许值：
	// CUSTOMIZE：自定义、 FEATURED：运营推荐、 SYSTEM： 系统推荐
	FlowPackageType enum.FlowPackageType `json:"flow_package_type,omitempty"`
}

// Encode implement GetRequest interface
func (r FlowPackageGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		bs, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(bs))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// FlowPackageGetResponse 获取穿山甲流量包 API Response
type FlowPackageGetResponse struct {
	model.BaseResponse
	// Data json 返回值`
	Data *FlowPackageGetData `json:"data,omitempty"`
}

// FlowPackageGetData json 返回值
type FlowPackageGetData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 流量包列表
	List []FlowPackage `json:"list,omitempty"`
}
