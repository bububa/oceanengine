package unipromotion

// CreativeSetting 创意设置
type CreativeSetting struct {
	// SmartSelectMaterial 智能优选视频
	// true 开启（默认）
	// false 不开启
	SmartSelectMaterial bool `json:"smart_select_material,omitempty"`
	// HideInAweme 抖音主页可见性设置，和抖音号关系类型相关，返回值参考【附录-抖音号授权类型】
	// 仅单次展示可见 true
	// 主页始终可见 false
	// 官方+自运营（bind_type为OFFICIAL或SELF）
	// 1、全是抖音号主页视频，无需传，传了亦无效
	// 2、存在非抖音号主页原生视频，支持设置
	// 达人（bind_type不为OFFICIAL或SELF）
	// 1、不支持设置，传了亦无效
	HideInAweme bool `json:"hide_in_aweme,omitempty"`
	// CreativeCombineTypelive 直播间画面是否开启
	CreativeCombineTypeLive bool `json:"creative_combine_type_live,omitempty"`
	// CreativeCombineType 自选投放视频
	CreativeCombineType bool `json:"creative_combine_type,omitempty"`
}
