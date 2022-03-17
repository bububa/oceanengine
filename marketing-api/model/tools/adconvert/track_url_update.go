package adconvert

import "encoding/json"

// TrackURLUpdateRequest 修改转化监测链接 API Request
type TrackURLUpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ConvertID 转化id
	ConvertID uint64 `json:"convert_id,omitempty"`
	// TrackURL 展示（监测链接）
	TrackURL string `json:"track_url,omitempty"`
	// ActionTrackURL 点击监测链接
	ActionTrackURL string `json:"action_track_url,omitempty"`
	// VideoPlayEffectiveTrackURL 视频有效播放监测链接
	VideoPlayEffectiveTrackURL string `json:"video_play_effective_track_url,omitempty"`
	// VideoPlayDoneTrackURL 视频播放完毕监测链接
	VideoPlayDoneTrackURL string `json:"video_play_done_track_url,omitempty"`
	// VideoPlayTrackURL 视频播放监测链接
	VideoPlayTrackURL string `json:"video_play_track_url,omitempty"`
}

// Encode implement GetRequest interface
func (r TrackURLUpdateRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
