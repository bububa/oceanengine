package spda

import (
	"encoding/xml"
)

// CDATA generate cdata field
type CDATA struct {
	Text string `xml:",cdata"`
}

// CDATAList list of cdata fields
type CDATAList struct {
	//Items
	Items []CDATA `xml:"e,omitempty"`
}

// DocIndex 索引文件
type DocIndex struct {
	XMLName xml.Name `xml:"docindex"`
	// FileItem 文件
	FileItem []FileItem `xml:"fileitem"`
}

// FileItem 单条索引文件
type FileItem struct {
	// File 包含商品信息的XML文件地址
	File string `xml:"file"`
	// Lastmod 最后修改时间，单位秒
	Lastmod int64 `xml:"lastmod"`
}

// Image 图片
type Image struct {
	// Url 图片链接
	Url string `xml:"url"`
	// Width 图片宽度
	Width int `xml:"width,omitempty"`
	// Height 图片高度
	Height int `xml:"height,omitempty"`
	// Desc 图片描述
	Desc *CDATA `xml:"description,omitempty"`
}

// ImageList
type ImageList struct {
	Items []Image `xml:"e,omitempty"`
}

// Video 视频
type Video struct {
	// Url 视频链接
	Url string `xml:"url"`
	// Width 视频宽度
	Width int `xml:"width,omitempty"`
	// Height 视频高度
	Height int `xml:"height,omitempty"`
	// Duration 视频时长
	Duration int `xml:"duration,omitempty"`
	// Ratio 视频码率
	Ratio *CDATA `xml:"ratio,omitempty"`
}

// Ext 扩展字段
type Ext struct {
	// Name
	Name CDATA `xml:"name"`
	// Value
	Value CDATA `xml:"value"`
}

type ExtList struct {
	Items []Ext `xml:"e"`
}

// GPS 投放范围
type GPS struct {
	// Lng 经度
	Lng float64 `xml:"longitude"`
	// Lat 维度
	Lat float64 `xml:"latitude"`
	// Range 缓存半径
	Range int `xml:"range"`
}

// Target 定向
type Target struct {
	// GPS 投放范围
	GPS *GPS `xml:"GPS,omitempty"`
	// Country 国家
	Country *CDATAList `xml:"country,omitempty"`
	// Province 省份
	Province *CDATAList `xml:"province,omitempty"`
	// City 城市
	City *CDATAList `xml:"city,omitempty"`
	// TradingArea 商圈
	TradingArea *CDATAList `xml:"tradingArea,omitempty"`
	// Age 年龄
	Age []int `xml:"age>e,omitempty"`
	// Gender 性别
	Gender int `xml:"gender,omitempty"`
	// ToutiaoTag 兴趣标签, https://bytedance.feishu.cn/sheets/shtcn7B174AiQDTsKrkeSNHiJ7c
	ToutiaoTag []int `xml:"toutiaoTag,omitempty"`
}

// Delivery 投放控制
type Delivery struct {
	// DeliverVideo 是否投放视频广告
	DeliverVideo int `xml:"deliverVideo"`
}

// File 商品信息文件
type File struct {
	XMLName xml.Name   `xml:"dataset"`
	Items   []FileData `xml:"data"`
}

type FileData struct {
	// ID 商品ID, 主键，用于区别商品
	ID uint64 `xml:"id"`
	// Name 商品名,
	Name CDATA `xml:"name"`
	// Title 商品标题
	Title *CDATA `xml:"title"`
	// Condition 商品新旧情况, 全新，翻新，二手
	Condition *CDATA `xml:"condition,omitempty"`
	// Desc 商品描述
	Desc *CDATA `xml:"description,omitempty"`
	// Stock 库存状态,用限制商品投放状态；如果投放时间在上线时间和下线时间之间，stock为1，且status为1时可以投放
	Stock int `xml:"stock"`
	// Status 投放状态, 限制商品投放状态；如果投放时间在上线时间和下线时间之间，stock为1，且status为1时可以投放
	Status int `xml:"status"`
	// Image 商品主图, 基础商品图，展示在信息流中的原始素材
	Image Image `xml:"image"`
	// Images 图片列表, 扩展商品图，商品图片的补充, "和image不一致的最好大于3张的商品图片，如果要投放视频广告建议至少5张图片，像素大小690*388以上（5K以上），图片的顺序默认为广告主认为按重要性传输的，与image比例一致，不超过10张"
	Images *ImageList `xml:"images,omitempty"`
	// Video 视频内容，7-30s之间, {""url"":""www.xxx.com"",""width"":1280,""height"":720,""duration"":15,""ratio"":4818//码率,单位kps}，抖音：1280*720（横板）或720*1280（竖版）URL应保证可用，浏览器可打开，比例一致，链接要保持长久有效，不能是临时链接"
	Video *Video `xml:"video,omitempty"`
	// BrandID 品牌ID, "应保证每一种品牌的id唯一性，用于相似商品推荐时强烈建议传，否则影响投放效果，如组图"
	BrandID *CDATA `xml:"brandID,omitempty"`
	// BrandName 品牌名
	BrandName *CDATA `xml:"brandName,omitempty"`
	// EnBrand 英文品牌名
	EnBrand *CDATA `xml:"enBrand,omitempty"`
	// BrandUrl PC端品牌落地页URL
	BrandUrl *CDATA `xml:"brandUrl,omitempty"`
	// BrandUrlMobile H5页面品牌落地页URL
	BrandUrlMobile *CDATA `xml:"brandUrlMobile,omitempty"`
	// BrandUrlAndroidApp Android应用品牌直达吊起链接
	BrandUrlAndroidApp *CDATA `xml:"brandUrlAndroidApp,omitempty"`
	// BrandUrlIosApp IOS应用品牌直达吊起链接
	BrandUrlIosApp *CDATA `xml:"brandUrlIosApp,omitempty"`
	// BrandUrlUniversalLink IOS应用品牌吊起ulink链接
	BrandUrlUniversalLink *CDATA `xml:"brandUrlUniversalLink,omitempty"`
	// ShopKeeperID 商户ID;"应保证每一个商户的id唯一性，用于相似商品推荐时强烈建议传，否则影响投放效果，如组图"
	ShopKeeperID *CDATA `xml:"shopKeeperID,omitempty"`
	// ShopKeeperName 商户名
	ShopKeeperName *CDATA `xml:"shopKeeperName,omitempty"`
	// ShopKeeperUrl PC端商户落地页URL
	ShopKeeperUrl *CDATA `xml:"shopKeeperUrl,omitempty"`
	// ShopKeeperUrlMobile H5页面商户落地页URL
	ShopKeeperUrlMobile *CDATA `xml:"shopKeeperUrlMobile,omitempty"`
	// ShopKeeperUrlAndroidApp Android应用商户直达吊起链接
	ShopKeeperUrlAndroidApp *CDATA `xml:"shopKeeperUrlAndroidApp,omitempty"`
	// ShopKeeperUrlIosApp IOS应用商户直达吊起链接
	ShopKeeperUrlIosApp *CDATA `xml:"shopKeeperUrlIosApp,omitempty"`
	// ShopKeeperUrlUniversalLink IOS应用商户吊起ulink链接
	ShopKeeperUrlUniversalLink *CDATA `xml:"shopKeeperUrlUniversalLink,omitempty"`
	// TargetUrl PC端商品落地页URL
	TargetUrl *CDATA `xml:"targetUrl,omitempty"`
	// TargetUrlMobile H5页面商品落地页URL
	TargetUrlMobile *CDATA `xml:"targetUrlMobile,omitempty"`
	// TargetUrlAndroidApp Android应用直达落地页
	TargetUrlAndroidApp *CDATA `xml:"targetUrlAndroidApp,omitempty"`
	// TargetUrlIosApp IOS应用商品直达吊起链接
	TargetUrlIosApp *CDATA `xml:"targetUrlIosApp,omitempty"`
	// TargetUrlUniversalLink IOS应用商品吊起ulink链接
	TargetUrlUniversalLink *CDATA `xml:"targetUrlUniversalLink,omitempty"`
	// FirstCategory 一级分类
	FirstCategory *CDATA `xml:"firstCategory,omitempty"`
	// SubCategory 二级分类
	SubCategory *CDATA `xml:"subCategory,omitempty"`
	// ThirdCategory 三级分类
	ThirdCategory *CDATA `xml:"thirdCategory,omitempty"`
	// FirstCategoryID 一级分类ID
	FirstCategoryID uint64 `xml:"firstCategryId,omitempty"`
	// SubCategoryID 二级分类ID
	SubCategoryID uint64 `xml:"subCategoryId,omitempty"`
	// ThirdCategoryID 三级分类ID
	ThirdCategoryID uint64 `xml:"thirdCategoryId,omitempty"`
	// Value 商品原价
	Value float64 `xml:"value,omitempty"`
	// PriceUnit 价格单位; RMB，万元两种格式都可
	PriceUnit string `xml:"priceUnit,omitempty"`
	// Saving 减价
	Saving float64 `xml:"saving,omitempty"`
	// Discount 商品折扣
	Discount float64 `xml:"discount,omitempty"`
	// Price 商品现价
	Price float64 `xml:"price,omitempty"`
	// SalesPromotion 促销活动
	SalesPromotion *CDATA `xml:"salesPromotion,omitempty"`
	// DownPayment 首付
	DownPayment *CDATA `xml:"downPayment,omitempty"`
	// Mortage 月付
	Mortage *CDATA `xml:"mortage,omitempty"`
	// DailyMortage 日付
	DailyMortage *CDATA `xml:"dailyMortage,omitempty"`
	// Address 地址; 可填商户地址
	Address *CDATA `xml:"address,omitempty"`
	// Feature 特色; 七天无条件退款
	Feature *CDATA `xml:"feature,omitempty"`
	// Mark 评分
	Mark float64 `xml:"mark,omitempty"`
	// Bought 购买量
	Bought int `xml:"bought,omitempty"`
	// Comments 评论数
	Comments int `xml:"comments,omitempty"`
	// Tag 商品标签
	Tag *CDATA `xml:"tag,omitempty"`
	// Ext 扩展字段; 广告主定制，通知头条方
	Ext *ExtList `xml:"ext,omitempty"`
	// Target 定向字段
	Target *Target `xml:"target,omitempty"`
	// Delivery 投放控制字段
	Delivery *Delivery `xml:"delivery,omitempty"`
	// OnlineTime 上线时间; 时间戳类型，限制商品投放状态，
	OnlineTime int64 `xml:"onlineTime,omitempty"`
	// OfflineTime 下线时间; 时间戳类型，限制商品投放状态；如果投放时间在上线时间和下线时间之间，stock为1，且status为1时可以投放
	OfflineTime int64 `xml:"offlineTime,omitempty"`
}
