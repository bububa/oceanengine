package tools

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeAuthRequest 广告主添加抖音号 API Request
type AwemeAuthRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 要授权的抖音号
	AwemeID string `json:"aweme_id,omitempty"`
	// AwemeShowID 抖音号，即用户在抖音端上的抖音号
	// aweme_id和aweme_show_id至少入参一个，如果两个都入参，以aweme_id为准
	AwemeShowID string `json:"aweme_show_id,omitempty"`
	// Code 达人合作码
	Code string `json:"code,omitempty"`
	// AuthType 授权类型，可选值
	// 合作达人：AWEME_COOPERATOR，支持选取主页视频投放、支持推广该抖音号直播间
	// 自运营：SELF，支持选取主页视频投放、支持推广该抖音号直播间、支持发布视频到主页、支持获取抖音号数据(包括被投放广告数据、粉丝数据)
	AuthType enum.AwemeBindType `json:"auth_type,omitempty"`
	// EndTime 授权结束时间，格式为yyyy-MM-dd HH:mm:ss
	// 如果想要不限，end_time=2099-12-31 23:59:59
	EndTime string `json:"end_time,omitempty"`
	// Notes 用户备注
	Notes string `json:"notes,omitempty"`
}

// Encode implemnet PostRequest interface
func (r AwemeAuthRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AwemeAuthResponse 广告主添加抖音号 API Response
type AwemeAuthResponse struct {
	model.BaseResponse
	Data struct {
		// AuthSuccess 授权状态，返回值：true 授权成功、false授权失败
		AuthSuccess bool `json:"auth_success,omitempty"`
	} `json:"data,omitempty"`
}
