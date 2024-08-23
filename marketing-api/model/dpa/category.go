package dpa

import "github.com/bububa/oceanengine/marketing-api/model"

// Category DPA分类
type Category struct {
	// Name 分类名称
	Name string `json:"name,omitempty"`
	// Subs 子级分类，嵌套递归
	Subs []Category `json:"subs,omitempty"`
	// ID 分类id
	ID model.Uint64 `json:"id,omitempty"`
	// Parent 父级分类id，没有父级则为-1
	Parent model.Uint64 `json:"parent,omitempty"`
}
