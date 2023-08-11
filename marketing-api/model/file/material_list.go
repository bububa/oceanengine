package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialListRequest 获取素材标签列表 API Request
type MaterialListRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialSource 素材来源，允许值：QIANCHUAN 千川
	// AD 广告平台
	MaterialSource enum.AccountType `json:"material_source,omitempty"`
	// PropertiesFilter 素材标签过滤项，允许值：
	// INEFFICIENT_MATERIAL低效素材；
	// SIMILAR_MATERIAL 同质化挤压严重素材；
	// SIMILAR_QUEUE_MATERIAL 同质化素材风险-排队投放素材
	// AD_HIGH_QUALITY_MATERIAL AD 优质素材
	// ECP_HIGH_QUALITY_MATERIAL 千川优质素材
	// FIRST_PUBLISH_MATERIAL  首发素材
	// 如果不传，则默认返回广告主ID下所有素材
	PropertiesFilter []enum.MaterialProperty `json:"properties_filter,omitempty"`
	// StartTime 素材创建时间，格式为yyyy-mm-dd HH:MM:SS，最早允许传入时间&默认时间：2022-01-01 00:00:00
	StartTime string `json:"start_time,omitempty"`
	// EndTime 素材结束时间，格式为yyyy-mm-dd HH:MM:SS，默认为当前时间
	EndTime string `json:"end_time,omitempty"`
	// Page 页数默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值:10, 最大值：100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r MaterialListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("material_source", string(r.MaterialSource))
	if len(r.PropertiesFilter) > 0 {
		values.Set("properties_filter", string(util.JSONMarshal(r.PropertiesFilter)))
	}
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// MaterialListResponse 获取素材标签列表 API Response
type MaterialListResponse struct {
	model.BaseResponse
	Data *MaterialListData `json:"data,omitempty"`
}

type MaterialListData struct {
	// Materials 低效素材id列表
	Materials []Material `json:"materials,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// Material 低效素材
type Material struct {
	// MaterialID 请求的素材 id
	MaterialID model.Uint64 `json:"material_id,omitempty"`
	// MaterialProperties 素材属性标签
	// INEFFICIENT_MATERIAL低效素材
	// SIMILAR_MATERIAL 同质化挤压严重素材
	// SIMILAR_QUEUE_MATERIAL 同质化素材风险-排队投放素材
	// AD_HIGH_QUALITY_MATERIAL AD 优质素材
	// ECP_HIGH_QUALITY_MATERIAL 千川优质素材
	// FIRST_PUBLISH_MATERIAL  首发素材
	MaterialProperties []enum.MaterialProperty `json:"material_properties,omitempty"`
	// IsInefficientMaterial 是否低效素材
	IsInefficientMaterial bool `json:"is_inefficient_material,omitempty"`
	// IsSimilarQueueMaterial 是否同质化素材风险-排队投放素材
	IsSimilarQueueMaterial bool `json:"is_similar_queue_material,omitempty"`
	// IsSimilarExpectedQueueMaterial 是否同质化素材风险-未投放预计排队素材
	IsSimilarExpectedQueueMaterial bool `json:"is_similar_expected_queue_material,omitempty"`
	// IsSimilarMaterial 是否同质化挤压严重素材
	IsSimilarMaterial bool `json:"is_similar_material,omitempty"`
	// IsEcpHighQuality 是否千川优质素材
	IsEcpHighQuality bool `json:"is_ecp_high_quality,omitempty"`
	// IsAdHighQuality 是否AD优质素材
	IsAdHighQuality bool `json:"is_ad_high_quality,omitempty"`
	// IsFirstPublishMaterial 是否是首发素材
	IsFirstPublishMaterial bool `json:"is_first_publish_matertial,omitempty"`
	// Aigc 素材是否是aigc生成
	Aigc bool `json:"aigc,omitempty"`
}
