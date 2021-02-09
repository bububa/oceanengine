package enum

type CampaignStatus string

const (
	CAMPAIGN_STATUS_ENABLE                   CampaignStatus = "CAMPAIGN_STATUS_ENABLE"                   // 启用
	CAMPAIGN_STATUS_DISABLE                  CampaignStatus = "CAMPAIGN_STATUS_DISABLE"                  // 暂停
	CAMPAIGN_STATUS_DELETE                   CampaignStatus = "CAMPAIGN_STATUS_DELETE"                   // 删除
	CAMPAIGN_STATUS_ALL                      CampaignStatus = "CAMPAIGN_STATUS_ALL"                      // 所有包含已删除
	CAMPAIGN_STATUS_NOT_DELETE               CampaignStatus = "CAMPAIGN_STATUS_NOT_DELETE"               // 所有不包含已删除（状态过滤默认值）
	CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED CampaignStatus = "CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED" // 超出广告主日预算
)
