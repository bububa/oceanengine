package site

// VideoBrick 视频组件
type VideoBrick struct {
	BaseBrick
	// LocalSource 本地视频信息，与在线视频信息传其中一个
	LocalSource *VideoSource `json:"local_source,omitempty"`
	// OnlineSource 在线视频信息，与本地视频信息传其中一个
	OnlineSource *VideoSource `json:"online_source,omitempty"`
}

// VideoSource 视频信息
type VideoSource struct {
	// Vid 视频云id, 用户自行上传，local_source不空时，必填
	Vid string `json:"vid,omitempty"`
	// Poster 视频封面图信息,local_source不空时，必填支持填写在【获取图片素材】返回的图片id（推荐）和自行上传的图片url
	Poster string `json:"poster,omitempty"`
	// OriginURL 视频原始地址，例如https://v.youku.com/v_show/id_xxx.html ，online_source不空时必填
	OriginURL string `json:"origin_url,omitempty"`
	// PosterURL 封面图url，用户自定义上传
	PosterURL string `json:"poster_url,omitempty"`
}

// Type implement Brick interface
func (b VideoBrick) Type() BrickType {
	return XrVideo
}
