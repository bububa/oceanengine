package customercenter

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TransferAccountType 广告账户类型
type TransferAccountType string

const (
	// TransferAccountType_PAYEE 可转入账户（我可向谁转账）
	TransferAccountType_PAYEE = "PAYEE"
	// TransferAccountType_REMITTER 可转出账户（谁可向我转账）
	TransferAccountType_REMITTER = "REMITTER"
)

// AdvertiserTransferableListRequest 获取可转账户列表（客户中心&广告主） API Request
type AdvertiserTransferableListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TransferAccountType 广告账户类型
	TransferAccountType TransferAccountType `json:"transfer_account_type,omitempty"`
	// Page 页码. 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量. 默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r AdvertiserTransferableListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.TransferAccountType != "" {
		values.Set("transfer_account_type", string(r.TransferAccountType))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AdvertiserTransferableListResponse 获取可转账户列表（客户中心&广告主）API Response
type AdvertiserTransferableListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdvertiserTransferableListData `json:"data,omitempty"`
}

// AdvertiserTransferableListData json返回值
type AdvertiserTransferableListData struct {
	// List 账户列表
	List []Advertiser `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
