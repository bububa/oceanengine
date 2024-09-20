package agent

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryRiskPromotionListRequest 【代理商】查询广告违规信息 API Request
type QueryRiskPromotionListRequest struct {
	// AgentID 代理商账户ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// BusinessType 业务线
	// AD 巨量广告（默认值），获取巨量广告账户下的违规广告信息
	BusinessType string `json:"business_type,omitempty"`
	// StartDate 推送开始时间，比如：2024-03-01
	StartDate string `json:"start_date,omitempty"`
	// EndDate 推送结束时间，比如：2024-03-01（最长跨度31天）
	EndDate string `json:"end_date,omitempty"`
	// Cursor 页码游标值，初始从Long.MAX开始，后续传入返回的cursor值，不传值相当于page=1，查询count条数据
	Cursor int `json:"cursor,omitempty"`
	// Count 页码游标值，最大支持500
	Count int `json:"count,omitempty"`
	// Filtering 过滤器
	Filtering *QueryRiskPromotionListFilter `json:"filtering,omitempty"`
}

type QueryRiskPromotionListFilter struct {
	//   PromotionStatus 广告状态 可选值:
	// ADV_OFFLINE_BUDGET adv账户超出预算
	// ADV_PRE_OFFLINE_BUDGET adv账户接近预算
	// AUDIT 广告新建审核中
	// AWEME_ACCOUNT_DISABLED 关联抖音账号不可投
	// AWEME_ANCHOR_DISABLED 关联锚点不可投
	// CREATE 广告新建
	// DELETE 已删除
	// DELIVERY_OK 投放中
	// ERROR_DEFAULT 补偿态
	// FROZEN 已终止
	// LIVE_ROOM_OFF 关联直播间不可投
	// NO_SCHEDULE 不在投放时间段内
	// OFFLINE_AUDIT 广告审核不通过
	// PRE_ONLINE 预上线（目前推广管理不披露，仅quota计算）
	// PRODUCT_OFFLINE 关联商品不可投
	// PROJECT_DISABLE 已被项目暂停
	// PROJECT_OFFLINE_BUDGET 项目超出预算
	// PROJECT_PRE_OFFLINE_BUDGET 项目接近预算
	// PROMOTION_DISABLE 广告暂停
	// PROMOTION_OFFLINE_BALANCE 广告余额不足
	// PROMOTION_OFFLINE_BUDGET 广告超出预算
	// PROMOTION_PRE_OFFLINE_BUDGET 广告接近预算
	// PROMOTION_QUOTA_LIMIT 因为quota限额暂停
	// RE_AUDIT 广告重新送审
	// TIME_DONE 已完成
	// TIME_NO_REACH 未达到投放时间
	PromotionStatus enum.PromotionStatus `json:"promotion_status,omitempty"`
	// IllegalMaterialIDs 违规素材ids，最多支持100个
	IllegalMaterialIDs []uint64 `json:"illegal_material_ids,omitempty"`
	// AdvertiserIDs 广告主账户ID，最多支持100个
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// AdvertiserName 广告主账户名称，模糊搜索，长度不能超过30
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// FinalOperatorTag 自走收综合标签（T+1后数据稳定） 可选值:
	// DECREASE_QUANTITY 走量
	// EMPTY 无标签
	// INCREASE_QUANTITY 收量
	// SELF_OPERATION 自运营
	FinalOperatorTag enum.AdvertiserReportType `json:"final_operator_tag,omitempty"`
}

// Encode implements GetRequest interface
func (r QueryRiskPromotionListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	values.Set("business_type", r.BusinessType)
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	if r.Cursor > 0 {
		values.Set("cursor", strconv.Itoa(r.Cursor))
	}
	if r.Count > 0 {
		values.Set("count", strconv.Itoa(r.Count))
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QueryRiskPromotionListResponse 查询广告违规信息 API Response
type QueryRiskPromotionListResponse struct {
	model.BaseResponse
	Data *QueryRiskPromotionListResult `json:"data,omitempty"`
}

type QueryRiskPromotionListResult struct {
	// CursorInfo 分页信息
	CursorInfo *model.PageInfo `json:"cursor_info,omitempty"`
	// Data 违规广告列表
	Data []RiskPromotion `json:"data,omitempty"`
}
