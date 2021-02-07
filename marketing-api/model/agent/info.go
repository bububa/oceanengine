package agent

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type InfoRequest struct {
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"` // 代理商ids
	Fields        []string `json:"fields,omitempty"`         // 查询字段集合, 如果指定, 则返回结果数组中, 每个元素是包含所查询字段的字典，默认全部指定.允许值: "agent_id", "agent_name", "customer_id", "customer_name","company_id", "company_name", "create_time", "role"
}

func (r InfoRequest) Encode() string {
	idsBytes, _ := json.Marshal(r.AdvertiserIDs)
	fieldsBytes, _ := json.Marshal(r.Fields)
	values := &url.Values{}
	values.Set("advertiser_ids", string(idsBytes))
	values.Set("fields", string(fieldsBytes))
	return values.Encode()
}

type InfoResponse struct {
	model.BaseResponse
	Data []Info `json:"data,omitempty"`
}

type Info struct {
	AgentID       uint64 `json:"agent_id,omitempty"`       // 代理商ID
	AgentName     string `json:"agent_name,omitempty"`     // 代理商名称
	CustomerID    uint64 `json:"customer_id,omitempty"`    // 客户id
	CompanyID     uint64 `json:"company_id,omitempty"`     // 公司id
	CompanyName   string `json:"company_name,omitempty"`   // 公司名称
	AccountStatus string `json:"account_status,omitempty"` // 用户状态
	CreateTime    string `json:"create_time,omitempty"`    // 注册时间
	Role          string `json:"role,omitempty"`           // 角色
}
