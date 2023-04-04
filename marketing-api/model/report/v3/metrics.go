package v3

import "github.com/bububa/oceanengine/marketing-api/model"

// Metrics 指标数据
type Metrics struct {
	// StatCost 表示广告在投放期内的预估花费金额。当天数据可能会有波动，次日稳定
	StatCost model.Float64 `json:"stat_cost,omitempty"`
	// ShowCnt 广告展示给用户的次数。计算方式：经平台判定有效且被计费的展示次数。
	ShowCnt model.Int64 `json:"show_cnt,omitempty"`
	// CpmPlatform 广告平均每一千次展现所付出的费用，计算公式是：总花费/展示数*1000。
	CpmPlatform model.Float64 `json:"cpm_platform,omitempty"`
	// ClickCnt 当头条用户点击广告素材时，触发点击事件，该事件被认为是一次有效的广告点击。
	ClickCnt model.Int64 `json:"click_cnt,omitempty"`
	// Ctr 广告被点击的次数占展示次数的百分比。计算方法：点击数/展示数*100%
	Ctr model.Float64 `json:"ctr,omitempty"`
	// CpcPlatform 广告主为每次点击付出的费用成本，计算公式是：总花费/点击数。
	CpcPlatform model.Float64 `json:"cpc_platform,omitempty"`
	// ClickStartCnt 用户在落地页中开始安装包下载进程的次数，仅安卓下载可以监测到该行为
	ClickStartCnt model.Int64 `json:"click_start_cnt,omitempty"`
	// ClickStartCost 计算方法：总花费/安卓下载开始数。
	ClickStartCost model.Float64 `json:"click_start_cost,omitempty"`
	// ClickStartRate 计算方法：安卓下载开始数/点击数
	ClickStartRate model.Float64 `json:"click_start_rate,omitempty"`
	// DownloadFinishCnt 用户在落地页中下载安装包进程完成的次数，仅安卓下载可以监测到该行为
	DownloadFinishCnt model.Int64 `json:"download_finish_cnt,omitempty"`
	// DownloadFinishCost 计算方式：总花费/安卓下载完成数。
	DownloadFinishCost model.Float64 `json:"download_finish_cost,omitempty"`
	// DownloadFinishRate 计算方式：安卓下载完成数/安卓下载开始数
	DownloadFinishRate model.Float64 `json:"download_finish_rate,omitempty"`
	// InstallFinishCnt 用户在落地页中将安装包安装完成的次数，仅安卓下载可以监测到该行为。如果您的计划开启了商店直投，则下载相关行为数据有部分无法被监测到，而安装数据不受影响，因此安装数可能会大于下载数。
	InstallFinishCnt model.Int64 `json:"install_finish_cnt,omitempty"`
	// InstallFinishCost 计算方式：总花费/安卓安装完成数
	InstallFinishCost model.Float64 `json:"install_finish_cost,omitempty"`
	// InstallFinishRate 计算方式：安卓安装完成数/安卓下载完成数
	InstallFinishRate model.Float64 `json:"install_finish_rate,omitempty"`
	// Active 如果您对接了API，激活数是您认可且回传成功的的激活数。如果您对接了SDK，则激活数是指用户下载您的APP后打开的次数。
	Active model.Int64 `json:"active,omitempty"`
	// ActiveCost 计算方式：总花费/激活数。
	ActiveCost model.Float64 `json:"active_cost,omitempty"`
	// ActiveRate 计算方式：激活数/点击数*100%
	ActiveRate model.Float64 `json:"active_rate,omitempty"`
	// ActiveRegister 如果您对接了API，注册数是您认可且回传成功的的注册数。如果您对接了SDK，则注册数是用户实现注册行为的次数，详情见「SDK对接文档」
	ActiveRegister model.Int64 `json:"active_register,omitempty"`
	// ActiveRegisterCost 广告主为每个注册所付出的成本，计算公式是：总花费/注册数，当天数据可能会有波动，次日早8点后稳定。
	ActiveRegisterCost model.Float64 `json:"active_register_cost,omitempty"`
	// ActiveRegisterRate 注册用户占激活用户的比例
	ActiveRegisterRate model.Float64 `json:"active_register_rate,omitempty"`
	// GameAddiction 有APP内关键行为的用户数量
	GameAddiction model.Int64 `json:"game_addiction,omitempty"`
	// GameAddictionCost 广告主为每个有APP内关键行为的用户所付出的成本，计算公式是总花费/关键行为数。当天数据可能会有波动，次日早8点后稳定。
	GameAddictionCost model.Float64 `json:"game_addiction_cost,omitempty"`
	// GameAddictionRate 关键行为用户占激活用户的比例
	GameAddictionRate model.Float64 `json:"game_addiction_rate,omitempty"`
	// AttributionNextDayOpenCnt 当日激活用户在第二天继续登录，则计为一个次留行为。该指标会将收到的次留回传数据，匹配到对应的激活时间上。 例如：8月1日有10个激活，其中5个在8月2日继续登录，则8月1日的次留数为5。
	AttributionNextDayOpenCnt model.Int64 `json:"attribution_next_day_open_cnt,omitempty"`
	// AttributionNextDayOpenCost 次留成本=消耗/次留数
	AttributionNextDayOpenCost model.Float64 `json:"attribution_next_day_open_cost,omitempty"`
	// AttributionNextDayOpenRate 次留率=次留数/激活数
	AttributionNextDayOpenRate model.Float64 `json:"attribution_next_day_open_rate,omitempty"`
	// NextDayOpen 根据您通过api或sdk回传给我们的时间为准，将次留数披露在回传时间。
	NextDayOpen model.Int64 `json:"next_day_open,omitempty"`
	// ActivePay 用户在应用内首次完成付费的次数
	ActivePay model.Int64 `json:"active_pay,omitempty"`
	// ActivePayCost 用户在应用内首次完成付费的成本，计算方式：消耗/首次付费数。
	ActivePayCost model.Float64 `json:"active_pay_cost,omitempty"`
	// ActivePayRate 计算方式：首次付费数/激活数。
	ActivePayRate model.Float64 `json:"active_pay_rate,omitempty"`
	// GamePayCount 当天用户在应用内完成付费的总次数，多次付费会重复计数。
	GamePayCount model.Int64 `json:"game_pay_count,omitempty"`
	// GamePayCost 当天用户在应用内完成付费的平均成本，计算方式：消耗/付费次数。
	GamePayCost model.Float64 `json:"game_pay_cost,omitempty"`
	// AttributionGamePay7dCount 用户激活应用后的7天内，在应用内完成付费的总次数，并将付费次数披露在每个用户的激活时间上。
	AttributionGamePay7dCount model.Int64 `json:"attribution_game_pay_7d_count,omitempty"`
	// AttributionGamePay7dCost 用户激活应用后的7天内，在应用内完成付费的平均成本，计算方式：消耗/7日付费次数(激活时间)。
	AttributionGamePay7dCost model.Float64 `json:"attribution_game_pay_7d_cost,omitempty"`
	// InAppUv 用户调起APP后到达的次数，一般在DPA广告中使用
	InAppUv model.Int64 `json:"in_app_uv,omitempty"`
	// InAppDetailUv 用户调起APP后到达指定详情页的次数，一般在DPA广告中使用
	InAppDetailUv model.Int64 `json:"in_app_detail_uv,omitempty"`
	// InAppCart 用户调起APP后加入购物车的次数，一般在DPA广告中使用
	InAppCart model.Int64 `json:"in_app_cart,omitempty"`
	// InAppPay 用户调起APP后完成付费的次数，一般在DPA广告中使用
	InAppPay model.Int64 `json:"in_app_pay,omitempty"`
	// InAppOrder 用户调起APP后提交订单的次数，一般在DPA广告中使用
	InAppOrder model.Int64 `json:"in_app_order,omitempty"`
	// AttributionGameInAppLtv1Day 所选时间范围内的激活用户，激活当日在APP内的付费金额。
	AttributionGameInAppLtv1Day model.Float64 `json:"attribution_game_in_app_ltv_1day,omitempty"`
	// AttributionGameInAppLtv2Days 所选时间范围内的激活用户，激活后一天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv2Days model.Float64 `json:"attribution_game_in_app_ltv_2days,omitempty"`
	// AttributionGameInAppLtv3Days 所选时间范围内的激活用户，激活后二天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv3Days model.Float64 `json:"attribution_game_in_app_ltv_3days,omitempty"`
	// AttributionGameInAppLtv4Days 所选时间范围内的激活用户，激活后三天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv4Days model.Float64 `json:"attribution_game_in_app_ltv_4days,omitempty"`
	// AttributionGameInAppLtv5Days 所选时间范围内的激活用户，激活后四天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv5Days model.Float64 `json:"attribution_game_in_app_ltv_5days,omitempty"`
	// AttributionGameInAppLtv6Days 所选时间范围内的激活用户，激活后五天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv6Days model.Float64 `json:"attribution_game_in_app_ltv_6days,omitempty"`
	// AttributionGameInAppLtv7Days 所选时间范围内的激活用户，激活后六天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv7Days model.Float64 `json:"attribution_game_in_app_ltv_7days,omitempty"`
	// AttributionGameInAppLtv8Days 所选时间范围内的激活用户，激活后七天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv8Days model.Float64 `json:"attribution_game_in_app_ltv_8days,omitempty"`
	// AttributionGameInAppRoi1Day 所选时间范围内的激活用户在激活当日的付费ROI，计算公式是：当日付费金额/所选时间的消耗。
	AttributionGameInAppRoi1Day model.Float64 `json:"attribution_game_in_app_roi_1day,omitempty"`
	// AttributionGameInAppRoi2Days 所选时间范围内的激活用户在激活后一日内的所有付费ROI，计算公式是：激活后一日付费金额/所选时间的消耗。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppRoi2Days model.Float64 `json:"attribution_game_in_app_roi_2days,omitempty"`
	// AttributionGameInAppRoi3Days 所选时间范围内的激活用户在激活后二日内的所有付费ROI，计算公式是：激活后二日付费金额/所选时间的消耗。
	AttributionGameInAppRoi3Days model.Float64 `json:"attribution_game_in_app_roi_3days,omitempty"`
	// AttributionGameInAppRoi4Days 所选时间范围内的激活用户在激活后三日内的所有付费ROI，计算公式是：激活后三日付费金额/所选时间的消耗。
	AttributionGameInAppRoi4Days model.Float64 `json:"attribution_game_in_app_roi_4days,omitempty"`
	// AttributionGameInAppRoi5Days 所选时间范围内的激活用户在激活后四日内的所有付费ROI，计算公式是：激活后四日付费金额/所选时间的消耗。
	AttributionGameInAppRoi5Days model.Float64 `json:"attribution_game_in_app_roi_5days,omitempty"`
	// AttributionGameInAppRoi6Days 所选时间范围内的激活用户在激活后五日内的所有付费ROI，计算公式是：激活后五日付费金额/所选时间的消耗。
	AttributionGameInAppRoi6Days model.Float64 `json:"attribution_game_in_app_roi_6days,omitempty"`
	// AttributionGameInAppRoi7Days 所选时间范围内的激活用户在激活后六日内的所有付费ROI，计算公式是：激活后六日付费金额/所选时间的消耗。
	AttributionGameInAppRoi7Days model.Float64 `json:"attribution_game_in_app_roi_7days,omitempty"`
	// AttributionGameInAppRoi8Days 所选时间范围内的激活用户在激活后七日内的整体付费ROI，计算公式是：激活后七日付费金额/所选时间的消耗。
	AttributionGameInAppRoi8Days model.Float64 `json:"attribution_game_in_app_roi_8days,omitempty"`
	// AttributionDayActivePayCount 广告计费当日激活且首次付费的次数
	AttributionDayActivePayCount model.Int64 `json:"attribution_day_acitve_pay_count,omitempty"`
	// AttributionDayActivePayCost 消耗/计费当日激活且首次付费数
	AttributionDayActivePayCost model.Float64 `json:"attribution_day_active_pay_cost,omitempty"`
	// AttributionDayActivePayRate 计费当日激活且首次付费数/激活数
	AttributionDayActivePayRate model.Float64 `json:"attribution_day_active_pay_rate,omitempty"`
	// ActivePayIntraDayCount 当日发生激活且首次付费的次数
	ActivePayIntroDayCount model.Int64 `json:"active_pay_intra_day_count,omitempty"`
	// ActivePayIntraDayCost 消耗/激活当日首次付费数
	ActivePayIntroDayCost model.Float64 `json:"active_pay_intra_day_cost,omitempty"`
	// ActivePayIntraDayRate 激活当日首次付费数/激活数
	ActivePayIntroDayRate model.Float64 `json:"active_pay_intra_day_rate,omitempty"`
	// ConvertCnt 将转化数记录在转化事件发生的时间上。建议广告主考核成本时参考“转化数据（计费时间）”例如您的广告在早上8点进行了展示和点击，用户晚上19点发生了激活行为，我们会把激活数披露在晚上19点。
	ConvertCnt model.Int64 `json:"convert_cnt,omitempty"`
	// ConversionCost 广告主为每个转化所付出的平均成本，计算方式：总花费/转化数。当天数据可能会有波动。
	ConversionCost model.Float64 `json:"conversion_cost,omitempty"`
	// ConversionRate 广告被用户转化的次数占点击次数的百分比。计算方式：转化数/点击数*100%
	ConversionRate model.Float64 `json:"conversion_rate,omitempty"`
	// DeepConvertCnt 将深度转化数记录在转化事件发生的时间上。建议广告主考核深度转化成本时参考“深度转化数（计费时间）”例如您的广告在早上8点进行了展示和点击，用户晚上19点发生了激活行为，我们会把激活数披露在晚上19点。
	DeepConvertCnt model.Int64 `json:"deep_convert_cnt,omitempty"`
	// DeepConvertCost 广告主为每个深度转化所付出的平均成本，计算方法：总花费/深度转化数。当天数据可能会有波动，次日早8点后稳定。
	DeepConvertCost model.Float64 `json:"deep_convert_cost,omitempty"`
	// DeepConvertRate 广告被用户进行深度转化的次数占转化次数的百分比。计算方式：深度转化数/转化数*100%
	DeepConvertRate model.Float64 `json:"deep_convert_rate,omitempty"`
}

type CustomMetrics struct {
	// StatCost 表示广告在投放期内的预估花费金额。当天数据可能会有波动，次日稳定
	StatCost model.Float64 `json:"stat_cost,omitempty"`
	// ShowCnt 广告展示给用户的次数。计算方式：经平台判定有效且被计费的展示次数。
	ShowCnt model.Int64 `json:"show_cnt,omitempty"`
	// CpmPlatform 广告平均每一千次展现所付出的费用，计算公式是：总花费/展示数*1000。
	CpmPlatform model.Float64 `json:"cpm_platform,omitempty"`
	// ClickCnt 当头条用户点击广告素材时，触发点击事件，该事件被认为是一次有效的广告点击。
	ClickCnt model.Int64 `json:"click_cnt,omitempty"`
	// Ctr 广告被点击的次数占展示次数的百分比。计算方法：点击数/展示数*100%
	Ctr model.Float64 `json:"ctr,omitempty"`
	// CpcPlatform 广告主为每次点击付出的费用成本，计算公式是：总花费/点击数。
	CpcPlatform model.Float64 `json:"cpc_platform,omitempty"`
	// ClickStartCnt 用户在落地页中开始安装包下载进程的次数，仅安卓下载可以监测到该行为
	ClickStartCnt model.Int64 `json:"click_start_cnt,omitempty"`
	// ClickStartCost 计算方法：总花费/安卓下载开始数。
	ClickStartCost model.Float64 `json:"click_start_cost,omitempty"`
	// ClickStartRate 计算方法：安卓下载开始数/点击数
	ClickStartRate model.Float64 `json:"click_start_rate,omitempty"`
	// DownloadFinishCnt 用户在落地页中下载安装包进程完成的次数，仅安卓下载可以监测到该行为
	DownloadFinishCnt model.Int64 `json:"download_finish_cnt,omitempty"`
	// DownloadFinishCost 计算方式：总花费/安卓下载完成数。
	DownloadFinishCost model.Float64 `json:"download_finish_cost,omitempty"`
	// DownloadFinishRate 计算方式：安卓下载完成数/安卓下载开始数
	DownloadFinishRate model.Float64 `json:"download_finish_rate,omitempty"`
	// InstallFinishCnt 用户在落地页中将安装包安装完成的次数，仅安卓下载可以监测到该行为。如果您的计划开启了商店直投，则下载相关行为数据有部分无法被监测到，而安装数据不受影响，因此安装数可能会大于下载数。
	InstallFinishCnt model.Int64 `json:"install_finish_cnt,omitempty"`
	// InstallFinishCost 计算方式：总花费/安卓安装完成数
	InstallFinishCost model.Float64 `json:"install_finish_cost,omitempty"`
	// InstallFinishRate 计算方式：安卓安装完成数/安卓下载完成数
	InstallFinishRate model.Float64 `json:"install_finish_rate,omitempty"`
	// Active 如果您对接了API，激活数是您认可且回传成功的的激活数。如果您对接了SDK，则激活数是指用户下载您的APP后打开的次数。
	Active model.Int64 `json:"active,omitempty"`
	// ActiveCost 计算方式：总花费/激活数。
	ActiveCost model.Float64 `json:"active_cost,omitempty"`
	// ActiveRate 计算方式：激活数/点击数*100%
	ActiveRate model.Float64 `json:"active_rate,omitempty"`
	// ActiveRegister 如果您对接了API，注册数是您认可且回传成功的的注册数。如果您对接了SDK，则注册数是用户实现注册行为的次数，详情见「SDK对接文档」
	ActiveRegister model.Int64 `json:"active_register,omitempty"`
	// ActiveRegisterCost 广告主为每个注册所付出的成本，计算公式是：总花费/注册数，当天数据可能会有波动，次日早8点后稳定。
	ActiveRegisterCost model.Float64 `json:"active_register_cost,omitempty"`
	// ActiveRegisterRate 注册用户占激活用户的比例
	ActiveRegisterRate model.Float64 `json:"active_register_rate,omitempty"`
	// GameAddiction 有APP内关键行为的用户数量
	GameAddiction model.Int64 `json:"game_addiction,omitempty"`
	// GameAddictionCost 广告主为每个有APP内关键行为的用户所付出的成本，计算公式是总花费/关键行为数。当天数据可能会有波动，次日早8点后稳定。
	GameAddictionCost model.Float64 `json:"game_addiction_cost,omitempty"`
	// GameAddictionRate 关键行为用户占激活用户的比例
	GameAddictionRate model.Float64 `json:"game_addiction_rate,omitempty"`
	// AttributionNextDayOpenCnt 当日激活用户在第二天继续登录，则计为一个次留行为。该指标会将收到的次留回传数据，匹配到对应的激活时间上。 例如：8月1日有10个激活，其中5个在8月2日继续登录，则8月1日的次留数为5。
	AttributionNextDayOpenCnt model.Int64 `json:"attribution_next_day_open_cnt,omitempty"`
	// AttributionNextDayOpenCost 次留成本=消耗/次留数
	AttributionNextDayOpenCost model.Float64 `json:"attribution_next_day_open_cost,omitempty"`
	// AttributionNextDayOpenRate 次留率=次留数/激活数
	AttributionNextDayOpenRate model.Float64 `json:"attribution_next_day_open_rate,omitempty"`
	// NextDayOpen 根据您通过api或sdk回传给我们的时间为准，将次留数披露在回传时间。
	NextDayOpen model.Int64 `json:"next_day_open,omitempty"`
	// ActivePay 用户在应用内首次完成付费的次数
	ActivePay model.Int64 `json:"active_pay,omitempty"`
	// ActivePayCost 用户在应用内首次完成付费的成本，计算方式：消耗/首次付费数。
	ActivePayCost model.Float64 `json:"active_pay_cost,omitempty"`
	// ActivePayRate 计算方式：首次付费数/激活数。
	ActivePayRate model.Float64 `json:"active_pay_rate,omitempty"`
	// GamePayCount 当天用户在应用内完成付费的总次数，多次付费会重复计数。
	GamePayCount model.Int64 `json:"game_pay_count,omitempty"`
	// GamePayCost 当天用户在应用内完成付费的平均成本，计算方式：消耗/付费次数。
	GamePayCost model.Float64 `json:"game_pay_cost,omitempty"`
	// AttributionGamePay7dCount 用户激活应用后的7天内，在应用内完成付费的总次数，并将付费次数披露在每个用户的激活时间上。
	AttributionGamePay7dCount model.Int64 `json:"attribution_game_pay_7d_count,omitempty"`
	// AttributionGamePay7dCost 用户激活应用后的7天内，在应用内完成付费的平均成本，计算方式：消耗/7日付费次数(激活时间)。
	AttributionGamePay7dCost model.Float64 `json:"attribution_game_pay_7d_cost,omitempty"`
	// AttributionGamePay7dPerCount 7天内用户的平均付费次数，计算方式：7日付费次数（激活时间）/7日首次付费数(激活时间)，对首次付费数的统计仅在计划内去重且披露在每个用户的激活时间上。
	AttributionGamePay7dPerCount model.Float64 `json:"attribution_active_pay_7d_per_count,omitempty"`
	// InAppUv 用户调起APP后到达的次数，一般在DPA广告中使用
	InAppUv model.Int64 `json:"in_app_uv,omitempty"`
	// InAppDetailUv 用户调起APP后到达指定详情页的次数，一般在DPA广告中使用
	InAppDetailUv model.Int64 `json:"in_app_detail_uv,omitempty"`
	// InAppCart 用户调起APP后加入购物车的次数，一般在DPA广告中使用
	InAppCart model.Int64 `json:"in_app_cart,omitempty"`
	// InAppPay 用户调起APP后完成付费的次数，一般在DPA广告中使用
	InAppPay model.Int64 `json:"in_app_pay,omitempty"`
	// InAppOrder 用户调起APP后提交订单的次数，一般在DPA广告中使用
	InAppOrder model.Int64 `json:"in_app_order,omitempty"`
	// InAppPayOrderGmv 引流电商订单GMV, 当您使用“in_app_order”事件回传订单金额时，对应的GMV金额
	InAppPayOrderGmv model.Float64 `json:"in_app_pay_order_gmv,omitempty"`
	// InAppPayOrderRoi 引流电商订单ROI, 引流电商订单GMV/消耗"
	InAppPayOrderRoi model.Float64 `json:"in_app_pay_order_roi,omitempty"`
	// InAppPayGmv 引流电商支付GMV, 当您使用“in_app_pay”事件回传支付金额时，对应的GMV金额"
	InAppPayGmv model.Float64 `json:"in_app_pay_gmv,omitempty"`
	// InAppPayRoi 引流电商支付ROI, 引流电商支付GMV/消耗
	InAppPayRoi model.Float64 `json:"in_app_pay_roi,omitempty"`
	// AttributionGameInAppLtv1Day 所选时间范围内的激活用户，激活当日在APP内的付费金额。
	AttributionGameInAppLtv1Day model.Float64 `json:"attribution_game_in_app_ltv_1day,omitempty"`
	// AttributionGameInAppLtv2Days 所选时间范围内的激活用户，激活后一天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv2Days model.Float64 `json:"attribution_game_in_app_ltv_2days,omitempty"`
	// AttributionGameInAppLtv3Days 所选时间范围内的激活用户，激活后二天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv3Days model.Float64 `json:"attribution_game_in_app_ltv_3days,omitempty"`
	// AttributionGameInAppLtv4Days 所选时间范围内的激活用户，激活后三天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv4Days model.Float64 `json:"attribution_game_in_app_ltv_4days,omitempty"`
	// AttributionGameInAppLtv5Days 所选时间范围内的激活用户，激活后四天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv5Days model.Float64 `json:"attribution_game_in_app_ltv_5days,omitempty"`
	// AttributionGameInAppLtv6Days 所选时间范围内的激活用户，激活后五天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv6Days model.Float64 `json:"attribution_game_in_app_ltv_6days,omitempty"`
	// AttributionGameInAppLtv7Days 所选时间范围内的激活用户，激活后六天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv7Days model.Float64 `json:"attribution_game_in_app_ltv_7days,omitempty"`
	// AttributionGameInAppLtv8Days 所选时间范围内的激活用户，激活后七天内在APP内的付费金额总和。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppLtv8Days model.Float64 `json:"attribution_game_in_app_ltv_8days,omitempty"`
	// AttributionGameInAppRoi1Day 所选时间范围内的激活用户在激活当日的付费ROI，计算公式是：当日付费金额/所选时间的消耗。
	AttributionGameInAppRoi1Day model.Float64 `json:"attribution_game_in_app_roi_1day,omitempty"`
	// AttributionGameInAppRoi2Days 所选时间范围内的激活用户在激活后一日内的所有付费ROI，计算公式是：激活后一日付费金额/所选时间的消耗。该指标隔日产出，并做了模糊化处理。
	AttributionGameInAppRoi2Days model.Float64 `json:"attribution_game_in_app_roi_2days,omitempty"`
	// AttributionGameInAppRoi3Days 所选时间范围内的激活用户在激活后二日内的所有付费ROI，计算公式是：激活后二日付费金额/所选时间的消耗。
	AttributionGameInAppRoi3Days model.Float64 `json:"attribution_game_in_app_roi_3days,omitempty"`
	// AttributionGameInAppRoi4Days 所选时间范围内的激活用户在激活后三日内的所有付费ROI，计算公式是：激活后三日付费金额/所选时间的消耗。
	AttributionGameInAppRoi4Days model.Float64 `json:"attribution_game_in_app_roi_4days,omitempty"`
	// AttributionGameInAppRoi5Days 所选时间范围内的激活用户在激活后四日内的所有付费ROI，计算公式是：激活后四日付费金额/所选时间的消耗。
	AttributionGameInAppRoi5Days model.Float64 `json:"attribution_game_in_app_roi_5days,omitempty"`
	// AttributionGameInAppRoi6Days 所选时间范围内的激活用户在激活后五日内的所有付费ROI，计算公式是：激活后五日付费金额/所选时间的消耗。
	AttributionGameInAppRoi6Days model.Float64 `json:"attribution_game_in_app_roi_6days,omitempty"`
	// AttributionGameInAppRoi7Days 所选时间范围内的激活用户在激活后六日内的所有付费ROI，计算公式是：激活后六日付费金额/所选时间的消耗。
	AttributionGameInAppRoi7Days model.Float64 `json:"attribution_game_in_app_roi_7days,omitempty"`
	// AttributionGameInAppRoi8Days 所选时间范围内的激活用户在激活后七日内的整体付费ROI，计算公式是：激活后七日付费金额/所选时间的消耗。
	AttributionGameInAppRoi8Days model.Float64 `json:"attribution_game_in_app_roi_8days,omitempty"`
	// AttributionDayActivePayCost 消耗/计费当日激活且首次付费数
	AttributionDayActivePayCost model.Float64 `json:"attribution_day_active_pay_cost,omitempty"`
	// AttributionDayActivePayRate 计费当日激活且首次付费数/激活数
	AttributionDayActivePayRate model.Float64 `json:"attribution_day_active_pay_rate,omitempty"`
	// ActivePayIntraDayCount 当日发生激活且首次付费的次数
	ActivePayIntroDayCount model.Int64 `json:"active_pay_intra_day_count,omitempty"`
	// ActivePayIntraDayCost 消耗/激活当日首次付费数
	ActivePayIntroDayCost model.Float64 `json:"active_pay_intra_day_cost,omitempty"`
	// ActivePayIntraDayRate 激活当日首次付费数/激活数
	ActivePayIntroDayRate model.Float64 `json:"active_pay_intra_day_rate,omitempty"`
	// ConvertCnt 将转化数记录在转化事件发生的时间上。建议广告主考核成本时参考“转化数据（计费时间）”例如您的广告在早上8点进行了展示和点击，用户晚上19点发生了激活行为，我们会把激活数披露在晚上19点。
	ConvertCnt model.Int64 `json:"convert_cnt,omitempty"`
	// ConversionCost 广告主为每个转化所付出的平均成本，计算方式：总花费/转化数。当天数据可能会有波动。
	ConversionCost model.Float64 `json:"conversion_cost,omitempty"`
	// ConversionRate 广告被用户转化的次数占点击次数的百分比。计算方式：转化数/点击数*100%
	ConversionRate model.Float64 `json:"conversion_rate,omitempty"`
	// DeepConvertCnt 将深度转化数记录在转化事件发生的时间上。建议广告主考核深度转化成本时参考“深度转化数（计费时间）”例如您的广告在早上8点进行了展示和点击，用户晚上19点发生了激活行为，我们会把激活数披露在晚上19点。
	DeepConvertCnt model.Int64 `json:"deep_convert_cnt,omitempty"`
	// DeepConvertCost 广告主为每个深度转化所付出的平均成本，计算方法：总花费/深度转化数。当天数据可能会有波动，次日早8点后稳定。
	DeepConvertCost model.Float64 `json:"deep_convert_cost,omitempty"`
	// DeepConvertRate 广告被用户进行深度转化的次数占转化次数的百分比。计算方式：深度转化数/转化数*100%
	DeepConvertRate model.Float64 `json:"deep_convert_rate,omitempty"`
	// AttributionConvertCnt 在转化行为发生（或回传）之后，将转化行为回记到过去30天内的扣费（消耗产生）时间上。 例如：广告在8月20日展示给用户，此时广告花费10元，用户点击广告后于8月23日产生1笔购买，则8月23日这笔购买将会展示在8月20日，8月23日没有转化数。
	AttributionConvertCnt model.Int64 `json:"attribution_convert_cnt,omitempty"`
	// AttributionConvertCost 转化成本(计费时间) = 消耗 / 转化数(计费时间)。例如：广告在8月20日展示给用户，此时广告花费10元，用户点击广告后，于8月23日产生2笔购买，则8月20日的转化成本（计费时间） = 5元（即10元除以2笔）。成本考核和系统赔付以该指标为准。
	AttributionConvertCost model.Float64 `json:"attribution_convert_cost,omitempty"`
	// AttributionDeepConvertCnt 在转化行为发生（或回传）之后，将转化行为回记到过去30天内的扣费（消耗产生）时间上。 例如：广告在8月20日展示给用户，此时广告花费10元，用户点击广告后于8月23日产生1笔购买，则8月23日这笔购买将会展示在8月20日。
	AttributionDeepConvertCnt model.Int64 `json:"attribution_deep_convert_cnt,omitempty"`
	// AttributionDeepConvertCost 是一个准确的深度转化成本评估指标。计算方式：消耗 / 深度转化数(计费时间)。例如：广告在8月20日展示给用户，此时广告花费10元，用户点击广告后，于 8 月 23 日产生2笔购买，则8月20日的深度转化成本（计费时间） = 5元（即10元除以2笔）。成本考核和系统赔付以该指标为准。
	AttributionDeepConvertCost model.Float64 `json:"attribution_deep_convert_cost,omitempty"`
	// AttributionDeepConvertRate 深度转化数（计费时间）/转化数（计费时间）*100%
	AttributionDeepConvertRate model.Float64 `json:"attribution_deep_convert_rate,omitempty"`
	// TotalPlay 播放量, 播放时间大于0S的数量，在某些蜂窝网络环境下，需要您手动点击开始才会开始播放，因此有时播放数小于展示数。
	TotalPlay model.Int64 `json:"total_play,omitempty"`
	// ValidPlay 有效播放数, 竞价广告播放时间大于等于10秒的数量，如果视频总时长不足10秒，则记录播放完成的次数。品牌广告在部分APP（头条、头条lite、抖音、西瓜、抖音火山版、皮皮虾）播放时间大于等于5秒的数量，在其他APP大于等于3秒的数量，如果视频总时长不足5秒/3秒时，则记录播放完成的次数。
	ValidPlay model.Int64 `json:"valid_play,omitempty"`
	// ValidPlayCost 有效播放成本, 计算公式：总花费/有效播放数，当天数据可能会有波动，次日早8点后稳定。
	ValidPlayCost model.Float64 `json:"valid_play_cost,omitempty"`
	// ValidPlayRate 有效播放率, 计算公式：有效播放数/展示数。
	ValidPlayRate model.Float64 `json:"valid_play_rate,omitempty"`
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
}
