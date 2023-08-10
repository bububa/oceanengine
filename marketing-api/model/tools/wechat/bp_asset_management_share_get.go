package wechat

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BpAssetManagementShareGetRequest 查看微信小游戏/小程序共享范围 API Request
type BpAssetManagementShareGetRequest struct {
	// OrganizationID 组织ID
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// InstanceID 资产ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// AssetType 资产类型 可选值:
	// APPLETS 微信小程序
	// WECHAT_GAME 微信小游戏
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	//	ShareType  可选值:
	// ACCOUNT 共享给账户
	// GROUP 共享给组织
	ShareType ShareType `json:"share_type,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值为10，最大值200
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r BpAssetManagementShareGetRequest) Encode() string {
	values := util.GetUrlValues()
	if r.OrganizationID > 0 {
		values.Set("organization_id", strconv.FormatUint(r.OrganizationID, 10))
	}
	values.Set("asset_type", string(r.AssetType))
	values.Set("instance_id", strconv.FormatUint(r.InstanceID, 10))
	if r.ShareType != "" {
		values.Set("share_type", string(r.ShareType))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// BpAssetManagementShareGetResponse 查看微信小游戏/小程序共享范围 API Response
type BpAssetManagementShareGetResponse struct {
	model.BaseResponse
	Data *BpAssetManagementShareList `json:"data,omitempty"`
}

type BpAssetManagementShareList struct {
	// SharedAccounts 共享账号列表
	SharedAccounts []ShareAccountInfo `json:"shared_accounts,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// ShareAccountInfo 共享账号信息
type ShareAccountInfo struct {
	// AccountType 业务线，枚举值：AD 广告、BP 巨量纵横
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// AccountID 账户ID
	// 共享到巨量纵横账号时仅支持合作组织账号，需通过巨量引擎工作台>>设置>>合作组织管理>>添加合作组织进行管理；
	// 且仅创建在纵横组织账号下的微信小游戏资产支持共享到其他纵横组织
	AccountID uint64 `json:"account_id,omitempty"`
	// CompanyID 公司主体ID，一次最多操作1个
	CompanyID uint64 `json:"company_id,omitempty"`
}

// ShareAccount 共享账号
type ShareAccount struct {
	// ShareMode 共享类型
	// 允许值：PART 指定账户共享、BP_ALL_ACCOUNTS组织内所有账户共享、COMPANY_ALL_ACCOUNTS 公司主体内所有账户共享
	// 共享类型为PART，读取account_infos；
	// 共享类型为BP_ALL_ACCOUNTS，读取all_account_by_bp，即表示该业务线下的全量账号；共享类型为COMPANY_ALL_ACCOUNTS，读取all_account_by_company
	ShareMode ShareMode `json:"share_mode,omitempty"`
	// AccountInfo 共享账号信息
	AccountInfo *ShareAccountInfo `json:"account_info,omitempty"`
	// AccountsByBp 组织内业务线 可选值:
	// AD 广告
	AccountsByBp string `json:"accounts_by_bp,omitempty"`
	// AllAccountsByCompany 主体信息
	AllAccountsByCompany *ShareAccountInfo `json:"all_accounts_by_company,omitempty"`
}
