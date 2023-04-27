package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoMaterialClearTaskResultGetRequest 下载清理任务结果 API Request
type VideoMaterialClearTaskResultGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ClearID 清理任务id
	ClearID uint64 `json:"clear_id,omitempty"`
	// Page 页码，默认1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认10，最大100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoMaterialClearTaskResultGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("clear_id", strconv.FormatUint(r.ClearID, 10))
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

// VideoMaterialClearTaskResultGetResponse 下载清理任务结果 API Response
type VideoMaterialClearTaskResultGetResponse struct {
	model.BaseResponse
	Data *VideoMaterialClearTaskResultGetData `json:"data,omitempty"`
}

// VideoMaterialClearTaskResultGetData
type VideoMaterialClearTaskResultGetData struct {
	// ClearResult 清理结果
	ClearResult []VideoMaterialClearResult `json:"clear_result,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// VideoMaterialClearResult 清理结果
type VideoMaterialClearResult struct {
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// ClearMaterialTypes 清理素材类型：
	// INEFFICIENT_MATERIAL低效素材；
	// SIMILAR_MATERIAL 同质化挤压严重素材；
	// SIMILAR_QUEUE_MATERIAL 同质化排队素材
	ClearMaterialTypes []enum.MaterialProperty `json:"clear_material_types,omitempty"`
	// CreativeSuccessCount 清理成功的创意数
	CreativeSuccessCount int64 `json:"creative_success_count,omitempty"`
	// CreativeFailCount 清理失败的创意数
	CreativeFailCount int64 `json:"creative_fail_count,omitempty"`
	// PromotionSuccessCount 清理成功的广告数
	PromotionSuccessCount int64 `json:"promotion_success_count,omitempty"`
	// PromotionFailCount 清理失败的广告数
	PromotionFailCount int64 `json:"promotion_fail_count,omitempty"`
	// ClearResult 清理结果：
	// SUCCESS 成功
	// PART_SUCCESS 部分成功
	// FAIL 失败
	ClearResult string `json:"clear_result,omitempty"`
}
