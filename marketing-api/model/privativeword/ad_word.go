package privativeword

// AdWord 否定词
type AdWord struct {
	// AdID 广告计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// CampaignID 广告组id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// PhraseWords 短语否定词，否定词列表不能超过100个否定词，单个否定词长度不能超过30个汉字（或60个英文字符），不能包含emoji，不能为空字符串。
	PhraseWords []string `json:"phrase_words,omitempty"`
	// PreciseWords 精确否定词，否定词列表不能超过100个否定词，单个否定词长度不能超过30个汉字（或60个英文字符），不能包含emoji，不能为空字符串。
	PreciseWords []string `json:"precise_words,omitempty"`
}
