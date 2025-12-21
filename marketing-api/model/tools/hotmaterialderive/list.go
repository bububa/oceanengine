package hotmaterialderive

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取账户下爆款裂变任务列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主账户 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤项目
	Filtering *ListFilter `json:"filtering,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 单页数
	PageSize int `json:"page_size,omitempty"`
}

type ListFilter struct {
	// OriginMaterialIDs 爆款素材 ID 过滤，最多支持 50 个
	OriginMaterialIDs []uint64 `json:"origin_material_ids,omitempty"`
	// TaskIDs 爆款裂变任务 ID 过滤，最多支持 50 个
	TaskIDs []uint64 `json:"task_ids,omitempty"`
	// Status 可选值:
	// FAILED 任务失败
	// INIT 任务初始化
	// PART_SUCCESS 任务部分成功
	// PPOCESSING 任务处理中
	// SUCCESS 任务成功
	Statuses []string `json:"statuses,omitempty"`
	// StartTime 创建开始时间筛选，格式 YYYY-MM-DD HH:MM:SS
	StartTime string `json:"start_time,omitempty"`
	// EndTime 创建结束时间筛选，格式 YYYY-MM-DD HH:MM:SS
	EndTime string `json:"end_time,omitempty"`
}

// Encode implements GetRequest interface
func (r ListRequest) Encode() string {
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

// ListResponse 获取账户下爆款裂变任务列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	Pagination *model.PageInfo `json:"pagination,omitempty"`
	Data       []Task          `json:"data,omitempty"`
}
