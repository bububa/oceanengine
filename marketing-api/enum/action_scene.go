package enum

// ActionScene 行为场景
type ActionScene string

const (
	// UNKNOWN_SCENE 未知场景
	UNKNOWN_SCENE ActionScene = ""
	// E_COMMERCE_SCENE 电商行为场景
	E_COMMERCE_SCENE ActionScene = "E-COMMERCE"
	// NEWS_SCENE 资讯行为场景
	NEWS_SCENE ActionScene = "NEWS"
	// APP_SCENE app行为场景
	APP_SCENE ActionScene = "APP"
)
