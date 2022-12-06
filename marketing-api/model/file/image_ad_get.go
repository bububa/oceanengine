package file

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/util"
)

// ImageAdGetRequest 获取同主体下广告主图片素材 API Request
type ImageAdGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ImageIDs 图片ids，长度限制1~100
	ImageIDs []string `json:"image_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r ImageAdGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.ImageIDs) > 0 {
		imageIDs, _ := json.Marshal(r.ImageIDs)
		values.Set("image_ids", string(imageIDs))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
