package adpreview

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// IDType 查询条件
type IDType string

const (
	// ID_TYPE_AD 按广告计划ID搜索
	ID_TYPE_AD IDType = "ID_TYPE_AD"
	// ID_TYPE_CREATIVE 按广告创意的ID搜索
	ID_TYPE_CREATIVE IDType = "ID_TYPE_CREATIVE"
)

// QrcodeGetRequest 获取广告预览二维码
type QrcodeGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// IDType 查询条件
	IDType IDType `json:"id_type,omitempty"`
	// ID 广告计划或者广告创意的ID，根据id_type进行相应id的填写
	ID uint64 `json:"id,omitempty"`
}

// Encode implement GetRequest interface
func (r QrcodeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("id_type", string(r.IDType))
	values.Set("id", strconv.FormatUint(r.ID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QrcodeGetResponse 获取广告预览二维码
type QrcodeGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *QrcodeGetResponseData `json:"data,omitempty"`
}

// QrcodeGetResponseData json返回值
type QrcodeGetResponseData struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// IDType 查询条件
	IDType IDType `json:"id_type,omitempty"`
	// ID 广告计划或者广告创意的ID，根据id_type进行相应id的填写
	ID uint64 `json:"id,omitempty"`
	// QrcodeMsgURL 广告预览的二维码信息url
	QrcodeMsgURL string `json:"qrcode_msg_url,omitempty"`
}
