package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QualificationCreateRequest 批量上传投放资质 API Request
type QualificationCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// QualificationType 资质类型，详见【附录-资质类型】.允许值: "QUALIFICATION_AD"
	QualificationType enum.QualificationType `json:"qualification_type"`
	// Qualifications 资质列表 min>=1 && max<=50
	Qualifications []QualificationForCreate `json:"qualifications,omitempty"`
}

type QualificationForCreate struct {
	// QualificationImgID 资质图片id，通过【上传广告主图片】获得（需指定 upload_to 为 AD）
	QualificationImgID string `json:"qualification_img_id,omitempty"`
	// AdQualificationType 资质图片类型
	// AUTHORIZATION 肖像、商标或其他授权
	// CERTIFY 专利证明
	// TRADEMARK_CERT 商标证
	// AD_DATA_CERT 广告内容中的数据证明
	// INSPECT_REPORT 质检报告
	AdQualificationType enum.AdQualificationType `json:"ad_qualification_type,omitempty"`
}

// Encode implement PostRequest interface
func (r QualificationCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
