package enum

// SiteTYpe 建站类型
type SiteType string

const (
	// SiteType_DEFAULT 默认类型
	SiteType_DEFAULT SiteType = "DEFAULT"
	// SiteType_CREATIVEFORM 附加创意表单
	SiteType_CREATIVEFORM SiteType = "CREATIVEFORM"
	// SiteType_CREATIVEFORMOLD 附加创意表单
	SiteType_CREATIVEFORMOLD SiteType = "CREATIVEFORMOLD"
	// SiteType_SITE_REJECTED 原生表单
	SiteType_SITE_REJECTED SiteType = "SITE_REJECTED"
	// SiteType_STORESITE 门店站点
	SiteType_STORESITE SiteType = "STORESITE"
	// SiteType_FORM 标准化落地页 普通表单
	SiteType_FORM SiteType = "FORM"
	// SiteType_SLIDE h5 翻页
	SiteType_SLIDE SiteType = "SLIDE"
	// SiteType_NOPUBLISHSITE 不可发布站点（比如DPA）
	SiteType_NOPUBLISHSITE SiteType = "NOPUBLISHSITE"
	// SiteType_SITE_TRASH 普通建站 使用应用下载模板
	SiteType_SITE_TRASH SiteType = "SITE_TRASH"
	// SiteType_PROGRAM 程序化落地页(存在于 V2 版建站)
	SiteType_PROGRAM SiteType = "PROGRAM"
	// SiteType_SELF 自助站点
	SiteType_SELF SiteType = "SELF"
	// SiteType_STREAMING 流式站点
	SiteType_STREAMING SiteType = "STREAMING"
	// SiteType_POLL 投票站点
	SiteType_POLL SiteType = "POLL"
	// SiteType_SUBCHAIN 搜索子链
	SiteType_SUBCHAIN SiteType = "SUBCHAIN"
	// SiteType_MINIAPP 小程序站点
	SiteType_MINIAPP SiteType = "MINIAPP"
	// SiteType_DPA 动态商品站点
	SiteType_DPA SiteType = "DPA"
	// SiteType_NATIVE 原生落地页
	SiteType_NATIVE SiteType = "NATIVE"
	// SiteType_AUTOGENNOVEL 小说自动生成
	SiteType_AUTOGENNOVEL SiteType = "AUTOGENNOVEL"
	// SiteType_STOREORANGESITE 门店橙子建站
	SiteType_STOREORANGESITE SiteType = "STOREORANGESITE"
	// SiteType_DEPOSITE 托管落地页
	SiteType_DEPOSITE SiteType = "DEPOSITE"
)
