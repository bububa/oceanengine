package diagnosis

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// SuggestionGetRequest 获取计划诊断建议 API Request
type SuggestionGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 广告计划ID列表，最多100个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// Scenes 希望获取建议的场景，允许值：CLEAN 清理低质计划场景、POTENTIAL 获取潜力计划场景
	Scenes []string `json:"scenes,omitempty"`
}

// Encode implement GetRequest interface
func (r SuggestionGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.AdIDs)
	values.Set("ad_ids", string(ids))
	scenes, _ := json.Marshal(r.Scenes)
	values.Set("scene", string(scenes))
	return values.Encode()
}

// SuggestionGetResponse 获取计划诊断建议 API Response
type SuggestionGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *SuggestionGetResponseData `json:"data,omitempty"`
}

type SuggestionGetResponseData struct {
	// DiagnosisID 诊断id
	DiagnosisID string `json:"diagnosis_id,omitempty"`
	// ExpireTimestamp 诊断id的过期时间，格式：YYYY-MM-DD HH:mm:ss
	ExpireTimestamp string `json:"expire_timestamp,omitempty"`
	// SuggestList 获取的诊断建议列表
	SuggestList []Suggest `json:"suggest_list,omitempty"`
}
