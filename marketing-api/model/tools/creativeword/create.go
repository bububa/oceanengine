package creativeword

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建动态创意词包 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 创意词包名称
	Name string `json:"name,omitempty"`
	// DefaultWord 默认词
	DefaultWord string `json:"default_word,omitempty"`
	// Words 替换词
	Words []string `json:"words,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建动态创意词包 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// CreativeWordID 创意词包ID
		CreativeWordID uint64 `json:"creative_word_id,omitempty"`
	} `json:"data,omitempty"`
}
