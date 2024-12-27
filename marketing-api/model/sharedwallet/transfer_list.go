package sharedwallet

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TransferListRequest 资金共享-查询转账列表 API Request
type TransferListRequest struct {
	// AccountID 鉴权账户
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 鉴权账户类型 可选值:
	// AD 巨量广告主
	// AGENT 代理
	// LOCAL 本地推广告主
	AccountType string `json:"account_type,omitempty"`
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// QueryBeginTime 查询开始时间(转账创建时间) yyyy-MM-dd HH:mm:ss
	QueryBeginTime string `json:"query_begin_time,omitempty"`
	// QueryEndTime 查询结束时间(转账创建时间) yyyy-MM-dd HH:mm:ss
	QueryEndTime string `json:"query_end_time,omitempty"`
	// QueryWalletIDList 需要查询的大小钱包id，加款方与减款方都需要包含在内
	QueryWalletIDList []uint64 `json:"query_wallet_id_list,omitempty"`
	// PayeeID 加款方id，可选，需要包含在query_wallet_id_list
	PayeeID uint64 `json:"payee_id,omitempty"`
	// RemitterID 减款方id，可选，需要包含在query_wallet_id_list
	RemitterID uint64 `json:"remitter_id,omitempty"`
	// PageInfo 分页信息
	PageInfo struct {
		// PageNum 页码，从1开始，仅支持page_size * page_num < 1w的深度分页
		PageNum int `json:"page_num,omitempty"`
		// PageSize 分页大小，范围：1-100
		PageSize int `json:"page_size,omitempty"`
	} `json:"page_info,omitempty"`
}

// Encode implements GetRequest interface
func (r TransferListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", r.AccountType)
	values.Set("biz_request_no", r.BizRequestNo)
	values.Set("query_begin_time", r.QueryBeginTime)
	values.Set("query_end_time", r.QueryEndTime)
	values.Set("query_wallet_id_list", string(util.JSONMarshal(r.QueryWalletIDList)))
	if r.PayeeID > 0 {
		values.Set("payee_id", strconv.FormatUint(r.PayeeID, 10))
	}
	if r.RemitterID > 0 {
		values.Set("remitter_id", strconv.FormatUint(r.RemitterID, 10))
	}
	values.Set("page_info", string(util.JSONMarshal(r.PageInfo)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// TransferListResponse 资金共享-查询转账列表 API Response
type TransferListResponse struct {
	model.BaseResponse
	Data struct {
		// RecordList 转账记录列表
		RecordList []TransferRecord `json:"record_list,omitempty"`
	} `json:"data,omitempty"`
}

// TransferRecord 转账记录
type TransferRecord struct {
	// TransferOrderSerial 转账单编号
	TransferOrderSerial string `json:"transfer_order_serial,omitempty"`
	// BizRequestNo 转账申请幂等id
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// RemitterID 减款方id
	RemitterID uint64 `json:"remitter_id,omitempty"`
	// RemitterCategory 减款方类别 可选值:
	// ACCOUNT 账户
	// WALLET 钱包
	RemitterCategory string `json:"remitter_category,omitempty"`
	// PayeeID 加款方id
	PayeeID uint64 `json:"payee_id,omitempty"`
	// PayeeCategory 加款方类别 可选值:
	// ACCOUNT 账户
	// WALLET 钱包
	PayeeCategory string `json:"payee_category,omitempty"`
	// TransferTargetAmount 加款方-减款方间转账金额(单位：分)
	TransferTargetAmount int64 `json:"transfer_target_amount,omitempty"`
	// TransferTargetStatus 加款方-减款方间转账状态 可选值:
	// NO_TRANSFER 未转账
	// TRANSFER_FAILED 转账失败(终态)
	// TRANSFER_ING 转账中
	// TRANSFER_PART 转账部分成功(终态)
	// TRANSFER_SUCCESS 转账成功(终态)
	TransferTargetStatus enum.TransferStatus `json:"transfer_target_status,omitempty"`
	// TransferCapitalRecordList 转账资金列表
	TransferCapitalRecordList []TransferCapitalRecord `json:"transfer_capital_record_list,omitempty"`
}

// TransferCapitalRecord 转账资金
type TransferCapitalRecord struct {
	// CapitalType 转账资金类型 可选值:
	// CREDIT_BIDDING 授信竞价
	// CREDIT_BRAND 授信品牌
	// CREDIT_GENERAL 授信通用
	// PREPAY_BIDDING 预付竞价
	// PREPAY_BRAND 预付品牌
	// PREPAY_GENERAL 预付通用
	CapitalType enum.CapitalType `json:"capital_type,omitempty"`
	// TransferAmount 转账资金金额(单位：分)
	TransferAmount int64 `json:"transfer_amount,omitempty"`
	// TransferStatus 转账资金状态 可选值:
	// NO_TRANSFER 未转账
	// TRANSFER_FAILED 转账失败(终态)
	// TRANSFER_ING 转账中
	// TRANSFER_PART 转账部分成功(终态)
	// TRANSFER_SUCCESS 转账成功(终态)
	TransferStatus enum.TransferStatus `json:"transfer_status,omitempty"`
	// FailReason 失败原因
	FailReason string `json:"fail_reason,omitempty"`
	// Platform 转账业务线 可选值:
	// AD 广告
	// AD_ALL 广告通用
	// BENDITUI 本地推
	// QIANCHUAN 千川
	// STAR 星图
	Platform string `json:"platform,omitempty"`
	// TransferTargetFinishTime 完成时间yyyy-MM-dd HH:mm:ss
	TransferTargetFinishTime string `json:"transfer_target_finish_time,omitempty"`
	// TransferTargetCreateTime 创建时间yyyy-MM-dd HH:mm:ss
	TransferTargetCreateTime string `json:"transfer_target_create_time,omitempty"`
	// Remark 转账备注
	Remark string `json:"remark,omitempty"`
}
