package agent

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryStatementRequest 查询项目关联结算单信息 API Request
type QueryStatementRequest struct {
	// AgentID 代理商ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// ProjectIDList 项目id列表，限制1000个
	ProjectIDList []uint64 `json:"project_id_list,omitempty"`
}

// Encode implements GetRequest interface
func (r QueryStatementRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	values.Set("project_id_list", string(util.JSONMarshal(r.ProjectIDList)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QueryStatementResponse 查询项目关联结算单信息 API Response
type QueryStatementResponse struct {
	model.BaseResponse
	Data struct {
		// ProjectRefStatmentList 结算单列表
		ProjectRefStatmentList []ProjectRefStatement `json:"project_ref_statement_list,omitempty"`
	} `json:"data,omitempty"`
}

// ProjectRefStatment 结算单
type ProjectRefStatement struct {
	// ProjectID 项目id
	ProjectID uint64 `json:"project_id,omitempty"`
	// StatementSerial 结算单编号
	StatementSerial string `json:"statement_serial,omitempty"`
	// RefStatementAmount 项目关联结算单金额
	RefStatementAmount model.Float64 `json:"ref_statement_amount,omitempty"`
	// StampStatus 结算单盖章状态（1：未盖章，2：审批中，3：已盖章，5：无需盖章，4: 电子签章审批中)
	StampStatus int `json:"stamp_status,omitempty"`
	// StampStatusName 结算单盖章状态名称
	StampStatusName string `json:"stamp_status_name,omitempty"`
	// ArchiveStatus 结算单归档状态
	ArchiveStatus int `json:"archive_status,omitempty"`
	// ArchiveStatusName 结算单归档状态名称
	ArchiveStatusName string `json:"archive_status_name,omitempty"`
	// AttachList 附件列表
	AttachList []ProjectRefStatmenAttach `json:"attach_list,omitempty"`
}

// ProjectRefStatmentAttach 附件
type ProjectRefStatmenAttach struct {
	// AttachID 附件id
	AttachID uint64 `json:"attach_id,omitempty"`
	// FileName 文件名称
	FileName string `json:"file_name,omitempty"`
}
