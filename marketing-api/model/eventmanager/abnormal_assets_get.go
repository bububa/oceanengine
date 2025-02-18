package eventmanager

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AbnormalAssetsGetRequest 获取异常应用资产列表 API Request
type AbnormalAssetsGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// Encode implements GetRequest interface
func (r AbnormalAssetsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if v := r.PageInfo; v != nil {
		values.Set("page_info", string(util.JSONMarshal(v)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AbnormalAssetsGetResponse 获取异常应用资产列表
type AbnormalAssetsGetResponse struct {
	model.BaseResponse
	Data *AbnormalAssetsGetResult `json:"data,omitempty"`
}

type AbnormalAssetsGetResult struct {
	// Date 数据更新时间（天级别更新）
	Date string `json:"date,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// SingleCase 分包的接入状态信息
	SingleCase *AbnormalAsset `json:"single_case,omitempty"`
}

// AbnormalAsset 分包的接入状态信息
type AbnormalAsset struct {
	// AppType 应用类型，枚举值：
	// ANDROID 安卓
	// IOS IOS
	AppType string `json:"app_type,omitempty"`
	// Package 包名
	Package string `json:"package,omitempty"`
	// Suggest 建议和明细信息
	Suggest []AbnormalAssetSuggest `json:"suggest,omitempty"`
}

// AbnormalAssetSuggest 建议和明细信息
type AbnormalAssetSuggest struct {
	// Message 建议
	Message string `json:"message,omitempty"`
	// URL 明细信息链接
	URL string `json:"url,omitempty"`
}
