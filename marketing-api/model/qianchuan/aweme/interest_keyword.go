package aweme

type InterestKeyword struct {
	// ID 兴趣ID，详见【附录-随心推兴趣标签】
	ID uint64 `json:"id,omitempty"`
	// Name 兴趣名称
	Name string `json:"name,omitempty"`
}
