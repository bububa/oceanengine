package eventmanager

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ShareRequest 事件管理资产共享 API Request
type ShareRequest struct {
	// OrganizationID 纵横组织id
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// AssetID 事件管理资产id
	AssetID uint64 `json:"asset_id,omitempty"`
	// ShareMode 共享模式，可选值:
	// ALL: 组织下某业务线账户PART: 指定账户共享
	ShareMode ShareMode `json:"share_mode,omitempty"`
	// AccountInfos 共享账户信息，一次最多操作10个
	// 当share_mode="PART"时必填
	AccountInfos []ShareAccount `json:"account_infos,omitempty"`
	// AllAccountType 共享账户业务线
	// 当share_mode="ALL"时必填，可选值：:
	// AD
	AllAccountType enum.AccountType `json:"all_account_type,omitempty"`
}

// Encode implement PostRequest interface
func (r ShareRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ShareResponse 事件管理资产共享 API Response
type ShareResponse struct {
	model.BaseResponse
	Data struct {
		// ErrorList 共享失败的账户ID及原因列表
		ErrorList []ShareError `json:"error_list,omitempty"`
	} `json:"data,omitempty"`
}

// ShareError .
type ShareError struct {
	ShareInfo
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message,omitempty"`
}

func (e ShareError) Error() string {
	builder := util.GetStringsBuilder()
	if e.AccountInfo != nil {
		builder.WriteString("AccountID: ")
		builder.WriteString(strconv.FormatUint(e.AccountInfo.AccountID, 10))
		builder.WriteString(", AccountType: ")
		builder.WriteString(string(e.AccountInfo.AccountType))
		builder.WriteString(", ")
	}
	builder.WriteString(e.ErrorMessage)
	ret := builder.String()
	util.PutStringsBuilder(builder)
	return ret
}
