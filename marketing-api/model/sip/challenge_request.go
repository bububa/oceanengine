package sip

// ChallengeRequest
type ChallengeRequest struct {
	// Challenge 随机数
	Challenge int64 `form:"challenge" binding:"required"`
	// Event 固定值，传递verify_webhook
	Event string `form:"event" binding:"required"`
}
