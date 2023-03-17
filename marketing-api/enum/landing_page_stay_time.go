package enum

// LandingPageStayTime 店铺停留时长，单位为毫秒，当external_action为AD_CONVERT_TYPE_STAY_TIME时有效且必填
type LandingPageStayTime int

const (
	LandingPageStayTime_1000  LandingPageStayTime = 1000
	LandingPageStayTime_5000  LandingPageStayTime = 5000
	LandingPageStayTime_12000 LandingPageStayTime = 12000
	LandingPageStayTime_20000 LandingPageStayTime = 20000
	LandingPageStayTime_30000 LandingPageStayTime = 30000
	LandingPageStayTime_40000 LandingPageStayTime = 40000
	LandingPageStayTime_50000 LandingPageStayTime = 50000
	LandingPageStayTime_60000 LandingPageStayTime = 60000
)
