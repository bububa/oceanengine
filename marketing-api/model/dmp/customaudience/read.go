package customaudience

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// ReadRequest 人群包详细信息API Request
type ReadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CustomAudienceIDs 人群包ID列表，长度取值范围：1-100
	CustomAudienceIds []uint64 `json:"custom_audience_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r ReadRequest) Encode() string {
	values := &url.Values{}
	values.Add("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.CustomAudienceIds != nil {
		ids, _ := json.Marshal(r.CustomAudienceIds)
		values.Add("custom_audience_ids", string(ids))
	}
	return values.Encode()
}

// ReadResponse 人群包详细信息API Response
type ReadResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *ReadResponseData `json:"data,omitempty"`
}

// ReadResponseData json返回值
type ReadResponseData struct {
	// CustomAudienceList 人群包列表数据
	CustomAudienceList []CustomAudience `json:"custom_audience_list,omitempty"`
}
