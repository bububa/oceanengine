package spi

// ChallengeRequest .
type ChallengeRequest struct {
	// Challenge 随机数
	Challenge int64 `form:"challenge" json:"challenge" binding:"required"`
	// Event 固定值，传递verify_webhook
	Event string `form:"event" json:"event" binding:"required"`
}
