package promotion

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取广告列表 API Request
type ListRequest struct {
	// Filtering 过滤条件
	Filtering *ListFilter `json:"filtering,omitempty"`
	// IncludingMaterialAttributes 如需查询一个广告下是否包含搬运打压状态的视频素材，以及一个视频素材是否存在搬运风险，请传入此参数，并定义枚举值为RETURN_CARRY_DATA
	// 可选值：RETURN_CARRY_DATA 返回视频素材的搬运属性
	// 如果不传此参数，应答参数将不会返回搬运相关的字段：has_carry_material、is_carry_material
	IncludingMaterialAttributes string `json:"including_material_attributes,omitempty"`
	// Fields 查询字段集合，如果指定则应答结果仅返回指定字段
	// 可参考应答参数返回的指标字段
	Fields []string `json:"fields,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页数默认值: 1，page必须大于0
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值:10，page_size范围为1-10
	PageSize int `json:"page_size,omitempty"`
	// Cursor 页码游标值：第一次拉取，传入0
	// page与cursor同时传入时，cursor优先级大于page；同时不传入默认走page逻辑
	// page+page_size与cursor+count为两种分页方式，返回参数只返回与入参对应的分页参数
	Cursor int `json:"cursor,omitempty"`
	// Count 页面数据量
	// page与cursor同时传入时，cursor优先级大于page；同时不传入默认走page逻辑
	// page+page_size与cursor+count为两种分页方式，返回参数只返回与入参对应的分页参数
	Count int `json:"count,omitempty"`
}

// ListFilter 过滤条件
type ListFilter struct {
	// Name 广告名称，长度是1-50个字（两个英文字符占1个字，该字段采取模糊查询的方式）
	Name string `json:"name,omitempty"`
	// DeliveryMode 投放模式，允许值：MANUAL手动投放 (默认值），PROCEDURAL自动投放
	DeliveryMode enum.DeliveryMode `json:"delivery_mode,omitempty"`
	// Status 广告状态过滤，允许值：
	// NOT_DELETED 不限 、ALL 不限（包含已删除）、OK 投放中、DELETED 已删除、PROJECT_OFFLINE_BUDGET 项目超出预算、PROJECT_PREOFFLINE_BUDGET 项目接近预算、TIME_NO_REACH 未到达投放时间、TIME_DONE 已完成、NO_SCHEDULE 不在投放时段、AUDIT 新建审核中、REAUDIT 修改审核中、FROZEN 已终止、AUDIT_DENY 审核不通过、OFFLINE_BUDGET 广告超出预算、OFFLINE_BALANCE 账户余额不足、PREOFFLINE_BUDGET 广告接近预算、DISABLED 已暂停、PROJECT_DISABLED 已被项目暂停、LIVE_ROOM_OFF 关联直播间不可投、PRODUCT_OFFLINE 关联商品不可投，、AWEME_ACCOUNT_DISABLED 关联抖音账号不可投、AWEME_ANCHOR_DISABLED 锚点不可投、DISABLE_BY_QUOTA 已暂停（配额达限）
	Status enum.PromotionStatus `json:"status,omitempty"`
	// StatusFirst 广告一级状态过滤，允许值：
	// PROMOTION_STATUS_ENABLE投放中
	// PROMOTION_STATUS_DISABLE未投放
	// PROMOTION_STATUS_FROZEN已终止
	// PROMOTION_STATUS_DONE已完成
	// PROMOTION_STATUS_DELETED已删除
	StatusFirst enum.PromotionStatusFirst `json:"status_first,omitempty"`
	// StatusSecond 广告二级状态过滤，允许值：
	// AUDIT_DENY 审核不通过
	// AUDIT 新建审核中
	// REAUDIT 修改审核中
	// DISABLED 已暂停
	// DISABLE_BY_QUOTA 已暂停（配额达限）
	// PROJECT_DISABLED 项目已被暂停
	// NO_SCHEDULE 不在投放时段
	// TIME_NO_REACH 未到达投放时间
	// OFFLINE_BALANCE 账户余额不足
	// BALANCE_OFFLINE_BUDGET 账户超出预算
	// PROJECT_OFFLINE_BUDGET 项目超出预算
	// PROMOTION_OFFLINE_BUDGET 广告超出预算
	// LIVE_ROOM_OFF 直播间不可投
	// PRODUCT_OFFLINE 商品不可投
	// AWEME_ACCOUNT_DISABLED 抖音账号不可投
	// AWEME_ANCHOR_DISABLED 锚点不可投
	// 当status_first = PROMOTION_STATUS_DISABLE时传入有效
	StatusSecond enum.PromotionStatus `json:"status_second,omitempty"`
	// PromotionCreateTime 广告创建时间，格式yyyy-mm-dd，表示过滤出当天创建的广告项目
	PromotionCreateTime string `json:"promotion_create_time,omitempty"`
	// PromotionModifyTime 广告更新时间，格式yyyy-mm-dd，表示过滤出当天更新的广告项目
	PromotionModifyTime string `json:"promotion_modify_time,omitempty"`
	// PromotionModifyStartTime 广告更新开始时间，该筛选支持您细分到秒级查询一段时间内修改过的广告信息
	// 格式yyyy-mm-dd hh:mm:ss
	// 注意
	// 必须与结束时间promotion_modify_end_time结合使用，注意时间跨度不能超过30天
	// 开始时间<=结束时间
	// 可以与promotion_modify_time同时使用，但不太有必要。直接使用promotion_modify_start_time、promotion_modify_end_time组合筛选即可
	PromotionModifyStartTime string `json:"promotion_modify_start_time,omitempty"`
	// PromotionModifyEndTime 广告更新结束时间，该筛选支持您细分到秒级查询一段时间内修改过的广告信息
	// 格式yyyy-mm-dd hh:mm:ss
	// 注意
	// 必须与开始时间promotion_modify_start_time结合使用，注意时间跨度不能超过30天
	// 开始时间<=结束时间
	// 可以与promotion_modify_time同时使用，但不太有必要。直接使用promotion_modify_start_time、promotion_modify_end_time组合筛选即可
	PromotionModifyEndTime string `json:"promotion_modify_end_time,omitempty"`
	// RejectReasonType 审核建议类型，允许值：NONE 无建议、REVIEW_REJECT 审核不通过、LOW_MATERAIL 低质素材、DISCOMFORT 引人不适、QUALITY_POOR 素材质量低、EXAGGERATION 夸大宣传、ELSE 其他
	RejectReasonType enum.PromotionRejectReasonType `json:"reject_reason_type,omitempty"`
	// HasCarryMaterial 按素材搬运打压状态过滤
	// TRUE：包含打压风险的广告
	// FALSE：不包含打压风险的广告
	HasCarryMaterial string `json:"has_carry_material,omitempty"`
	// LearningPhase 学习期状态 允许值：LEARNING（学习期中）、LEARNED（学习期结束）、LEARN_FAILED（学习期失败)
	LearningPhase []enum.LearningPhase `json:"learning_phase,omitempty"`
	// IDs 按广告ID过滤，范围为1-100
	IDs []uint64 `json:"ids,omitempty"`
	// ProjectID 按项目id过滤
	ProjectID uint64 `json:"project_id,omitempty"`
}

func (f ListFilter) GetIDs() []uint64 {
	return f.IDs
}

func (f ListFilter) GetName() string {
	return f.Name
}

func (f ListFilter) GetCreateTime() string {
	return f.PromotionCreateTime
}

func (f ListFilter) GetModifyTime() string {
	return f.PromotionModifyTime
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if len(r.Fields) > 0 {
		fields, _ := json.Marshal(r.Fields)
		values.Set("fields", string(fields))
	}
	if r.IncludingMaterialAttributes != "" {
		values.Set("including_material_attributes", r.IncludingMaterialAttributes)
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Cursor > 0 {
		values.Set("cursor", strconv.Itoa(r.Cursor))
	}
	if r.Count > 0 {
		values.Set("count", strconv.Itoa(r.Count))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ListResponse 获取广告列表 API Response
type ListResponse struct {
	Data *ListResponseData `json:"data,omitempty"`
	model.BaseResponse
}

type ListResponseData struct {
	// CursorInfo 游标分页信息，当分页方式为cursor+count时返回
	CursorInfo *model.PageInfo `json:"cursor_info,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 项目列表
	List []Promotion `json:"list,omitempty"`
}
