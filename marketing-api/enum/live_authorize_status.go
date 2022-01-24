package enum

// LiveAuthorizeStatus 直播达人授权状态
type LiveAuthorizeStatus string

const (
	// LiveAuthorizeStatus_APPLYING 申请中
	LiveAuthorizeStatus_APPLYING LiveAuthorizeStatus = "APPLYING"
	// LiveAuthorizeStatus_APPLY_OVERDUE 申请过期
	LiveAuthorizeStatus_APPLY_OVERDUE LiveAuthorizeStatus = "APPLY_OVERDUE"
	// LiveAuthorizeStatus_AUTHORIZING 授权中
	LiveAuthorizeStatus_AUTHORIZING LiveAuthorizeStatus = "AUTHORIZING"
	// LiveAuthorizeStatus_AUTHORIZE_OVERDUE 授权过期
	LiveAuthorizeStatus_AUTHORIZE_OVERDUE LiveAuthorizeStatus = "AUTHORIZE_OVERDUE"
)
