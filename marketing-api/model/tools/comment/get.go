package comment

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取评论列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 计划id列表，一次最多10个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// InventoryType 广告位 允许值："INVENTORY_AWEME_FEED"（抖音）
	InventoryType string `json:"inventory_type,omitempty"`
	// StartTime 查询起始时间，格式：yyyy-MM-dd，若不填，默认6天前（即获取最近七天的内容）
	StartTime string `json:"start_time,omitempty"`
	// EndTime 查询截止时间，格式：yyyy-MM-dd，若不填，默认当天
	EndTime string `json:"end_time,omitempty"`
	// Page 页数 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.AdIDs) > 0 {
		ids, _ := json.Marshal(r.AdIDs)
		values.Set("ad_ids", string(ids))
	}
	values.Set("inventory_type", r.InventoryType)
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
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

// GetResponse 获取评论列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json 返回值
type GetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// CommentsList 评论列表
	CommentsList []Comment `json:"comments_list,omitempty"`
}
