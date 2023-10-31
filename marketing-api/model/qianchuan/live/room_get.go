package live

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RoomGetRequest 获取今日直播间列表 API Request
type RoomGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音号ID
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// DateTime 开播日期，格式 2021-04-05
	DateTime string `json:"date_time,omitempty"`
	// RoomStatus 直播间状态，可选值，默认全部
	// 全部ALL
	// 直播中LIVING
	// 直播结束FINISH
	RoomStatus enum.LiveRoomStatus `json:"room_status,omitempty"`
	// AdStatus 投放状态，可选值，默认不限
	// 不限ALL
	// 广告在投DELIVERY_OK
	// 暂无投放NO_DELIVERY
	AdStatus string `json:"ad_status,omitempty"`
	// Fields 要查询的消耗指标，具体可参考返回值
	Fields []string `json:"fields,omitempty"`
	// Page 页码，默认为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认为20，取值范围：1-100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r RoomGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	values.Set("date_time", r.DateTime)
	if r.RoomStatus != "" {
		values.Set("room_status", string(r.RoomStatus))
	}
	if r.AdStatus != "" {
		values.Set("ad_status", r.AdStatus)
	}
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RoomGetResponse 获取今日直播间列表 API Response
type RoomGetResponse struct {
	model.BaseResponse
	Data *RoomGetData `json:"data,omitempty"`
}

type RoomGetData struct {
	// RoomList 直播间列表
	RoomList []RoomStat `json:"room_list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// RoomStat 直播间统计数据
type RoomStat struct {
	Room
	report.LiveStat
}
