package clue

import (
	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// Clue 线索
type Clue struct {
	// ClueID 线索ID
	ClueID string `json:"clue_id,omitempty"`
	// LocalAccountID 广告主ID
	LocalAccountID model.Uint64 `json:"local_account_id,omitempty"`
	// AdvertiserName 广告主名
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// PromotionID 广告ID
	PromotionID model.Uint64 `json:"promotion_id,omitempty"`
	// PromotionName 广告名称
	PromotionName string `json:"promotion_name,omitempty"`
	// ContentID 内容ID
	ContentID string `json:"content_id,omitempty"`
	// ToolID 线索工具ID
	ToolID string `json:"tool_id,omitempty"`
	// CreateTimeDetail 线索创建时间，如：2020-04-29 00:00:00
	CreateTimeDetail string `json:"create_time_detail,omitempty"`
	// ModifyTime 线索修改时间，如：2020-04-29 00:00:00
	ModifyTime string `json:"modify_time,omitempty"`
	// Name 姓名
	Name string `json:"name,omitempty"`
	// Telephone 客户留资手机号。当团购订单退款后，不可获取明文手机号。 默认值: ""
	Telephone string `json:"telephone,omitempty"`
	// Gender 性别 可选值:
	// UNKNOWN 未知
	// MALE 男
	// FEMALE 女
	// 默认值: UNKNOWN
	Gender local.Gender `json:"gender,omitempty"`
	// Age 年龄
	Age int `json:"age,omitempty"`
	// ProvinceName 用户填写省份
	ProvinceName string `json:"province_name,omitempty"`
	// CityName 用户填写城市
	CityName string `json:"city_name,omitempty"`
	// CountyName 用户填写区县
	CountyName string `json:"county_name,omitempty"`
	// Address 用户填写详细地址
	Address string `json:"address,omitempty"`
	// Remark 商家备注
	Remark string `json:"remark,omitempty"`
	// RemarkDict 商家表单自定义的字段信息，及其他线索相关信息
	RemarkDict string `json:"remark_dict,omitempty"`
	// FlowType 流量类型 可选值:
	// NATURE 自然流量
	// AD 广告流量
	FlowType local.ClueFlowType `json:"flow_type,omitempty"`
	// ActionType 互动场景 可选值:
	// SHORT_VIDEO 短视频
	// LIVE_VIDEO 直播
	// HOME_PAGE 企业主页
	// IM_MESSAGE 消息列表
	// GROUPON_ORDER 团购tab
	// ALIEN_CARD 异形卡
	// OTHERS 其他
	ActionType local.ClueActionType `json:"action_type,omitempty"`
	// LeadsPage 留资页面 可选值:
	// POI POI页
	// PRODUCT_DETAIL 商详页
	// OTHER 其他
	LeadsPage local.LeadsPage `json:"leads_page,omitempty"`
	// ClueType 留资组件 可选值:
	// FORM 表单提交
	// CONSULT 在线咨询
	// SMARTPHONE 智能电话
	// GROUP_BUYING 团购留资
	ClueType local.ClueType `json:"clue_type,omitempty"`
	// FollowLifeAccountType 跟进账户类型 可选值:
	// HEAD 总部
	// REGION 区域
	// SINGLE 门店
	FollowLifeAccountType local.FollowLifeAccountType `json:"follow_life_account_type,omitempty"`
	// FollowLifeAccountID 跟进账户ID
	FollowLifeAccountID string `json:"follow_life_account_id,omitempty"`
	// FollowLifeAccountName 跟进账户名称
	FollowLifeAccountName string `json:"follow_life_account_name,omitempty"`
	// OrderID 订单ID
	OrderID uint64 `json:"order_id,omitempty"`
	// EffectiveState 线索阶段 可选值:
	// 0 新线索
	// 1 有意向
	// 2 成交
	// 3 无效
	// 6 已加微信
	// 7 待再次沟通
	// 204 到店
	EffectiveState int `json:"effective_state,omitempty"`
	// ClueOwnerName 所属人姓名
	ClueOwnerName string `json:"clue_owner_name,omitempty"`
	// FollowStateName 线索通话状态 可选值:
	// NOT_CALLED 待联系
	// NOT_ANSWERED 未接通
	// SHORT_ANSWERED 已接通
	// ANSWERED 有效沟通
	// DEEP_ANSWERED 深度沟通
	FollowStateName local.FollowStateName `json:"follow_state_name,omitempty"`
	// SystemTags 线索被打上的系统标签，是一个标签项的数组
	SystemTags []string `json:"system_tags,omitempty"`
	// Tags 线索被打上的人工标签，是一个标签项的数组，包括自定义标签和行业标签
	Tags []string `json:"tags,omitempty"`
	// AllocationStatus  分配状态 可选值:
	// NOT_ASSIGN 待分配
	// ASSIGNED 已分配
	AllocationStatus local.ClueAllocationStatus `json:"allocation_status,omitempty"`
	// ReqID 当前线索对应广告的请求id
	ReqID string `json:"req_id,omitempty"`
}
