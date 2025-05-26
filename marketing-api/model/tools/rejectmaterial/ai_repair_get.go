package rejectmaterial

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AIRepairGetRequest 获取拒审素材修复建议 API Request
type AIRepairGetRequest struct {
	// AdvertiserID 广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤器，多个条件之间是「与」的关系
	Filtering *AIRepairGetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值1
	// 注意：
	// page 范围[1,9999999]
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值10
	// 注意：
	// page_size范围[1,100]
	PageSize int `json:"page_size,omitempty"`
}

type AIRepairGetFilter struct {
	// MaterialIDs material_ids不可与video_ids同时筛选
	// 拒审素材id列表，上限50
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// PromotionIDs 广告id列表，上限50
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
	// AIRepairIDs 修复建议id列表，上限50
	AIRepairIDs []uint64 `json:"ai_repair_ids,omitempty"`
	// RepairStartTime 修复时间筛选-起始时间
	// 需与ai_repair_end_time筛选搭配使用
	// 格式：yyyy-mm-dd hh:mm:ss
	// 起始时间必须≤结束时间
	RepairStartTime string `json:"repair_start_time,omitempty"`
	// RepairEndTime 修复时间筛选-结束时间
	// 需与ai_repair_start_time筛选搭配使用
	// 格式：yyyy-mm-dd hh:mm:ss
	// 起始时间必须≤结束时间
	RepairEndTime string `json:"repair_end_time,omitempty"`
}

// Encode implements GetRequest interface
func (r AIRepairGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
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

// AIRepairGetResponse 获取拒审素材修复建议 API Response
type AIRepairGetResponse struct {
	model.BaseResponse
	Data *AIRepairGetResult `json:"data,omitempty"`
}

type AIRepairGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// RepairMaterialInfo 原素材及修复素材列表
	RepairMaterialInfo []RepairMaterialInfo `json:"repair_material_info,omitempty"`
}
