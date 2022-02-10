package creative

// TitleMaterial 标题素材
type TitleMaterial struct {
	// Title 创意标题，如果要使用动态词包，格式：“XXX{词包名}XXX{词包名}XXX”。请注意当您使用动态词包时，需在word_list字段中按顺序传入词包ID，并且在一个标题中最多使用两个动态词包。长度为5-30个字, 两个英文字符占1位。
	Title string `json:"title,omitempty"`
	// SubTitle APP 副标题素材。仅推广目标为APP，4到24个字符，填写Android下载链接时可设置,每个创意下只需传入一个副标题
	SubTitle string `json:"sub_title,omitempty"`
	// WordList 动态词包ID，可使用 【查询动态词包接口】 获得，结合标题中的词包格式您需要填写相同个数与顺序的词包ID，如果实际ID顺序与标题中词包名顺序不一致我们将以词包ID顺序为准。
	WordList []WordListItem `json:"word_list,omitempty"`
	// DpaWordList 创建DPA创意时可以选择传入DPA词包ID列表，动态词包与DPA词包总数最多为2。可通过【获取DPA词包】接口获取，结合标题中的词包格式您需要填写相同个数与顺序的词包ID，如果实际ID顺序与标题中词包名顺序不一致我们将以词包ID顺序为准
	DpaWordList []WordListItem `json:"dpa_word_list,omitempty"`
	// BidwordList 搜索关键词列表，创建搜索广告创意时选择传入，对应title格式如下：“XXX{#关键词#}XXX”
	BidwordList []WordListItem `json:"bidword_list,omitempty"`
	// CreativeWordIDs 动态词包ID，最多支持两个词包。可使用【查询动态创意词包】获得，结合标题中的词包格式您需要填写相同个数与顺序的词包ID，如果实际ID顺序与标题中词包名顺序不一致我们将以词包ID顺序为准
	CreativeWordIDs []uint64 `json:"creative_word_ids,omitempty"`
	// DpaDictIDs DPA词包ID列表，动态词包与DPA词包总数最多为2。可通过【获取DPA词包】接口获取，结合标题中的词包格式您需要填写相同个数与顺序的词包ID，如果实际ID顺序与标题中词包名顺序不一致我们将以词包ID顺序为准
	DpaDictIDs []uint64 `json:"dpa_dict_ids,omitempty"`
}

// WordListItem 动态词包
type WordListItem struct {
	WordID      uint64 `json:"word_id,omitempty"`
	DefaultWord string `json:"default_word,omitempty"`
}
