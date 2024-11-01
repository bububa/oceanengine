package security

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ScoreViolationEventGetRequest 查询违规积分明细 API Request
type ScoreViolationEventGetRequest struct {
	// Filtering 过滤器
	Filtering *ScoreViolationEventGetFilter `json:"filtering,omitempty"`
	// BusinessLine 业务线，允许值：
	// AD 巨量广告
	BusinessLine string `json:"business_line,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页数，默认值：1，page范围为[1,99999]
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值：10，page_size范围为[1,100]
	PageSize int `json:"page_size,omitempty"`
}

type ScoreViolationEventGetFilter struct {
	// Status 生效状态，允许值：
	// APPEAL 已申诉（失效）
	// FAILAPPEAL 申诉失败
	// ONAPPEAL 申诉中
	// VALID 生效
	Status enum.SecurityScoreViolationEventStatus `json:"status,omitempty"`
	// StartTime 开始时间
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间
	EndTime string `json:"end_time,omitempty"`
	// IllegalType 违规类型
	// GENERAL 一般违规（AD）
	// SERIOUS 严重违规（AD）
	IllegalType enum.SecurityScoreIllegalType `json:"illegal_type,omitempty"`
	// EventID 积分违规单
	EventID uint64 `json:"event_id,omitempty"`
}

func (r ScoreViolationEventGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("page", strconv.Itoa(r.Page))
	values.Set("page_size", strconv.Itoa(r.PageSize))
	if r.BusinessLine != "" {
		values.Set("business_line", r.BusinessLine)
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ScoreViolationEventGetResponse 查询违规积分明细 API Response
type ScoreViolationEventGetResponse struct {
	Data *ScoreViolationEventGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type ScoreViolationEventGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AdvScoreEvent
	AdvScoreEvent []ScoreEvent `json:"adv_score_event,omitempty"`
}

type ScoreEvent struct {
	// MaterialID 素材id
	MaterialID string `json:"material_id,omitempty"`
	// ViolationEvidenceImg 违规证据截图
	ViolationEvidenceImg string `json:"violation_evidence_img,omitempty"`
	// RejectReason 拒绝理由
	RejectReason string `json:"reject_reason,omitempty"`
	// CreateTIme 创建时间
	CreateTIme string `json:"create_time,omitempty"`
	// Status 状态，枚举值：
	// APPEAL 已申诉（失效）
	// FAILAPPEAL 申诉失败
	// ONAPPEAL 申诉中
	// VALID 生效
	Status enum.SecurityScoreViolationEventStatus `json:"status,omitempty"`
	// IllegalType 违规类型
	IllegalType enum.SecurityScoreIllegalType `json:"illegal_type,omitempty"`
	// EventID 违规单id
	EventID uint64 `json:"event_id,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划id（AD2.0为广告id）
	AdID uint64 `json:"ad_id,omitempty"`
	// Score 扣罚分值
	Score int64 `json:"score,omitempty"`
}
