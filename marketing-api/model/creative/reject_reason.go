package creative

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RejectReasonRequest 获取创意审核建议 API Request
type RejectReasonRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreativeIDs 广告创意ID，长度限制：1～10。创意ID需要属于当前广告主，否则会报错。只有审核不通过的创意才有审核建议，审核通过的创意没有审核建议。（所有的程序化创意都是审核通过的）
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r RejectReasonRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.CreativeIDs)
	values.Set("creative_ids", string(ids))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RejectReasonResponse 获取创意审核建议 API Response
type RejectReasonResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data struct {
		// List 审核建议数据
		List []RejectReason `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// RejectReason 审核建议
type RejectReason struct {
	// CreativeID 创意ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// RejectData 审核建议
	RejectData []RejectData `json:"reject_data,omitempty"`
	// MaterialReject 自定义创意类型的素材审核建议
	MaterialReject []MaterialReject `json:"material_reject,omitempty"`
}

// RejectData 审核建议
type RejectData struct {
	// RejectItem 审核项
	RejectItem string `json:"reject_item,omitempty"`
	// RejectReasion 审核建议，审核中/审核通过创意也存在返回审核建议的情况
	RejectReason string `json:"reject_reason,omitempty"`
}

// MaterialReject 自定义创意类型的素材审核建议
type MaterialReject struct {
	// MaterialType 素材类型。1-图片，2-标题，3-视频，4-副标题，5-头图，6-摘要
	MaterialType int `json:"material_type,omitempty"`
	// Title 标题
	Title string `json:"title,omitempty"`
	// ImageID 图片ID
	ImageID []string `json:"image_id,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// RejectReason 审核建议
	RejectReason string `json:"reject_reason,omitempty"`
}
