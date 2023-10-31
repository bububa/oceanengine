package analyse

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CompareStatsDataRequest 商品竞争分析详情-效果对比 API Request
type CompareStatsDataRequest struct {
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
func (r CompareStatsDataRequest) Encode() string {
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

// CompareStatsDataResponse 商品竞争分析详情-效果对比 API Response
type CompareStatsDataResponse struct {
	model.BaseResponse
	Data *CompareStatsData `json:"data,omitempty"`
}

type CompareStatsData struct {
	// OwnProductData 我的商品数据
	OwnProductData *Metrics `json:"own_product_data,omitempty"`
	// SimilarProductData 相似商品数据
	SimilarProductData *Metrics `json:"similar_product_data,omitempty"`
	// CompeteProductData 全部商品数据
	CompeteProductData *Metrics `json:"compete_product_data,omitempty"`
}
