package customaudience

import "github.com/bububa/oceanengine/marketing-api/enum/local"

// CustomAudience 人群包
type CustomAudience struct {
	// CustomAudienceID 人群包ID
	CustomAudienceID uint64 `json:"custom_audience_id,omitempty"`
	// Name 人群包名称
	Name string `json:"name,omitempty"`
	// TagsType 人群包属性
	TagsType local.CustomAudienceTagsType `json:"tags_type,omitempty"`
	// CreateTime 人群包创建时间
	CreateTime string `json:"create_time,omitempty"`
}
