package aweme

// LiveAuthorize 达人授权信息
type LiveAuthorize struct {
	// IesID 抖音号id
	IesID string `json:"ies_id,omitempty"`
	// IesUserName 抖音昵称
	IesUserName string `json:"ies_user_name,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreateTime 发起时间，格式yyyy-mm-dd hh:MM:ss
	CreateTime string `json:"create_time,omitempty"`
	// AuthorizeStarttime 直播授权开始时间，格式yyyy-mm-dd hh:MM:ss
	AuthorizeStarttime string `json:"authorize_starttime,omitempty"`
	// AuthorizeEndtime 直播授权结束时间，格式yyyy-mm-dd hh:MM:ss
	AuthorizeEndtime string `json:"authorize_endtime,omitempty"`
	// LimitedPromotionTypes 限制抖音号使用场景信息
	LimitedPromotionTypes []LimitedPromotionType `json:"limited_promotion_types,omitempty"`
}

// LimitedPromotionType 限制抖音号使用场景信息
type LimitedPromotionType struct {
	// Msg 场景信息
	Msg string `json:"msg,omitempty"`
}
