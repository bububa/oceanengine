package interestaction

import "github.com/bububa/oceanengine/marketing-api/model"

// Object 类目/关键词对象
type Object struct {
	// ID 类目/关键词IID
	ID model.Uint64 `json:"id,omitempty"`
	// Name 名称
	Name string `json:"name,omitempty"`
	// Num 覆盖人数
	Num string `json:"num,omitempty"`
	// Children 子类目
	Children []Object `json:"children,omitempty"`
}
