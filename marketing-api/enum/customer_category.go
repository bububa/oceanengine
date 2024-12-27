package enum

// CustomerCategory 客户类型
type CustomerCategory string

const (
	// CustomerCategory_AGENT 代理商
	CustomerCategory_AGENT CustomerCategory = "AGENT"
	// CustomerCategory_DIRECT 直客
	CustomerCategory_DIRECT CustomerCategory = "DIRECT"
	// CustomerCategory_PARTNER 合作方
	CustomerCategory_PARTNER CustomerCategory = "PARTNER"
	// CustomerCategory_SELF_SERVICE 自助客户
	CustomerCategory_SELF_SERVICE CustomerCategory = "SELF_SERVICE"
	// CustomerCategory_SVIP_CUSTOMER VIP大型客户
	CustomerCategory_SVIP_CUSTOMER CustomerCategory = "SVIP_CUSTOMER"
	// CustomerCategory_VIRTUAL_CUS 虚客
	CustomerCategory_VIRTUAL_CUS CustomerCategory = "VIRTUAL_CUS"
)
