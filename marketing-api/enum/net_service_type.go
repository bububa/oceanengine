package enum

// NetServiceType 推广内容
type NetServiceType string

const (
	// NetServiceType_GENERAL 常规应用下载
	NetServiceType_GENERAL NetServiceType = "GENERAL"
	// NetServiceType_WECHAT_PACKAGE 跳转场景=添加微信号
	NetServiceType_WECHAT_PACKAGE NetServiceType = "WECHAT_PACKAGE"
	// NetServiceType_MICRO_APP 跳转场景=进入微信小程序
	NetServiceType_MICRO_APP NetServiceType = "MICRO_APP"
	// NetServiceType_WECOM_PACKAGE 跳转场景=企业微信客服
	NetServiceType_WECOM_PACKAGE NetServiceType = "WECOM_PACKAGE"
	// NetServiceType_QUICK_APP 快应用
	NetServiceType_QUICK_APP NetServiceType = "QUICK_APP"
	// NetServiceType_WECHAT_EXTERNAL_URL 跳转场景= 跳转微信链接
	NetServiceType_WECHAT_EXTERNAL_URL NetServiceType = "WECHAT_EXTERNAL_URL"
)
