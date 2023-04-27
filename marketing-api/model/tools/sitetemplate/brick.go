package sitetemplate

// BrickType 组件类型
type BrickType string

const (
	// BrickType_BUTTON 按钮组件
	BrickType_BUTTON BrickType = "BUTTON"
	// BrickType_COUPON 卡券组件
	BrickType_COUPON BrickType = "COUPON"
	// BrickType_FORM 表单组件
	BrickType_FORM BrickType = "FORM"
	// BrickType_PICTURE 图片组件
	BrickType_PICTURE BrickType = "PICTURE"
	// BrickType_PICTURE_GROUP 组图组件
	BrickType_PICTURE_GROUP BrickType = "PICTURE_GROUP"
	// BrickType_RICH_TEXT 富文本组件
	BrickType_RICH_TEXT BrickType = "RICH_TEXT"
	// BrickType_SIMPLE_TEXT 文本组件
	BrickType_SIMPLE_TEXT BrickType = "SIMPLE_TEXT"
	// BrickType_VIDEO 视频组件
	BrickType_VIDEO BrickType = "VIDEO"
	// BrickType_WECHAT_GAME 微信小游戏组件
	BrickType_WECHAT_GAME = "WECHAT_GAME"
)

// Brick 组件
type Brick struct {
	// Type 组件类型，枚举值：BUTTON 按钮组件、COUPON 卡券组件、FORM 表单组件、PICTURE 图片组件、PICTURE_GROUP 组图组件、RICH_TEXT 富文本组件、SIMPLE_TEXT 文本组件、VIDEO 视频组件
	Type BrickType `json:"type,omitempty"`
	// Index 组件在模板中的位置描述
	Index string `json:"index,omitempty"`
	// Video 视频组件描述
	Video *VideoBrick `json:"video,omitempty"`
	// Picture 图片组件描述
	Picture *PictureBrick `json:"picture,omitempty"`
	// PictureGroup 组图组件描述
	PictureGroup *PictureGroupBrick `json:"picture_group,omitempty"`
	// Text 文本组件描述
	Text *TextBrick `json:"text,omitempty"`
	// Button 按钮组件描述
	Button *ButtonBrick `json:"button,omitempty"`
	// Form 表单组件描述
	Form *FormBrick `json:"form,omitempty"`
	// Coupon 发券组件描述
	Coupon *CouponBrick `json:"coupon,omitempty"`
	// WechatGame 微信小游戏组件
	WechatGame *WechatGameBrick `json:"wechat_game,omitempty"`
}

// VideoBrick 视频组件
type VideoBrick struct {
	// LocalVideo 本地视频内容
	LocalVideo *Video
	// OnlineVideo 在线视频内容
	OnlineVideo *Video
}

// Video 视频信息
type Video struct {
	// VideoID 视频ID，本地视频上传后可得到视频ID，可通过【获取视频素材】接口获取，当local_video不为空时，有返回值
	VideoID string `json:"video_id,omitempty"`
	// PosterURL 视频封面图片URL，当local_video不为空时，有返回值
	PosterURL string `json:"poster_url,omitempty"`
	// OriginURL 视频原始链接，例如https://v.youku.com/v_show/id_xxx.html，当online_video不为空时，有返回值
	OriginURL string `json:"origin_url,omitempty"`
}

// PictureBrick 图片组件描述
type PictureBrick struct {
	// URL 图片url，当picture不为空时，有返回值
	URL string `json:"url,omitempty"`
	// Link 图片跳转链接信息
	Link *Link `json:"link_dto,omitempty"`
	// Tag 标签，用户自定义标注
	Tag string `json:"tag,omitempty"`
}

// PictureGroupBrick 组图组件描述
type PictureGroupBrick struct {
	// Content 组图内容列表
	Content []PictureBrick `json:"content,omitempty"`
}

// TextBrick 文本组件描述
type TextBrick struct {
	// Content 文本内容，长度至少为1
	Content string `json:"content,omitempty"`
	// Link 跳转链接信息
	Link *Link `json:"link_dto,omitempty"`
}

// ButtonBrick 按钮组件描述
type ButtonBrick struct {
	// EventType 事件行为类型，当button不为空时，有返回值，枚举值：APPOINT_EVENT 预约事件行为、DOWNLOAD_EVENT 下载事件行为、LINK_ENENT 链接事件行为、TELEPHONE_EVENT 电话事件行为
	EventType EventType `json:"event_type,omitempty"`
	// DownloadEvent downloadEvent事件行为描述
	DownloadEvent *DownloadEvent `json:"download_event,omitempty"`
	// LinkEvent linkEvent事件行为描述
	LinkEvent *LinkEvent `json:"link_event,omitempty"`
	// PhoneEvent phoneEvent事件行为描述
	PhoneEvent *PhoneEvent `json:"phone_event,omitempty"`
	// AppointEvent appointEvent事件行为描述
	AppointEvent *AppointEvent `json:"appoint_event,omitempty"`
}

// FormBrick 表单组件描述
type FormBrick struct {
	// InstanceID 表单ID，当form不为空时，有返回值。用户可以通过【获取表单列表】接口或【青鸟线索通平台】获取表单ID
	InstanceID uint64 `json:"instance_id,omitempty"`
}

// CouponBrick 发券组件描述
type CouponBrick struct {
	// ActivityID 活动ID，当coupon不为空时，有返回值。用户可以通过【获取卡券列表】接口或【青鸟线索通平台】获取活动ID
	ActivityID uint64 `json:"activity_id,omitempty"`
}

// WechatGameBrick 微信小游戏组件
type WechatGameBrick struct {
	// InstanceID 微信小游戏组件ID，当wechat_game不为空时，有返回值。用户可以通过【青鸟线索通平台】获取微信小游戏组件ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// GamePath 微信小游戏的路径参数
	GamePath string `json:"game_path,omitempty"`
	// Items 标签，个数不超过2，字数不超过5个中文字符
	Items []string `json:"items,omitempty"`
	// Introduction 简介，长度不超过40个中文字符
	Introduction string `json:"introduction,omitempty"`
	// Logo logo链接地址
	Logo string `json:"logo,omitempty"`
}
