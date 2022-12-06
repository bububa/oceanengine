package customaudience

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SelectRequest 人群包列表API Request
type SelectRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SelectType 查询类型,枚举值："0"：该广告主创建的人群包和被推送给该广告主的人群包，"1"：状态为可投放的人群包
	SelectType int `json:"select_type,omitempty"`
	// Offset 偏移,类似于SQL中offset(起始为0,翻页时new_offset=old_offset+limit），默认值：0，取值范围:≥ 0
	Offset int `json:"offset,omitempty"`
	// Limit 返回数据量，默认值：100，取值范围：1-100
	Limit int `json:"limit,omitempty"`
}

// Encode implement GetRequest interface
func (r SelectRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("select_type", strconv.Itoa(r.SelectType))
	if r.Offset >= 0 {
		values.Set("offset", strconv.Itoa(r.Offset))
	}
	if r.Limit > 0 {
		values.Set("limit", strconv.Itoa(r.Limit))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// SelectResponse 人群包列表API Response
type SelectResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *SelectResponseData `json:"data,omitempty"`
}

// SelectResponseData json返回值
type SelectResponseData struct {
	// CustomAudienceList 人群包列表数据
	CustomAudienceList []CustomAudience `json:"custom_audience_list,omitempty"`
	// Offset 偏移,类似于SQL中offset(起始为0,翻页时new_offset=old_offset+limit）
	Offset int `json:"offset,omitempty"`
	// TotalNum 总的人群包数量
	TotalNum int `json:"total_num,omitempty"`
}
