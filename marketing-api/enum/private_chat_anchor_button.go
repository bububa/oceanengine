package enum

// PrivateChatAnchorButton 按钮文案，可选私信商家、立即咨询、咨询顾问、咨询设计师、问问老师
type PrivateChatAnchorButton string

const (
	// PrivateChatAnchorButton_PRIVATE_MESSAGE 私信商家
	PrivateChatAnchorButton_PRIVATE_MESSAGE PrivateChatAnchorButton = "PRIVATE_MESSAGE"
	// PrivateChatAnchorButton_CONSULT_NOW 立即咨询
	PrivateChatAnchorButton_CONSULT_NOW PrivateChatAnchorButton = "CONSULT_NOW"
	// PrivateChatAnchorButton_CONSULT_ADVISOR 咨询顾问
	PrivateChatAnchorButton_CONSULT_ADVISOR PrivateChatAnchorButton = "CONSULT_ADVISOR"
	// PrivateChatAnchorButton_CONSULT_DESIGNER 咨询设计师
	PrivateChatAnchorButton_CONSULT_DESIGNER PrivateChatAnchorButton = "CONSULT_DESIGNER"
	// PrivateChatAnchorButton_ASK_TEACHER 问问老师
	PrivateChatAnchorButton_ASK_TEACHER PrivateChatAnchorButton = "ASK_TEACHER"
)
