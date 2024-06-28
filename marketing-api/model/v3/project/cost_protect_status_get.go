package project

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CostProtectStatusGetRequest 批量获取项目成本保障状态 API Request
type CostProtectStatusGetRequest struct {
	// AdvertiserID 巨量广告平台广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectIDs 项目id列表，每次最多传入50个
	// 注意：仅允许传入广告账户下存在的project_id，否则应答参数中compensate_status会返回UNSUPPORTED表示查询失败，因为该项目在advertiser_id下不存在
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r CostProtectStatusGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("project_ids", string(util.JSONMarshal(r.ProjectIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CostProtectStatusGetResponse 批量获取项目成本保障状态 API Response
type CostProtectStatusGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CostProtectStatusGetResult `json:"data,omitempty"`
}

type CostProtectStatusGetResult struct {
	// CompensationStatusInfoList 项目成本保障信息列表
	CompensationStatusInfoList []CompensationStatusInfo `json:"compensation_status_info_list,omitempty"`
}

// CompensationStatusInfo 项目成本保障信息
type CompensationStatusInfo struct {
	// ProjectID 项目id
	ProjectID uint64 `json:"project_id,omitempty"`
	// CompensateStatus 项目成本保障状态 可选值:
	// CONFIRMING 成本保障确认中
	// ENDED 成本保障已结束
	// INVALID 成本保障已失效
	// IN_EFFECT 成本保障生效中
	// PAID 成本保障已到账
	// DEFAULT 无成本保障状态，表示该项目不支持成本保障，此时在巨量广告升级版平台页面上也不会展示相关标志
	// 其他异常情况返回：
	// FAILED 成本保障查询超时，请重试
	// UNSUPPORTED 查询失败，传入参数project信息有误，多为project_id在账户下不存在
	CompensateStatus enum.CompensateStatus `json:"compensate_status,omitempty"`
	// CompensateInvalidReasons 成本保障失效原因，仅当compensate_status = INVALID 时&&仅在以下情况会返回。原因可能有多个，因此可能会返回多个枚举
	// AUD_CHANGES单日修改定向超过2次,
	// BID_CHANGES单日修改出价超过2次
	// ROI_CHANGES单日修改roi_goal超过2次
	// ANTI_SPAM命中反作弊
	// BID_TYPE_EXPIRED选择的出价产品暂不支持成本保障
	// MANUAL_JUDGE_SPAM有异常的作弊行为
	// AUD_BID_CHANGES单日修改定向/出价超过2次
	// AUD_ROI_CHANGES单日修改定向/ROI目标超过2次
	// ACCOUNT_TRANSFER_APPLICATION 申请转户
	CompensateInvalidReasons []enum.CompensateInvalidReason `json:"compensate_invalid_reasons,omitempty"`
	// CompensateEndReasons 成本保障结束原因，仅当compensate_status = ENDED 时&&仅在以下情况会返回
	// 原因可能有多个，因此可能会返回多个枚举
	// UN_OBERCOST 超成本比例没有达到1.2倍
	// ROI_REAL_EXPECTED实际roi大于目标roi的80%
	// CONVERT_UNDER_THRESHOLD转化数没有达到门槛
	// CURRENCY_PRECISION赔付金额小于0.01元
	CompensateEndReasons []enum.CompensateEndReason `json:"compensate_end_reasons,omitempty"`
	// CompensateAmount 赔付金额（元），最多2位小数
	CompensateAmount float64 `json:"compensate_amount,omitempty"`
	// CompensateURL 赔付规则链接，主要为飞书文档地址，如：https://bytedance.larkoffice.com/docx/UWifd88wfou3bmxHAq7ciOgPn9e
	CompensateURL string `json:"compensate_url,omitempty"`
}
