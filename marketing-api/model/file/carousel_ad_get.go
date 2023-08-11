package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CarouselAdGetRequest 获取同主体下广告主图集素材 API Request
type CarouselAdGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CarouselIDs 图集id
	CarouselIDs []uint64 `json:"carousel_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r CarouselAdGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("carousel_ids", string(util.JSONMarshal(r.CarouselIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CarouselAdGetResponse 获取同主体下广告主图集素材 API Response
type CarouselAdGetResponse struct {
	model.BaseResponse
	Data *CarouselAdGetResult `json:"data,omitempty"`
}

type CarouselAdGetResult struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Carousels 图集信息
	Carousels []Carousel `json:"carousels,omitempty"`
}
