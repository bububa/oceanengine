package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoGetRequest 获取千川素材库视频 API Request
type VideoGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 图片过滤条件
	Filtering *ImageGetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值20
	PageSize int `json:"page_size,omitempty"`
}

// ImageVideoFilter 视频过滤条件
type VideoGetFilter struct {
	// VideoIDs 视频ids，示例: ["86adb23eaa21229fc04ef932b5089bb8"] 数量限制：<=100
	// 注意：video_ids、material_ids、signatures只能选择一个进行过滤
	VideoIDs []string `json:"video_ids,omitempty"`
	// MaterialIDs 素材id列表，可以根据material_ids（素材报表使用的id，一个素材唯一对应一个素材id）进行过滤数量限制：<=100
	// 注意：image_ids、material_ids、signatures只能选择一个进行过滤
	MaterialIDs []string `json:"material_ids,omitempty"`
	// Sigatures md5值列表，可以根据素材的md5进行过滤数量限制：<=100
	// 注意：image_ids、material_ids、signatures只能选择一个进行过滤
	Sigatures []string `json:"signatures,omitempty"`
	// ImageMode 素材类型
	ImageMode []enum.MaterialMode `json:"image_mode,omitempty"`
	// Tags 素材标签
	Tags []string `json:"tags,omitempty"`
	// Sources 素材来源，允许值：
	// ARTHUR 亚瑟共享素材
	// BP 巨量纵横共享素材
	// CREATIVE_CENTER 巨量创意PC共享素材
	// E_COMMERCE 本地上传
	// LIVE_HIGHLIGHT 直播剪辑素材
	// STAR 星图&即合共享素材
	// TADA tada共享素材
	// VIDEO_CAPTURE 易拍APP共享素材
	Sources []enum.MaterialSource `json:"sources,omitempty"`
	// StartTime 根据图片上传时间进行过滤的起始时间，与end_time搭配使用，格式："yyyy-mm-dd"
	StartTIme string `json:"start_time,omitempty"`
	// EndTime 根据图片上传时间进行过滤的截止时间，与start_time搭配使用，格式："yyyy-mm-dd"
	EndTime string `json:"end_time,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoGetRequest) Encode() string {
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

// VideoGetResponse 获取视频素材 API Response
type VideoGetResponse struct {
	model.BaseResponse
	Data *VideoGetResult `json:"data,omitempty"`
}

// VideoGetResult json返回值
type VideoGetResult struct {
	// List 视频列表
	List []Video `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
