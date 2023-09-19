package enterprise

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// InfoRequest 获取企业号信息 API Request
type InfoRequest struct {
	// EDouyinIDs 企业号ids，一次传入id数量不超过10个
	EDouyinIDs []string `json:"e_douyin_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r InfoRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("e_douyin_ids", string(util.JSONMarshal(r.EDouyinIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// InfoResponse 获取企业号信息 API Response
type InfoResponse struct {
	model.BaseResponse
	Data struct {
		List []Enterprise `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
