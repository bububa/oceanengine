package agent

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdvertiserInfoQueryRequest 广告主账户信息查询 API Request
type AdvertiserInfoQueryRequest struct {
	// AccountIDs 广告主账户id，最多50个
	AccountIDs []uint64 `json:"account_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r AdvertiserInfoQueryRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_ids", string(util.JSONMarshal(r.AccountIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AdvertiserInfoQueryResponse 广告主账户信息查询 API Response
type AdvertiserInfoQueryResponse struct {
	model.BaseResponse
	Data struct {
		// AccountDetailList 广告主账户详情
		AccountDetailList []AccountDetail `json:"account_detail_list,omitempty"`
	} `json:"data,omitempty"`
}

// AccountDetail 广告主账户详情
type AccountDetail struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdvertiserName 广告主账户名称
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// AdvertiserStatus 广告主账户状态，枚举值：
	// STATUS_CONFIRM_FAIL 审核失败/可再次申请
	// STATUS_CONFIRM_FAIL_END 审核失败/最终状态
	// STATUS_CONFIRM_MODIFY_FAIL 修改审核失败
	// STATUS_DISABLE 已禁用
	// STATUS_ENABLE 已审核
	// STATUS_LIMIT 限制
	// STATUS_PENDING_CONFIRM 待审核
	// STATUS_PENDING_CONFIRM_MODIFY 修改审核
	// STATUS_PENDING_VERIFIED 待验证
	// STATUS_SELF_SERVICE_UNAUDITED 自助开户待验证资质
	// STATUS_WAIT_FOR_BPM_AUDIT 等待CRM审核
	// STATUS_WAIT_FOR_PUBLIC_AUTH 待对公验证
	AdvertiserStatus enum.AdvertiserStatus `json:"advertiser_status,omitempty"`
	// FirstIndustryName 一级行业名称
	FirstIndustryName string `json:"first_industry_name,omitempty"`
	// SecondIndustryName 二级行业名称
	SecondIndustryName string `json:"second_industry_name,omitempty"`
	// CreateTime 广告主账户创建时间
	CreateTime string `json:"create_time,omitempty"`
	// BindTime 广告主账户与代理商账户的绑定时间
	BindTime string `json:"bind_time,omitempty"`
	// SelfOperationTag 账户报备类型，枚举值：
	// DECREASE_QUANTITY 走量
	// EMPTY 无标签
	// INCREASE_QUANTITY 收量
	// SELF_OPERATION 自运营
	SelfOperationTag enum.AdvertiserReportType `json:"self_operation_tag,omitempty"`
	// FirstAgentCompanyName 一代代理商公司名称
	FirstAgentCompanyName string `json:"first_agent_company_name,omitempty"`
	// FirstAgentCompantyID 一代代理商公司ID
	FirstAgentCompantyID uint64 `json:"first_agent_companty_id,omitempty"`
	// FirstAgentName 一级代理商账户名称
	FirstAgentName string `json:"first_agent_name,omitempty"`
	// FirstAgentID 一级代理商账户ID
	FirstAgentID uint64 `json:"first_agent_id,omitempty"`
	// AdvCompanyName 广告主公司名称
	AdvCompanyName string `json:"adv_company_name,omitempty"`
	// AdvCompanyID 广告主公司ID
	AdvCompanyID uint64 `json:"adv_company_id,omitempty"`
	// CustomerSaleName 直客销售名称
	CustomerSaleName string `json:"customer_sale_name,omitempty"`
	// AuthExpireDate 对公认证过期时间
	AuthExpireDate string `json:"auth_expire_date,omitempty"`
	// ContactName 联系人姓名
	ContactName string `json:"contact_name,omitempty"`
	// OptimizerID 竞价权限负责人id
	OptimizerID uint64 `json:"optimizer_id,omitempty"`
	// OptimizerName 竞价权限负责人姓名
	OptimizerName string `json:"optimizer_name,omitempty"`
	// Collaborators
	Collaborators []Collaborator `json:"collaborators,omitempty"`
}

type Collaborator struct {
	// BrandOptimizerID 品牌权限负责人id
	BrandOptimizerID uint64 `json:"brand_optimizer_id,omitempty"`
	// BrandOptimizerName 品牌权限负责人姓名     a
	BrndOptimizerName string `json:"brand_optimizer_name,omitempty"`
	// BrandCollaborators
	BrandCollaborators []BrandCollaborators `json:"brand_collaborators,omitempty"`
}

type BrandCollaborators struct {
	// EmployeID 品牌权限协作者id
	EmployeID uint64 `json:"employe_id,omitempty"`
	// EmployeName 品牌权限协作者姓名
	EmployeeName string `json:"employee_name,omitempty"`
}
