package enterprise

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OperationLogGetRequest 获取企业号推广操作记录 API Request
type OperationLogGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// OpenID 企业号open_id
	OpenID string `json:"open_id,omitempty"`
	// StartTime 查询时间，格式为yyyy-mm-dd ，默认为今天，仅支持查询2020年11月1日之后的数据
	StartTime string `json:"start_time,omitempty"`
	// Offset 偏移，类似于SQL中offset（起始为0，翻页时new_offset=old_offset+limit），默认值：0，取值范围:≥ 0
	Offset int `json:"offset,omitempty"`
	// Limit 返回数据量，默认值：100，取值范围：1-100
	Limit int `json:"limit,omitempty"`
}

// Encode implement GetRequest interface
func (r OperationLogGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("open_id", r.OpenID)
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.Offset > 0 {
		values.Set("offset", strconv.Itoa(r.Offset))
	}
	if r.Limit > 0 {
		values.Set("limit", strconv.Itoa(r.Limit))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// OperationLogGetResponse 获取企业号推广操作记录 API Response
type OperationLogGetResponse struct {
	model.BaseResponse
	Data *OperationLogGetResult `json:"data,omitempty"`
}

type OperationLogGetResult struct {
	// ActionCount 推广操作次数
	ActionCount int `json:"action_count,omitempty"`
	// Offset 偏移
	Offset int `json:"offset,omitempty"`
	// List 推广操作记录列表
	List []OperationLog `json:"list,omitempty"`
}
