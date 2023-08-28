package agent

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdvertiserCopyRequest 广告主账户复制 API Request
type AdvertiserCopyRequest struct {
	// AgentID 代理商账户ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// AdvertiserID 被复制广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Item 复制账户信息
	Item []AdvertiserCopyItem `json:"item,omitempty"`
	// Copytag 是否复制原广告主账户的账户标签
	CopyTag bool `json:"copy_tag,omitempty"`
	// ReportType 自运营报备类型 可选值:
	// EMPTY 不报备
	// INCREASE_QUANTITY 走量报备
	// SELF_OPERATION 自运营报备
	//  默认值: EMPTY
	Reportype enum.AdvertiserReportType `json:"reportype,omitempty"`
}

// AdvertiserCopyItem 复制账户信息
type AdvertiserCopyItem struct {
	// Name 新广告主账户名称。不可重复
	Name string `json:"name,omitempty"`
	// AdvertiserID 广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Success 复制是否成功
	Success bool
	// ErrorMsg 当失败时，返回的失败原因
	ErrorMsg string `json:"error_msg,omitempty"`
}

// IsError returns true if item is error
func (i AdvertiserCopyItem) IsError() bool {
	return !i.Success
}

// Error implement error interface
func (i AdvertiserCopyItem) Error() string {
	return util.StringsJoin(i.ErrorMsg, ", name:", i.Name)
}

// Encode implement PostRequest interface
func (r AdvertiserCopyRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdvertiserCopyResponse 广告主账户复制 API Response
type AdvertiserCopyResponse struct {
	model.BaseResponse
	Data *AdvertiserCopyResult `json:"data,omitempty"`
}

type AdvertiserCopyResult struct {
	// CopyStatus 复制结果状态码 可选值:
	// 1 全部成功
	// 2 部分成功
	// 3 全部失败
	CopyStatus enum.AdvertiserCopyStatus `json:"copy_status,omitempty"`
	// Item 复制结果，包括单个账户的失败原因
	Item []AdvertiserCopyItem `json:"item,omitempty"`
}
