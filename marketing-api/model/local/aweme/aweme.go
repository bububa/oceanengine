package aweme

import "github.com/bububa/oceanengine/marketing-api/enum"

// Aweme 抖音号
type Aweme struct {
	// AwemeID 抖音号
	AwemeID string `json:"aweme_id,omitempty"`
	// AwemeName 抖音号名称
	AwemeName string `json:"aweme_name,omitempty"`
	// AwemeAvatar 抖音头像
	AwemeAvatar string `json:"aweme_avatar,omitempty"`
	// AuthType 抖音号授权类型，枚举值：
	// OFFICIAL 官方
	// SELF 自运营
	// AWEME_COOPERATOR 合作达人
	// 不同授权类型的抖音号权限说明：
	// 官方：支持发布视频到主页（政媒号不支持）、支持选取主页视频投放、支持推广该抖音号的直播间
	// 自运营：支持发布视频到主页（政媒号不支持）、支持选取主页视频投放、支持推广该抖音号的直播间
	// 合作达人：支持选取主页视频投放、支持推广该抖音号的直播间
	AuthType enum.AwemeBindType `json:"auth_type,omitempty"`
	// AwemeHasUniProm 该抖音号是否有直播roi2计划投放
	// 注意：如果该抖音号有直播roi2计划投放，那么该抖音号不可以创建 【标准推广-推直播间-直播间团购购买】 的直播广告
	AwemeHasUniProm bool `json:"aweme_has_uni_prom,omitempty"`
	// CanCreateRoi2Ad 该抖音号是否能创建roi2广告
	// marketing_goal=LIVE时返回
	CanCreateRoi2Ad bool `json:"can_create_roi2_ad,omitempty"`
}
