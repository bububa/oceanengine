package dpa

// ProductForUpdate 商品详情
type ProductForUpdate struct {
	// Name 商品名称
	Name string `json:"name,omitempty"`
	// Title 商品标题
	Title string `json:"title,omitempty"`
	// Description 商品标题
	Description string `json:"description,omitempty"`
	// OfflineTime 下线时间, 格式为十位unix时间戳类型，默认为当前时间+3年
	OfflineTime int64 `json:"offline_time,omitempty"`
	// OnlineTime 上线时间, 格式为十位unix时间戳类型，默认为当前时间
	OnlineTime int64 `json:"online_time,omitempty"`
	// Status 商品投放状态，0代表不可投放，1代表可投放
	Status int `json:"status"`
	// Stock 商品库存状态，0代表无库存，1代表有库存
	Stock int `json:"stock"`
	// FirstCategory 商品所处一级行业
	FirstCategory string `json:"first_category,omitempty"`
	// SubCategory 商品所处二级行业
	SubCategory string `json:"sub_category,omitempty"`
	// ThirdCategory 商品所处三级行业
	ThirdCategory string `json:"third_category,omitempty"`
	// FirstCategoryID 商品所处一级行业 ID
	FirstCategoryID string `json:"first_category_id,omitempty"`
	// SubCategoryID 商品所处二级行业 ID
	SubCategoryID string `json:"sub_category_id,omitempty"`
	// ThirdCategoryID 商品所处三级行业 ID
	ThirdCategoryID string `json:"third_category_id,omitempty"`
	// SpuID 商品spu_id
	SpuID string `json:"spu_id,omitempty"`
	// OuterID 商品外部id
	OuterID string `json:"outer_id,omitempty"`
	// ImageURL 商品封面图片链接
	ImageURL string `json:"image_url,omitempty"`
	// ImageURLs 扩展商品图，商品图片的补充
	ImageURLs []Link `json:"image_urls,omitempty"`
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

func CopyProductForUpdateFromProduct(src *Product, dist *ProductForUpdate) {
	dist.Name = src.Name
	dist.Title = src.Title
	dist.Description = src.Description
	dist.OfflineTime = src.OfflineTime.Value()
	dist.Status = src.Status.Value()
	dist.Stock = src.Stock
	dist.FirstCategory = src.FirstCategory
	dist.SubCategory = src.SubCategory
	dist.ThirdCategory = src.ThirdCategory
	dist.FirstCategoryID = src.FirstCategoryID.String()
	dist.SubCategoryID = src.SubCategoryID.String()
	dist.ThirdCategoryID = src.ThirdCategoryID.String()
	dist.SpuID = src.SpuID
	dist.OuterID = src.OuterID
	dist.ImageURL = src.ImageURL
	dist.ImageURLs = src.ImageURLs
	dist.LandingInfo = src.LandingInfo
	dist.BrandInfo = src.BrandInfo
	dist.ShopKeeperInfo = src.ShopKeeperInfo
	dist.PriceInfo = src.PriceInfo
	dist.Feature = src.Feature
	dist.Mark = src.Mark
	dist.Bought = src.Bought
	dist.Comments = src.Comments
	dist.Province = src.Province
	dist.City = src.City
	dist.Age = src.Age
	dist.Label = src.Label
	dist.ExternalURL = src.ExternalURL
	dist.BrandName = src.BrandName
	dist.Tags = src.Tags
	dist.Video = src.Video
	dist.Videos = src.Videos
	dist.Profession = src.Profession
}
