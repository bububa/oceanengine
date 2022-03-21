package creative

// AbstractMaterial 摘要素材
type AbstractMaterial struct {
	// TextAbstractInfo 文本摘要信息; 注意：与标签摘要struct_abstract_info选其一
	TextAbstractInfo *TextAbstractInfo `json:"text_abstract_info,omitempty"`
	// StructAbstractInfo 标签摘要信息列表;注意：与文本摘要text_abstract_info选其一
	StructAbstractInfo *StructAbstractInfo `json:"struct_abstract_info,omitempty"`
}

// TextAbstractInfo 文本摘要信息
type TextAbstractInfo struct {
	// AbstractText 文本摘要内容,长度限制为25-45个字, 两个英文字符占1位。使用文本摘要(text_abstract_info)时必填; 如果要使用动态词包，格式如下：“XXX{词包名}XXX{词包名}XXX”，请注意当您使用动态词包需在下方 word_list 字段中按顺序传入词包ID，并且在一个文本摘要内容中最多使用两个动态词包。如果要使用搜索关键词，格式如下：“XXX{#关键词#}XXX”，请注意当您使用关键词需在下方 bidword_list 字段中传入关键词，并且在一个文本摘要内容中最多使用一个关键词
	AbstractText string `json:"abstract_text,omitempty"`
	// BidwordList 搜索关键词列表
	BidwordList []WordListItem `json:"bidword_list,omitempty"`
	// WordList 动态词包ID，可使用 【查询动态词包接口】 获得，结合文本摘要内容中的词包格式您需要填写相同个数与顺序的词包ID，如果实际ID顺序与文本摘要内容中词包名顺序不一致我们将以词包ID顺序为准。
	WordList []WordListItem `json:"word_list,omitempty"`
	// CreativeWordIDs 动态词包ID
	CreativeWordIDs []WordListItem `json:"creative_word_ids,omitempty"`
}

// StructAbstractInfo 标签摘要信息
type StructAbstractInfo struct {
	// AbstractLabel 摘要标签,单个标签限2-4字，标签需文本不同。使用标签摘要(struct_abstract_info)时必填
	AbstractLabel string `json:"abstract_label,omitempty"`
	// AbstractText 摘要内容,内容限2-10字。使用标签摘要(struct_abstract_info)时必填
	AbstractText string `json:"abstract_text,omitempty"`
}
