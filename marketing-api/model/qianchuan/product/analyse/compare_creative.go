package analyse

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CompareCreativeRequest 商品竞争分析详情-创意比对 API Request
type CompareCreativeRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProductID 商品ID
	ProductID uint64 `json:"product_id,omitempty"`
	// TimeRange 商品数据范围，可选值
	// 30：30天
	// 15：15天
	// 7：7天（默认）
	TimeRange int `json:"time_range,omitempty"`
}

// Encode implement GetRequest interface
func (r CompareCreativeRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("product_id", strconv.FormatUint(r.ProductID, 10))
	if r.TimeRange > 0 {
		values.Set("time_range", strconv.Itoa(r.TimeRange))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CompareCreativeResponse 商品竞争分析详情-创意比对 API Response
type CompareCreativeResponse struct {
	model.BaseResponse
	Data *CompareCreative `json:"data,omitempty"`
}

type CompareCreative struct {
	// OwnProductCreative 我的商品创意数据
	OwnProductCreative []Creative `json:"own_product_creative,omitempty"`
	// SimilarProductCreative 相似商品创意数据
	SimilarProductCreative []Creative `json:"similar_product_creative,omitempty"`
}
