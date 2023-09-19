package enterprise

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BindListGetRequest 获取广告主关联的企业号列表 API Request
type BindListGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r BindListGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// BindListGetResponse 获取广告主关联的企业号列表 API Response
type BindListGetResponse struct {
	model.BaseResponse
	Data struct {
		// List 关联的抖音号列表（含企业号）
		List []BindItem `json:"list"`
	} `json:"data"`
}

// BindItem 关联的抖音号
type BindItem struct {
	// OpenID 抖音号open_id（含企业号open_id）
	OpenID string `json:"open_id,omitempty"`
	// AwemeID 抖音号id（含企业号id）
	AwemeID string `json:"aweme_id,omitempty"`
	// AwemeName 抖音号名称（含企业号名称）
	AwemeName string `json:"aweme_name,omitempty"`
	// AwemeAvatar 抖音号头像（含企业号头像）
	AwemeAvatar string `json:"aweme_avatar,omitempty"`
	// AuthorizeTimes 授权有效期
	AuthorizeTimes []AuthorizeTime `json:"authorize_times,omitempty"`
}

// AuthorizeTime 授权有效期
type AuthorizeTime struct {
	// AuthorizeStartTime 授权开始时间
	AuthorizeStartTime string `json:"authorize_start_time,omitempty"`
	// AuthorizeEndTime 授权结束时间
	AuthorizeEndTime string `json:"authorize_end_time,omitempty"`
}
