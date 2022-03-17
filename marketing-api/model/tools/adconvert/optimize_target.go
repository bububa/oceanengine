package adconvert

import "github.com/bububa/oceanengine/marketing-api/enum"

// OptimizeTarget 转化数据
type OptimizeTarget struct {
	// MarketingPurpose 营销目的，允许值：UNLIMITED不限，CONVERSION行动转化， INTENTION用户意向，ACKNOWLEDGE品牌认知
	MarketingPurpose enum.MarketingPurpose `json:"marketing_purpose,omitempty"`
	// Disabled 是否禁用, true 表示已经禁用，false 表示可用
	Disabled bool `json:"disabled,omitempty"`
	// Converts 优化来源下的转化目标列表
	Converts []AdConvert `json:"converts,omitempty"`
}
