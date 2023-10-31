package keyword

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

type IListFilter interface {
	GetID() uint64
}

// GetRequest 获取关键词列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤结构
	Filtering *GetFiltering `json:"filtering,omitempty"`
}

// GetFiltering 过滤结构
type GetFiltering struct {
	// AdID 待过滤的广告计划ID
	AdID uint64 `json:"ad_id,omitempty"`
}

func (f GetFiltering) GetID() uint64 {
	return f.AdID
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filter, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filter))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 获取关键词列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List 关键词列表
		List []Keyword `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
