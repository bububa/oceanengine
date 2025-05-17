package agent

import "github.com/bububa/oceanengine/marketing-api/enum"

// RiskPromotion 违规广告
type RiskPromotion struct {
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// PromotionName 广告名称
	PromotionName string `json:"promotion_name,omitempty"`
	// PromotionStatus 广告状态，同入参promotion_status枚举 可选值:
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
	// MaterialList 违规素材列表，包含广告下投前+投中拒审的素材信息
	MaterialList []RiskMaterial `json:"material_list,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdvertiserName 广告主账户名称
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// CompanyID 广告主公司ID
	CompanyID uint64 `json:"company_id,omitempty"`
	// CompanyName 广告主公司名称
	CompanyName string `json:"company_name,omitempty"`
	// AgentID 代理商账户ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// AgentName 代理商账户名称
	AgentName string `json:"agen_name,omitempty"`
	// AgentCompanyID 代理商公司ID
	AgentCompanyID uint64 `json:"agent_company_id,omitempty"`
	// AgentCompanyName 代理商公司名称
	AgentCompanyName string `json:"agent_company_name,omitempty"`
	// FirstAgentCompanyID 一代代理商公司ID
	FirstAgentCompanyID uint64 `json:"first_agent_company_id,omitempty"`
	// FirstAgentCompanyName 一代代理商公司名称
	FirstAgentCompanyName string `json:"first_agent_company_name,omitempty"`
	// BusinessType 业务线，同入参business_type枚举 可选值:
	// AD 巨量广告
	BusinessType string `json:"business_type,omitempty"`
	// FinalOperatorTag 自走收综合标签，同入参final_operator_tag枚举 可选值:
	// DECREASE_QUANTITY 走量
	// EMPTY 无标签
	// INCREASE_QUANTITY 收量
	// SELF_OPERATION 自运营
	FinalOperatorTag enum.AdvertiserReportType `json:"final_operator_tag,omitempty"`
	// OptimizerID 优化师ID
	OptimizerID uint64 `json:"optimizer_id,omitempty"`
	// OptimizerName 优化师姓名
	OptimizerName string `json:"optimizer_name,omitempty"`
	// CollaboratorIDs 协作者ID
	CollaboratorIDs []uint64 `json:"collaborator_ids,omitempty"`
	// CollaboratorNames 协作者姓名
	CollaboratorNames []string `json:"collaborator_names,omitempty"`
	// SendTime 推送时间
	SendTime string `json:"send_time,omitempty"`
	// ID 记录的唯一ID
	ID uint64 `json:"id,omitempty"`
}

// RiskMaterial 违规素材
type RiskMaterial struct {
	// ID 素材ID（落地页站点ID）
	ID uint64 `json:"id,omitempty"`
	// Type 素材类型 可选值:
	// IMAGE 图片
	// VIDEO 视频
	// SITE 落地页
	Type string `json:"type,omitempty"`
	// RiskContent 素材违规原因，比如：["违规内容1", "违规内容2"]
	RiskContent []string `json:"risk_content,omitempty"`
	// RefPromotionIDs 同代理商账户下的其他关联广告ID
	RefPromotionIDs []uint64 `json:"ref_promotion_ids,omitempty"`
	// RefPromotionList 同代理商账户下的其他关联广告信息，只披露近七天有投放消耗的关联广告
	RefPromotionList []RiskPromotion `json:"ref_promotion_list,omitempty"`
}
