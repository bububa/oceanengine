package report

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
)

// Dimension 维度数据
type Dimension struct {
	// StatDatetime 数据起始时间，分组条件包含 STAT_GROUP_BY_FIELD_STAT_TIME 时返回，格式为：yyyy-MM-dd HH:mm:ss
	StatDatetime string `json:"stat_datetime,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告组id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 广告组名称
	CampaignName string `json:"campaign_name,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// AdName 计划名称
	AdName string `json:"ad_name,omitempty"`
	// CreativeID 创意id
	CreativeID uint64 `json:"creative_id,omitempty"`
	// MaterialType 素材类型，允许值：
	// video ：视频素材
	// image ：图片素材
	MaterialType qianchuan.MaterialType `json:"material_type,omitempty"`
	// MarketingGoal 营销目标
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// MaterialID 素材id，一个素材唯一对应一个素材id，相同素材上传多次对应一个material_id
	MaterialID uint64 `json:"material_id,omitempty"`
	// CreateTime 素材创建时间，即该素材最近一次上传素材库的时间
	CreateTime string `json:"create_time,omitempty"`
	// Duration 视频时长，仅素材类型为视频素材有效
	Duration float64 `json:"duration,omitempty"`
	// RelatedAdCnt 关联计划数
	RelatedAdCnt int64 `json:"related_ad_cnt,omitempty"`
	// RelatedAdIDs 关联计划id
	RelatedAdIDs []uint64 `json:"related_ad_ids,omitempty"`
	// RelatedCreativeCnt 关联创意数
	RelatedCreativeCnt int64 `json:"related_creative_cnt,omitempty"`
	// RelatedCreativeIDs 关联创意id
	RelatedCreativeIDs []uint64 `json:"related_creative_ids,omitempty"`
	// Tags 素材标签，枚举值：
	// LQ：低质
	// Null：没有
	Tags string `json:"tags,omitempty"`
	// MaterialMode 素材样式筛选
	// 素材类型为视频素材时允许值：
	// VIDEO_LARGE：横版视频
	// VIDEO_VERTICAL：竖版视频
	// 素材类型为图片素材时允许值：
	// SMALL：横版小图
	// LARGE：横版大图
	// LARGE_VERTICAL：竖版图片
	// SQUARE ：商品卡方图
	MaterialMode enum.MaterialMode `json:"material_mode,omitempty"`
	// VideoSource 视频来源筛选，以平台素材库接口的视频来源枚举值为准，允许值：
	// E_COMMERCE：本地上传
	// LIVE_HIGHLIGHT：直播剪辑素材
	// BP：巨量纵横共享素材
	// VIDEO_CAPTURE：易拍APP共享素材
	// ARTHUR：亚瑟共享素材
	// STAR：星图&即合共享素材
	// TADA：tada共享素材
	// CREATIVE_CENTER：巨量创意PC共享素材
	// 注意：仅素材类型为视频素材时，支持
	VideoSource enum.MaterialSource `json:"video_source,omitempty"`
	// ImageSource 图片来源，可选值:
	// CREATIVE_CENTER 巨量创意PC共享素材
	// E_COMMERCE 巨量千川本地上传
	// JI_CHUANG 即创共享素材
	// SQUARE 商品图
	ImageSource enum.MaterialSource `json:"image_source,omitempty"`
	// CarouselSource 图文来源，可选值:
	// AWEME 抖音号主页视频
	// E_COMMERCE 巨量千川本地上传
	// JI_CHUANG 即创共享素材
	CarouselSource enum.MaterialSource `json:"carousel_source,omitempty"`
	// StartDate 开始日期
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束日期
	EndDate string `json:"end_date,omitempty"`
	// SearchWord 搜索词
	SearchWord string `json:"search_word,omitempty"`
	// Keyword 关键词
	Keyword string `json:"key_word,omitempty"`
	// KeywordMatchType 触发（匹配）方式:
	// EXTENSIVE: 广泛
	// PHRASE: 短语
	// PRECISION: 精确
	KeywordMatchType enum.KeywordMatchType `json:"keyword_match_type,omitempty"`
	// HSec 秒。指标数据对应的视频时间进度，例如：h_sec=3表示视频第三秒时产生的指标数据
	HSec int `json:"h_sec,omitempty"`
}
