package file

import "encoding/json"

// Audio 音频信息
type Audio struct {
	// Vid 音频id
	Vid string `json:"vid,omitempty"`
	// VideoID 音频id
	VideoID string `json:"video_id,omitempty"`
	// URL 音频地址
	URL string `json:"url,omitempty"`
}

func (a *Audio) UnmarshalJSON(data []byte) error {
	type A Audio
	if err := json.Unmarshal(data, (*A)(a)); err != nil {
		return err
	}
	if a.Vid == "" {
		a.Vid = a.VideoID
	} else if a.VideoID == "" {
		a.VideoID = a.Vid
	}
	return nil
}
