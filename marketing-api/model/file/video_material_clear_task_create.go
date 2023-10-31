package file

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoMaterialClearTaskCreateRequest 创建素材清理任务 API Request
type VideoMaterialClearTaskCreateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ClearTaskParams 任务参数，若同时填写多项任务的参数，则最终返回各个参数的交集
	ClearTaskParams *VideoMaterialClearTaskParameters `json:"clear_task_params,omitempty"`
}

// VideoMaterialClearTaskParameters 任务参数，若同时填写多项任务的参数，则最终返回各个参数的交集
type VideoMaterialClearTaskParameters struct {
	// MaterialSource 清理素材来源
	// 允许值：QIANCHUAN 千川、AD 广告平台
	MaterialSource enum.AccountType `json:"material_source,omitempty"`
	// MaterialIDs 待清理素材列表，单次最多支持100个
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// ClearMaterialTypes 清理素材类型，允许值：INEFFICIENT_MATERIAL低效素材；SIMILAR_MATERIAL 同质化挤压严重素材；SIMILAR_QUEUE_MATERIAL同质化排队素材
	// 如果不传清理素材类型，默认清理这三个标签的全部素材
	ClearMaterialTypes []enum.MaterialProperty `json:"clear_material_types,omitempty"`
	// CreateTimeUpper 清理创建时间"yyyy-mm-dd"之前的素材，不包括这一天
	CreateTimeUpper string `json:"create_time_upper,omitempty"`
	// StartTime 账户下累积转化数和累积消耗数的数据统计开始时间，格式为"yyyy-mm-dd"，且包含该日期
	StartTime string `json:"start_time,omitempty"`
	// EndTime 账户下累积转化数和累积消耗数的数据统计结束时间，格式为"yyyy-mm-dd"，且包含该日期
	EndTime string `json:"end_time,omitempty"`
	// Convert 账户下累积转化数，在start_time和end_time期间发生的累积转化数，清理小于该累积转化数的素材
	// 当start_time和end_time传入时有效
	Convert int64 `json:"convert,omitempty"`
	// Cost 账户下累积消耗，在start_time和end_time期间发生的累积消耗，清理小于该累积消耗的素材
	// 当start_time和end_time传入时有效
	Cost float64 `json:"cost,omitempty"`
}

// Encode implement PostRequest interface
func (r VideoMaterialClearTaskCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// VideoMaterialClearTaskCreateResponse 创建素材清理任务 API Response
type VideoMaterialClearTaskCreateResponse struct {
	model.BaseResponse
	// Data
	Data struct {
		// ClearID 创建成功的素材清理任务id
		ClearID uint64 `json:"clear_id,omitempty"`
	}
}
