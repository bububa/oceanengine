package aweme

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeSimilarAuthorSearchRequest 查询抖音类似帐号 API Request
type AwemeSimilarAuthorSearchRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音id
	AwemeID string `json:"aweme_id,omitempty"`
	// Behaviors 抖音用户行为类型
	Behaviors []enum.Behavior `json:"behaviors,omitempty"`
}

// Encode implement GetRequest interface
func (r AwemeSimilarAuthorSearchRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", r.AwemeID)
	if len(r.Behaviors) > 0 {
		buf, _ := json.Marshal(r.Behaviors)
		values.Set("behaviors", string(buf))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AwemeSimilarAuthorSearchResponse 查询抖音帐号和类目信息  API Response
type AwemeSimilarAuthorSearchResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		List []Author `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
