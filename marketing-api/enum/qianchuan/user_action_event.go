package qianchuan

// UserActionEvent  用户来源
type UserActionEvent string

const (
	// UserActionEvent_ENTER 进入直播间
	UserActionEvent_ENTER UserActionEvent = "ENTER"
	// UserActionEvent_PAY 支付成功
	UserActionEvent_PAY UserActionEvent = "PAY"
)
