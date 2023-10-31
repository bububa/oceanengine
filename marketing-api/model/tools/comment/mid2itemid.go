package comment

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// Mid2ItemIDRequest 获取评论视频ID列表 API Request
type Mid2ItemIDRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartTime 查询起始时间，格式：yyyy-MM-dd，时间跨度最大三个月，仅支持获取2021-01-01之后的评论
	StartTime string `json:"start_time,omitempty"`
	// EndTime 查询截止时间，格式：yyyy-MM-dd，时间跨度最大三个月，仅支持获取2021-01-01之后的评论
	EndTime string `json:"end_time,omitempty"`
	// MaterialID 素材ID
	MaterialID uint64 `json:"material_id,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值10，最大值100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r Mid2ItemIDRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	values.Set("material_id", strconv.FormatUint(r.MaterialID, 10))
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

// Mid2ItemIDResponse 获取评论视频ID列表 API Response
type Mid2ItemIDResponse struct {
	model.BaseResponse
	Data *Mid2ItemIDList `json:"data,omitempty"`
}

type Mid2ItemIDList struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ItemIDs 抖音视频id
	ItemIDs []uint64 `json:"item_ids,omitempty"`
}
