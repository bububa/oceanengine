package appmanagement

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AndroidBasicPackageUpdateRequest 更新安卓应用母包 API Request
type AndroidBasicPackageUpdateRequest struct {
	// AccountID 账户id指可以接的账号体系如广告主id、巨量纵横组织id等
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账号类型，允许值：
	// AD 广告主类型、BP 巨量纵横账号类型
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// PackageID 应用包id
	PackageID string `json:"package_id,omitempty"`
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
	// PermissionsDescription 隐私权限说明
	PermissionsDescription string `json:"permissions_description,omitempty"`
	// UpdateDescription 版本更新说明，最长200字符
	UpdateDescription string `json:"update_description,omitempty"`
	// Recommend 推荐语，最长30字符
	Recommend string `json:"recommend,omitempty"`
	// AppDescription 应用介绍，最长500字符
	AppDescription string `json:"app_description,omitempty"`
	// PaymentType 付费类型，允许值：
	// FREE 免费下载，且不包含付费内容
	// PRODUCT 免费下载，但应用内包含付费内容，如道具等虚拟物品
	// TRIAL_AND_LIMIT_FUNCTIONS 付费使用完整功能，且应用内包含付费内容，如道具等虚拟物品
	// TRIAL_AND_PURCHASE 付费使用完整功能，但应用内不包含任何付费内容
	PaymentType enum.AppPaymentType `json:"payment_type,omitempty"`
	// AppDeveloperName 应用开发者名称，最长100字符
	AppDeveloperName string `json:"app_developer_name,omitempty"`
	// AutoPublish 是否自动发布
	AutoPublish bool `json:"auto_publish,omitempty"`
	// IndustryID 分类id， 通过 【获取应用细分分类及题材标签】 接口获取
	IndustryID string `json:"industry_id,omitempty"`
	// ThemeTagID 应用题材标签id ，通过 【获取应用细分分类及题材标签】 接口获取，仅游戏行业应用填写且必填
	ThemeTagID string `json:"theme_tag_id,omitempty"`
	// FileOption 文件选项，允许值：
	// USE_LAST_IMAGE_VIDEO 使用上一次的图片、视频、USE_NEW 使用新图片、视频
	FileOption string `json:"file_option,omitempty"`
	// Files 文件内容
	// 当file_option为USE_NEW时，必须传入一个APK APK包文件，且file_tag文件标识为DEFAULT 默认
	// 非游戏行业应用
	// 图片文件：1≤图片文件数量 ≤ 5，files_tag文件标识为MATERIAL_SCREENSHOT 素材截图
	// 视频文件：视频文件数量 ≤ 1，文件标识为DEFAULT 默认
	// 游戏行业应用
	// 图片文件：
	// 必填图片：AGE_REMINDER 适龄提醒、ANTI_ADDICTION_TIPS 防沉迷提示、REAL_NAME_VERIFIED 实名注册，三类图片文件（每类图片各一张）
	// 选填图片：素材截图文件数量 ≤ 5（文件标识选择MATERIAL_SCREENSHOT 素材截图）
	// 视频文件：视频文件数量 ≤ 1，文件标识为DEFAULT 默认
	Files []File `json:"files,omitempty"`
}

// Encode implement PostRequest interface
func (r AndroidBasicPackageUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
