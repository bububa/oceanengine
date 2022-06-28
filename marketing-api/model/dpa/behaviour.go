package dpa

// Behaviour 行为
type Behaviour struct {
	// CodeID 行为id
	CodeID uint64 `json:"code_id,omitempty"`
	// Name 行为名称
	Name string `json:"name,omitempty"`
}
