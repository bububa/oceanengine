package microgame

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ConvertWindowGetRequest 查询字节小游戏归因激活时间窗 API Request
type ConvertWindowGetRequest struct {
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
}

// Encode implements GetRequest interface
func (r ConvertWindowGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	values.Set("asset_id", r.AssetID)
	values.Set("asset_type", string(r.AssetType))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ConvertWindowGetResponse 查询字节小游戏归因激活时间窗 API Response
type ConvertWindowGetResponse struct {
	model.BaseResponse
	Data struct {
		// ConvertWindow 时间窗口设置：（枚举及含义同编辑接口）
		// ONE ：1天
		//  THREE ：3天
		//  SEVEN ：7天
		//  FOURTEEN ：14天
		//  THIRTY ：30天
		//  SIXTY ：60天
		//  NINETY ：90天
		//  ONE_HUNDRED_EIGHTY：180天
		//  PERMANENT ：永久
		ConvertWindow enum.ConvertWindow `json:"convert_window,omitempty"`
	} `json:"data,omitempty"`
}
