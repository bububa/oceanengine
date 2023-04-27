package enum

// StrategyType 策略类型
type StrategyType string

const (
	// StrategyType_Horizontal2Horizontal 横转横
	StrategyType_Horizontal2Horizontal StrategyType = "Horizontal2Horizontal"
	// StrategyType_Horizontal2Vertical 横转竖
	StrategyType_Horizontal2Vertical StrategyType = "Horizontal2Vertical"
	// StrategyType_Vertical2Horizontal 竖转横
	StrategyType_Vertical2Horizontal StrategyType = "Vertical2Horizontal"
	// StrategyType_Vertical2Vertical 竖转竖
	StrategyType_Vertical2Vertical StrategyType = "Vertical2Vertical"
)

// StrategyStateType 配置项类型
type StrategyStateType string

const (
	// StrategyStateType_Image 图片
	StrategyStateType_Image StrategyStateType = "Image"
	// StrategyStateType_Text 文案
	StrategyStateType_Text StrategyStateType = "Text"
)
