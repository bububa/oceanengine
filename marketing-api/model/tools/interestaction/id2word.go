package interestaction

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// Id2WordRequest 兴趣行为类目关键词id转词 API Request
type Id2WordRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// IDs 类目或关键词id列表
	IDs []uint64 `json:"ids,omitempty"`
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
func (r Id2WordRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.IDs)
	values.Set("ids", string(ids))
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

// Id2WordResponse 兴趣行为类目关键词id转词 API Response
type Id2WordResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *Id2WordResponseData `json:"data,omitempty"`
}

// Id2WordResponseData json返回值
type Id2WordResponseData struct {
	// Categories 类目列表
	Categories []Object `json:"categories,omitempty"`
	// Keywords 关键词列表
	Keywords []Object `json:"keywords,omitempty"`
}
