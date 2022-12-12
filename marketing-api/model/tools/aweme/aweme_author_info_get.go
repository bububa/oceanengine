package aweme

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeAuthorInfoGetRequest 查询抖音号id对应的达人信息 API Request
type AwemeAuthorInfoGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LabelIDs 抖音号id列表，取值大小：1～50；label_id即计划中设置的抖音达人账号的id
	LabelIDs []uint64 `json:"label_ids,omitempty"`
	// Behaviors 抖音用户行为类型
	Behaviors []enum.Behavior `json:"behaviors,omitempty"`
}

// Encode implement GetRequest interface
func (r AwemeAuthorInfoGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.LabelIDs)
	values.Set("label_ids", string(ids))
	if len(r.Behaviors) > 0 {
		buf, _ := json.Marshal(r.Behaviors)
		values.Set("behaviors", string(buf))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AwemeAuthorInfoGetResponse 查询抖音号id对应的达人信息 API Response
type AwemeAuthorInfoGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		List []Author `json:"authors,omitempty"`
	} `json:"data,omitempty"`
}
