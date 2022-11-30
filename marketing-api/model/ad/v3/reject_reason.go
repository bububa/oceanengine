package v3

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// RejectReasonRequest 获取计划审核建议 API Request
type RejectReasonRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionIDs 广告计划 ID，最多传10个广告计划ID
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r RejectReasonRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.PromotionIDs)
	values.Set("promotion_ids", string(ids))
	return values.Encode()
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
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// PromotionReject  广告维度审核建议
	PromotionReject *PromotionReject `json:"promotion_reject,omitempty"`
	// MaterialReject 素材维度审核建议
	MaterialReject []MaterialReject `json:"material_reject,omitempty"`
}

// RejectData 审核建议数据
type RejectData struct {
	// RejectItem 审核项
	RejectItem string `json:"reject_item,omitempty"`
}

// PromotionReject 广告维度审核建议
type PromotionReject struct {
	// Content 审核项
	Content string `json:"content,omitempty"`
	// RejectReason 审核建议
	RejectReason []string `json:"reject_reason,omitempty"`
	// Suggestion 审核建议
	Suggestion []string `json:"suggestion,omitempty"`
}

// MaterialReject 素材维度审核建议
type MaterialReject struct {
	// Type 素材类型，枚举值：
	//  CREATIVE_IMAGE 图片、CREATIVE_VIDEO 视频、CREATIVE_TITLE 标题、CALL_TO_ACTION 行动号召、CREATIVE_URL 创意详情页、PRODUCT_IMAGE 产品主图、PRODUCT_SELLING_POINTS 产品卖点、PRODUCT_DESCRIBE 产品名称 CREATIVE_COMPONENT 创意组件
	Type enum.MaterialType `json:"type,omitempty"`
	// Item 根据不同素材类型返回相应的值
	// - CREATIVE_URL 创意详情页：web_url;external_url;playable_url
	// - CREATIVE_TITLE 标题：title
	// - CREATIVE_IMAGE 图片：image_id
	// - CREATIVE_VIDEO 视频：video_id
	// - CALL_TO_ACTION 行动号召：号召文案
	// - PRODUCT_IMAGE 产品主图：image_id
	// - PRODUCT_SELLING_POINTS 产品卖点：产品卖点文案
	// - PRODUCT_DESCRIB 产品名称：产品名称文案
	// CREATIVE_COMPONENT 创意组件：component_id
	Item string `json:"item,omitempty"`
	// RejectReason 审核建议
	RejectReason []string `json:"reject_reason,omitempty"`
	// Suggestion 审核建议
	Suggestion []string `json:"suggestion,omitempty"`
}
