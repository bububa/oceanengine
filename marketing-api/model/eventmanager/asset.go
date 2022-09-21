package eventmanager

// AssetBaseInfo 资产基本数据``
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
	// PackageName 快应用包名
	PackageName string `json:"package_name,omitempty"`
}

// App 应用数据
type App struct {
	AssetBaseInfo
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

// MiniProgram 字节小程序快应用资产
type MiniProgram struct {
	AssetBaseInfo
	// MiniProgramID 字节小程序id
	MiniProgramID string `json:"mini_program_id,omitempty"`
}
