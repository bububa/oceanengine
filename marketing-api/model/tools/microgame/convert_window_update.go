package microgame

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ConvertWindowUpdateRequest 修改字节小游戏归因激活时间窗 API Request
type ConvertWindowUpdateRequest struct {
	// AccountID 账户ID，account_type对应的账户ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型，允许值：
	// AD 广告主账户
	// 本次字节小游戏归因激活时间窗传AD
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// AssetID 资产id
	// 字节小游戏传instance_id的值（对应创建接口出参的值）
	AssetID string `json:"asset_id,omitempty"`
	// AssetType 资产类型，允许值：
	// BYTE_GAME：字节小游戏
	AssetType enum.MiniProgramType `json:"asset_type,omitempty"`
	// ConvertWindow 时间窗口设置，允许值：
	// ONE ：1天
	// THREE :3天
	// SEVEN ：7天
	// FOURTEEN ：14天
	// THIRTY ：30天
	// SIXTY ：60天
	// NINETY ：90天
	// ONE_HUNDRED_EIGHTY：180天
	// PERMANENT ：永久
	ConvertWindow enum.ConvertWindow `json:"convert_window,omitempty"`
}

// Encode implement PostRequest interface
func (r ConvertWindowUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
