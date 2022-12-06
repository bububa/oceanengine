package tools

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdStatExtraInfoGetRequest 查询广告计划学习期状态 API Request
type AdStatExtraInfoGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 广告id列表, 最多传100个广告计划id
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r AdStatExtraInfoGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	buf, _ := json.Marshal(r.AdIDs)
	values.Set("ad_ids", string(buf))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AdStatExtraInfoGetResponse 查询广告计划学习期状态 API Request
type AdStatExtraInfoGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data []AdStatExtraInfo `json:"data,omitempty"`
}

// AdStatExtraInfo 广告计划学习期
type AdStatExtraInfo struct {
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// LearningPhase  学习期状态
	LearningPhase enum.LearningPhase `json:"learning_phase,omitempty"`
}
