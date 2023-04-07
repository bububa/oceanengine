package appmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BookingRecordsGetRequest 查询应用预约记录 API Request
type BookingRecordsGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PackageID 应用id
	// 如果是Anroid应用，package_id从【查询应用信息】接口获取；
	// 如果是IOS应用，package_id为App Store中的ItunesID；
	PackageID string `json:"package_id,omitempty"`
	// HostType 应用是否寄存在应用管理中:
	// HOST_IN:寄存
	// HOST_OUT:未寄存
	// 如果是Anroid应用，必须先上传至应用管理中，请选择寄存；
	// 如果是IOS应用,请选择未寄存;
	HostType HostType `json:"host_type,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值为10，最大不超过200
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r BookingRecordsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("package_id", r.PackageID)
	values.Set("host_type", string(r.HostType))
	if r.Page != 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// BookingRecordsGetResponse 查询应用预约记录 API Response
type BookingRecordsGetResponse struct {
	model.BaseResponse
	Data *BookingRecordsGetData `json:"data,omitempty"`
}

type BookingRecordsGetData struct {
	// List 预约记录数据
	List []BookingRecord `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// BookingRecord 预约记录
type BookingRecord struct {
	// AdID 广告计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// CreativeID 广告创意ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// OrderID 预约ID
	OrderID string `json:"order_id,omitempty"`
	// PackageID 应用ID
	PackageID string `json:"package_id,omitempty"`
	// ReqID 请求ID
	ReqID string `json:"req_id,omitempty"`
	// CreateTime 创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
}
