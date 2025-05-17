package promotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取广告列表 API Request
type ListRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// Filtering 筛选项
	Filtering *ListFilter `json:"filtering,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，最大值100，默认10
	PageSize int `json:"page_size,omitempty"`
}

type ListFilter struct {
	// PromotionIDs 按广告IDs筛选，单次最多100个
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
	// PromotionName 按广告名称模糊搜索
	PromotionName string `json:"promotion_name,omitempty"`
	// ProjectID 项目ID筛选
	ProjectID uint64 `json:"project_id,omitempty"`
	// PromotionStatusFirst 广告一级状态过滤，默认不限（不包含已删除），允许值：
	// PROMOTION_STATUS_ALL 不限（包含已删除）
	// PROMOTION_STATUS_DELETED 已删除:
	// PROMOTION_STATUS_DISABLE 未投放
	// PROMOTION_STATUS_DONE 已完成
	// PROMOTION_STATUS_ENABLE 投放中
	// PROMOTION_STATUS_FROZEN 已终止
	// PROMOTION_STATUS_NOT_DELETE 不限（不包含已删除）
	// 默认值: PROMOTION_STATUS_NOT_DELETE
	PromotionStatusFirst local.PromotionStatusFirst `json:"promotion_status_first,omitempty"`
	// PromotionStatusSecond 广告二级状态过滤，允许值：
	// 仅当promotion_status_first=PROMOTION_STATUS_DISABLE未投放时传入有效且必填，不传会报错；其他情况下传入该字段无效。
	// AUDIT_DENY 审核不通过
	// AUDIT 新建审核中
	// REAUD/IT 修改审核中
	// DISABLED 已暂停
	// DISABLE_BY_QUOTA 配额达限
	// PROJECT_DISABLED 项目已被暂停
	// NO_SCHEDULE 不在投放时段
	// TIME_NO_REACH 未到达投放时间
	// OFFLINE_BALANCE 账户余额不足
	// BALANCE_OFFLINE_BUDGET 账户超出预算
	// PROJECT_OFFLINE_BUDGET 项目超出预算
	// PROMOTION_OFFLINE_BUDGET 广告超出预算
	// LIVE_ROOM_OFF 直播间不可投放
	// AWEME_ACCOUNT_DISABLED 抖音账号不可投放
	// PRODUCT_OR_POI_OFFLINE 商品/门店不可投
	PromotionStatusSecond local.PromotionStatusSecond `json:"promotion_status_second,omitempty"`
	// AdType 广告类型筛选，默认不限，允许值：
	// ALL 不限
	// GENERAL 通投广告
	// SEARCHING 搜索广告
	// 默认值: ALL
	AdType local.AdType `json:"ad_type,omitempty"`
	// MarketingGoal 营销场景筛选，默认不限，允许值：
	// ALL 不限
	// LIVE 直播
	// VIDEO_IMAGE 短视频/图文
	// 默认值: ALL
	MarketingGoal local.MarketingGoal `json:"marketing_goal,omitempty"`
	// PromotionCreateTimeStart 广告创建开始时间，格式yyyy-MM-dd HH:mm:ss与time_end搭配使用
	PromotionCreateTimeStart string `json:"promotion_create_time_start,omitempty"`
	// PromotionCreateTimeEnd 广告创建结束时间，格式yyyy-MM-dd HH:mm:ss与time_start搭配使用
	PromotionCreateTimeEnd string `json:"promotion_create_time_end,omitempty"`
	// PromotionModifyTimeStart 广告更新开始时间，格式yyyy-MM-dd HH:mm:ss与time_end搭配使用
	PromotionModifyTimeStart string `json:"promotion_modify_time_start,omitempty"`
	// PromotionModifyTimeEnd 广告更新结束时间，格式yyyy-MM-dd HH:mm:ss与time_start搭配使用
	PromotionModifyTimeEnd string `json:"promotion_modify_time_end,omitempty"`
	// RejectReasonType 审核建议类型筛选，默认不限，允许值:
	// ALL 不限
	// NONE 无建议
	// LOW_MATERAIL 低俗素材
	// QUALITY_POOR 素材质量低
	//  EXAGGERATION 夸大宣传
	// ELSE 其他
	// DISCOMFORT 引人不适
	// REVIEW_REJECT 审核不通过
	// 默认值: ALL
	RejectReasonType enum.PromotionRejectReasonType `json:"reject_reason_type,omitempty"`
	// LearningPhase 学习期状态筛选，默认不限，允许值:
	// ALL 不限
	// LEARNED 学习期结束
	// LEARNING 学习中
	// LEARN_FAILED 学习失败
	//  默认值: ALL
	LearningPhase enum.LearningPhase `json:"learning_phase,omitempty"`
	// BudgetMode 预算类型筛选，默认不限，允许值:
	// ALL 不限
	// BUDGET_MODE_DAY 日预算
	// BUDGET_MODE_TOTAL 总预算
	//  默认值: ALL
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// BidType 出价方式筛选，默认不限，允许值:
	// ALL 不限
	// MANUAL 手动出价
	// SMART 智能出价
	//  默认值: ALL
	BidType local.BidType `json:"bid_type,omitempty"`
}

// Encode implements GetRequest interface
func (r ListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// ListResponse 获取广告列表 API Response
type ListResponse struct {
	Data *ListResult `json:"data,omitempty"`
	model.BaseResponse
}

type ListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// PromotionList 广告列表
	Promotion []Promotion `json:"promotion_list,omitempty"`
}
