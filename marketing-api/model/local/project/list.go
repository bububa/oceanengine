package project

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取项目列表 API Request
type ListRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// Filtering 过滤字段
	Filtering *ListRequestFilter `json:"filtering,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，最大值100，默认值20
	PageSize int `json:"page_size,omitempty"`
}

type ListRequestFilter struct {
	// ProjectIDs 项目IDs筛选，最多100个
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
	// ProjectStatusFirst 项目一级状态筛选，允许值：
	// PROJECT_STATUS_ALL 不限（包含已删除）
	// PROJECT_STATUS_DELETE 已删除
	// PROJECT_STATUS_DISABLE 未投放
	// PROJECT_STATUS_DONE 已完成
	// PROJECT_STATUS_ENABLE 启用中
	// PROJECT_STATUS_NOT_DELETE 不限（不包含已删除）
	// 默认值: PROJECT_STATUS_NOT_DELETE 不限（不包含已删除）
	ProjectStatusFirst local.ProjectStatus `json:"project_status_first,omitempty"`
	// ProjectStatusSecond 项目二级状态筛选，允许值：
	// PROJECT_STATUS_BUDGET_EXCEED 项目超出预算
	// PROJECT_STATUS_DISABLE 已暂停
	// PROJECT_STATUS_NOT_SCHEDULE 不在投放时段
	// PROJECT_STATUS_NOT_START 未达投放时间
	// 仅当status_first = PROJECT_STATUS_DISABLE 未投放时传入有效
	ProjectStatusSecond local.ProjectStatus `json:"project_status_second,omitempty"`
	// ShopIDs 按门店IDs筛选，单次限制最多10个
	ShopIDs []uint64 `json:"shop_ids,omitempty"`
	// ProductIDs 按商品IDs筛选，单次限制最多10个
	ProductIDs []uint64 `json:"product_ids,omitempty"`
	// LocalDeliveryScene 推广目的筛选，默认不限，允许值：
	// ALL 不限
	// CONTENT_HEAT 内容加热
	// POI_RECOMMEND 门店引流
	// PRODUCT_PAY 团购成交
	// 默认值: ALL
	LocalDeliveryScene local.LocalDeliveryScene `json:"local_delivery_scene,omitempty"`
	// MarketingGoal 营销场景筛选，默认不限 ，允许值：
	// ALL 不限
	// LIVE 直播
	// VIDEO_IMAGE 短视频/图文
	// 默认值: ALL
	MarketingGoal local.MarketingGoal `json:"marketing_goal,omitempty"`
	// AdType 广告类型筛选，默认不限，允许值：
	// ALL 不限
	// GENERAL 通投广告
	// SEARCHING 搜索广告
	// 默认值: ALL
	AdType local.AdType `json:"ad_type,omitempty"`
	// ProjectName 项目名称，模糊搜索
	ProjectName string `json:"project_name,omitempty"`
	// ProjectCreateTimeStart 项目创建开始时间，格式yyyy-MM-dd HH:mm:ss，与project_create_time_end搭配使用
	ProjectCreateTimeStart string `json:"project_create_time_start,omitempty"`
	// ProjectCreateTimeEnd 项目创建结束时间，格式yyyy-MM-dd HH:mm:ss，与project_create_time_start搭配使用
	ProjectCreateTimeEnd string `json:"project_create_time_end,omitempty"`
	// ProjectModifyTimeStart 项目更新开始时间，格式yyyy-MM-dd HH:mm:ss，与project_modify_time_end搭配使用
	ProjectModifyTimeStart string `json:"project_modify_time_start,omitempty"`
	// ProjectModifyTimeEnd 项目更新结束时间，格式yyyy-MM-dd HH:mm:ss，与project_modify_time_start搭配使用
	ProjectModifyTimeEnd string `json:"project_modify_time_end,omitempty"`
	// BidType 出价方式，默认不限，允许值：
	// ALL 不限
	// MANUAL 手动出价
	// SMART 智能出价
	// 默认值: ALL
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

// ListResponse 获取项目列表 API Response
type ListResponse struct {
	Data *ListResult `json:"data,omitempty"`
	model.BaseResponse
}

type ListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ProjectList 项目列表
	ProjectList []Project `json:"project_list,omitempty"`
}
