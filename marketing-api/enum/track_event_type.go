package enum

// TrackEventType 转化事件类型
type TrackEventType = int

const (
	// Track_ACTIVE 激活;用户下载安装完毕应用之后，在联网环境下打开应用
	Track_ACTIVE TrackEventType = 0
	// Track_REGISTER 注册; 完成应用下载并且在联网环境打开应用后，完成个人账号/游戏角色注册信息提交
	Track_REGISTER TrackEventType = 1
	// Track_PAY 付费; 用户在推广的落地页场景下发生交易并完成至少一笔付款，具体支付形式取决于广告主业务模式
	Track_PAY TrackEventType = 2
	// Track_FORM 表单; 页面内完成表单填写并提交
	Track_FORM TrackEventType = 3
	// Track_CONSULT 在线咨询; 用户点击在线咨询按钮
	Track_CONSULT TrackEventType = 4
	// Track_CONSULT_EFFECTIVE 有效咨询; 点击页面“在线咨询”，且在咨询对话框中内容达到≥1条以上对话记录的，记录一次转化。广告主可自定义有效的标准。
	Track_CONSULT_EFFECTIVE TrackEventType = 5
	// Track_NEXT_DAY_ACTIVE 次留; 用户激活后次日联网环境下打开应用;
	Track_NEXT_DAY_ACTIVE TrackEventType = 6
	// Track_VIEW_CONTENT 关键页面浏览
	Track_VIEW_CONTENT TrackEventType = 7
	// Track_LEADS 有效获客; 用户完成了一次有价值的动作，如预约到店，完成授权等，支持广告主根据业务场景自定义
	Track_LEADS TrackEventType = 19
	// Track_IN_APP_PURCHASE app内下单; 在应用内完成一次订单提交，例如：点击“立即下单”
	Track_IN_APP_PURCHASE TrackEventType = 20
	// Track_IN_APP_VISIT app内访问;用户成功打开访问应用
	Track_IN_APP_VISIT TrackEventType = 21
	// Track_IN_APP_CART app内添加购物车; 在应用内成功将商品加入购物车，例如：点击“加入购物车”
	Track_IN_APP_CART TrackEventType = 22
	// Track_IN_APP_PAY app内付费; 在应用内完成一次订单付费。目前主要是电商行业使用，常规建议使用付费事件
	Track_IN_APP_PAY TrackEventType = 23
	// Track_KEY_ACTION 关键行为; 用户在应用内发生的关键行为/行为集合，若是关键行为集合一般是有关联的行为路径。（举例：某直播类客户以注册+发送弹幕作为关键行为转化目标，电商用注册+收藏商品+加入购物车+下单等）
	Track_KEY_ACTION TrackEventType = 25
	// Track_AUTHORIZE 授权; 完成授权电商/支付/社交等账号登陆
	Track_AUTHORIZE TrackEventType = 28
	// Track_IN_APP_KEY_PAGEVIEW app内详情页到站uv; 成功访问应用内指定页面的UV数
	Track_IN_APP_KEY_PAGEVIEW TrackEventType = 29
	// Track_CLICK_PRODUCT 点击商品; 从多商品页点击某一商品，进入商品详情页，例如：点击某一商品，表达兴趣
	Track_CLICK_PRODUCT TrackEventType = 179
	// Track_WISHLIST 加入收藏/心愿单
	Track_WISHLIST TrackEventType = 128
	// Track_COUPON 领取优惠券
	Track_COUPON TrackEventType = 213
	// Track_PURCHASE 立即购买
	Track_PURCHASE TrackEventType = 175
	// Track_ADD_RECEIVER 添加/选定收货信息、电话
	Track_ADD_RECEIVER TrackEventType = 212
	// Track_ADD_PAYMENT 添加/选定支付信息，绑定支付宝、微信、银行卡等
	Track_ADD_PAYMENT TrackEventType = 127
	// Track_CONFIRM_ORDER 提交订单
	Track_CONFIRM_ORDER TrackEventType = 176
	// Track_CONFIRM_RECEIVE 订单提交/确认收货
	Track_CONFIRM_RECEIVE TrackEventType = 214
	// Track_ENTER_LIVE_ROOM 进入直播间
	Track_ENTER_LIVE_ROOM TrackEventType = 202
	// Track_FOLLOW_LIVE_ROOM 直播间内点击关注按钮; 点击关注/收藏直播间
	Track_FOLLOW_LIVE_ROOM TrackEventType = 204
	// Track_COMMENT_LIVE_ROOM 直播间内评论
	Track_COMMENT_LIVE_ROOM TrackEventType = 205
	// Track_GIFT_LIVE_ROOM 直播间内打赏
	Track_GIFT_LIVE_ROOM TrackEventType = 206
	// Track_CLICK_CART_LIVE_ROOM 直播间内点击购物车按钮
	Track_CLICK_CART_LIVE_ROOM TrackEventType = 207
	// Track_PRODUCT_LIVE_ROOM 直播间内商品点击
	Track_PRODUCT_LIVE_ROOM TrackEventType = 208
	// Track_REDIRECT_LIVE_ROOM 直播间进入种草页跳转到第三方
	Track_REDIRECT_LIVE_ROOM TrackEventType = 209
	// Track_ADD_CART_LIVE_ROOM 直播-加购; 由直播间路径加购
	Track_ADD_CART_LIVE_ROOM TrackEventType = 210
	// Track_PURCHASE_LIVE_ROOM 直播-下单; 由直播间路径下单
	Track_PURCHASE_LIVE_ROOM TrackEventType = 211
	// Track_INFO_CONFIRM 回访_信息确认; 线索经联系确认是本人提交的信息，或者是本人有初步意向了解
	Track_INFO_CONFIRM TrackEventType = 194
	// Track_ADD_CONTACT 回访_加为好友; 线索和销售建立了交流，比如互加好友，建立联系，可以持续跟进
	Track_ADD_CONTACT TrackEventType = 195
	// Track_POTENTIAL_CUSTOMER 回访_高潜成交; 线索有较强意向成交或者处于成交流程，尚未完结
	Track_POTENTIAL_CUSTOMER TrackEventType = 196
	// Track_PAY_POTENTIAL 支付_存在意向; 在表单提交成功（获取用户手机号）之后在落地页的支付行为-支付成功
	Track_PAY_POTENTIAL TrackEventType = 218
	// Track_ECOM_CANCEL_ORDER 电商取消订单
	Track_ECOM_CANCEL_ORDER TrackEventType = 450
	// Track_ECOM_SUCCESSFUL_PAY 电商支付成功
	Track_ECOM_SUCCESSFUL_PAY TrackEventType = 449
	// Track_ECOM_REFUND 电商退款
	Track_ECOM_REFRUND TrackEventType = 448
)
