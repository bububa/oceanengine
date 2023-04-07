package appmanagement

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BpShareRequest 设置应用共享 API Request
type BpShareRequest struct {
	// OrganizationID 巨量纵横组织id（即管家账户）
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// PackageID 应用包ID， 可通过「查询应用信息」接口查询获取
	PackageID uint64 `json:"package_id,omitempty"`
	// ShareMode 共享类型
	// 允许值：PART 指定账户共享、 ALL 组织内某类型所有账户共享、COMPANY 公司主体内某类型所有账户共享
	ShareMode ShareMode `json:"share_mode,omitempty"`
	// AllAccounts 共享账号类别，一次最多操作10个
	// 当share_mode 为ALL时可用且必填
	AllAccounts []ShareAccountInfo `json:"all_accounts,omitempty"`
	// AccountInfos 共享账号信息，一次最多操作10个
	// 当share_mode为 PART时可用且必填
	AccountInfso []ShareAccountInfo `json:"account_infos,omitempty"`
	// AllAccountByCompany 公司主体信息，一次最多操作1个
	// 当share_mode为 COMPANY时可用且必填
	AllAccountByCompany []ShareAccountInfo `json:"all_account_by_company,omitempty"`
}

// Encode implement PostRequest interface
func (r BpShareRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BpShareResponse 设置应用共享 API Response
type BpShareResponse struct {
	model.BaseResponse
	Data *BpShareData `json:"data,omitempty"`
}

type BpShareData struct {
	// SuccessList 成功共享列表
	SuccessList []ShareAccount `json:"success_list,omitempty"`
	// ErrorList 共享失败列表
	ErrorList []ShareAccount `json:"error_list,omitempty"`
}
