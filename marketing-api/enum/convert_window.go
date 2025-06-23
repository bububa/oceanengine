package enum

// ConvertWindow 时间窗口设置
type ConvertWindow string

const (
	// ConvertWindow_ONE ：1天
	ConvertWindow_ONE ConvertWindow = "ONE"
	// ConvertWindow_THREE :3天
	ConvertWindow_THREE ConvertWindow = "THREE"
	// ConvertWindow_SEVEN ：7天
	ConvertWindow_SEVEN ConvertWindow = "SEVEN"
	// ConvertWindow_FOURTEEN ：14天
	ConvertWindow_FOURTEEN ConvertWindow = "FOURTEEN"
	// ConvertWindow_THIRTY ：30天
	ConvertWindow_THIRTY ConvertWindow = "THIRTY"
	// ConvertWindow_SIXTY ：60天
	ConvertWindow_SIXTY ConvertWindow = "SIXTY"
	// ConvertWindow_NINETY ：90天
	ConvertWindow_NINETY ConvertWindow = "NINETY"
	// ConvertWindow_ONE_HUNDRED_EIGHTY：180天
	ConvertWindow_ONE_HUNDRED_EIGHTY ConvertWindow = "ONE_HUNDRED_EIGHTY"
	// ConvertWindow_PERMANENT ：永久
	ConvertWindow_PERMANENT ConvertWindow = "PERMANENT"
)
