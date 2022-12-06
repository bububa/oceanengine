package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AuthorizedGetRequest 获取千川账户下已授权抖音号 API Request
type AuthorizedGetRequest struct {
	// AdvertiserID 千川广告主账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码.默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量.默认值: 10， 最大值：100
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

// AuthorizedGetResponse 获取千川账户下已授权抖音号 API Response
type AuthorizedGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AuthorizedGetResponseData `json:"data,omitempty"`
}

// AuthorizedGetResponseData json返回值
type AuthorizedGetResponseData struct {
	// PageInfo 分页信息
	PageInfo model.PageInfo `json:"page_info,omitempty"`
	// AwemeList 抖音号列表
	AwemeList []Aweme `json:"aweme_id_list,omitempty"`
}
