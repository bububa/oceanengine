package enum

type AccountRole string

const (
	ADVERTISER        AccountRole = "ADVERTISER"        // 广告主
	CUSTOMER_ADMIN    AccountRole = "CUSTOMER_ADMIN"    // 管家-管理员
	AGENT             AccountRole = "AGENT"             // 代理商
	CHILD_AGENT       AccountRole = "CHILD_AGENT"       // 二级代理商
	CUSTOMER_OPERATOR AccountRole = "CUSTOMER_OPERATOR" // 管家-操作者
	// PLATFORM_ROLE_QIANCHUAN_AGENT 代理商账户
	PLATFORM_ROLE_QIANCHUAN_AGENT AccountRole = "PLATFORM_ROLE_QIANCHUAN_AGENT"
	// PLATFORM_ROLE_SHOP_ACCOUNT 店铺账户
	PLATFORM_ROLE_SHOP_ACCOUNT AccountRole = "PLATFORM_ROLE_SHOP_ACCOUNT"
)
