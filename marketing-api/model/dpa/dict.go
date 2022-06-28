package dpa

// Dict 词包
type Dict struct {
	// ID 词包id
	ID uint64 `json:"id,omitempty"`
	// PdaID 商品库id
	PdaID uint64 `json:"pda_id,omitempty"`
	// DefaultWord 默认词
	DefaultWord string `json:"default_word,omitempty"`
	// Name 词包名称
	Name string `json:"name,omitempty"`
}
