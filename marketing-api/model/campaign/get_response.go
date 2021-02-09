package campaign

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

type GetResponse struct {
	model.BaseResponse
	Data *GetResponseData `json:"data,omitempty"`
}

type GetResponseData struct {
	List     []GetResponseList `json:"list,omitempty"`
	PageInfo *model.PageInfo   `json:"page_info,omitempty"`
}

type GetResponseList struct {
	ID                 uint64              `json:"id,omitempty"`                   // 广告组ID
	Name               string              `json:"name,omitempty"`                 // 广告组名称
	Budget             float64             `json:"budget,omitempty"`               // 广告组预算
	BudgetMode         enum.BudgetMode     `json:"budget_mode,omitempty"`          // 广告组预算类型
	LandingType        enum.LandingType    `json:"landing_type,omitempty"`         // 广告组推广目的
	Status             enum.CampaignStatus `json:"status,omitempty"`               // 广告组状态
	CampaignCreateTime string              `json:"campaign_create_time,omitempty"` // 广告组创建时间, 格式：yyyy-mm-dd hh:MM:ss
	CampaignModifyTime string              `json:"campaign_modify_time,omitempty"` // 广告组修改时间, 格式：yyyy-mm-dd hh:MM:ss
	UniqueFk           string              `json:"unique_fk,omitempty"`            // 第三方唯一键，传该值时保证接口重试的幂等性，带有相同unique_fk的请求服务端会视为同一个广告组处理。仅在创建接口传入且无法修改，如果创建时传入了已存在的唯一键值，那么会返回该唯一键值所对应的广告组ID。该值可用于内部系统会生成的唯一ID与头条ID做关联的场景，避免超时重试实际上一次创建请求又成功导致的重复创建问题，通过unique_fk可与内部系统ID实现关联并避免重复创建，可结合实际场景选择使用，广告组中的unique_fk要求不重复，与计划中的unique_fk无相关。
	DeliveryRelatedNum enum.CampaignDPA    `json:"delivery_related_num,omitempty"` // 广告组商品类型
	DeliveryMode       string              `json:"delivery_mode,omitempty"`        // 投放类型，允许值：MANUAL（手动）、PROCEDURAL（自动，投放管家）
}
