package tools

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdQualityGetRequest 查询广告质量度 API Request
type AdQualityGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 广告id列表
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r AdQualityGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	buf, _ := json.Marshal(r.AdIDs)
	values.Set("ad_ids", string(buf))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AdQualityGetResponse 查询广告质量度 API Response
type AdQualityGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		List []AdQuality `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// AdQuality 广告质量度
type AdQuality struct {
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// QualityScore 计划综合质量得分
	QualityScore float64 `json:"quality_score,omitempty"`
	// CtrScore 创意质量得分
	CtrScore float64 `json:"ctr_score,omitempty"`
	// WebScore 落地页响应得分
	WebScore float64 `json:"web_score,omitempty"`
	// CvrScore 落地页素材得分
	CvrScore float64 `json:"cvr_score,omitempty"`
}
