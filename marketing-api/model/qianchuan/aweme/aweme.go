package aweme

import "github.com/bububa/oceanengine/marketing-api/enum"

// Aweme 抖音号
type Aweme struct {
	// ID 抖音id，用于创建计划，拉取抖音号视频素材时入参
	ID uint64 `json:"aweme_id,omitempty"`
	// Avatar 抖音头像
	Avatar string `json:"aweme_avatar,omitempty"`
	// ShowID 抖音号，即客户在手机端上看到的抖音号，若向客户披露抖音号请使用该字段
	ShowID string `json:"aweme_show_id,omitempty"`
	// Name 抖音号名称
	Name string `json:"aweme_name,omitempty"`
	// Status 抖音号带货状态，返回值：NORMAL 可以正常投放;ANCHOR_FORBID 带货口碑分过低，暂时无法创建计划;ANCHOR_REACH_UPPER_LIMIT_TODAY 带货分过低或暂无带货分，可以创建计划，但无法产生消耗，带货分恢复正常后可正常消耗
	Status enum.AwemeStatus `json:"aweme_status,omitempty"`
	// BindingType 抖音号关系类型
	BindingType []enum.AwemeBindingType `json:"binding_type,omitempty"`
}
