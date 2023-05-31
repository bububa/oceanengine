package clue

import "github.com/bububa/oceanengine/marketing-api/model"

// Clue 线索
type Clue struct {
	// ClueID 线索ID
	ClueID model.Uint64 `json:"clue_id,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID model.Uint64 `json:"advertiser_id,omitempty"`
	// AdvertiserName 广告主名
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// AdID 计划ID
	AdID model.Uint64 `json:"ad_id,omitempty"`
	// AdName 计划名
	AdName string `json:"ad_name,omitempty"`
	// CreativeID 广告创意id
	CreativeID model.Uint64 `json:"creative_id,omitempty"`
	// PromotionName 体验版广告名称
	PromotionName string `json:"promotion_name,omitempty"`
	// PromotionID 体验版广告ID
	// 注意：体验版广告信息与计划/创意信息互斥，不会同时返回
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// SiteID 站点ID
	SiteID model.Uint64 `json:"site_id,omitempty"`
	// SiteName 站点名称
	SiteName string `json:"site_name,omitempty"`
	// MidInfo 升级版素材组(投放2.0)
	// 注意：投放1.0为空，mid_info是一个jsonstring，示例
	//  {    "titleId" : 123,    "videoId" : 234,    "imageId" : 135}
	MidInfo string `json:"mid_info,omitempty"`
	// IntentionEstimation 线索意向。拉取到线索后，intention_estimation的获取会有约 10 分钟延迟；建议 10 分钟后重新拉取并和该条线索进行匹配
	IntentionEstimation string `json:"intention_estimation,omitempty"`
	// ExternalURL 落地页 url
	ExternalURL string `json:"external_url,omitempty"`
	// ClueType 组件类型，允许值：
	// 0: 表单提交；
	// 1: 在线咨询;
	// 2: 智能电话;
	// 3: 网页回呼;
	// 4: 卡券;
	// 5：抽奖
	ClueType int `json:"clue_type,omitempty"`
	// ModuleName 组件名字
	ModuleName string `json:"module_name,omitempty"`
	// ModuleID 组件ID
	ModuleID model.Uint64 `json:"module_id,omitempty"`
	// CreateTime 线索创建时间，如：2020-04-29
	CreateTime string `json:"create_time,omitempty"`
	// CreateTimeDetail 线索创建时间，如：2020-04-29 00:00:00
	CreateTimeDetail string `json:"create_time_detail,omitempty"`
	// Date 日期
	Date string `json:"date,omitempty"`
	// ClueSource 线索来源，目前几种线索来源分别是，允许值：
	// 0: 外部流量；
	// 1: 正常投放;
	// 2: 外部导入;
	// 3: 异常提交;
	// 4: 广告预览;
	// 5: 抖音私信;
	// 6: 鲁班线索
	ClueSource int `json:"clue_source,omitempty"`
	// Name 姓名
	Name string `json:"name,omitempty"`
	// Telephone 电话
	Telephone string `json:"telephone,omitempty"`
	// Gender 性别
	// 允许值：
	// 0：未知；
	// 1：男 ；
	// 2：女
	Gender int `json:"gender,omitempty"`
	// Age 年龄
	Age int `json:"age,omitempty"`
	// Email 邮箱
	Email string `json:"email,omitempty"`
	// Weixin 微信
	Weixin string `json:"weixin,omitempty"`
	// QQ QQ
	QQ string `json:"qq,omitempty"`
	// ProvinceName 省份
	ProvinceName string `json:"province_name,omitempty"`
	// CityName 城市
	CityName string `json:"city_name,omitempty"`
	// Address 详细地址
	Address string `json:"address,omitempty"`
	// Location 用户地址信息
	Location string `json:"location,omitempty"`
	// Remark 用户备注留言
	Remark string `json:"remark,omitempty"`
	// RemarkDict 广告主自定义的其它字段信息，可自定义获取; 例：若创建表单时自定义了"test_case"字段的设置，则产生的remark_dict线索为："test_case": "xxx"
	RemarkDict map[string]string `json:"remark_dict,omitempty"`
	// FormRemark 广告主表单自定义字段，无业务逻辑，仅作表单标志
	FormRemark string `json:"form_remark,omitempty"`
	// AppName 流量来源，比如「今日头条 」、「抖音」 拉取到线索后，app_name 的获取会有约 10 分钟延迟；建议 10 分钟后重新拉取并和该条线索进行匹配。
	AppName string `json:"app_name,omitempty"`
	// ConvertStatus 转化状态;允许值：数据获取中、无、合法转化、不匹配转化、广告预览、其他不合法转化、外部流量;拉取到线索后，convert_status 的获取会有约 10 分钟延迟；建议 10 分 钟后重新拉取并和该条线索进行匹配。
	ConvertStatus string `json:"convert_status,omitempty"`
	// ReqID 当前线索对应广告的请求id
	ReqID string `json:"req_id,omitempty"`
	// ClueState 线索状态值; 枚举值：1 新线索、2 明确意向、3 转商机、4 无效、5 加微信
	ClueState int `json:"clue_state,omitempty"`
	// ClueStateName 线索状态名称，“clue_state”对应的名称
	ClueStateName string `json:"clue_state_name,omitempty"`
	// ClueOwnerName 所属人姓名
	ClueOwnerName string `json:"clue_owner_name,omitempty"`
	// FollowStateName 线索通话状态名称; 包含值：未联系、未接通、已接通、有效沟通
	FollowStateName string `json:"follow_state_name,omitempty"`
	// CountryName 所在城市中的【区】字段
	CountryName string `json:"country_name,omitempty"`
	// SystemTags 线索相关的系统标签，是一个标签项的数组
	SystemTags []string `json:"system_tags,omitempty"`
	// Tags 线索被打上的人工标签，，是一个标签项的数组，包括自定义标签和行业标签
	Tags []string `json:"tags,omitempty"`
	// AllocationStatus 分配状态值; 枚举值：0 待分配、1 已分配
	AllocationStatus int `json:"allocation_status,omitempty"`
	// Store 门店信息
	Store *Store `json:"store,omitempty"`
}
