package carousel

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取千川素材库图文 API Request
type GetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 图文过滤信息
	Filtering *GetFilter `json:"filtering,omitempty"`
	// OrderFields 排序字段，默认不传为create_time，见返回参数中metrics
	OrderFields string `json:"order_fields,omitempty"`
	// OrderType 排序方式，允许值：
	// ASC 升序（默认）
	// DESC 降序
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值20
	PageSize int `json:"page_size,omitempty"`
}

type GetFilter struct {
	// MaterialIDs 素材id列表，可以根据material_ids（素材报表使用的id，一个素材唯一对应一个素材id）进行过滤
	// 数量限制：<=100
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// Sources 图文素材来源，允许值
	// JICHAUNG 即创
	Sources []qianchuan.CarouselSource `json:"sources,omitempty"`
	// ImageMode 素材类型，允许值
	// CAROUSEL 图文
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// QueryString 支持根据图文名称/ID搜索
	QueryString string `json:"query_string,omitempty"`
	// StartTime 根据图片上传时间进行过滤的起始时间，与end_time搭配使用，格式："yyyy-mm-dd"
	StartTime string `json:"start_time,omitempty"`
	// EndTime 根据图片上传时间进行过滤的截止时间，与start_time搭配使用，格式："yyyy-mm-dd"
	EndTime string `json:"end_time,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.OrderFields != "" {
		values.Set("order_fields", r.OrderFields)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
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

// GetResponse 获取千川素材库图文 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	PageInfo  *model.PageInfo `json:"page_info,omitempty"`
	Carousels []Carousel      `json:"carousels,omitempty"`
}

// Carousel 素材信息
type Carousel struct {
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// ImageMode 素材类型:
	// CAROUSEL 图文
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// Images 图片信息
	Images []Image `json:"images,omitempty"`
	// Audio 音频信息
	Audio *Audio `json:"audio,omitempty"`
	// Filename 素材的文件名
	Filename string `json:"filename,omitempty"`
	// CreateTime 图文创建时间
	CreateTime string `json:"create_time,omitempty"`
	// Description 素材描述
	Description string `json:"description,omitempty"`
}
