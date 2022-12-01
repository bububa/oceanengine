package enum

// CampaignType 广告组类型
type CampaignType string

const (
	// CampaignType_FEED 信息流
	CampaignType_FEED CampaignType = "FEED"
	// CampaignType_SEARCH 搜索广告
	CampaignType_SEARCH CampaignType = "SEARCH"
	// CampaignType_ALL
	CampaignType_ALL CampaignType = "ALL"
)
