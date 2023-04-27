package appmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ShareAccountListRequest 查看应用共享范围 API Request
type ShareAccountListRequest struct {
	// OrganizationID 巨量纵横组织id（即管家账户）
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// PackageID 应用包ID
	PackageID string `json:"package_id,omitempty"`
	// SearchType 共享类型
	// 允许值：ORGANIZATION_SHARE 共享的组织范围、OTHER 共享的其他账号
	// 默认：ORGANIZATION_SHARE
	SearchType ShareSearchType `json:"search_type,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值为10，最大值200
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r ShareAccountListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("organization_id", strconv.FormatUint(r.OrganizationID, 10))
	values.Set("package_id", r.PackageID)
	if r.SearchType != "" {
		values.Set("search_type", string(r.SearchType))
	}
	if r.Page != 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ShareAccountListResponse 查看应用共享范围 API Response
type ShareAccountListResponse struct {
	model.BaseResponse
	Data *ShareAccountListData `json:"data,omitempty"`
}

type ShareAccountListData struct {
	// List 共享账号列表
	List []ShareAccount `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// ShareAccount 应用共享账户
type ShareAccount struct {
	// ShareMode 共享类型
	// 如果共享类型为PART，读取account_info
	// 如果共享类型为ALL，读取all_account，即表示该业务线下的全量账号
	// 如果共享类型为COMPANY ，读取all_account_by_company，即表示组织下所属主体的全量账号。
	// 枚举值：PART 指定账户共享、 ALL 组织内某类账户所有共享、COMPANY 公司主体内某类型所有账户共享
	ShareMode ShareMode `json:"share_mode,omitempty"`
	// AccountInfo 共享账号信息
	AccountInfo *ShareAccountInfo `json:"account_info,omitempty"`
	// AllAccount 共享账号类别
	AllAccount *ShareAccountInfo `json:"all_account,omitempty"`
	// AllAccountByCompany 主体信息
	AllAccountByCompany *ShareAccountInfo `json:"all_account_by_company,omitempty"`
}

// ShareAccountInfo 共享账号信息
type ShareAccountInfo struct {
	// AccountID 账号id
	AccountID uint64 `json:"account_id,omitempty"`
	// CompanyID 公司id
	CompanyID uint64 `json:"company_id,omitempty"`
	// AccountType 账号类型
	// 枚举值：AD 广告主、STAR 星图、BP 巨量纵横
	AccountType enum.AccountType `json:"account_type,omitempty"`
}
