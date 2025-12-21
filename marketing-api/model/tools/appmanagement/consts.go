package appmanagement

// AppSearchType 应用搜索类型
type AppSearchType string

const (
	// ALL:查询全部，包括创建和被共享的应用（默认）
	AppSearchType_ALL AppSearchType = "ALL"
	// CREATE_ONLY:只查询广告主创建的应用
	AppSearchType_CREATE_ONLY AppSearchType = "CREATE_ONLY"
	// SHARED_ONLY:只查询被共享的应用
	AppSearchType_SHARED_ONLY AppSearchType = "SHARED_ONLY"
)

// AppStatus 应用状态
type AppStatus string

const (
	// Status_ALL 所有状态
	AppStatus_ALL AppStatus = "ALL"
	// AppStatus_AUDIT_DOING 审核中
	AppStatus_AUDIT_DOING AppStatus = "AUDIT_DOING"
	// AppStatus_AUDIT_REJECTED 审核失败
	AppStatus_AUDIT_REJECTED AppStatus = "AUDIT_REJECTED"
	// AppStatus_AUDIT_ACCEPTED 审核成功
	AppStatus_AUDIT_ACCEPTED AppStatus = "AUDIT_ACCEPTED"
	// AppStatus_BOOKING 预约中
	AppStatus_BOOKING AppStatus = "BOOKING"
	// AppStatus_ENABLE 已发布（默认）
	AppStatus_ENABLE AppStatus = "ENABLE"
	// AppStatus_PAST_DUE 已逾期
	AppStatus_PAST_DUE AppStatus = "PAST_DUE"
)

// ExtendPackageStatus 分包状态
type ExtendPackageStatus string

const (
	// ExtendPackageStatus_ALL 全部
	ExtendPackageStatus_ALL ExtendPackageStatus = "ALL"
	// ExtendPackageStatus_NOT_UPDATE 未更新
	ExtendPackageStatus_NOT_UPDATE ExtendPackageStatus = "NOT_UPDATE"
	// ExtendPackageStatus_CREATING 创建中
	ExtendPackageStatus_CREATING ExtendPackageStatus = "CREATING"
	// ExtendPackageStatus_UPDATING 更新中
	ExtendPackageStatus_UPDATING ExtendPackageStatus = "UPDATING"
	// ExtendPackageStatus_PUBLISHED 已发布
	ExtendPackageStatus_PUBLISHED ExtendPackageStatus = "PUBLISHED"
	// ExtendPackageStatus_ENABLE 已发布（默认值）
	ExtendPackageStatus_ENABLE ExtendPackageStatus = "ENABLE"
	// ExtendPackageStatus_CREATION_FAILED 创建失败
	ExtendPackageStatus_CREATION_FAILED ExtendPackageStatus = "CREATION_FAILED"
	// ExtendPackageStatus_UPDATE_FAILED 更新失败
	ExtendPackageStatus_UPDATE_FAILED ExtendPackageStatus = "UPDATE_FAILED"
)

// HostType 应用是否寄存在应用管理中
type HostType string

const (
	// HOST_IN 寄存
	HOST_IN HostType = "HOST_IN"
	// HOST_OUT 未寄存
	HOST_OUT HostType = "HOST_OUT"
)

// DownloadPackageStatus 包解析状态
type DownloadPackageStatus string

const (
	// DownloadPackageStatus_NEW_PACKAGE 新建
	DownloadPackageStatus_NEW_PACKAGE DownloadPackageStatus = "NEW_PACKAGE"
	// DownloadPackageStatus_PARSING 下载中
	DownloadPackageStatus_PARSING DownloadPackageStatus = "PARSING"
	// DownloadPackageStatus_SUCCESS 成功
	DownloadPackageStatus_SUCCESS DownloadPackageStatus = "SUCCESS"
	// DownloadPackageStatus_FAIL 失败
	DownloadPackageStatus_FAIL DownloadPackageStatus = "FAIL"
)

// ShareSearchType 共享类型
type ShareSearchType string

const (
	// ShareSearchType_ORGANIZATION_SHARE 共享的组织范围
	ShareSearchType_ORGANIZATION_SHARE ShareSearchType = "ORGANIZATION_SHARE"
	// ShareSearchType_OTHER 共享的其他账号
	ShareSearchType_OTHER ShareSearchType = "OTHER"
)

// ShareMode 共享类型
type ShareMode string

const (
	// ShareMode_PART 指定账户共享
	ShareMode_PART ShareMode = "PART"
	// ShareMode_ALL 组织内某类账户所有共享
	ShareMode_ALL ShareMode = "ALL"
	// ShareMode_COMPANY 公司主体内某类型所有账户共享
	ShareMode_COMPANY ShareMode = "COMPANY"
)
