package dmp

import (
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// Orientation 定向包
type Orientation struct {
	ad.Audience
	// OrientationName 定向包名称
	OrientationName string `json:"orientation_name,omitempty"`
	// AdNum 关联计划数
	AdNum int `json:"ad_num,omitempty"`
	// OrientationInfo 定向包描述
	OrientationInfo string `json:"orientation_info,omitempty"`
	// InActiveRetargetingTags 失效人群包
	InActiveRetargetingTags []InActiveRetargetingTag `json:"InActive_retargeting_tags,omitempty"`
}

// InActiveRetargetingTag 失效人群包
type InActiveRetargetingTag struct {
	// RetargetingTag 失效人群包id
	RetargetingTag uint64 `json:"retargeting_tag,omitempty"`
	// RetargetingName 失效人群包名称
	RetargetingName string `json:"retargeting_name,omitempty"`
	// InActiveType 失效类型，枚举值：
	// EXPIRE: 人群包过期
	// TAG_OFFLINE: 人群包tag下线
	// MANUAL_OFFLINE: 精选人群包手动下线
	InActiveType qianchuan.InActiveRetargetingType `json:"InActive_type,omitempty"`
}
