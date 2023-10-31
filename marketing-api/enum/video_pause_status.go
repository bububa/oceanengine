package enum

// VideoPauseStatus 暂停结果
type VideoPauseStatus string

const (
	// VideoPauseStatus_SUCCESS 成功
	VideoPauseStatus_SUCCESS VideoPauseStatus = "SUCCESS"
	// VideoPauseStatus_PART_SUCCESS 部分成功
	VideoPauseStatus_PART_SUCCESS VideoPauseStatus = "PART_SUCCESS"
	// VideoPauseStatus_FAIL 失败
	VideoPauseStatus_FAIL VideoPauseStatus = "FAIL"
)
