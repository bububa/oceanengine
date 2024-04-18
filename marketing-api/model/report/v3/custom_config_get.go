package v3

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CustomConfigGetRequest 获取自定义报表可用指标和维度 API Request
type CustomConfigGetRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DataTopics 数据主题列表
	// 枚举值：BASIC_DATA广告基础数据、QUERY_DATA搜索词数据、BIDWORD_DATA关键词数据、MATERIAL_DATA素材数据, PRODUCT_DATA用于表达产品数据主题
	DataTopics []string `json:"data_topics,omitempty"`
}

// Encode implement GetRequest interface
func (r CustomConfigGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.DataTopics) > 0 {
		bs, _ := json.Marshal(r.DataTopics)
		values.Set("data_topics", string(bs))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CustomConfigGetResponse 获取自定义报表可用指标和维度 API Response
type CustomConfigGetResponse struct {
	model.BaseResponse
	Data struct {
		// List 数据主题配置列表
		List []CustomConfig `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// CustomConfig 数据主题配置
type CustomConfig struct {
	// DataTopic 数据主题
	DataTopic string `json:"data_topic,omitempty"`
	// Dimensions 维度列表
	Dimensions []Dimension `json:"dimensions,omitempty"`
	// Metrics 指标列表
	Metrics []Dimension `json:"metrics,omitempty"`
}
