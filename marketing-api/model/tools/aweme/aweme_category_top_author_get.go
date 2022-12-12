package aweme

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeCategoryTopAuthorGetRequest 查询抖音类目下的推荐达人 API Request
type AwemeCategoryTopAuthorGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CategoryID 类目id，一级，二级，三级类目id均可
	CategoryID uint64 `json:"category_id,omitempty"`
	// Behaviors 抖音用户行为类型
	Behaviors []enum.Behavior `json:"behaviors,omitempty"`
}

// Encode implement GetRequest interface
func (r AwemeCategoryTopAuthorGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("category_id", strconv.FormatUint(r.CategoryID, 10))
	if len(r.Behaviors) > 0 {
		buf, _ := json.Marshal(r.Behaviors)
		values.Set("behaviors", string(buf))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AwemeCategoryTopAuthorGetResponse 查询抖音类目下的推荐达人 API Response
type AwemeCategoryTopAuthorGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		List []Author `json:"authors,omitempty"`
	} `json:"data,omitempty"`
}
