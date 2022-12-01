package enum

// District 地域类型
type District string

const (
	// District_CITY 省市
	District_CITY District = "CITY"
	// District_COUNTY 区县
	District_COUNTY District = "COUNTY"
	// District_REGION 行政区域
	District_REGION District = "REGION"
	// District_OVERSEA 海外区域
	District_OVERSEA District = "OVERSEA"
	// District_BUSINESS_DISTRICT 商圈
	District_BUSINESS_DISTRICT District = "BUSINESS_DISTRICT"
	// District_NONE 不限
	District_NONE District = "NONE"
)
