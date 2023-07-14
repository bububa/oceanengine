package enum

// WechatAuthorizationStatus 授权状态
type WechatAuthorizationStatus string

const (
	// WechatAuthorizationStatus_AUTHORIZED 已授权
	WechatAuthorizationStatus_AUTHORIZED WechatAuthorizationStatus = "AUTHORIZED"
	// WechatAuthorizationStatus_UNAUTHORIZED 未授权
	WechatAuthorizationStatus_UNAUTHORIZED WechatAuthorizationStatus = "UNAUTHORIZED"
	// WechatAuthorizationStatus_AUTHORIZATION_FAILED 授权失败
	WechatAuthorizationStatus_AUTHORIZATION_FAILED WechatAuthorizationStatus = "AUTHORIZATION_FAILED"
)
