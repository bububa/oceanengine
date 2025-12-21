package project

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取项目列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *ListFilter `json:"filtering,omitempty"`
	// Fields 查询字段集合，如果指定则应答结果仅返回指定字段
	// 可参考应答参数返回的指标字段（不支持audience下字段筛选）
	Fields []string `json:"fields,omitempty"`
	// Page 页数默认值: 1，page必须大于0
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值:10，page_size范围为1-10
	PageSize int `json:"page_size,omitempty"`
}

// ListFilter 过滤条件
type ListFilter struct {
	// IDs 按广告项目ID过滤，范围为1-100
	IDs []uint64 `json:"ids,omitempty"`
	// DeliveryMode 投放模式，允许值：MANUAL手动投放、PROCEDURAL自动投放
	DeliveryMode enum.DeliveryMode `json:"delivery_mode,omitempty"`
	// DeliveryType 按投放类型过滤（当前过滤查询仅支持搜索广告），必须同时传入ad_type = SEARCH，可选值：
	// - NORMAL 常规投放
	// -DURATION周期稳投
	DeliveryType enum.DeliveryType `json:"delivery_type,omitempty"`
	// LandingType 推广目的，允许值：APP 应用推广、LINK 销售线索推广、MICRO_GAME小游戏
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// AppPromotionType 子目标，允许值：DOWNLOAD 应用下载、LAUNCH 应用调用、RESERVE 预约下载
	AppPromotionType enum.AppPromotionType `json:"app_promotion_type,omitempty"`
	// MarketingGoal 营销场景，允许值：VIDEO_AND_IMAGE 短视频/图片
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// AdType 广告类型，允许值：ALL 所有广告
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// Name 广告项目名称，长度是1-50个字（两个英文字符占1个字，该字段采取模糊查询的方式）
	Name string `json:"name,omitempty"`
	// Status 广告项目状态过滤，允许值：
	// PROJECT_STATUS_ENABLE 启用
	// PROJECT_STATUS_DISABLE 暂停
	// PROJECT_STATUS_DELETE 删除
	// PROJECT_STATUS_ALL 所有包含已删除
	// PROJECT_STATUS_NOT_DELETE 所有不包含已删除
	// PROJECT_STATUS_BUDGET_EXCEED 项目超出预算
	// PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET 项目接近预算
	// PROJECT_STATUS_NOT_START 未达投放时间
	// PROJECT_STATUS_DONE 已完成
	// PROJECT_STATUS_NO_SCHEDULE 不在投放时段
	Status enum.ProjectStatus `json:"status,omitempty"`
	// StatusFirst 项目一级状态过滤，允许值：
	// PROJECT_STATUS_DELETE已删除
	// PROJECT_STATUS_DONE已完成
	// PROJECT_STATUS_DISABLE未投放
	// PROJECT_STATUS_ENABLE启用中
	StatusFirst enum.ProjectStatus `json:"status_first,omitempty"`
	// StatusSecond 项目二级状态过滤，允许值：
	// PROJECT_STATUS_STOP 已暂停
	// PROJECT_STATUS_BUDGET_EXCEED 项目超出预算
	// PROJECT_STATUS_NOT_START 未达投放时间
	// PROJECT_STATUS_NO_SCHEDULE 不在投放时段
	// 当status_first = PROJECT_STATUS_DISABLE时传入有效
	StatusSecond enum.ProjectStatus `json:"status_second,omitempty"`
	// ProjectCreateTime 项目创建时间，格式yyyy-mm-dd，表示过滤出当天创建的广告项目
	ProjectCreateTime string `json:"project_create_time,omitempty"`
	// ProjectModifyTime 项目更新时间，格式yyyy-mm-dd，表示过滤出当天更新的广告项目
	ProjectModifyTime string `json:"project_modify_time,omitempty"`
	// ProjectModifyStartTime 项目更新开始时间，该筛选支持您细分到秒级查询一段时间内修改过的项目信息
	// 格式yyyy-mm-dd hh:mm:ss
	// 注意
	// 必须与结束时间project_modify_end_time结合使用，注意时间跨度不能超过30天
	// 开始时间<=结束时间
	// 可以与project_modify_time同时使用，但不太有必要。直接使用project_modify_start_time、project_modify_end_time组合筛选即可
	ProjectModifyStartTime string `json:"project_modify_start_time,omitempty"`
	// ProjectModifyEndTime 项目更新结束时间，该筛选支持您细分到秒级查询一段时间内修改过的项目信息
	// 格式yyyy-mm-dd hh:mm:ss
	// 注意
	// 必须与开始时间project_modify_start_time结合使用，注意时间跨度不能超过30天
	// 开始时间<=结束时间
	// 可以与project_modify_time同时使用，但不太有必要。直接使用project_modify_start_time、project_modify_end_time组合筛选即可
	ProjectModifyEndTime string `json:"project_modify_end_time,omitempty"`
	// Pricing 按计费方式过滤，允许值：PRCING_OCPM OCPM
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// InventoryType 按首选位置过滤，允许值
	// INVENTORY_FEED 今日头条
	// INVENTORY_VIDEO_FEED 西瓜视频
	// INVENTORY_AWEME_FEED 抖音短视频
	// INVENTORY_TOMATO_NOVEL 番茄小说
	// INVENTORY_UNION_SLOT 穿山甲
	// UNION_BOUTIQUE_GAME ohayoo精品游戏
	// INVENTORY_SEARCH 搜索广告位
	InventoryType enum.StatInventoryType `json:"inventory_type,omitempty"`
	// Platform 按平台过滤，允许值：IOS、ANDROID
	Platform enum.AudiencePlatform `json:"platform,omitempty"`
	// BudgetGroupID 按预算组ID过滤，仅允许传入1个，该功能白名单开放，如需使用请联系销售
	// 您需要首先通过「创建预算组」/「获取预算组列表」接口拿到此id
	BudgetGroupID uint64 `json:"budget_group_id,omitempty"`
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

// ListResponse 获取项目列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResponseData `json:"data,omitempty"`
}

type ListResponseData struct {
	// List 项目列表
	List []Project `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
