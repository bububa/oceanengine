package eventmanager

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// EventsCreateRequest 资产下创建事件 API Request
type EventsCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetID 资产ID
	AssetID uint64 `json:"asset_id,omitempty"`
	// EventID 事件ID
	EventID uint64 `json:"event_id,omitempty"`
	// TrackTypes 事件回传方式列表，允许值：JSSDK JS埋码 、EXTERNAL_API API回传 、XPATH XPath圈选
	TrackTypes []enum.EventTrackType `json:"track_types,omitempty"`
}

// Encode implement PostRequest interface
func (r EventsCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
