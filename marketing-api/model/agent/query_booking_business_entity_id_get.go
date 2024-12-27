package agent

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryBookingBusinessEntityIDGetRequest 排期—查询业务实体ID API Request
type QueryBookingBusinessEntityIDGetRequest struct {
	// AgentID 代理商账户ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// OrderIDs 主订单ID
	OrderIDs []uint64 `json:"order_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r QueryBookingBusinessEntityIDGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	values.Set("order_ids", string(util.JSONMarshal(r.OrderIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QueryBookingBusinessEntityIDGetResponse 排期—查询业务实体ID API Response
type QueryBookingBusinessEntityIDGetResponse struct {
	model.BaseResponse
	Data struct {
		// BusinessEntityIDInfos 业务实体信息
		BusinessEntityIDInfos []BusinessEntityIDInfo `json:"business_entity_id_infos,omitempty"`
	} `json:"data,omitempty"`
}

// BusinessEntityIDInfo 业务实体信息
type BusinessEntityIDInfo struct {
	// CartID 排期ID
	CartID uint64 `json:"cart_id,omitempty"`
	// CartName 排期名称
	CartName string `json:"cart_name,omitempty"`
	// BusinessEntityID 业务实体ID
	BusinessEntityID uint64 `json:"business_entity_id,omitempty"`
	// OrderID 主订单ID
	OrderID uint64 `json:"order_id,omitempty"`
}
