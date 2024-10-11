package poi

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MultiPoiIDsGetRequest 根据多门店ID拉取门店ID API Request
type MultiPoiIDsGetRequest struct {
	// LocalAccountID 本地推广告主ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// MultiPoiIDs 多门店id，可通过【获取项目详情】接口获取项目推广的多门店ID
	// 数量上限：50
	MultiPoiIDs []uint64 `json:"multi_poi_ids,omitempty"`
	// NeedEnable 是否仅查询当前在投门店，即过滤不在投放中的门店（可能存在部分门店，因门店已下线/门店下无团购/被全域推广转移，被从计划中剔除，不在投放中），允许值：
	// true 仅查询在投放中的门店
	// false查询多门店ID对应的所有门店（默认值）
	NeedEnable bool `json:"need_enable,omitempty"`
}

// Encode implements GetRequest interface
func (r MultiPoiIDsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	values.Set("multi_poi_ids", string(util.JSONMarshal(r.MultiPoiIDs)))
	if r.NeedEnable {
		values.Set("need_enable", "true")
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// MultiPoiIDsGetResponse 根据多门店ID拉取门店ID API Response
type MultiPoiIDsGetResponse struct {
	model.BaseResponse
	Data *MultiPoiIDsGetResult `json:"data,omitempty"`
}

type MultiPoiIDsGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// MultiPoiInfo 多门店信息
	MultiPoiInfo []MultiPoi `json:"multi_poi_info,omitempty"`
}
