package eventmanager

import "github.com/bububa/oceanengine/marketing-api/enum"

// AssetBaseInfo 资产基本数据“
type AssetBaseInfo struct {
	// AssetID 快应用资产ID
	AssetID uint64 `json:"asset_id,omitempty"`
	// AssetName 快应用名称
	AssetName string `json:"asset_name,omitempty"`
}

// LandingPage 三方落地页数据
type LandingPage struct {
	AssetBaseInfo
	// Name  落地页名称，长度限制为25，一个字符长度为1
	Name string `json:"name,omitempty"`
	// Description 落地页名称，长度限制为150，一个字符长度为1
	Description string `json:"description,omitempty"`
}

// QuickApp 快应用数据
type QuickApp struct {
	AssetBaseInfo
	// Name 快应用名称，长度限制为20，一个字符长度为1
	Name string `json:"name,omitempty"`
	// PackageName 快应用包名
	PackageName string `json:"package_name,omitempty"`
	// QuickAppID 快应用ID
	QuickAppID uint64 `json:"quick_app_id,omitempty"`
}

// App 应用数据
type App struct {
	AssetBaseInfo
	// Name 快应用名称，长度限制为20，一个字符长度为1
	Name string `json:"name,omitempty"`
	// AppName 应用名
	AppName string `json:"app_name,omitempty"`
	// AppType 应用类型
	AppType string `json:"app_type,omitempty"`
	// DownloadURL 应用下载链接
	DownloadURL string `json:"download_url,omitempty"`
	// AppID 应用ID
	AppID uint64 `json:"app_id,omitempty"`
	// PackageID 母包ID
	PackageID string `json:"package_id,omitempty"`
	// PackageName 应用包名
	PackageName string `json:"package_name,omitempty"`
	// AppPackageName 应用包名，精确搜索
	AppPackageName string `json:"app_package_name,omitempty"`
	// Role 资产来源
	Role string `json:"role,omitempty"`
}

// Site 橙子落地页信息
type Site struct {
	AssetBaseInfo
	// SiteID 橙子建站站点ID，橙子落地页必填
	SiteID uint64 `json:"site_id,omitempty"`
	// SiteName 橙子建站站点名称，橙子落地页必填
	SiteName string `json:"site_name,omitempty"`
}

// MiniProgram 字节小程序快应用资产
type MiniProgram struct {
	AssetBaseInfo
	// MiniProgramID 字节小程序id
	MiniProgramID string `json:"mini_program_id,omitempty"`
	// MicroProgramName 字节小程序的名称，需要与资产id后面信息一致
	MicroProgramName string `json:"micro_program_name,omitempty"`
	// InstanceID 字节小程序资产ID，通过【工具】-【获取字节小程序/小游戏】获取
	InstanceID uint64 `json:"instance_id,omitempty"`
	// MiniProgramType 字节小程序类型：
	// BYTE_APP：字节小程序
	// BYTE_GAME：字节小游戏
	MiniProgramType enum.MiniProgramType `json:"mini_program_type,omitempty"`
}
