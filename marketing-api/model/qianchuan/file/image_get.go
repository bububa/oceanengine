package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ImageGetRequest 获取千川素材库图片 API Request
type ImageGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 图片过滤条件
	Filtering *ImageGetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值20
	PageSize int `json:"page_size,omitempty"`
}

// ImageGetFilter 图片过滤条件
type ImageGetFilter struct {
	// ImageIDs 图片ids，可以根据图片ids（创意中使用的图片key，存在一张图片对应多个image_ids的情况）进行过滤数量限制：<=100
	// 注意：image_ids、material_ids、signatures只能选择一个进行过滤
	ImageIDs []string `json:"image_ids,omitempty"`
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
	// StartTime 根据图片上传时间进行过滤的起始时间，与end_time搭配使用，格式："yyyy-mm-dd"
	StartTIme string `json:"start_time,omitempty"`
	// EndTime 根据图片上传时间进行过滤的截止时间，与start_time搭配使用，格式："yyyy-mm-dd"
	EndTime string `json:"end_time,omitempty"`
}

// Encode implement GetRequest interface
func (r ImageGetRequest) Encode() string {
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

// ImageGetResponse 获取图片素材 API Response
type ImageGetResponse struct {
	model.BaseResponse
	Data *ImageGetResult `json:"data,omitempty"`
}

// ImageGetResult json返回值
type ImageGetResult struct {
	// List 图片列表
	List []Image `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
