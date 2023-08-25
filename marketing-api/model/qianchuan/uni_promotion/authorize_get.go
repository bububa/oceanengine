package unipromotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AuthorizedGetRequest 获取可投全域推广抖音号列表 API Request
type AuthorizedGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码
	// 默认值： 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	// 默认值：10，最大值：100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r AuthorizedGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AuthorizedGetResponse 获取可投全域推广抖音号列表 API Response
type AuthorizedGetResponse struct {
	model.BaseResponse
	Data *AuthorizedGetResult `json:"data,omitempty"`
}

type AuthorizedGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AwemeIDList 抖音号列表
	AwemeIDList []aweme.Aweme `json:"aweme_id_list,omitempty"`
}
