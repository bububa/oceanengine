package enum

// 类目词级别/抖音达人级别
type AudienceLevel string

const (
	FIRST_LEVEL   AudienceLevel = "FIRST_LEVEL"   // 一级类目词/一级达人分类
	SECOND_LEVEL  AudienceLevel = "SECOND_LEVEL"  // 二级类目词/二级达人分类
	THIRD_LEVEL   AudienceLevel = "THIRD_LEVEL"   // 三级类目词/三级达人分类
	FOURTH_LEVEL  AudienceLevel = "FOURTH_LEVEL"  // 四级类目词
	KEYWORD_AWEME AudienceLevel = "KEYWORD_AWEME" // 关键词/抖音达人
)
