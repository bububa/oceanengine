package enum

// 行为场景
type ActionScene string

const (
	E_COMMERCE_SCENE ActionScene = "E-COMMERCE" // 电商行为场景
	NEWS_SCENE       ActionScene = "NEWS"       // 资讯行为场景
	APP_SCENE        ActionScene = "APP"        // app行为场景
)
