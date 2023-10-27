package rta

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetInfoRequest 获取RTA策略数据 API Request
type GetInfoRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告组id，若传入，则拉取的是组维度的RTA策略
	CampaignID uint64 `json:"campaign_id,omitempty"`
}

// Encode implement GetRequest interface
func (r GetInfoRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("campaign_id", strconv.FormatUint(r.CampaignID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetInfoResponse 获取RTA策略数据 API Response
type GetInfoResponse struct {
	model.BaseResponse
	Data *GetInfoData `json:"data,omitempty"`
}

type GetInfoData struct {
	// InterfaceInfo RTA配置数据
	InterfaceInfo *InterfaceInfo `json:"interface_info,omitempty"`
	// RtaInfo RTA策略信息
	RtaInfo *RtaInfo `json:"rta_info,omitempty"`
}

// InterfaceInfo RTA配置数据
type InterfaceInfo struct {
	// Status 接口地址状态
	// 1：生效 0：失效
	Status int `json:"status,omitempty"`
	// DeliveryRange 适用流量范围:
	// LOCAL_ONLY: 站内
	// UNION_ONLY: 穿山甲
	// UNIVERSAL_DELIVERY: 全部
	DeliveryRange string `json:"delivery_range,omitempty"`
	// LocalQPS 站内QPS
	LocalQPS int64 `json:"local_qps,omitempty"`
	// UnionQPS 穿山甲QPS
	UnionQPS int64 `json:"union_qps,omitempty"`
	// URL 接口地址
	URL string `json:"url,omitempty"`
}

// RtaInfo RTA策略信息
type RtaInfo struct {
	// RtaID RTA策略ID
	RtaID uint64 `json:"rta_id,omitempty"`
	// Remark 备注，即RTA策略描述
	Remark string `json:"remark,omitempty"`
}
