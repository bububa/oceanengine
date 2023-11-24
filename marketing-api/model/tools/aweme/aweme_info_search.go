package aweme

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeInfoSearchRequest 查询抖音帐号和类目信息 API Request
type AwemeInfoSearchRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// QueryWord 搜索词
	QueryWord string `json:"query_word,omitempty"`
	// Behaviors 抖音用户行为类型
	Behaviors []enum.Behavior `json:"behaviors,omitempty"`
}

// Encode implement GetRequest interface
func (r AwemeInfoSearchRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("query_word", r.QueryWord)
	if len(r.Behaviors) > 0 {
		buf, _ := json.Marshal(r.Behaviors)
		values.Set("behaviors", string(buf))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AwemeInfoSearchResponse 查询抖音帐号和类目信息 API Response
type AwemeInfoSearchResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		List []AwemeInfoSearchResult `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// AwemeInfoSearchResult 抖音账号类目信息
type AwemeInfoSearchResult struct {
	Categories []Category `json:"categories,omitempty"`
	Authors    []Author   `json:"authors,omitempty"`
}
