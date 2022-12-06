package interestaction

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// KeywordSuggestRequest 获取行为兴趣推荐关键词
type KeywordSuggestRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ID 类目或关键词id
	ID uint64 `json:"ids,omitempty"`
	// TagType 查询类型：类目还是关键词; 允许值：CATEGORY（类目）、KEYWORD（关键词）
	TagType string `json:"tag_type,omitempty"`
	// TargetingType 查询目标：兴趣还是行为; 允许值：ACTION（行为）、INTEREST（兴趣）
	TargetingType string `json:"targeting_type,omitempty"`
	// ActionScene 行为场景，查询目标为行为时必填，兴趣不生效;允许值：E-COMMERCE、NEWS、APP
	ActionScene []enum.ActionScene `json:"action_scene,omitempty"`
	// ActionDays 行为天数，查询目标为行为时必填，兴趣不生效; 允许值：7, 15, 30, 60, 90, 180, 365
	ActionDays int `json:"action_days,omitempty"`
}

// Encode implement GetRequest interface
func (r KeywordSuggestRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("id", strconv.FormatUint(r.ID, 10))
	values.Set("tag_type", r.TagType)
	values.Set("targeting_type", r.TargetingType)
	if len(r.ActionScene) > 0 {
		scene, _ := json.Marshal(r.ActionScene)
		values.Set("action_scene", string(scene))
	}
	if r.ActionDays > 0 {
		values.Set("action_days", strconv.Itoa(r.ActionDays))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// KeywordSuggestResponse 获取行为兴趣推荐关键词 API Response
type KeywordSuggestResponse struct {
	model.BaseResponse
	// Data json返回值
	Data KeywordSuggestResponseData `json:"data,omitempty"`
}

// KeywordSuggestResponseData json返回值
type KeywordSuggestResponseData struct {
	// Keywords 关键词列表
	Keywords []Object `json:"keywords,omitempty"`
}
