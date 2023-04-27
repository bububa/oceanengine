package report

// Metrics 指标数据
type Metrics struct {
	// 展现数据-总花费; 表示广告在投放期内的预估花费金额,当天数据可能会有波动，次日稳定
	Cost float64 `json:"cost,omitempty"`
	// 展现数据-展示数; 广告展示给用户的次数。计算方式：经平台判定有效且被计费的展示次数
	Show int64 `json:"show,omitempty"`
	// 展现数据-平均千次展现费用; 广告平均每一千次展现所付出的费用，计算公式是：总花费/展示数*1000
	AvgShowCost float64 `json:"avg_show_cost,omitempty"`
	// 展现数据-点击数; 当头条用户点击广告素材时，触发点击事件，该事件被认为是一次有效的广告点击
	Click int64 `json:"click,omitempty"`
	// 展现数据-平均点击单价; 广告主为每次点击付出的费用成本，计算公式是：总花费/点击数
	AvgClickCost float64 `json:"avg_click_cost,omitempty"`
	// 展现数据-点击率; 广告被点击的次数占展示次数的百分比。计算方法：点击数/展示数*100%
	Ctr float64 `json:"ctr,omitempty"`
	// 转化数据-转化数;将转化数记录在转化事件发生的时间上。建议广告主考核成本时参考“转化数据（计费时间）”例如您的广告在早上8点进行了展示和点击，用户晚上19点发生了激活行为，我们会把激活数披露在晚上19点
	Convert int64 `json:"convert,omitempty"`
	// 转化数据-转化成本; 广告主为每个转化所付出的平均成本，计算方式：总花费/转化数。当天数据可能会有波动
	ConvertCost float64 `json:"convert_cost,omitempty"`
	// 转化数据-转化率; 广告被用户转化的次数占点击次数的百分比。计算方式：转化数/点击数*100%
	ConvertRate float64 `json:"convert_rate,omitempty"`
	// 转化数据-深度转化数; 将深度转化数记录在转化事件发生的时间上。建议广告主考核深度转化成本时参考“深度转化数（计费时间）”例如您的广告在早上8点进行了展示和点击，用户晚上19点发生了激活行为，我们会把激活数披露在晚上19点
	DeepConvert int64 `json:"deep_convert,omitempty"`
	// 转化数据-深度转化成本; 广告主为每个深度转化所付出的平均成本，计算方法：总花费/深度转化数。当天数据可能会有波动，次日早8点后稳定
	DeepConvertCost float64 `json:"deep_convert_cost,omitempty"`
	// 转化数据-深度转化率; 广告被用户进行深度转化的次数占转化次数的百分比。计算方式：深度转化数/转化数*100%
	DeepConvertRate float64 `json:"deep_convert_rate,omitempty"`
	// 转化数据（计费时间）-转化数（计费时间）; 在转化行为发生（或回传）后，将转化行为回记到过去7天内的扣费（消耗产生）时间上。 例如：广告在8月20日展示给用户，此时广告花费10元，用户点击广告后于8月23日产生1笔购买，则8月23日这笔购买将会展示在8月20日，8月23日没有转化数。暂不支持查看当日数据
	AttributionConvert int64 `json:"attribution_convert,omitempty"`
	// 转化数据（计费时间）-转化成本（计费时间）;转化成本(计费时间) = 消耗 / 转化数(计费时间)。例如：广告在8月20日展示给用户，此时广告花费10元，用户点击广告后，于8月23日产生2笔购买，则8月20日的转化成本（计费时间） = 5元（即10元除以2笔）。成本考核和系统赔付以该指标为准，暂不支持查看当日数据
	AttributionConvertCost float64 `json:"attribution_convert_cost,omitempty"`
	// 转化数据（计费时间）-深度转化数（计费时间）;在转化行为发生（或回传）后，将转化行为回记到过去7天内的扣费（消耗产生）时间上。 例如：广告在8月20日展示给用户，此时广告花费10元，用户点击广告后于8月23日产生1笔购买，则8月23日这笔购买将会展示在8月20日。暂不支持查看当日数据
	AttributionDeepConvert int64 `json:"attribution_deep_convert,omitempty"`
	// 转化数据（计费时间）-深度转化成本（计费时间）; 是一个准确的深度转化成本评估指标。计算方式：消耗 / 深度转化数(计费时间)。例如：广告在8月20日展示给用户，此时广告花费10元，用户点击广告后，于 8 月 23 日产生2笔购买，则8月20日的深度转化成本（计费时间） = 5元（即10元除以2笔）。成本考核和系统赔付以该指标为准，暂不支持查看当日数据
	AttributionDeepConvertCost float64 `json:"attribution_deep_convert_cost,omitempty"`
	// 应用下载广告数据-安卓下载开始数; 用户在落地页中开始安装包下载进程的次数，仅安卓下载可以监测到该行为
	DownloadStart int64 `json:"download_start,omitempty"`
	// 应用下载广告数据-安卓下载开始成本; 计算方法：总花费/安卓下载开始数
	DownloadStartCost float64 `json:"download_start_cost,omitempty"`
	// 应用下载广告数据-安卓下载开始率
	DownloadStartRate float64 `json:"download_start_rate,omitempty"`
	// 应用下载广告数据-安卓下载完成数; 用户在落地页中下载安装包进程完成的次数，仅安卓下载可以监测到该行为
	DownloadFinish int64 `json:"download_finish,omitempty"`
	// 应用下载广告数据-安卓下载完成成本; 计算方式：总花费/安卓下载完成数
	DownloadFinishCost float64 `json:"download_finish_cost,omitempty"`
	// 应用下载广告数据-安卓下载完成率; 计算方式：安卓下载完成数/安卓下载开始数
	DownloadFinishRate float64 `json:"download_finish_rate,omitempty"`
	// 应用下载广告数据-点击安装数; 用户点击安装完成的次数
	ClickInstall int64 `json:"click_install,omitempty"`
	// 应用下载广告数据-安卓安装完成数; 用户在落地页中将安装包安装完成的次数，仅安卓下载可以监测到该行为
	InstallFinish int64 `json:"install_finish,omitempty"`
	// 应用下载广告数据-安卓安装完成成本; 计算方式：总花费/安卓安装完成数
	InstallFinishCost float64 `json:"install_finish_cost,omitempty"`
	// 应用下载广告数据-安卓安装完成率;计算方式：安卓安装完成数/安卓下载完成数
	InstallFinishRate float64 `json:"install_finish_rate,omitempty"`
	// 应用下载广告数据-激活数; 如果您对接了API，激活数是您认可且回传成功的的激活数。如果您对接了SDK，则激活数是指用户下载您的APP后打开的次数
	Active int64 `json:"active,omitempty"`
	// 应用下载广告数据-激活成本; 计算方式：总花费/激活数
	ActiveCost float64 `json:"active_cost,omitempty"`
	// 应用下载广告数据-激活率; 计算方式：激活数/点击数*100%
	ActiveRate float64 `json:"active_rate,omitempty"`
	// 应用下载广告数据-注册数; 如果您对接了API，注册数是您认可且回传成功的的注册数。如果您对接了SDK，则注册数是用户实现注册行为的次数
	Register int64 `json:"register,omitempty"`
	// 应用下载广告数据-注册成本; 广告主为每个注册所付出的成本，计算公式是：总花费/注册数，当天数据可能会有波动，次日早8点后稳定
	ActiveRegisterCost float64 `json:"active_register_cost,omitempty"`
	// 应用下载广告数据-注册率; 注册用户占激活用户的比例
	ActiveRegisterRate float64 `json:"active_register_rate,omitempty"`
	// 应用下载广告数据-次留（未对回）; 在次留行为发生或回传后，将次留数匹配到相应的激活时间上，用来查看当天在第二天的次留数。当天数据不产出（不精准）
	NextDayOpen int64 `json:"next_day_open,omitempty"`
	// 应用下载广告数据-次留率（未对回）; 次留率=次留数/激活数（不精准）
	NextDayOpenRate float64 `json:"next_day_open_rate,omitempty"`
	// 应用下载广告数据-次留成本（未对回）; 次留成本=消耗/次留数（不精准）
	NextDayOpenCost float64 `json:"next_day_open_cost,omitempty"`
	// 应用下载广告数据-次留数; 在次留行为发生或回传后，将次留数匹配到相应的激活时间上，用来查看当天在第二天的次留数。当天数据不产出
	AttributionNextDayOpenCnt int64 `json:"attribution_next_day_open_cnt,omitempty"`
	// 应用下载广告数据-次留成本; 次留成本=消耗/次留数
	AttributionNextDayOpenCost float64 `json:"attribution_next_open_cost,omitempty"`
	// 应用下载广告数据-次留率; 次留率=次留数/激活数
	AttributionNextDayOpenRate float64 `json:"attribution_next_day_open_rate,omitempty"`
	// 应用下载广告数据-关键行为数; 有APP内关键行为的用户数量
	GameAddiction int64 `json:"game_addiction,omitempty"`
	// 应用下载广告数据-关键行为成本; 广告主为每个有APP内关键行为的用户所付出的成本，计算公式是总花费/关键行为数。当天数据可能会有波动，次日早8点后稳定
	GameAddictionCost float64 `json:"game_addiction_cost,omitempty"`
	// 应用下载广告数据-关键行为率; 关键行为用户占激活用户的比例
	GameAddictionRate float64 `json:"game_addiction_rate,omitempty"`
	// 应用下载广告数据-首次付费次数; 用户在应用内首次完成付费的次数
	PayCount int64 `json:"pay_count,omitempty"`
	// 应用下载广告数据-首次付费成本; 用户在应用内首次完成付费的成本，计算方式：消耗/首次付费数
	ActivePayCost float64 `json:"active_pay_cost,omitempty"`
	// 应用下载广告数据-首次付费率; 计算方式：首次付费数/激活数
	ActivePayRate float64 `json:"active_pay_rate,omitempty"`
	// 应用下载广告数据-完件数; 互联网金融-贷款行业中，用户成功提交贷款额度申请的行为
	LoanCompletion int64 `json:"loan_completion,omitempty"`
	// 应用下载广告数据-完件成本; 计算方式：总花费／完件数。当天数据可能会有波动，次日早8点后稳定
	LoanCompletionCost float64 `json:"loan_completion_cost,omitempty"`
	// 应用下载广告数据-完件率; 计算方式：完件数／注册数
	LoanCompletionRate float64 `json:"loan_completion_rate,omitempty"`
	// 应用下载广告数据-预授信数; 互联网金融-贷款行业中，用户提交一部分个人信息后，广告主初步审批通过，并引导用户进行更加详细的信息填写以完成最终授信
	PreLoanCredit int64 `json:"pre_loan_credit,omitempty"`
	// 应用下载广告数据-预授信成本; 计算方式：总花费／预授信数。当天数据可能会有波动，次日早8点后稳定
	PreLoanCreditCost float64 `json:"pre_loan_credit_cost,omitempty"`
	// 应用下载广告数据-授信数; 互联网金融-贷款行业中，用户提交贷款额度申请后，广告主审批通过，给予用户可贷款的额度
	LoanCredit int64 `json:"loan_credit,omitempty"`
	// 应用下载广告数据-授信成本; 计算方式：总花费／授信数。当天数据可能会有波动，次日早8点后稳定
	LoanCreditCost float64 `json:"loan_credit_cost,omitempty"`
	// 应用下载广告数据-授信率; 计算方式：授信数／完件数
	LoanCreditRate float64 `json:"loan_credit_rate,omitempty"`
	// 应用下载广告数据-APP内访问; 用户调起APP后到达的次数，一般在DPA广告中使用
	InAppUv int64 `json:"in_app_uv,omitempty"`
	// 应用下载广告数据-APP内访问详情页;用户调起APP后到达指定详情页的次数，一般在DPA广告中使用
	InAppDetailUv int64 `json:"in_app_detail_uv,omitempty"`
	// 应用下载广告数据-APP内加入购物车; 用户调起APP后到达指定详情页的次数，一般在DPA广告中使用
	InAppCart int64 `json:"in_app_cart,omitempty"`
	// 应用下载广告数据-APP内付费; 用户调起APP后完成付费的次数，一般在DPA广告中使用
	InAppPay int64 `json:"in_app_pay,omitempty"`
	// 应用下载广告数据-APP内下单; 用户调起APP后提交订单的次数，一般在DPA广告中使用
	InAppOrder int64 `json:"in_app_order,omitempty"`
	// 应用下载广告数据-7日付费次数（激活时间）; 用户激活应用后的7天内，在应用内完成付费的总次数，并将付费次数披露在每个用户的激活时间上
	AttributionGamePay7dCount int64 `json:"attribution_game_pay_7d_count,omitempty"`
	// 应用下载广告数据-7日付费成本（激活时间）; 用户激活应用后的7天内，在应用内完成付费的平均成本，计算方式：消耗/7日付费次数(激活时间)
	AttributionGamePay7dCost float64 `json:"attribution_game_pay_7d_cost,omitempty"`
	// 应用下载广告数据-7日人均付费次数（激活时间）; 7天内用户的平均付费次数，计算方式：7日付费次数（激活时间）/7日首次付费数(激活时间)，对首次付费数的统计仅在计划内去重且披露在每个用户的激活时间上
	AttributionActivePay7dPerCount int64 `json:"attribution_active_pay_7d_per_count,omitempty"`
	// 应用下载广告数据-付费成本; 当天用户在应用内完成付费的平均成本，计算方式：消耗/付费次数
	GamePayCost float64 `json:"game_pay_cost,omitempty"`
	// 应用下载广告数据-付费次数; 当天用户在应用内完成付费的总次数，多次付费会重复计数
	GamePayCount int64 `json:"game_pay_count,omitempty"`
	// 落地页转化数据-点击电话按钮;用户点击POI页面的电话按钮的次数
	Phone int64 `json:"phone,omitempty"`
	// 落地页转化数据-表单提交; 用户在门店落地页多线沟通提交表单的次数
	Form int64 `json:"form,omitempty"`
	// 落地页转化数据-地图搜索; 用户进行地图搜索的次数
	MapSearch int64 `json:"map_search,omitempty"`
	// 落地页转化数据-按钮button; 用户点击按钮button的次数
	Button int64 `json:"button,omitempty"`
	// 落地页转化数据-关键页面浏览; 用户在关键页面的浏览次数
	View int64 `json:"view,omitempty"`
	// 落地页转化数据-下载开始; 用户点击下载开始的次数
	Download int64 `json:"download,omitempty"`
	// 落地页转化数据-QQ咨询; 用户点击QQ咨询按钮的次数
	QQ int64 `json:"qq,omitempty"`
	// 落地页转化数据-抽奖;用户点击抽奖按钮的次数
	Lottery int64 `json:"lottery,omitempty"`
	// 落地页转化数据-投票; 用户点击投票按钮的次数
	Vote int64 `json:"vote,omitempty"`
	// 落地页转化数据-短信咨询; 用户点击短信咨询的次数
	Message int64 `json:"message,omitempty"`
	// 落地页转化数据-页面跳转; 用户跳转至其他页面的次数
	Redirect int64 `json:"redirect,omitempty"`
	// 落地页转化数据-商品购买; 用户购买商品的次数
	Shopping int64 `json:"shopping,omitempty"`
	// 落地页转化数据-在线咨询;用户点击在线咨询按钮的次数
	Consult int64 `json:"consult,omitempty"`
	// 落地页转化数据-微信复制; 用户通过长按复制微信号的次数
	Wechat int64 `json:"wechat,omitempty"`
	// 落地页转化数据-智能电话-确认拨打; 用户点击确认拨打电话按钮的次数
	PhoneConfirm int64 `json:"phone_confirm,omitempty"`
	// 落地页转化数据-智能电话-确认接通; 用户在门店落地页点击多线沟通的电话按钮并接通的次数
	PhoneConnect int64 `json:"phone_connect,omitempty"`
	// 落地页转化数据-有效咨询; 用户在门店落地页多线沟通的在线咨询中有效咨询的次数
	ConsultEffective int64 `json:"consult_effective,omitempty"`
	// 落地页转化数据-建站卡券领取; 用户点击POI页面卡券领取按钮的次数
	Coupon int64 `json:"coupon,omitempty"`
	// 落地页转化数据-卡券页领取; 用户点击卡券页领取按钮的次数
	CouponSinglePage int64 `json:"coupon_single_page,omitempty"`
	// 落地页及门店数据-调起店铺; 用户调起第三方店铺的次数
	RedirectToShop int64 `json:"redirect_to_shop,omitempty"`
	// 落地页及门店数据-店铺收藏; 用户点击POI页面的收藏按钮的次数
	PoiCollect int64 `json:"poi_collect,omitempty"`
	// 落地页及门店数据-查看店铺地址; 用户点击顶部地图图标区域展开详细地图/点击图片下方地址展开详细地图/点击图片下方地址展开详细地图的次数
	PoiAddressClick int64 `json:"poi_address_click,omitempty"`
	// 落地页及门店数据-鲁班订单量; 该条广告产生的订单数量
	LubanOrderCnt int64 `json:"luban_order_cnt,omitempty"`
	// 落地页及门店数据-鲁班订单金额; 该条广告产生订单的实际销售金额，不含优惠券和退款。当前仅支持商品推广的计划
	LubanOrderStatAmount float64 `json:"luban_order_stat_amount,omitempty"`
	// 落地页及门店数据-鲁班ROI; 计算公式：订单金额/广告消耗。当前仅支持商品推广的计划
	LubanOrderRoi float64 `json:"luban_order_roi,omitempty"`
	// 落地页及门店数据-直播间观看数; 用户进入您直播间的次数
	LubanLiveEnterCnt int64 `json:"luban_live_enter_cnt,omitempty"`
	// 落地页及门店数据-直播间超过1分钟观看数; 用户在您的直播间单次观看时长超过1分钟的次数，目前仅支持鲁班直播、抖音号推广直播、品牌Feeds Live直播和品牌TopLive直播
	LubanWatchOneMinuteCount int64 `json:"live_watch_one_minute_count,omitempty"`
	// 落地页及门店数据-直播间关注数; 您的直播间新增的关注数，目前仅支持鲁班直播、抖音号推广直播、品牌Feeds Live直播和品牌TopLive直播
	LubanLiveFollowCnt int64 `json:"luban_live_follow_cnt,omitempty"`
	// 落地页及门店数据-直播间加入粉丝团; 用户在您的直播间成功加入粉丝团的次数，目前仅支持鲁班直播、抖音号推广直播、品牌Feeds Live直播和品牌TopLive直播
	LubanFansClubJoinCnt int64 `json:"luban_fans_club_join_cnt,omitempty"`
	// 落地页及门店数据-直播间评论数; 用户在您的直播间输入评论并点击发送消息的次数，目前仅支持鲁班直播、抖音号推广直播、品牌Feeds Live直播和品牌TopLive直播
	LubanLiveCommentCnt int64 `json:"luban_live_comment_cnt,omitempty"`
	// 落地页及门店数据-直播间分享数; 用户在您的直播间点击分享到其他社交媒体的次数，目前仅支持鲁班直播、抖音号推广直播、品牌Feeds Live直播和品牌TopLive直播
	LubanLiveShareCnt int64 `json:"luban_live_share_cnt,omitempty"`
	// 落地页及门店数据-直播间打赏次数; 用户在您的直播间成功打赏的次数，目前仅支持鲁班直播、抖音号推广直播、品牌Feeds Live直播和品牌TopLive直播
	LubanLiveGiftCnt int64 `json:"luban_live_gift_cnt,omitempty"`
	// 落地页及门店数据-直播间礼物总金额; 用户在您的直播间打赏的总金额，单位：音浪。目前仅支持鲁班直播、抖音号推广直播、品牌Feeds Live直播和品牌TopLive直播
	LubanLiveGiftAmount float64 `json:"luban_live_gift_amount,omitempty"`
	// 落地页及门店数据-直播间查看购物车数; 用户在您的直播间点击查看购物车的次数，目前仅支持鲁班直播、抖音号推广直播、品牌Feeds Live直播和品牌TopLive直播
	LubanLiveSidecartClickCnt int64 `json:"luban_live_slidecart_click_cnt,omitempty"`
	// 落地页及门店数据-直播间点击商品数; 用户在您的直播间点击商品的次数，目前仅支持鲁班直播、抖音号推广直播、品牌Feeds Live直播和品牌TopLive直播
	LubanLiveClickProductCnt int64 `json:"luban_live_click_product_cnt,omitempty"`
	// 落地页及门店数据-直播间订单量; 用户在您直播间的支付成功的订单量
	LubanLivePayOrderCount int64 `json:"luban_live_pay_order_count,omitempty"`
	// 落地页及门店数据-直播间订单金额; 用户在您直播间的支付成功的订单金额，单位：元。
	LubanLivePayOrderStatCost float64 `json:"luban_live_pay_order_stat_cost,omitempty"`
	// 落地页及门店数据-微信内注册数; 当天用户在微信内完成注册的总次数，将注册数记录在事件发生的时间上。例如，您的广告在8月20日展示给用户，用户点击广告后于8月23日在微信内产生注册行为，则这个注册行为将会展示在8月23日
	WechatLoginCount int64 `json:"wechat_login_count,omitempty"`
	// 落地页及门店数据-微信内注册数(计费时间); 微信内注册行为发生（或回传）后，将注册行为记录到过去30天内的扣费（消耗产生）时间上。例如，广告在8月20日展示给用户，用户点击广告后于8月23日在微信内产生注册行为，则这个注册行为将会展示在8月20日，8月23日没有注册数
	AttributionWechatLogin30dCount int64 `json:"attribution_wechat_login_30d_count,omitempty"`
	// 落地页及门店数据-微信内注册成本; 当天用户在微信内完成注册的平均成本，计算公式：消耗/微信内注册数
	WechatLoginCost float64 `json:"wechat_login_cost,omitempty"`
	// 落地页及门店数据-微信内注册成本(计费时间); 计算方式： 消耗 / 微信内注册数(计费时间)。例如，广告在8月20日展示给2位用户，此时广告计划10元，2位用户点击广告后于8月23日在微信内产生注册行为，则8月20日的微信内注册成本(计费时间) = 10/2=5元
	AttributionWechatLogin30Cost float64 `json:"attribution_wechat_login_30d_cost,omitempty"`
	// 落地页及门店数据-微信内首次付费数; 当天用户在微信内首次完成付费的总次数，将付费数记录在事件发生的时间上。例如，用户于8月20日产生注册行为，8月23日在微信内产生首次付费行为，则这个付费行为将会展示在8月23日
	WechatFirstPayCount int64 `json:"wechat_first_pay_count,omitempty"`
	// 落地页及门店数据-微信内首次付费数(计费时间); 微信内付费行为发生（或回传）后，将付费行为记录到过去30天内的扣费（消耗产生）时间上。例如，广告在8月20日展示给用户，用户点击广告后于8月23日在微信内产生首次付费行为，则这个付费行为将会展示在8月20日，8月23日没有付费数
	AttributionWechatFirstPay30dCount int64 `json:"attribution_wechat_first_pay_30d_count,omitempty"`
	// 落地页及门店数据-微信内首次付费成本; 当天用户在微信内完成首次付费的平均成本，计算公式：消耗/微信内首次付费数
	WechatFirstPayCost float64 `json:"wechat_first_pay_cost,omitempty"`
	// 落地页及门店数据-微信内首次付费成本(计费时间); 计算方式：消耗 / 微信内首次付费数（计费时间）。例如，广告在8月20日展示给2位用户，此时广告计划10元，2位用户点击广告后于8月23日在微信内产生首次付费行为，则8月20日的微信内首次付费成本(计费时间) = 10/2=5元
	AttributionWechatFirstPay30dCost float64 `json:"attribution_wechat_first_pay_30d_cost,omitempty"`
	// 落地页及门店数据-微信内首次付费率; 计算公式：微信内首次付费数/微信内注册数
	WechatFirstPayRate float64 `json:"wechat_first_pay_rate,omitempty"`
	// 落地页及门店数据-微信内首次付费率(计费时间); 计算公式：微信内首次付费数(计费时间)/微信内注册数(计费时间)
	AttributionWechatFirstPay30dRate float64 `json:"attribution_wechat_first_pay_30d_rate,omitempty"`
	// 落地页及门店数据-微信内首次付费金额; 当天用户在微信内完成首次付费的总金额
	WechatPayAmount float64 `json:"wechat_pay_amount,omitempty"`
	// 落地页及门店数据-微信内首次付费金额(计费时间); 当天用户在微信内完成首次付费的总金额，将付费金额记录到过去30天内的扣费（消耗产生）时间上。例如，广告在8月20日展示给用户，用户点击广告后于8月23日在微信内产生首次付费行为，则这个首次付费金额将会展示在8月20日，8月23日没有首次付费金额
	AttributionWechatPay30dAmount float64 `json:"attribution_wechat_pay_30d_amount,omitempty"`
	// 落地页及门店数据-微信内首次付费ROI; 所选时间范围内的微信内首次付费ROI，计算公式：微信内首次付费金额（计费时间）/消耗
	AttributionWechatPay30dRoi float64 `json:"attribution_wechat_pay_30d_roi,omitempty"`
	// 落地页及门店数据-智能电话-有效接通; 用户在门店落地页点击多线沟通的电话按钮并有效接通的次数
	PhoneEffective int64 `json:"phone_effective,omitempty"`
	// 视频数据-播放数; 播放时间大于0S的数量，在某些蜂窝网络环境下，需要您手动点击开始才会开始播放，因此有时播放数小于展示数
	TotalPlay int64 `json:"total_play,omitempty"`
	// 视频数据-有效播放数;竞价广告播放时间大于等于10秒的数量，如果视频总时长不足10秒，则记录播放完成的次数。品牌广告在部分APP（头条、头条lite、抖音、西瓜、抖音火山版、皮皮虾）播放时间大于等于5秒的数量，在其他APP大于等于3秒的数量，如果视频总时长不足5秒/3秒时，则记录播放完成的次数
	ValidPlay int64 `json:"valid_play,omitempty"`
	// 视频数据-有效播放成本;计算公式：总花费/有效播放数，当天数据可能会有波动，次日早8点后稳定
	ValidPlayCost float64 `json:"valid_play_cost,omitempty"`
	// 视频数据-有效播放率; 计算公式：有效播放数/展示数
	ValidPlayRate float64 `json:"valid_play_rate,omitempty"`
	// 视频数据-25%进度播放数; 用户播放至视频长度25%及以上的次数，包括跳跃播放至此长度的播放次数
	Play25FeedBreak int64 `json:"play_25_feed_break,omitempty"`
	// 视频数据-50%进度播放数; 用户播放至视频长度50%及以上的次数，包括跳跃播放至此长度的播放次数
	Play50FeedBreak int64 `json:"play_50_feed_break,omitempty"`
	// 视频数据-75%进度播放数; 用户播放至视频长度75%及以上的次数，包括跳跃播放至此长度的播放次数
	Play75FeedBreak int64 `json:"play_75_feed_break,omitempty"`
	// 视频数据-99%进度播放数; 用户播放至视频长度99%及以上的次数，包括跳跃播放至此长度的播放次数
	Play100FeedBreak int64 `json:"play_100_feed_break,omitempty"`
	// 视频数据-平均单次播放时长,单位：秒; 计算方法：视频播放总实际时长／播放总次数（不包含跳跃的时长）
	AveragePlayTimePerPlay float64 `json:"average_play_time_per_play,omitempty"`
	// PlayOver 播放完成数
	PlayOver int64 `json:"play_over,omitempty"`
	// 视频数据-播完率; 计算公式：播放完成数/播放数
	PlayOverRate float64 `json:"play_over_rate,omitempty"`
	// 视频数据-WiFi播放占比; 在wifi环境下视频的播放数/视频播放总数
	WifiPlayRate float64 `json:"wifi_play_rate,omitempty"`
	// 视频数据-WiFi播放量;在wifi环境下视频的播放数
	WifiPlay int64 `json:"wifi_play,omitempty"`
	// 视频数据-播放时长，单位ms; 计算方法：视频播放总实际时长／播放总次数（不包含跳跃的时长）
	PlayDurationSum int64 `json:"play_duration_sum,omitempty"`
	// 附加创意-附加创意电话按钮点击; 用户查看创意的附加创意后，点击电话拨打按钮的次数
	AdvancedCreativePhoneClick int64 `json:"advanced_creative_phone_click,omitempty"`
	// 附加创意-附加创意在线咨询点击; 用户查看创意的附加创意后，点击在线咨询的次数
	AdvancedCreativeConsuelClick int64 `json:"advanced_creative_counsel_click,omitempty"`
	// 附加创意-附加创意表单按钮点击; 用户查看创意的附加创意后，点击表单提交按钮的次数
	AdvancedCreativeFormClick int64 `json:"advanced_creative_form_click,omitempty"`
	// 附加创意-附加创意卡券领取; 用户查看创意的附加创意后，领取卡券的次数
	AdvancedCreativeCouponAddition int64 `json:"advanced_creative_coupon_addition,omitempty"`
	// 附加创意-附加创意表单提交; 用户查看创意的附加创意后，提交表单的次数
	AdvancedCreativeFormSubmit int64 `json:"advanced_creative_form_submit,omitempty"`
	// 视频数据3秒卡片展现; 对于视频卡片类广告，在视频播放到3秒时进行卡片展现的数量
	CardShow int64 `json:"card_show,omitempty"`
	// 互动数据-分享数; 用户把分享广告到其他社交媒体的次数，成功完成一次分享行为记一次分享数
	Share int64 `json:"share,omitempty"`
	// 互动数据-评论数; 用户在评论区对广告创意输入评论并点击提交的次数，仅限有评论区的APP（西瓜目前没有评论区）
	Comment int64 `json:"comment,omitempty"`
	// 互动数据-点赞数; 广告被用户点赞的数量
	Like int64 `json:"like,omitempty"`
	// 互动数据-新增关注数; 广告被受众新增关注的数量，目前抖音、头条号外、微头条、抖音火山版可新增关注
	Follow int64 `json:"follow,omitempty"`
	// 互动数据-主页访问量;点击广告头像进入主页进行访问的次数，目前抖音、头条号外、抖音火山版可进行主页访问
	HomeVisited int64 `json:"home_visited,omitempty"`
	// 互动数据-挑战赛查看数; 抖音广告中，用户点击挑战赛查看的数量
	IesChallengeClick int64 `json:"ies_challenge_click,omitempty"`
	// 互动数据-音乐查看数; 抖音广告中，用户点击音乐查看的数量
	IesMusicClick int64 `json:"ies_music_click,omitempty"`
	// 互动数据-POI点击数; 抖音广告中，用户点击POI组件的次数，一般出现在信息流或评论页面。POI为Point of Interest缩写
	LocationClick int64 `json:"location_click,omitempty"`
	// 互动数据-私信数; 用户在广告主抖音号主页中发起私信的次数，按人次记数
	MessageAction int64 `json:"message_action,omitempty"`
	// 互动数据-推广页访问量; 用户在广告主抖音号主页内访问落地页的次数
	ClickLandingPage int64 `json:"click_landing_page,omitempty"`
	// 互动数据-主页商品橱窗访问量;用户在广告主抖音号主页中访问商品橱窗的次数
	ClickShopwindow int64 `json:"click_shopwindow,omitempty"`
	// 互动数据-主页内落地页访问量（主页官网访问量）; 用户在广告主抖音号主页内访问落地页的次数
	ClickWebsite int64 `json:"click_website,omitempty"`
	// 互动数据-主页下载链接点击量; 用户在广告主抖音号主页内点击下载的次数
	ClickDownload int64 `json:"click_download,omitempty"`
	// 互动数据-主页内电话拨打点击量; 用户在广告主抖音号主页内点击电话拨打的次数
	ClickCallDy int64 `json:"click_call_dy,omitempty"`
	// Cost Per Click（每次点击的花费），按点击计费
	Cpc float64 `json:"cpc,omitempty"`
	// Cost Per Mille（每千次曝光的花费），按展示计费
	Cpm float64 `json:"cpm,omitempty"`
	// Cost Per Action（每次转化的花费），按转化计费
	Cpa float64 `json:"cpa,omitempty"`
	// 互动数据-单次互动成本
	InteractPerCost float64 `json:"interact_per_cost,omitempty"`
	// 互动数据-转化展现率
	ConverShowRate float64 `json:"convert_show_rate,omitempty"`
	// 互动数据-举报数;用户认为广告质量较差而点击举报的行为数量
	Report int64 `json:"report,omitempty"`
	// 互动数据-不感兴趣数; 用户点击对广告不感兴趣的反馈的行为数量
	Dislike int64 `json:"dislike,omitempty"`
	// 视频数据-10秒播放率; 计算公式：播放时长超过10秒的次数/播放数，播放时长以播放终止时所在的视频时长来计算，包含拖拽等行为
	PlayDuration10sRate float64 `json:"play_duration_10s_rate,omitempty"`
	// 视频数据-5秒播放率; 计算公式：播放时长超过5秒的次数/播放数，播放时长以播放终止时所在的视频时长来计算，包含拖拽等行为
	PlayDuration5sRate float64 `json:"play_duration_5s_rate,omitempty"`
	// 视频数据-3秒播放率; 计算公式：播放时长超过3秒的次数/播放数，播放时长以播放终止时所在的视频时长来计算，包含拖拽等行为
	PlayDuration3sRate float64 `json:"play_duration_3s_rate,omitempty"`
	// 视频数据-2秒播放率; 计算公式：播放时长超过2秒的次数/播放数，播放时长以播放终止时所在的视频时长来计算，包含拖拽等行为
	PlayDuration2sRate float64 `json:"play_duration_2s_rate,omitempty"`
	// 视频数据-100%进度播放率; 计算公式：播进度超过100%的次数/播放数
	Play100FeedBreakRate float64 `json:"play_100_feed_break_rate,omitempty"`
	// 视频数据-75%进度播放率; 计算公式：播进度超过75%的次数/播放数
	Play75FeedBreakRate float64 `json:"play_75_feed_break_rate,omitempty"`
	// 视频数据-50%进度播放率; 计算公式：播进度超过50%的次数/播放数
	Play50FeedBreakRate float64 `json:"play_50_feed_break_rate,omitempty"`
	// 视频数据-25%进度播放率; 计算公式：播进度超过50%的次数/播放数
	Play25FeedBreakRate float64 `json:"play_25_feed_break_rate,omitempty"`
	// 视频数据-平均播放进度; 计算方法：视频播放总进度/播放总次数
	AveragePlayProgress float64 `json:"average_play_progress,omitempty"`
	// 播放时长; 计算公式：总的播放时长超，播放时长以播放终止时所在的视频时长来计算，包含拖拽等行为
	PlayDuration float64 `json:"play_duration,omitempty"`
	// 播放3s; 计算公式：播放时长超过3秒的次数，播放时长以播放终止时所在的视频时长来计算，包含拖拽等行为
	PlayDuration3s float64 `json:"play_duration_3s,omitempty"`
	// 播放2s; 计算公式：播放时长超过2s数，播放时长以播放终止时所在的视频时长来计算，包含拖拽等行为
	PlayDuration2s float64 `json:"play_duration_2s,omitempty"`
	// 播放10s; 计算公式：播放时长超过10s数，播放时长以播放终止时所在的视频时长来计算，包含拖拽等行为
	PlayDuration10s float64 `json:"play_duration_10s,omitempty"`
	// 平均视频播放量
	AverageVideoPlay float64 `json:"average_video_play,omitempty"`
	// 关键词排名，关键词报表专属字段，返回-1.0对应在ad后台页面上中展示“-”，表示无效值
	AverageRank float64 `json:"average_rank,omitempty"`
	// AttributionActivePayIntraOneDayCount 游戏行业-激活后24h首次付费数
	AttributionActivePayIntraOneDayCount int64 `json:"attribution_active_pay_intra_one_day_count,omitempty"`
	// AttributionActivePayIntraOneDayCost 游戏行业-激活后24h首次付费成本
	AttributionActivePayIntraOneDayCost float64 `json:"attribution_active_pay_intra_one_day_cost,omitempty"`
	// AttributionActivePayIntraOneDayRate 游戏行业-激活后24h首次付费率
	AttributionActivePayIntraOneDayRate float64 `json:"attribution_active_pay_intra_one_day_rate,omitempty"`
	// AttributionActivePayIntraOneDayAmount 游戏行业-激活后24h付费金额
	AttributionActivePayIntraOneDayAmount float64 `json:"attribution_active_pay_intra_one_day_amount,omitempty"`
	// AttributionActivePayIntraOneDayRoi 游戏行业-激活后24小时付费roi
	AttributionActivePayIntraOneDayRoi float64 `json:"attribution_active_pay_intra_one_day_roi,omitempty"`
	// AttributionMicroGame0dLtv 小游戏当日LTV-所选时间范围内的激活用户在激活当日的变现金额
	AttributionMicroGame0dLtv float64 `json:"attribution_micro_game_0d_ltv,omitempty"`
	// AttributionMicroGame3dLtv 小游戏激活后三日LTV-所选时间范围内的激活用户在激活后三日内的变现金额
	AttributionMicroGame3dLtv float64 `json:"attribution_micro_game_3d_ltv,omitempty"`
	// AttributionMicroGame7dLtv 小游戏激活后七日LTV-所选时间范围内的激活用户在激活后七日内的变现金额
	AttributionMicroGame7dLtv float64 `json:"attribution_micro_game_7d_ltv,omitempty"`
	// AttributionMicroGame0dRoi 小游戏当日广告变现ROI-所选时间范围内的激活用户在激活当日的广告变现ROI，计算公式是：当日LTV / 所选时间的消耗
	AttributionMicroGame0dRoi float64 `json:"attribution_micro_game_0d_roi,omitempty"`
	// AttributionMicroGame3dRoi 小游戏激活后三日广告变现ROI-所选时间范围内的激活用户在激活后三日内的广告变现ROI，计算公式是：三日LTV / 所选时间的消耗
	AttributionMicroGame3dRoi float64 `json:"attribution_micro_game_3d_roi,omitempty"`
	// AttributionMicroGame7dRoi 小游戏激活后七日广告变现ROI-所选时间范围内的激活用户在激活后七日内的广告变现ROI，计算公式是：七日LTV / 所选时间的消耗
	AttributionMicroGame7dRoi float64 `json:"attribution_micro_game_7d_roi,omitempty"`
	// AttributionGameInAppLtv1Day 当日付费金额-所选时间范围内的激活用户，激活当日在APP内的付费金额
	AttributionGameInAppLtv1Day float64 `json:"attribution_game_in_app_ltv_1day,omitempty"`
	// AttributionGameInAppRoi1Day 当日付费ROI
	AttributionGameInAppRoi1Day float64 `json:"attribution_game_in_app_roi_1day,omitempty"`
}
