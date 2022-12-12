package ad

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RejectReasonRequest 获取计划审核建议 API Request
type RejectReasonRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 广告计划 ID，最多传10个广告计划ID
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r RejectReasonRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.AdIDs)
	values.Set("ad_ids", string(ids))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RejectReasonResponse 获取计划审核建议 API Response
type RejectReasonResponse struct {
	model.BaseResponse
	Data struct {
		List []RejectReason `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// RejectReason 计划审核意见
type RejectReason struct {
	// AdReject 计划维度审核建议
	AdReject *AdReject `json:"ad_reject,omitempty"`
	// CreativeReject 创意维度审核建议
	CreativeReject []CreativeReject `json:"creative_reject,omitempty"`
	// MaterialReject 程序化创意的素材维度审核建议
	MaterialReject []MaterialReject `json:"material_reject,omitempty"`
	// IsProcedualAd 1 表示程序化创意，0 表示自定义创意
	IsProcedualAd int `json:"is_procedual_ad,omitempty"`
}

// RejectData 审核建议数据
type RejectData struct {
	// RejectItem 审核项
	RejectItem string `json:"reject_item,omitempty"`
	// RejectReason 审核建议
	RejectReason string `json:"reject_reason,omitempty"`
}

// AdReject 计划维度审核建议
type AdReject struct {
	// AdID 广告计划 id
	AdID uint64 `json:"ad_id,omitempty"`
	// RejectData 审核建议数据
	RejectData []RejectData `json:"reject_data,omitempty"`
}

// CreativeReject 创意维度审核建议
type CreativeReject struct {
	// CreativeID 创意 id
	CreativeID uint64 `json:"creative_id,omitempty"`
	// RejectData 审核建议数据
	RejectData []RejectData `json:"reject_data,omitempty"`
	// MaterialReject 自定义创意的素材维度审核建议。字段和下面程序化创意素材维度审核建议完全相同。
	MaterialReject []MaterialReject `json:"material_reject,omitempty"`
}

// MaterialReject 创意的素材维度审核建议
type MaterialReject struct {
	// Title 标题
	Title string `json:"title,omitempty"`
	// ImageID 图片 id
	ImageID []string `json:"image_id,omitempty"`
	// VideoID 视频 id
	VideoID string `json:"video_id,omitempty"`
	// RejectReason 审核建议
	RejectReason string `json:"reject_reason,omitempty"`
}
