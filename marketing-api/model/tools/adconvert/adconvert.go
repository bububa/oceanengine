package adconvert

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// AdConvert 转化数据
type AdConvert struct {
	// ConvertType 跟踪方式, 即原转化接口中转化来源
	ConvertType enum.AdConvertSource `json:"convert_type,omitempty"`
	// Disabled 是否禁用, true 表示已经禁用，false 表示可用
	Disabled bool `json:"disabled,omitempty"`
	// ExternalActions 转化来源下的转化目标列表
	ExternalActions []ExternalAction `json:"external_actions,omitempty"`
}

// ExternalAction 转化来源下的转化目标
type ExternalAction struct {
	// ConvertID 转化目标ID，返回自定义转化目标ID，预定义转化目标返回为Null，对应数字值可根据external_action参考【枚举值-转化类型】
	ConvertID model.Uint64 `json:"convert_id,omitempty"`
	// ActionTrackURL 转化监测连接
	ActionTrackURL string `json:"action_track_url,omitempty"`
	// Disabled 转化是否禁用,true 表示已经禁用，false 表示可用
	Disabled bool `json:"disabled,omitempty"`
	// ExternalAction 转化目标，返回预定义的转化类型，详见【附录-枚举值-转化类型】，即原转化接口中 convert_type
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// ExternalActionName 转化类型名称
	ExternalActionName string `json:"external_action_name,omitempty"`
	// ExternalActions 多转化目标
	ExternalActions []enum.AdConvertType `json:"external_actions,omitempty"`
	// ExternalName 自定义转化名称
	ExternalName string `json:"external_name,omitempty"`
	// Source 转化目标创建来源
	Source enum.ExternalActionSource `json:"source,omitempty"`
	// Belong 线索通来源
	Belong []enum.ExternalActionBelong `json:"belong,omitempty"`
	// ConvertDataType 转化统计方式
	ConvertDataType []enum.AdConvertDataType `json:"convert_data_type,omitempty"`
	// DeepExternalActions 深度转化转化目标
	DeepExternalActions []DeepExternalAction `json:"deep_external_actions,omitempty"`
}

// DeepExternalAction 深度转化转化目标
type DeepExternalAction struct {
	// DeepExternalAction 深度转化的转化目标
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
	// DeepExternalName 深度转化转化名称
	DeepExternalName string `json:"deep_external_name,omitempty"`
	// Disabled 深度转化是否禁用,true 表示已经禁用，false 表示可用
	Disabled bool `json:"disabled,omitempty"`
}
