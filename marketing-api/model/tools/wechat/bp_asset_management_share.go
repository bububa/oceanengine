package wechat

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BpAssetManagementShareRequest 设置微信小游戏/小程序共享 API Request
type BpAssetManagementShareRequest struct {
	// OrganizationID 组织ID
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// InstanceID 资产ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// AssetType 资产类型 可选值:
	// APPLETS 微信小程序
	// WECHAT_GAME 微信小游戏
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// ShareMode 共享类型 可选值:
	// BP_ALL_ACCOUNTS 组织内所有账户共享
	// COMPANY_ALL_ACCOUNTS 公司主体内所有账户共享
	// PART 指定账户共享
	ShareMode ShareMode `json:"share_mode,omitempty"`
	// AllAccountsByBp 共享账号类别，当share_mode 为BP_ALL_ACCOUNTS时有效且必填，一次最多操作1个，枚举值：AD 广告
	AllAccountsByBp []string `json:"all_accounts_by_bp,omitempty"`
	// AccountInfos 共享账号信息，当share_mode= PART时有效且必填，一次共享最多支持100个账号信息
	AccountInfos []ShareAccountInfo `json:"account_infos,omitempty"`
	// AllAccountsByCompany 公司主体信息，一次最多操作1个
	// 当share_mode为 COMPANY_ALL_ACCOUNTS时有效且必填
	AllAccountsByCompany []ShareAccountInfo `json:"all_accounts_by_company,omitempty"`
}

// Encode implement PostRequest interface
func (r BpAssetManagementShareRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BpAssetManagementShareResponse 设置微信小游戏/小程序共享 API Response
type BpAssetManagementShareResponse struct {
	model.BaseResponse
	Data struct {
		// ErrorList 共享失败列表，列表为空则代表全部共享成功
		ErrorList []BpAssetManagementShareError `json:"error_list,omitempty"`
	} `json:"data,omitempty"`
}

// BpAssetManagementShareError 共享失败
type BpAssetManagementShareError struct {
	// ErrorMessage 共享失败原因
	ErrorMessage string `json:"error_message,omitempty"`
	ShareAccount
}

// Error implement error interface
func (e BpAssetManagementShareError) Error() string {
	return e.ErrorMessage
}
