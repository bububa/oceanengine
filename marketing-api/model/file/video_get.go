package file

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// VideoGetRequest 获取视频素材 API Request
type VideoGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 视频过滤条件
	Filtering *VideoGetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值20
	PageSize int `json:"page_size,omitempty"`
}

// VideoGetFilter 视频过滤条件
type VideoGetFilter struct {
	// VideoIDs 视频ids，示例: ["86adb23eaa21229fc04ef932b5089bb8"]
	// 数量限制：<=100
	// 注意：video_ids、material_ids、signatures只能选择一个进行过滤
	VideoIDs []string `json:"video_ids,omitempty"`
	// MaterialIDs 素材id列表，可以根据material_ids（素材报表使用的id，一个素材唯一对应一个素材id）进行过滤
	// 数量限制：<=100
	// 注意：image_ids、material_ids、signatures只能选择一个进行过滤
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// Signatures md5值列表，可以根据素材的md5进行过滤
	// 数量限制：<=100
	// 注意：image_ids、material_ids、signatures只能选择一个进行过滤
	Signatures []string `json:"signatures,omitempty"`
	// Width 图片宽度
	Width int `json:"width,omitempty"`
	// Height 图片高度
	Height int `json:"height,omitempty"`
	// Ratio 图片宽高比，eg: [1.7, 2.5]，输入1.7则搜索满足宽高比介于1.65-1.75之间的图片，即精度上下浮动0.05
	Ratio []float64 `json:"ratio,omitempty"`
	// StartTime 根据视频上传时间进行过滤的起始时间，与end_time搭配使用，格式：yyyy-mm-dd
	StartTime string `json:"start_time,omitempty"`
	// EndTime 根据视频上传时间进行过滤的截止时间，与start_time搭配使用，格式：yyyy-mm-dd
	EndTime string `json:"end_time,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filter, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filter))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// VideoGetResponse 获取视频素材 API Response
type VideoGetResponse struct {
	model.BaseResponse
	Data *VideoGetResponseData `json:"data,omitempty"`
}

// VideoGetResponseData json返回值
type VideoGetResponseData struct {
	// List 视频列表
	List []Video `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
