package product

import "github.com/bububa/oceanengine/marketing-api/enum/qianchuan"

// Product 商品
type Product struct {
	// ID 商品id
	ID uint64 `json:"id,omitempty"`
	// Img 主图
	Img string `json:"img,omitempty"`
	// Inventory 库存
	Inventory int64 `json:"inventory,omitempty"`
	// Name 商品名称
	Name string `json:"name,omitempty"`
	// CategoryName 分类
	CategoryName string `json:"category_name,omitempty"`
	// ProductRate 好评率
	ProductRate float64 `json:"product_rate,omitempty"`
	// SaleTime 上架时间
	SaleTime string `json:"sale_time,omitempty"`
	// DiscountPrice 售价，单位：元，已废弃
	DiscountPrice float64 `json:"discount_price,omitempty"`
	// MarketPrice 原价，单位为元
	MarketPrice float64 `json:"market_price,omitempty"`
	// DiscountLowerPrice 折扣价区间最小值，单位为元
	DiscountLowerPrice float64 `json:"discount_lower_price,omitempty"`
	// DiscountHigherPrice 折扣价区间最大值，单位为元
	DiscountHigherPrice float64 `json:"discount_higher_price,omitempty"`
	// Tags 是否可投猜喜可选值:
	// 1: 可投猜喜
	// 0: 不可投猜喜
	Tags int `json:"tags,omitempty"`
	// SupportProductNewOpen 当前商品是否支持开启新品加速
	// 支持：true
	// 不支持：false
	SupportProductNewOpen bool `json:"support_product_new_open,omitempty"`
	// ImgList 商品卡方图url列表
	ImgList []struct {
		// ImgURL 商品卡方图url
		// 注意：在使用商品卡投放广告时，需要先将商品卡方图下载，然后上传至素材库
		ImgURL string `json:"img_url,omitempty"`
	} `json:"img_list,omitempty"`
	// ChannelID 渠道ID，如果渠道品生效，价格、销量等信息需要返回渠道品信息
	// 注意：在抖音号入參的情况下，如果当前商品在该抖音号下存在渠道品，会返回渠道品信息。在创编链路，需要使用渠道品来创建计划
	// 渠道品相关介绍见《【抖店】销售渠道品功能操作手册》
	ChannelID uint64 `json:"channel_id,omitempty"`
	// ChannelType 渠道类型，枚举值
	// SHOP_SELL 商家自卖
	// STAR_SELL 达人自播
	ChannelType qianchuan.ProductChannelType `json:"channel_type,omitempty"`
}
