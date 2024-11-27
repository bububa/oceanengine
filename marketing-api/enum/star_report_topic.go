package enum

// StarReportTopic 数据主题
type StarReportTopic string

const (
	// StarReportTopic_BASIC_DATA 基础信息
	StarReportTopic_BASIC_DATA StarReportTopic = "BASIC_DATA"
	// StarReportTopic_FLOW_DATA 流量表现
	StarReportTopic_FLOW_DATA StarReportTopic = "FLOW_DATA"
	// StarReportTopic_CONVERT_DATA 转化表现
	StarReportTopic_CONVERT_DATA StarReportTopic = "CONVERT_DATA"
	// StarReportTopic_SEARCH_DATA 搜索表现
	StarReportTopic_SEARCH_DATA StarReportTopic = "SEARCH_DATA"
	// StarReportTopic_RECOMMEND_DATA 种草表现
	StarReportTopic_RECOMMEND_DATA StarReportTopic = "RECOMMEND_DATA"
	// StarReportTopic_DY_SHOP_DATA 抖音进店
	StarReportTopic_DY_SHOP_DATA StarReportTopic = "DY_SHOP_DATA"
	// StarReportTopic_USER_DISTRIBUTION_DATA 用户画像、
	StarReportTopic_USER_DISTRIBUTION_DATA StarReportTopic = "USER_DISTRIBUTION_DATA"
	// StarReportTopic_INDEX_SCORE_DATA 指数得分
	StarReportTopic_INDEX_SCORE_DATA StarReportTopic = "INDEX_SCORE_DATA"
	// StarReportTopic_COMMENT_DATA 评论数据
	StarReportTopic_COMMENT_DATA StarReportTopic = "COMMENT_DATA"
)
