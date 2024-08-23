package blueflow

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PackageListRequest 获取蓝海流量包 API Request
type PackageListRequest struct {
	// Filtering 过滤器
	Filtering *PackageListFilter `json:"filtering,omitempty"`
	// AdvertiserID 广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

type PackageListFilter struct {
	// BlueFlowPackageName 蓝海流量包名称，支持模糊搜索
	BlueFlowPackageName string `json:"blue_flow_package_name,omitempty"`
	// BlueFlowPackageID 蓝海流量包ID，名称与ID同时传入则ID忽略名称筛选
	BlueFlowPackageID uint64 `json:"blue_flow_package_id,omitempty"`
}

// Encode implements GetRequest interface
func (r PackageListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// PackageListResponse 获取蓝海流量包 API Response
type PackageListResponse struct {
	model.BaseResponse
	Data struct {
		// BlueFlowPackages 蓝海流量包信息列表
		BlueFlowPackages []Package `json:"blue_flow_packages,omitempty"`
	} `json:"data,omitempty"`
}
