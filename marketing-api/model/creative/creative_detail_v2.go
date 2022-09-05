package creative

// CreativeDetailV2 创意详情 (新)
type CreativeDetailV2 struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划ID，计划ID要属于广告主ID，且非删除计划，否则会报错
	AdID uint64 `json:"ad_id,omitempty"`
	// CreativeList 自定义素材信息
	CreativeList []CreativeInfo `json:"creative_list,omitempty"`
	// Creative 程序化素材信息，投放位置和创意类型决定素材规格。
	Creative *CreativeInfo `json:"creative,omitempty"`
	// AdData 广告计划数据
	AdData *AdData `json:"ad_data,omitempty"`
}
