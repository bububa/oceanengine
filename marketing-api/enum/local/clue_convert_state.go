package local

// ClueConvertState 线索事件状态
type ClueConvertState string

const (
	//   ClueConvertState_ARRIVAL 顾客到店/销售人员成功上门
	ClueConvertState_ARRIVAL ClueConvertState = "ARRIVAL"
	// ClueConvertState_CLUE_CONFIRM 顾客表达有意向
	ClueConvertState_CLUE_CONFIRM ClueConvertState = "CLUE_CONFIRM"
	// ClueConvertState_CLUE_HIGH_INTENTION 定金或钩子品支付
	ClueConvertState_CLUE_HIGH_INTENTION ClueConvertState = "CLUE_HIGH_INTENTION"
	// ClueConvertState_CONVERSION_CLASS 正价支付
	ClueConvertState_CONVERSION_CLASS ClueConvertState = "CONVERSION_CLASS"
	// ClueConvertState_INVALID_EVENT 无效
	ClueConvertState_INVALID_EVENT ClueConvertState = "INVALID_EVENT"
)
