package star

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// InfoRequest 获取星图账户信息 API Request
type InfoRequest struct {
	// StarIDs 星图id
	StarIDs []uint64 `json:"star_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r InfoRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("star_ids", string(util.JSONMarshal(r.StarIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// InfoResponse 获取星图账户信息 API Response
type InfoResponse struct {
	model.BaseResponse
	Data struct {
		// InfoList 信息列表
		InfoList []Info `json:"info_list,omitempty"`
	} `json:"data,omitempty"`
}

// Info 星图账户信息
type Info struct {
	// StarID 星图id
	StarID uint64 `json:"star_id,omitempty"`
	// StarName 星图名称
	StarName string `json:"star_name,omitempty"`
	// Status 账户状态
	// DELETED 已删除
	// ENABLE 有效
	// FROZEN 禁用
	// NEW_PROTOCOL 待同意新协议
	// PUNISH 惩罚
	// QUALIFICATION_VERIFICATION 资质验证
	// UN_PROTOCOL 未同意协议
	Status enum.StarInfoStatus `json:"status,omitempty"`
	// CompanyName 公司名称
	CompanyName string `json:"company_name,omitempty"`
	// CategoryID 行业分类id
	CategoryID uint64 `json:"categorty_id,omitempty"`
	// CompanyID 公司id
	CompanyID uint64 `json:"company_id,omitempty"`
	// FirstInfo 一级行业
	FirstInfo *IndustryInfo `json:"first_info,omitempty"`
	// SecondInfo 二级行业
	SecondInfo *IndustryInfo `json:"second_info,omitempty"`
	// CreateTime 创建时间
	CreateTime int64 `json:"create_time,omitempty"`
}

// IndustryInfo 行业信息
type IndustryInfo struct {
	// IndustryName 行业名称
	IndustryName string `json:"industry_name,omitempty"`
	// IndustryID 行业id
	IndustryID uint64 `json:"industry_id,omitempty"`
}
