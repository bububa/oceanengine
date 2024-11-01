package report

// Report 报表
type Report struct {
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// ProjectName 项目名称
	ProjectName string `json:"project_name,omitempty"`
	// PromotionID 广告id
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// PromotionName 广告名称
	PromotionName string `json:"promotion_name,omitempty"`
	// MaterialID 素材ID
	MaterialID uint64 `json:"material_id,omitempty"`
	// MaterialName 素材名称
	MaterialName string `json:"material_name,omitempty"`
	// StatTimeDay 时间-天
	StatTimeDay string `json:"stat_time_day,omitempty"`
	// StatTimeHour 时间-小时
	StatTimeHour string `json:"stat_time_hour,omitempty"`
	// StatCost 消耗(元)
	StatCost float64 `json:"stat_cost,omitempty"`
	// ShowCnt 展示次数
	ShowCnt int64 `json:"show_cnt,omitempty"`
	// ClickCnt 点击次数
	ClickCnt int64 `json:"click_cnt,omitempty"`
	// Ctr 点击率
	Crt float64 `json:"ctr,omitempty"`
	// CpcPlatform 点击均价(元)
	CpcPlatform float64 `json:"cpc_platform,omitempty"`
	// CpmPlatform 平均千次展示费用(元)
	CpmPlatform float64 `json:"cpm_platform,omitempty"`
	// ConvertCnt 转化数
	ConvertCnt int64 `json:"convert_cnt,omitempty"`
	// ConversionRate 转化率
	ConversionRate float64 `json:"conversion_rate,omitempty"`
	// ConversionCost 转化成本(元)
	ConversionCost float64 `json:"conversion_cost,omitempty"`
	// AttributionConvertCnt 转化数(计费时间)
	AttributionConvertCnt int64 `json:"attribution_convert_cnt,omitempty"`
	// AttributionConversionRate 转化率(计费时间)
	AttributionConversionRate float64 `json:"attribution_conversion_rate,omitempty"`
	// AttributionConversionCost 转化成本(计费时间)
	AttributionConversionCost float64 `json:"attribution_conversion_cost,omitempty"`
	// FormCnt 表单提交数
	FormCnt int64 `json:"form_cnt,omitempty"`
	// CluePayOrderCnt 团购线索数
	CluePayOrderCnt int64 `json:"clue_pay_order_cnt,omitempty"`
	// ClueMessageCount 私信留资数
	ClueMessageCount int64 `json:"clue_message_count,omitempty"`
	// PhoneConfirmCnt 电话拨打数
	PhoneConfirmCnt int64 `json:"phone_confirm_cnt,omitempty"`
	// PhoneConnectCnt 电话接通数
	PhoneConnectCnt int64 `json:"phone_connect_cnt,omitempty"`
	// MessageActionCnt 私信咨询数
	MessageActionCnt int64 `json:"message_action_cnt,omitempty"`
	// IntentionFormCnt 意向表单数
	IntentionFormCnt int64 `json:"intention_form_cnt,omitempty"`
	// IntentionPhoneCnt 意向话单数
	IntentionPhoneCnt int64 `json:"intention_phone_cnt,omitempty"`
	// IntentionMessageClueCnt 意向咨询数
	IntentionMessageClueCnt int64 `json:"intention_message_clue_cnt,omitempty"`
	// AttributionFormCnt 表单提交数(计费时间)
	AttributionFormCnt int64 `json:"attribution_form_cnt,omitempty"`
	// AttributionIntentionMessageClueCnt 团购线索数(计费时间)
	AttributionCluePayOrderCnt int64 `json:"attribution_clue_pay_order_cnt,omitempty"`
	// AttributeClueMessageCnt 私信留资数(计费时间)
	AttributeClueMessageCnt int64 `json:"attribute_clue_message_count,omitempty"`
	// AttributionPhoneConfirmCnt 电话拨打数(计费时间)
	AttributionPhoneConfirmCnt int64 `json:"attribution_phone_confirm_cnt,omitempty"`
	// AttributePhoneConnectCnt 电话接通数(计费时间)
	AttributePhoneConnectCnt int64 `json:"attribute_phone_connect_cnt,omitempty"`
	// AttributeMessactionActionCnt 私信咨询数(计费时间)
	AttributeMessactionActionCnt int64 `json:"attribute_messaction_action_cnt,omitempty"`
	// AttributeIntentionFormCnt 意向表单数(计费时间)
	AttributionIntentionFormCnt int64 `json:"attribution_intention_form_cnt,omitempty"`
	// AttributionIntentionPhoneCnt 意向话单数(计费时间)
	AttributionIntentionPhoneCnt int64 `json:"attribution_intention_phone_cnt,omitempty"`
	// AttributionIntentionMessageClueCnt 意向咨询数(计费时间)
	AttributionIntentionMessageClueCnt int64 `json:"attribution_intention_message_clue_cnt,omitempty"`
	// DyLike 视频点赞次数
	DyLike int64 `json:"dy_like,omitempty"`
	// DyComment 视频评论次数
	DyComment int64 `json:"dy_comment,omitempty"`
	// DyShare 视频分享次数
	DyShare int64 `json:"dy_share,omitempty"`
	// DyCollect 视频收藏次数
	DyCollect int64 `json:"dy_collect,omitempty"`
	// DyFollow 粉丝量
	DyFollow int64 `json:"dy_follow,omitempty"`
	// TotalPlay 视频播放次数
	TotalPlay int64 `json:"total_play,omitempty"`
	// PlayDuration3s 视频3s播放次数
	PlayDuration3s int64 `json:"play_duration_3s,omitempty"`
	// PlayDuration5s 视频5s播放次数
	PlayDuration5s int64 `json:"play_duration_5s,omitempty"`
	// Play25FeedBreak 视频25%进度播放次数
	Play25FeedBreak int64 `json:"play_25_feed_break,omitempty"`
	// Play50FeedBreak 视频50%进度播放次数
	Play50FeedBreak int64 `json:"play_50_feed_break,omitempty"`
	// Play75FeedBreak 视频75%进度播放次数
	Play75FeedBreak int64 `json:"play_75_feed_break,omitempty"`
	// PlayOver 视频播放完成次数
	PlayOver int64 `json:"play_over,omitempty"`
	// PlayDuration5sShowCntRate 视频5s播放率
	PlayDuration5sShowCntRate float64 `json:"play_duration_5s_show_cnt_rate,omitempty"`
	// PlayOverRate 视频完播率
	PlayOverRate float64 `json:"play_over_rate,omitempty"`
	// DyLikeRate 视频点赞率
	DyLikeRate float64 `json:"dy_like_rate,omitempty"`
	// LubanLiveEnterCnt 直播间观看次数
	LubanLiveEnterCnt int64 `json:"luban_live_enter_cnt,omitempty"`
	// LiveWatchOneMinuteCount 直播间超1分钟停留次数
	LiveWatchOneMinuteCount int64 `json:"live_watch_one_minute_count,omitempty"`
	// LubanLiveCommentCnt 直播间评论次数
	LubanLiveCommentCnt int64 `json:"luban_live_comment_cnt,omitempty"`
	// LubanLiveShareCnt 直播间分享次数
	LubanLiveShareCnt int64 `json:"luban_live_share_cnt,omitempty"`
}