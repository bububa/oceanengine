package file

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoAdGetRequest 获取同主体下广告主视频素材API Request
type VideoAdGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// VideoIDs 视频ids，长度限制 1~60
	VideoIDs []string `json:"video_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoAdGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.VideoIDs) > 0 {
		ids, _ := json.Marshal(r.VideoIDs)
		values.Set("video_ids", string(ids))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
