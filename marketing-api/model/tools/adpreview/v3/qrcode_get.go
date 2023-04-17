package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// IDType 查询条件
type IDType string

const (
	// ID_TYPE_PROJECT 按广告项目ID搜索
	ID_TYPE_PROJECT IDType = "ID_TYPE_PROJECT"
	// ID_TYPE_PROMOTION 按广告计划ID搜索
	ID_TYPE_PROMOTION IDType = "ID_TYPE_PROMOTION"
)

// QrcodeGetRequest 获取广告预览二维码
type QrcodeGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// IDType 查询条件
	IDType IDType `json:"id_type,omitempty"`
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// MaterialID 素材ID，查询素材预览时使用
	MaterialID uint64 `json:"material_id,omitempty"`
}

// Encode implement GetRequest interface
func (r QrcodeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("id_type", string(r.IDType))
	values.Set("promotion_id", strconv.FormatUint(r.PromotionID, 10))
	if r.MaterialID > 0 {
		values.Set("material_id", strconv.FormatUint(r.MaterialID, 10))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QrcodeGetResponse 获取广告预览二维码
type QrcodeGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		Data *QrcodeGetResponseData `json:"data,omitempty"`
	} `json:"data,omitempty"`
}

// QrcodeGetResponseData json返回值
type QrcodeGetResponseData struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// IDType 查询条件
	IDType IDType `json:"id_type,omitempty"`
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// MaterialID 素材ID，查询素材预览时使用
	MaterialID uint64 `json:"material_id,omitempty"`
	// QrcodeMsgURL 广告预览的二维码信息url
	QrcodeMsgURL string `json:"qrcode_msg_url,omitempty"`
}
