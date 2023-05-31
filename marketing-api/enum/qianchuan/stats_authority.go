package qianchuan

// StatsAuthority 需要查询的广告账户维度，直客账户支持查看官方抖音号下同主体全部广告账户的数据表现
type StatsAuthority string

const (
	// StatsAuthority_QUALIFICATION 同主体账户, 仅直客账户支持
	StatsAuthority_QUALIFICATION StatsAuthority = "QUALIFICATION"
	// StatsAuthority_CURRENT 当前账户
	StatsAuthority_CURRENT StatsAuthority = "CURRENT"
)
