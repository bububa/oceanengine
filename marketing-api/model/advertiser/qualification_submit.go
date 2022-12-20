package advertiser

import "github.com/bububa/oceanengine/marketing-api/util"

// QualificationSubmitRequest 提交广告主资质（新版） API Request
type QualificationSubmitRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Subject 主体资质
	Subject *QualificationSubject `json:"subject,omitempty"`
	// Industries 行业资质
	Industries []QualificationIndustry `json:"industries,omitempty"`
}

// Encode implement PostRequest interface
func (r QualificationSubmitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
