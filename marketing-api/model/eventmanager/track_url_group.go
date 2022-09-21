package eventmanager

// TrackURLGroup 监测链接组信息
type TrackURLGroup struct {
	// ActionTrackURL 点击（监测链接），只允许传入1个
	ActionTrackURL string `json:"action_track_url,omitempty"`
	// TrackURL 展示（监测链接），只允许传入1个
	TrackURL string `json:"track_url,omitempty"`
	// VideoPlayTrackURL 视频播放（监测链接），只允许传入1个
	VideoPlayTrackURL string `json:"video_play_track_url,omitempty"`
	// VideoPlayDoneTrackURL 视频播完（监测链接），只允许传入1个
	VideoPlayDoneTrackURL string `json:"video_play_done_track_url,omitempty"`
	// VideoPlayEffectiveTrackURL 视频有效播放（监测链接），只允许传入1个
	VideoPlayEffectiveTrackURL string `json:"video_play_effective_track_url,omitempty"`
	// TrackURLGroupName 监测链接组名称，应用型资产必填
	TrackURLGroupName string `json:"track_url_group_name,omitempty"`
	// TrackURLGroupID 监测链接组ID
	TrackURLGroupID uint64 `json:"track_url_group_id,omitempty"`
}
