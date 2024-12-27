package enum

// TransactionBizType 流水类型(充值、退款等)
type TransactionBizType string

const (
	// TransactionBizType_ADJUST_DECREASE 调差减款
	TransactionBizType_ADJUST_DECREASE TransactionBizType = "ADJUST_DECREASE"
	// TransactionBizType_ADJUST_FREEZE 调差冻结/解冻
	TransactionBizType_ADJUST_FREEZE TransactionBizType = "ADJUST_FREEZE"
	// TransactionBizType_ADJUST_INCREASE 调差加款
	TransactionBizType_ADJUST_INCREASE TransactionBizType = "ADJUST_INCREASE"
	// TransactionBizType_CREDIT_RECHARGE 信控充值
	TransactionBizType_CREDIT_RECHARGE TransactionBizType = "CREDIT_RECHARGE"
	// TransactionBizType_CREDIT_REFUND 信控减款
	TransactionBizType_CREDIT_REFUND TransactionBizType = "CREDIT_REFUND"
	// TransactionBizType_INIT 钱包初始化
	TransactionBizType_INIT TransactionBizType = "INIT"
	// TransactionBizType_ORDER_PAY 订单支付
	TransactionBizType_ORDER_PAY TransactionBizType = "ORDER_PAY"
	// TransactionBizType_ORDER_WITHDRAW 订单还款
	TransactionBizType_ORDER_WITHDRAW TransactionBizType = "ORDER_WITHDRAW"
	// TransactionBizType_REFUND_FREEZE 冻结冻结/解冻
	TransactionBizType_REFUND_FREEZE TransactionBizType = "REFUND_FREEZE"
	// TransactionBizType_SHARED_CANCEL_RECHARGE 共享钱包充值撤销
	TransactionBizType_SHARED_CANCEL_RECHARGE TransactionBizType = "SHARED_CANCEL_RECHARGE"
	// TransactionBizType_SHARED_FROZEN_REFUND 共享钱包退款
	TransactionBizType_SHARED_FROZEN_REFUND TransactionBizType = "SHARED_FROZEN_REFUND"
	// TransactionBizType_SHARED_RECHARGE 共享钱包充值
	TransactionBizType_SHARED_RECHARGE TransactionBizType = "SHARED_RECHARGE"
	// TransactionBizType_TRANSFER 转账
	TransactionBizType_TRANSFER TransactionBizType = "TRANSFER"
	// TransactionBizType_TRANSFER_IN 转入
	TransactionBizType_TRANSFER_IN TransactionBizType = "TRANSFER_IN"
	// TransactionBizType_TRANSFER_OUT 转出
	TransactionBizType_TRANSFER_OUT TransactionBizType = "TRANSFER_OUT"
)
