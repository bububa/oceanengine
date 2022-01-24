package creativeword

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// SelectRequest 查询动态创意词包 API Request
type SelectRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreativeWordIDs 创意词包id列表，如不填默认返回所有创意词包
	CreativeWordIDs []uint64 `json:"creative_word_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r SelectRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	buf, _ := json.Marshal(r.CreativeWordIDs)
	values.Set("creative_word_ids", string(buf))
	return values.Encode()
}

// SelectResponse 查询动态创意词包 API Response
type SelectResponse struct {
	model.BaseResponse
	Data struct {
		// List 创意词包列表
		List []CreativeWord `json:"creative_word,omitempty"`
	} `json:"data,omitempty"`
}
