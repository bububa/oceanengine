package live

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RoomUserGetRequest 获取直播间用户洞察 API Request
type RoomUserGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// RoomID 直播间id
	RoomID uint64 `json:"room_id,omitempty"`
	// ActionEvent  用户来源，
	// 进入直播间：ENTER
	// 支付成功：PAY
	ActionEvent qianchuan.UserActionEvent `json:"action_event,omitempty"`
	// Dimensions 洞察维度，可选择 "CITY", "GENDER", "AGE"
	Dimensions []string `json:"dimensions,omitempty"`
	// FlowSource  广告类型，默认全部广告
	// 全部广告：ALL
	// 全部千川PC端广告：PC
	FlowSource string `json:"flow_source,omitempty"`
}

// Encode implement GetRequest interface
func (r RoomUserGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("room_id", strconv.FormatUint(r.RoomID, 10))
	if len(r.Dimensions) > 0 {
		values.Set("dimensions", string(util.JSONMarshal(r.Dimensions)))
	}
	if r.FlowSource != "" {
		values.Set("flow_source", r.FlowSource)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RoomUserGetResponse 获取直播间用户洞察 API Response
type RoomUserGetResponse struct {
	model.BaseResponse
	Data *RoomUser `json:"data,omitempty"`
}

// RoomUser 直播间用户洞察
type RoomUser struct {
	// City 地域
	City *City `json:"city,omitempty"`
	// Gender 性别
	Gender *Gender `json:"gender,omitempty"`
	// Age 年龄
	Age *Age `json:"age,omitempty"`
}

// City 地域
type City struct {
	// CityName 城市名称
	CityName string `json:"city_name,omitempty"`
	// Value 人数
	Value int64 `json:"value,omitempty"`
}

// Gender 性别
type Gender struct {
	// GenderType 性别类型
	GenderType string `json:"gender_type,omitempty"`
	// Value 人数
	Value int64 `json:"value,omitempty"`
}

// Age 年龄
type Age struct {
	// AgeType 年龄阶段
	AgeType string `json:"age_type,omitempty"`
	// Value 人数
	Value int64 `json:"value,omitempty"`
}
