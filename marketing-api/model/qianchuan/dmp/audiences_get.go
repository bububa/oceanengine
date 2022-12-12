package dmp

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AudiencesGetRequest 查询人群包列表 API Request
type AudiencesGetRequest struct {
	// AdvertiserID 千川广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// RetargetingTagsType 人群包类型，枚举值：0：不限营销目标的平台精选人群包，1：自定义人群包
	RetargetingTagsType int `json:"retargeting_tags_type,omitempty"`
	// Offset 偏移,类似于SQL中offset(起始为0,翻页时new_offset=old_offset+limit），默认值：0，取值范围:≥ 0
	Offset int `json:"offset,omitempty"`
	// Limit 返回数据量，默认值：100，取值范围：1-100
	Limit int `json:"limit,omitempty"`
}

// Encode implement GetRequest interface
func (r AudiencesGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("retargeting_tags_type", strconv.Itoa(r.RetargetingTagsType))
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

// AudiencesGetResponse 查询人群包列表 API Response
type AudiencesGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AudiencesGetResponseData `json:"data,omitempty"`
}

// AudiencesGetResponseData json返回值
type AudiencesGetResponseData struct {
	// Offset 下一次查询的偏移,类似于SQL中offset(起始为0,翻页时new_offset=old_offset+limit），返回0时，代表已查询到最后一页
	Offset int `json:"offset,omitempty"`
	// TotalNum 总的人群包数量
	TotalNum int `json:"total_num,omitempty"`
	// RetargetingTags 人群包列表
	RetargetingTags []RetargetingTag `json:"retargeting_tags,omitempty"`
}
