package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取否定词列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectIDs 项目id列表，最多100个
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.ProjectIDs) > 0 {
		values.Set("project_ids", string(util.JSONMarshal(r.ProjectIDs)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 获取否定词列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List 项目否定词列表
		List []Word `json:"projects_privative,omitempty"`
	} `json:"data,omitempty"`
}
