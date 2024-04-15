package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// IesAccountSearchRequest 获取绑定的抖音号 API Request
type IesAccountSearchRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r IesAccountSearchRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// IesAccountSearchResponse 获取绑定的抖音号 API Response
type IesAccountSearchResponse struct {
	model.BaseResponse
	Data []IesAccount `json:"data,omitempty"`
}

// IesAccount 绑定的抖音号
type IesAccount struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdvertiserName 广告主名
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// AwemeID 抖音id
	AwemeID string `json:"aweme_id,omitempty"`
	// AwemeName 抖音昵称
	AwemeName string `json:"aweme_name,omitempty"`
	// AwemeAvatar 抖音头像
	AwemeAvatar string `json:"aweme_avatar,omitempty"`
	// BindType 绑定类型，“1”：竞价，“2”：品牌
	BindType string `json:"bind_type,omitempty"`
	// Status 绑定状态
	Status enum.IesAccountStatus `json:"status,omitempty"`
	// StartTime 绑定起始日期，格式：YYYY-MM-DD
	StartTime string `json:"start_time,omitempty"`
	// CommerceUserLevel 企业主页等级
	CommerceUserLevel string `json:"commerce_user_level,omitempty"`
	// SharedAdvertisers 共享的广告主列表
	SharedAdvertisers []IesSharedAdvertiser `json:"shared_advertisers,omitempty"`
}

// IesSharedAdvertiser 共享的广告主
type IesSharedAdvertiser struct {
	// SharedAdvertiserID 共享的广告主id
	SharedAdvertiserID uint64 `json:"shared_advertiser_id,omitempty"`
	// SharedAdvertiserName 共享的广告主名
	SharedAdvertiserName string `json:"shared_advertiser_name,omitempty"`
}
