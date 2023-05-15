package advertiser

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeliveryQualificationListRequest 投放资质查询
type DeliveryQualificationListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// QualificationType 投放资质类型
	QualificationType enum.DeliveryQualificationType `json:"qualification_type,omitempty"`
	// Status 投放资质状态
	Status enum.DeliveryQualificationStatus `json:"status,omitempty"`
	// Page 页数默认值：1，page范围为[1,1000]
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值：10，page_size范围为[1,100]
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r DeliveryQualificationListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.QualificationType != "" {
		values.Set("qualification_type", string(r.QualificationType))
	}
	if r.Status != "" {
		values.Set("status", string(r.Status))
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

type DeliveryQualificationListResponse struct {
	model.BaseResponse
	Data *DeliveryQualificationListData `json:"data,omitempty"`
}

type DeliveryQualificationListData struct {
	model.PageInfo
	// List 投放资质列表
	List []DeliveryQualification `json:"list,omitempty"`
}

// DeliveryQualification 投放资质
type DeliveryQualification struct {
	// QualificationID
	QualificationID uint64 `json:"qualification_id,omitempty"`
	// QualificationType 投放资质类型
	QualificationType enum.DeliveryQualificationType `json:"qualification_type,omitempty"`
	// Status 投放资质状态
	Status enum.DeliveryQualificationStatus `json:"status,omitempty"`
	// RejectReason
	RejectReason string `json:"reject_reason,omitempty"`
	// AuditTime string
	AuditTime string `json:"audit_time,omitempty"`
	// Images
	Images []DeliveryQualificationImage `json:"images,omitempty"`
}

type DeliveryQualificationImage struct {
	// AttachmentID
	AttachmentID uint64 `json:"attachment_id,omitempty"`
	// ImageDownloadURL
	ImageDownloadURL string `json:"image_download_url,omitempty"`
	// ImagePreviewURL
	ImagePreviewURL string `json:"image_preview_url,omitempty"`
}
