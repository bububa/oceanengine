package dmp

import "github.com/bububa/oceanengine/marketing-api/enum/qianchuan"

// Audience 人群
type Audience struct {
	// AudienceID 人群ID
	AudienceID uint64 `json:"audience_id,omitempty"`
	// AudienceName 人群名称
	AudienceName string `json:"audience_name,omitempty"`
	// AudienceSource 人群来源，枚举值:
	// 0: 云图
	// 1：千川DMP
	// 2：用户360
	// 3：抖点罗盘
	// 4：抖点用户中心
	// 5：巨量引擎DMP
	AudienceSource int `json:"audience_source,omitempty"`
	// Status 人群状态，枚举值：
	// 1：计算完成
	// 2：等待计算
	// 3：计算中
	// 4：计算失败
	// 5：即将过期
	// 6：已过期
	// 7：推送中
	// 8：推送完成
	// 9：推送失败
	Status int `json:"status,omitempty"`
	// AudienceType 人群类型，枚举值:
	// BASIC: 基础
	// EXTEND: 扩展
	// UPLOAD: 上传
	AudienceType qianchuan.AudienceType `json:"audience_type,omitempty"`
	// AudienceGroup 人群分组
	AudienceGroup string `json:"audience_group,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// PushProduct 推送产品
	PushProduct string `json:"push_product,omitempty"`
	// StockNum 库存数量，品牌计划才有返回值
	StockNum int64 `json:"stock_num,omitempty"`
	// StockStatus 库存状态，品牌计划才有返回值
	// 0：无库存信息
	// 1：可定库存
	// 2：不可定库存（可更新）
	// 3：不可定库存（不可更新）
	StockStatus int `json:"stock_status,omitempty"`
}
