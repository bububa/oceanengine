package dpa

// Product 商品详情
type Product struct {
	// Name 商品名称
	Name string `json:"name,omitempty"`
	// Title 商品标题
	Title string `json:"title,omitempty"`
	// Description 商品标题
	Description string `json:"description,omitempty"`
	// OfflineTime 下线时间，格式"YYYY-MM-DD"
	OfflineTime string `json:"offline_time,omitempty"`
	// OnlineTime 上线时间，格式"YYYY-MM-DD"
	OnlineTime string `json:"online_time,omitempty"`
	// PlatformID 商品库ID
	PlatformID uint64 `json:"platform_id,omitempty"`
	// ProductID 商品ID
	ProductID uint64 `json:"product_id,omitempty"`
	// SpuID 商品spu_id
	SpuID string `json:"spu_id,omitempty"`
	// OuterID 商品外部id
	OuterID string `json:"outer_id,omitempty"`
	// ImageURL 商品封面图片链接
	ImageURL string `json:"image_url,omitempty"`
	// ImageURLs 扩展商品图，商品图片的补充
	ImageURLs []Link `json:"image_urls,omitempty"`
	// Status 商品投放状态，0代表不可投放，1代表可投放
	Status int `json:"status,omitempty"`
	// Stock 商品库存状态，0代表无库存，1代表有库存
	Stock int `json:"stock,omitempty"`
	// LandingInfo 落地页信息
	LandingInfo *LandingInfo `json:"landing_info,omitempty"`
	// BrandInfo 品牌信息
	BrandInfo *BrandInfo `json:"brand_info,omitempty"`
	// ShopKeeperInfo 商户信息
	ShopKeeperInfo *ShopKeeperInfo `json:"shop_keeper_info,omitempty"`
	// PriceInfo 价格信息
	PriceInfo *PriceInfo `json:"price_info,omitempty"`
	// Feature 特色信息
	Feature string `json:"feature,omitempty"`
	// Mark 评分
	Mark float64 `json:"mark,omitempty"`
	// Bought 购买量
	Bought int64 `json:"bought,omitempty"`
	// Comments 评论数
	Comments int64 `json:"comments,omitempty"`
	// Province 省份，用于定向人群，默认不限，示例：["江苏","浙江"]
	Province []string `json:"province,omitempty"`
	// City 定向城市
	City []string `json:"city,omitempty"`
	// Age 年龄段，用于定向人群，默认不限，数组项允许值如下：
	// 1 2 3 4 5 6
	// 1代表年龄段<18
	// 2 代表年龄段在18~23之间
	// 3代表年龄段在24~30之间
	// 4代表年龄段在31~40之间
	// 5代表年龄段在41~49之间
	// 6代表年龄段>50
	// 如：[2,4]代表年龄段在18~23之间或31~40之间
	Age []int `json:"age,omitempty"`
	// Label 商品标签，小说库特有字段
	Label string `json:"label,omitempty"`
	// ExternalURL 落地页链接
	ExternalURL string `json:"external_url,omitempty"`
	// FirstCategory 商品所处一级行业
	FirstCategory string `json:"first_category,omitempty"`
	// SubCategory 商品所处二级行业
	SubCategory string `json:"sub_category,omitempty"`
	// ThirdCategory 商品所处三级行业
	ThirdCategory string `json:"third_category,omitempty"`
	// FirstCategoryID 商品所处一级行业 ID
	FirstCategoryID uint64 `json:"first_category_id,omitempty"`
	// SubCategoryID 商品所处二级行业 ID
	SubCategoryID uint64 `json:"sub_category_id,omitempty"`
	// ThirdCategoryID 商品所处三级行业 ID
	ThirdCategoryID uint64 `json:"third_category_id,omitempty"`
	// BrandName 商品名称
	BrandName string `json:"brand_name,omitempty"`
	// Tags 商品标签
	Tags []string `json:"tags,omitempty"`
	// Video 视频链接url
	Video string `json:"video,omitempty"`
	// Videos 视频内容，小说库特有字段
	Videos []Link `json:"videos,omitempty"`
	// HasVideo 当前商品是否有商品视频 0：没有，1：有
	HasVideo int `json:"has_video,omitempty"`
	// Profession 额外信息
	Profession *Profession `json:"profession,omitempty"`
}

type Link struct {
	URL string `json:"string,omitempty"`
}

// LandingInfo 落地页信息
type LandingInfo struct {
	// TargetURL PC端商品落地页URL
	TargetURL string `json:"target_url,omitempty"`
	// TargetURLMobile H5页面商品落地页URL
	TargetURLMobile string `json:"target_url_mobile,omitempty"`
	// TargetURLAndroidApp Android应用直达落地页
	TargetURLAndroidApp string `json:"target_url_android_app,omitempty"`
	// TargetURLIosApp IOS应用商品直达调起链接
	TargetURLIosApp string `json:"target_url_ios_app,omitempty"`
	// TargetURLUniversalLink IOS应用商品调起ulink链接
	TargetURLUniversalLink string `json:"target_url_universal_link,omitempty"`
}

// BrandInfo 品牌信息
type BrandInfo struct {
	// BrandID 品牌ID
	BrandID string `json:"brand_id,omitempty"`
	// BrandName 品牌名称
	BrandName string `json:"brand_name,omitempty"`
	// EnBrand 英文品牌名
	EnBrand string `json:"en_brand,omitempty"`
	// BrandLogo 品牌logo
	BrandLogo string `json:"brand_logo,omitempty"`
	// BrandURL PC端品牌落地页URL
	BrandURL string `json:"brand_url,omitempty"`
	// BrandURLMobile H5页面品牌落地页URL
	BrandURLMobile string `json:"brand_url_mobile,omitempty"`
	// BrandURLAndroidApp Android应用品牌直达调起链接
	BrandURLAndroidApp string `json:"brand_url_android_app,omitempty"`
	// BrandURLIosApp IOS应用品牌直达调起链接
	BrandURLIosApp string `json:"brand_url_ios_app,omitempty"`
	// BrandURLUniversalLink IOS应用品牌调起ulink链接
	BrandURLUniversalLink string `json:"brand_url_universal_link,omitempty"`
}

// ShopKeeperInfo 商户信息
type ShopKeeperInfo struct {
	// ShopKeeperID 商户ID
	ShopKeeperID string `json:"shop_keeper_id,omitempty"`
	// ShopKeeperName 商户名称
	ShopKeeperName string `json:"shop_keeper_name,omitempty"`
	// ShopKeeperURL PC端商户落地页URL
	ShopKeeperURL string `json:"shop_keeper_url,omitempty"`
	// ShopKeeperURLMobile H5页面商户落地页URL
	ShopKeeperURLMobile string `json:"shop_keeper_url_mobile,omitempty"`
	// ShopKeeperURLAndroidApp Android应用商户直达调起链接
	ShopKeeperURLAndroidApp string `json:"shop_keeper_url_android_app,omitempty"`
	// ShopKeeperURLIosApp IOS应用商户直达调起链接
	ShopKeeperURLIosApp string `json:"shop_keeper_url_ios_app,omitempty"`
	// ShopKeeperURLUniversalLink IOS应用商户调起ulink链接
	ShopKeeperURLUniversalLink string `json:"shop_keeper_url_universal_link,omitempty"`
	// Address 商户地址
	Address string `json:"address,omitempty"`
}

// PriceInfo 价格信息
type PriceInfo struct {
	// Value 商品原价，可用于素材拼接，以及动态创意标题或者素材
	Value float64 `json:"value,omitempty"`
	// PriceUnit 价格单位
	PriceUnit string `json:"price_unit,omitempty"`
	// Saving 减价
	Saving float64 `json:"saving,omitempty"`
	// Discount 折扣
	Discount float64 `json:"discount,omitempty"`
	// Price 商品现价
	Price float64 `json:"price,omitempty"`
	// SalesPromotion 促销活动，关于商品促销活动的描述信息
	SalesPromotion string `json:"sales_promotion,omitempty"`
	// DownPayment 首付
	DownPayment string `json:"down_payment,omitempty"`
	// Montage 月付
	Montage string `json:"montage,omitempty"`
	// DailyMontage 日付
	DailyMontage string `json:"daily_montage,omitempty"`
}

// Profession 额外信息
type Profession struct {
	// Platform 平台信息，枚举值：淘宝 天猫 京东 其他
	Platform string `json:"platform,omitempty"`
	// Chapter 章节信息，结构为一个json字符串，章节名称name长度小于100；章节详细内容content长度小于5000
	// 示例："[{"id":"0","name":"莫欺少年穷","content":"莫欺少年穷"},{"id":"1","name":"莫欺中年穷","content":"莫欺中年穷"}]"
	Chapter string `json:"chapter,omitempty"`
	// NovelLength 书籍长短篇，枚举值：长篇书 短篇书
	NovelLength string `json:"novel_length,omitempty"`
	// NovelGender 书籍男女频，枚举值：男频书 女频书 其他
	NovelGender string `json:"novel_gender,omitempty"`
	// NovelAuthor 书籍作者
	NovelAuthor string `json:"novel_author,omitempty"`
	// NovelStar 书籍收藏数
	NovelStar string `json:"novel_star,omitempty"`
	// NovelType 书籍付费/免费信息，枚举值：免费小说 付费小说
	NovelType string `json:"novel_type,omitempty"`
	// StartPayChapter 起始付费章节
	StartPayChapter string `json:"start_pay_chapter,omitempty"`
	// MinPayMoney 最低付费金额
	MinPayMoney string `json:"min_pay_money,omitempty"`
	// AccPayPeople 累计付费人数
	AccPayPeople string `json:"acc_pay_people,omitempty"`
	// AccPayMoney 累计付费金额
	AccPayMoney string `json:"acc_pay_money,omitempty"`
}
