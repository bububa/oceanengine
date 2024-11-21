package task

import "github.com/bububa/oceanengine/marketing-api/enum"

// Demand 任务信息
type Demand struct {
	// StarID 客户的星图id
	StarID uint64 `json:"star_id,omitempty"`
	// DemandID 任务id
	DemandID uint64 `json:"demand_id,omitempty"`
	// DemandName 任务名称
	DemandName string `json:"demand_name,omitempty"`
	// CreateTime 任务创建时间，格式：yyyy-mm-dd HH:MM:SS
	CreateTime string `json:"create_time,omitempty"`
	// EvaluateType 任务优化目标 可选值:
	// ACTIVE 激活
	// ACTIVE_PAY 首次付费
	// DEEP_PURCHASE 每次付费
	// INSTALL_FINISH 安装完成
	// REGISTER 注册
	EvaluateType enum.StarAdUniteTaskEvaluateType `json:"evaluate_type,omitempty"`
	// Budget 任务预算，单位：元*1000，建议获取后除以1000展示为元单位
	Budget int64 `json:"budget,omitempty"`
	// WeekSchedule 投放时段(没返回就是全时段投放)
	WeekSchedule [][]int `json:"week_schedule,omitempty"`
	// Status 任务状态 可选值:
	// BILLING 计费中
	// BULLETIN 公示中
	// CANCELED 已取消
	// CLOSED 已关闭
	// ENROLL 预热中
	// FINISHED 已完成
	// PROVIDER_ACCEPTING 服务商接单中
	// STARTED 投稿中
	Status enum.StarAdUniteTaskStatus `json:"status,omitempty"`
	// AuditStatus 任务审核状态 可选值:
	// CONFIRM 审核通过
	// CONFIRM_FAIL 审核失败
	// PENDING_CONFIRM 审核中
	AuditStatus enum.StarAdUniteTaskAuditStatus `json:"audit_status,omitempty"`
	// StatInfo 分日数据列表
	StatInfo []StatInfo `json:"stat_info,omitempty"`
}

// StatInfo 分日数据
type StatInfo struct {
	// StatDate 数据日期，只和安卓/iOS消耗、转化数、深度转化数相关
	StatDate string `json:"stat_date,omitempty"`
	// AndroidConvertUnitAmount 安卓最新出价，单位：元*1000，建议获取后除以1000展示为元单位
	AndroidConvertUnitAmount int64 `json:"android_convert_unit_amount,omitempty"`
	// IosConvertUnitAmount iOS最新出价，单位：元*1000，建议获取后除以1000展示为元单位
	IosConvertUnitAmount int64 `json:"ios_convert_unit_amount,omitempty"`
	// AndroidCost 安卓消耗，单位：元*100000，建议获取后除以100000展示为元单位
	AndroidCost int64 `json:"android_cost,omitempty"`
	// IosCost iOS消耗，单位：元*100000，建议获取后除以100000展示为元单位
	IosCost int64 `json:"ios_cost,omitempty"`
	// AndroidConvert 安卓转化数
	AndroidConvert int64 `json:"android_convert,omitempty"`
	// IosConvert iOS转化数
	IosConvert int64 `json:"ios_convert,omitempty"`
	// AndroidDeepConvert 深度转化数
	AndroidDeepConvert int64 `json:"android_deep_convert,omitempty"`
	// IosDeepConvert 深度转化数
	IosDeepConvert int64 `json:"ios_deep_convert,omitempty"`
}

// ItemStatInfo 分日数据
type ItemStatInfo struct {
	// StatDate 数据日期，只和安卓/iOS消耗、转化数、深度转化数相关
	StatDate string `json:"stat_date,omitempty"`
	// StarID 客户的星图id
	StarID uint64 `json:"star_id,omitempty"`
	// AuthorID 达人id
	AuthorID uint64 `json:"author_id,omitempty"`
	// AuthorNickName 达人昵称
	AuthorNickName string `json:"author_nick_name,omitempty"`
	// DemandID 任务id
	DemandID uint64 `json:"demand_id,omitempty"`
	// ItemID 视频id
	ItemID uint64 `json:"item_id,omitempty"`
	// ItemURL 视频链接
	ItemURL string `json:"item_url,omitempty"`
	// ItemTitle 视频标题
	ItemTitle string `json:"item_title,omitempty"`
	// ItemStatus 视频状态 可选值:
	// AUDIT_PASS 审核通过
	// CREATED 视频已发布
	// PRIVATE_AREA 仅个人可见
	// USER_DELETED 用户删除
	ItemStatus enum.StarAdUniteTaskItemStatus `json:"item_status,omitempty"`
	// ReleaseTime 视频发布日期
	ReleaseTime int64 `json:"release_time,omitempty"`
	// PlayCnt 播放量，T+1更新，第二天10点后可获取稳定数据
	PlayCnt int64 `json:"play_cnt,omitempty"`
	// Convert 转化数，实时更新
	Convert int64 `json:"convert,omitempty"`
	// Cost 预估实时消耗，实时更新，单位：元*100000，建议获取后除以100000展示为元单位
	Cost int64 `json:"cost,omitempty"`
	// SettledCost 累计已结算消耗，T+1更新，单位：元*100000，建议获取后除以100000展示为元单位，第二天10点后可获取稳定数据
	SettledCost int64 `json:"settled_cost,omitempty"`
}
