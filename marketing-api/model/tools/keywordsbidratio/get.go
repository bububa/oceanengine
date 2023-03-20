package keywordsbidratio

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 查询优词提量系数信息 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 筛选项
	Filtering *GetFilter `json:"filtering,omitempty"`
}

// GetFilter 筛选项
type GetFilter struct {
	// ProjectIDs 生效的项目id，最多50个
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
	// Words 关键词，最多50个
	Words []string `json:"words,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		bs, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(bs))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 查询优词提量系数信息 API Response
type GetResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data struct {
		// List
		List []Keyword `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// Keyword 优词信息
type Keyword struct {
	// PromotionWordID 优词计划ID
	PromotionWordID string `json:"promotion_word_id,omitempty"`
	// Word 优词
	Word string `json:"word,omitempty"`
	// Dimension 生效维度
	// ADVERTISER: 账户维度
	// PROJECT: 项目维度
	Dimension string `json:"dimension,omitempty"`
	// BidRatio 优词提量系数
	BidRatio float64 `json:"bid_ratio,omitempty"`
	// ProjectNum 生效的项目数量，账户维度不返回数据
	ProjectNum int `json:"project_num,omitempty"`
}
