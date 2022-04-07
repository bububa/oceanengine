package spi

// ChallengeResponse .
type ChallengeResponse struct {
	BaseResp BaseResponse `json:"BaseResp"`
	// 回传challenge数据
	Challenge int64 `json:"challenge"`
}
