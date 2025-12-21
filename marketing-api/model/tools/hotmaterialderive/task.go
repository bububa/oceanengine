package hotmaterialderive

// Task 任务
type Task struct {
	// TaskID 任务 ID
	TaskID uint64 `json:"task_id,omitempty"`
	// MaterialID 原爆款素材 ID
	MaterialID uint64 `json:"material_id,omitempty"`
	// OriginMaterialID 原爆款素材 ID
	OriginMaterialID uint64 `json:"origin_material_id,omitempty"`
	// DeriveMaterials 爆款素材信息
	DeriveMaterials []DeriveMaterial `json:"derive_materials,omitempty"`
	// CreateTime 任务创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ModifyTime 任务更新时间
	ModifyTime string `json:"modify_time,omitempty"`
	// Status 任务状态，枚举值
	// RENDERING   执行中
	// RENDER_SUCCESS 执行成功
	// RENDER_FAILED 执行失败
	Status string `json:"status,omitempty"`
	// StatusCode 任务返回码，详见 【应答 status_code 和 message】
	StatusCode int `json:"status_code,omitempty"`
	// StatusMessage 任务返回信息，详见 【应答 status_code 和 message】
	StatusMessage string `json:"status_message,omitempty"`
}

// DeriveMaterial 爆款素材信息
type DeriveMaterial struct {
	// VideoID 视频 ID
	VideoID string `json:"video_id,omitempty"`
	// VideoURL 预览链接，存在有效期 7 天
	VideoURL string `json:"video_url,omitempty"`
	// Title 视频标题
	Title string `json:"title,omitempty"`
	// StrategyDetail 策略详情
	StrategyDetail *StrategyDetail `json:"strategy_detail,omitempty"`
}

// StrategyDetail 策略详情
type StrategyDetail struct {
	// Strategy 策略 可选值:
	// CLIP_REPLACE
	// DERIVE_FROM_CHOSEN_HOT_MID
	// DERIVE_FROM_INDUSTRY_HOT_PATTERN
	// HOT_PRE_VIDEO
	// MIX_CUT
	// PRE_VIDEO_CLIP_REPLACE
	// ROBOT_REPLACE
	Strategy string `json:"strategy,omitempty"`
	// StrategyName 策略名称
	StrategyName string `json:"strategy_name,omitempty"`
	// StrategyDescription 策略描述
	StrategyDescription string `json:"strategy_description,omitempty"`
	// ApplyTimes 策略生效的视频时间戳
	ApplyTimes []ApplyTime `json:"apply_times,omitempty"`
}

// ApplyTime 策略生效的视频时间戳
type ApplyTime struct {
	// StartTime 策略生效开始时间戳（毫秒级别），以下生效策略不返回
	// DERIVE_FROM_CHOSEN_HOT_MID - 自有爆款套路
	// DERIVE_FROM_INDUSTRY_HOT_PATTERN - 行业爆款套路
	StartTime int64 `json:"start_time,omitempty"`
	// EndTime 策略生效结束时间戳（毫秒级别），以下生效策略不返回
	// DERIVE_FROM_CHOSEN_HOT_MID - 自有爆款套路
	// DERIVE_FROM_INDUSTRY_HOT_PATTERN - 行业爆款套路
	EndTime int64 `json:"end_time,omitempty"`
}
