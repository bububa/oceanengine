package live

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RoomDetailGetRequest 获取直播间详情 API Request
type RoomDetailGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// RoomID 直播间id
	RoomID uint64 `json:"room_id,omitempty"`
}

// Encode implement GetRequest interface
func (r RoomDetailGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("room_id", strconv.FormatUint(r.RoomID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RoomDetailGetResponse 获取直播间详情 API Response
type RoomDetailGetResponse struct {
	model.BaseResponse
	Data *Room `json:"data,omitempty"`
}
