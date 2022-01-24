package dmp

import "github.com/bububa/oceanengine/marketing-api/enum"

// RetargetingTag 人群包
type RetargetingTag struct {
	// ID 人群包id
	ID uint64 `json:"retargeting_tags_id,omitempty"`
	// Name 人群包名称
	Name string `json:"name,omitempty"`
	// Source 人群包来源，自定义类详见【附录-DMP相关-人群包来源】，平台精选类返回空值
	Source enum.DmpSource `json:"source,omitempty"`
	// Status 人群包来源，自定义类详见【附录-DMP相关-人群包来源】，平台精选类返回空值
	Status int `json:"status,omitempty"`
	// Op 人群包可选的定向规则，枚举值：INCLUDE只支持定向，EXCLUDE只支持排除，ALL支持两种规则。当source为RETARGETING_TAGS_TYPE_PLATFORM时，只支持INCLUDE或EXCLUDE；当source为RETARGETING_TAGS_TYPE_CUSTOM时，支持ALL
	Op string `json:"retargeting_tags_op,omitempty"`
	// CoverNum 预估人群包覆盖人群数目
	CoverNum int64 `json:"cover_num,omitempty"`
	// Tip 人群包说明
	Tip string `json:"retargeting_tags_tip,omitempty"`
	// IsCommon 0 该人群包不支持通投，1 该人群包支持通投，注意：不支持通投的人群包不能在千川平台创建计划，否则会报错。
	IsCommon int `json:"is_common,omitempty"`
}
