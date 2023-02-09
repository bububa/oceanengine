package site

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建橙子建站站点 API Request
type CreateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 站点名称，范围：长度 >= 1
	Name string `json:"name,omitempty"`
	// Bricks 站点业务数据。具体的格式见下方的7大组件格式，list中每个object在站点页面上按照顺序从上到下排列。
	Bricks []Brick `json:"bricks,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建橙子建站站点 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// SiteID 站点id
		SiteID string `json:"site_id,omitempty"`
	} `json:"data,omitempty"`
}
