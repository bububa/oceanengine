package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/util"
)

// SuggestRequest 搜索快投关键词推荐 API Request
type SuggestRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectID 项目id
	ProjectID uint64 `json:"project_id,omitempty"`
	// QueryList 以词推词，最多只传入10个，长度要求30字内 ,可传入行业名称获取行业推词
	QueryList []string `json:"query_list,omitempty"`
}

// Encode implement GetRequest interface
func (r SuggestRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("project_id", strconv.FormatUint(r.ProjectID, 10))
	if len(r.QueryList) > 0 {
		values.Set("query_list", string(util.JSONMarshal(r.QueryList)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
