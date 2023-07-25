package duoplus

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// Order  Duo+订单
type Order struct {
	// Order 订单信息
	Order *OrderInfo `json:"order,omitempty"`
	// AdList 计划列表
	AdList []Ad `json:"ad_list,omitempty"`
	// LiveRoomInfo 直播间信息
	LiveRoomInfo *LiveRoomInfo `json:"live_room_info,omitempty"`
	// ItemInfoList 视频信息
	ItemInfoList []AwemeItem `json:"item_info_list,omitempty"`
}

// OrderInfo 订单信息
type OrderInfo struct {
	// OrderID 订单id
	OrderID uint64 `json:"order_id,omitempty"`
	// TaskID 订单id（页面上显示）
	TaskID uint64 `json:"task_id,omitempty"`
	// SceneType 营销目标
	SceneType enum.MarketingGoal `json:"scene_type,omitempty"`
	// OrderCreateTime 订单创建时间
	OrderCreateTime string `json:"order_create_time,omitempty"`
	// TaskStatus 任务状态 可选值：
	// UNPAID未支付
	// AUDITING 审核中
	// OFFLINE_AUDIT审核不通过
	// TIME_NO_REACH 待开播
	// DELIVERING 投放中
	// UNDELIVERIED 投放终止
	// DELIVERIED 投放完成/结束
	// AUDIT_PAUSE 审核暂停
	// WAIT_TO_START 等待开始
	TaskStatus enum.DuoplusOrderStatus `json:"task_status,omitempty"`
	// Budget 投放金额，单位：分
	Budget int64 `json:"budget,omitempty"`
	// LiveScene 可选值:
	// 混合投放
	// 直推直播间
	// 视频导流直播间
	LiveScene string `json:"live_scene,omitempty"`
}

// Ad 计划
type Ad struct {
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// AdStatus 状态 可选值：
	// UNPAID未支付
	// AUDITING 审核中
	// OFFLINE_AUDIT审核不通过
	// TIME_NO_REACH 待开播
	// DELIVERING 投放中
	// UNDELIVERIED 投放终止
	// DELIVERIED 投放完成/结束
	// AUDIT_PAUSE 审核暂停
	// WAIT_TO_START 等待开始
	AdStatus enum.DuoplusOrderStatus `json:"ad_status,omitempty"`
	// IsFans 是否粉丝必见计划
	IsFans bool `json:"is_fans,omitempty"`
	// Budget 投放金额，单位：分
	Budget int64 `json:"budget,omitempty"`
	// ExternalAction 优化目标
	ExternalAction enum.TrackEventType `json:"external_action,omitempty"`
	// DeliveryTime 期望投放时长（小时）
	DeliveryTime float64 `json:"delivery_time,omitempty"`
	// BidMode 可选值:
	// 保播放
	// 保转化
	// 自动出价
	BidMode string `json:"bid_mode,omitempty"`
	// CpaBid cpa出价，单位：元
	CpaBId float64 `json:"cpa_bid,omitempty"`
	// Audience 定向信息
	Audience *Audience `json:"audience,omitempty"`
}

// Audience 定向信息
type Audience struct {
	// Gender 性别定向
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// Age 年龄定向
	Age []enum.AudienceAge `json:"age,omitempty"`
	// District 可选值:
	// 区县
	// 商圈
	// 省市
	District string `json:"district,omitempty"`
	// City 定向城市信息
	City []string `json:"city,omitempty"`
	// Province 定向省份信息
	Province []string `json:"province,omitempty"`
	// InterestCategories 兴趣类目词
	InterestCategories []string `json:"interest_categories,omitempty"`
	// AuthorPkgs 抖音达人ID列表
	AuthorPkgs []string `json:"author_pkgs,omitempty"`
	// DeliveryType 定向模式
	// 智能定向
	// 自定义定向
	DeliveryType string `json:"delivery_type,omitempty"`
	// Business 定向商圈
	Business []string `json:"business,omitempty"`
}

// LiveRoomInfo 直播间信息
type LiveRoomInfo struct {
	// RoomID 直播间id
	RoomID uint64 `json:"room_id,omitempty"`
	// RoomTitle 直播间标题
	RoomTitle string `json:"room_title,omitempty"`
	// RoomCover 直播间封面
	RoomCover []string `json:"room_cover,omitempty"`
	// RoomStatus 直播间状态，可选值:
	// FINISH: 结束
	// LIVING: 直播中
	// PAUSE: 暂停
	// PREPARE: 未开播
	// UNKNOW: 未知
	RoomStatus enum.LiveRoomStatus `json:"room_status,omitempty"`
}

// AwemeItem 视频信息
type AwemeItem struct {
	// AwemeAuthorName 抖音号昵称
	AwemeAuthorName string `json:"aweme_author_name,omitempty"`
	// AwemeAuthorUniqueID 抖音号
	AwemeAuthorUniqueID string `json:"aweme_author_unique_id,omitempty"`
	// AwemeAuthorAvatar 抖音号头像
	AwemeAuthorAvatar []string `json:"aweme_author_avatar,omitempty"`
	// AwemeItemID 视频id
	AwemeItemID model.Uint64 `json:"aweme_item_id,omitempty"`
	// AwemeItemTitle 视频标题
	AwemeItemTitle string `json:"aweme_item_title,omitempty"`
	// AwemeItemCover 视频封面
	AwemeItemCover []string `json:"aweme_item_cover,omitempty"`
}
