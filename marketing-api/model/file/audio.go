package file

import "encoding/json"

// Audio 音频信息
type Audio struct {
	// Vid 音频id
	Vid string `json:"vid,omitempty"`
	// VideoID 音频id（旧版）（该字段将在12月19日下线，暂不对您的调用产生影响，请及时调整，使用下方audio_id）
	VideoID string `json:"video_id,omitempty"`
	// URL 音频地址（旧版）（该字段将在12月19日下线，暂不对您的调用产生影响，请及时调整，使用下方audio_url）
	URL string `json:"url,omitempty"`
	// AudioID 音频id
	AudioID string `json:"audio_id,omitempty"`
	// AudioURL 音频地址
	AudioURL string `json:"audio_url,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// AudioSignature 音频素材md5值
	AudioSignature string `json:"audio_signature,omitempty"`
}

func (a *Audio) UnmarshalJSON(data []byte) error {
	type A Audio
	if err := json.Unmarshal(data, (*A)(a)); err != nil {
		return err
	}
	if a.Vid == "" {
		if a.AudioID != "" {
			a.Vid = a.AudioID
		} else {
			a.Vid = a.VideoID
		}
	} else if a.VideoID == "" {
		if a.AudioID != "" {
			a.VideoID = a.AudioID
		} else {
			a.VideoID = a.Vid
		}
	} else if a.AudioID == "" {
		if a.Vid == "" {
			a.Vid = a.VideoID
		} else if a.VideoID == "" {
			a.VideoID = a.Vid
		}
	}
	if a.URL == "" {
		a.URL = a.AudioURL
	} else if a.AudioURL == "" {
		a.AudioURL = a.URL
	}
	return nil
}
