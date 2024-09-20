package sitetemplate

// LinkType 链接类型
type LinkType string

const (
	// LinkType_URL 链接地址
	LinkType_URL LinkType = "URL"
	// LinkType_SCHEME scheme地址
	LinkType_SCHEME LinkType = "SCHEME"
	// LinkType_QUICK_APP 快应用地址
	LinkType_QUICK_APP LinkType = "QUICK_APP"
)

// Link 跳转链接信息
type Link struct {
	// AssetID 快应用资产id，当link_type为QUICK_APP时，有返回值
	AssetID string `json:"asset_id,omitempty"`
	// URL 链接地址，当link_type为url时，必填
	URL string `json:"url,omitempty"`
	// Scheme scheme地址，当link_type为scheme时，必填
	Scheme string `json:"scheme,omitempty"`
	// QuickApp 快应用地址，当link_type为quick_app时，必填;获取快应用链接可参考【快应用官方网站】
	QuickApp string `json:"quick_app,omitempty"`
	// LinkType 链接类型，可传值有:url：链接地址,sheme：scheme地址,quickApp：快应用地址,当传递链接信息参数时，必填
	LinkType LinkType `json:"link_type,omitempty"`
	// Description 应用描述，为了展示效果，推荐12个中文字符长度
	Description string `json:"description,omitempty"`
}
