package enum

// AbTestConclusions 根据实验结论过滤，允许值："FAILED"：不满足实验条件，"TMP_OBVIOUS"：有初步结论。
type AbTestConclusion string

const (
	// AbTestConclusion_NOT_START 未开始
	AbTestConclusion_NOT_START AbTestConclusion = "NOT_START"
	// AbTestConclusion_OBVIOUS 有明显结论
	AbTestConclusion_OBVIOUS AbTestConclusion = "OBVIOUS"
	// AbTestConclusion_INSUFFICIENT 数据量不足
	AbTestConclusion_INSUFFICIENT AbTestConclusion = "INSUFFICIENT"
	// AbTestConclusion_FAILED 不满足实验条件
	AbTestConclusion_FAILED AbTestConclusion = "FAILED"
	// AbTestConclusion_TMP_OBVIOUS 有初步结论
	AbTestConclusion_TMP_OBVIOUS AbTestConclusion = "TMP_OBVIOUS"
)
