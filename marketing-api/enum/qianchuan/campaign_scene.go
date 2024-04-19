package qianchuan

// CampaignScene 营销场景
type CampaignScene string

const (
	// CampaignScene_DAILY_SALE 日常销售（默认）
	CampaignScene_DAILY_SALE CampaignScene = "DAILY_SALE"
	// CampaignScene_NEW_CUSTOMER_TRANSFORMATION 新客转化
	CampaignScene_NEW_CUSTOMER_TRANSFORMATION CampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
	// CampaignScene_LIVE_HEAT 直播间加热
	CampaignScene_LIVE_HEAT CampaignScene = "LIVE_HEAT"
	// CampaignScene_PLANT_GRASS 人群种草
	CampaignScene_PLANT_GRASS CampaignScene = "PLANT_GRASS"
	// CampaignScene_PRODUCT_HEAT 商品加热
	CampaignScene_PRODUCT_HEAT CampaignScene = "PRODUCT_HEAT"
	// CampaignScene_NEW_PRODUCT_BOOST 新品起量
	CampaignScene_NEW_PRODUCT_BOOST CampaignScene = "NEW_PRODUCT_BOOST"
)
