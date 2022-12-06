package agent

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// InfoRequest 获取代理商信息 API Request
type InfoRequest struct {
	// AdvertiserIDs 代理商ids
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// Fields 查询字段集合, 如果指定, 则返回结果数组中, 每个元素是包含所查询字段的字典，默认全部指定.允许值: "agent_id", "agent_name", "customer_id", "customer_name","company_id", "company_name", "create_time", "role"
	Fields []string `json:"fields,omitempty"`
}

// Encode implement GetRequest inteface
func (r InfoRequest) Encode() string {
	idsBytes, _ := json.Marshal(r.AdvertiserIDs)
	fieldsBytes, _ := json.Marshal(r.Fields)
	values := util.GetUrlValues()
	values.Set("advertiser_ids", string(idsBytes))
	values.Set("fields", string(fieldsBytes))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// InfoResponse 获取代理商信息 API Response
type InfoResponse struct {
	model.BaseResponse
	// Data json返回值
	Data []Info `json:"data,omitempty"`
}

// Info 代理商信息
type Info struct {
	// AgentID 代理商ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// AgentName 代理商名称
	AgentName string `json:"agent_name,omitempty"`
	// CustomerID 客户id
	CustomerID uint64 `json:"customer_id,omitempty"`
	// CompanyID 公司id
	CompanyID uint64 `json:"company_id,omitempty"`
	// CompanyName 公司名称
	CompanyName string `json:"company_name,omitempty"`
	// AccountStatus 用户状态
	AccountStatus string `json:"account_status,omitempty"`
	// CreateTime 注册时间
	CreateTime string `json:"create_time,omitempty"`
	// Role 角色
	Role string `json:"role,omitempty"`
}
