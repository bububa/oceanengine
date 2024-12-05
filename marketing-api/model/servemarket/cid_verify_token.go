package servemarket

import "github.com/bububa/oceanengine/marketing-api/model"

// CidVerifyTokenResponse 获取App Access Token校验信息 API Response
type CidVerifyTokenResponse struct {
	Data *CidVerifyTokenResult `json:"data,omitempty"`
	model.BaseResponse
}

type CidVerifyTokenResult struct {
	// AppAccessToken 传入的app_access_token
	AppAccessToken string `json:"app_access_token,omitempty"`
	// DevStatus 开放平台开发者状态：
	// 1 开发者身份待申请
	// 2 主体资质待认证
	// 4 对公验证待认证
	// 5 合同创建中
	// 3 合同审核中
	// 6 合同待签署
	// 7 正常
	// 8 其它
	// 仅“7-正常”状态可校验通过
	DevStatus int `json:"dev_status,omitempty"`
	// DevQualificationName 开放平台开发者资质名称
	DevQualificationName string `json:"dev_qualification_name,omitempty"`
	// FacilitatorSubStatus 群峰服务商状态：
	// 10 主体资质待认证
	// 30 对公验证待认证
	// 40 CA签章待申请
	// 45 合同创建中
	// 50 合同已签署
	// 60 保证金已缴纳
	// 100 其它
	// 仅“60-保证金已缴纳”状态可校验通过
	FacilitatorSubStatus int `json:"facilitator_sub_status,omitempty"`
	// CidCapacityStatus 群峰服务商「电商技术服务」能力开通状态：
	// 1 审核中
	// 2 审核拒绝
	// 3 审核通过
	// 6 服务商退出
	// 100 未知
	// 仅“3-审核通过”状态可校验通过
	CidCapacityStatus int `json:"cid_capacity_status,omitempty"`
	// TokenIsValid app_access_token是否有效，如果为true，则表示，转化数据回传时，携带的app_access_token动态身份标识，可校验通过
	TokenIsValid bool `json:"token_is_valid,omitempty"`
}
