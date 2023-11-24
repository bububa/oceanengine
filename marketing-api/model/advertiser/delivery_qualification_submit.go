package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeliveryQualificationSubmitRequest 投放资质提交 API Request
type DeliveryQualificationSubmitRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Qualifications
	Qualifications []DeliverySubmitQualification `json:"qualifications,omitempty"`
}

// DeliverySubmitQualification
type DeliverySubmitQualification struct {
	// QualificationType
	QualificaitonType enum.DeliveryQualificationType `json:"qualification_type,omitempty"`
	// AttachmentIDs 附件ids，通过【批量上传资质附件】接口获取
	AttachmentIDs []uint64 `json:"attachment_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r DeliveryQualificationSubmitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeliveryQualificationSubmitResponse 提交投放资质（新版）API Response
type DeliveryQualificationSubmitResponse struct {
	model.BaseResponse
	Data struct {
		// QualificationIDs 资质id列表
		QualificationIDs []uint64 `json:"qualification_ids,omitempty"`
	} `json:"data,omitempty"`
}
