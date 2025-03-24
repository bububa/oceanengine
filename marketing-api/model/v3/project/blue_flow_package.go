package project

// BlueFlowPackage 搜索蓝海流量投放相关参数
type BlueFlowPackage struct {
	// BlueFlowPackageSetting 蓝海流量设置允许值：ON 启用、OFF不启用
	//
	// 仅当ad_type = SEARCH、delivery_type= NORMAL、landing_type=APP 应用推广、LINK 销售线索推广、MICRO_GAME 小程序、SHOP 电商店铺推广、NATIVE_ACTION 原生互动时生效，创建后不可修改
	// 如果不需要蓝海流量投放，无需传入此参数
	// 是否开启蓝海投放：
	// 搜索&手动模式投放通过blue_flow_package_setting = ON/OFF
	// 进行设置开启/关闭蓝海投放
	// 自动链路通过设置blue_flow_keyword_name，进行蓝海投放，若不设置blue_flow_keyword_name即不进行蓝海投放
	//
	//
	// 白名单说明
	//
	// 【注意】仅支持白名单广告账户传入蓝海流量包，您可通过「查询白名单能力」接口查询广告主是否命中白名单
	// 【注意】当前蓝海流量投放为白名单功能，如当前广告主账号没有权限，传参会报错
	BlueFlowPackageSetting string `json:"blue_flow_package_setting,omitempty"`
	// BlueFlowPackageID 蓝海流量包ID，当启用蓝海流量投放blue_flow_package_setting = ON时必填，可通过【查询蓝海流量包】查询，创建后不可修改
	// 注：仅支持搜索&手动模式设置，搜索极速智投设置不生效
	BlueFlowPackageID uint64 `json:"blue_flow_package_id,omitempty"`
	// BlueFlowKeywordName 蓝海关键词，当项目设置ad_type = SEARCH && delivery_type = DURATION时，支持设置蓝海关键词，蓝海关键词不支持在更新接口中删除
	// 白名单说明
	//
	// 【注意】仅支持白名单广告账户传入蓝海流量包，您可通过「查询白名单能力」接口查询广告主是否命中白名单
	// 【注意】当前蓝海流量投放为白名单功能，如当前广告主账号没有权限 或 账号下没有推送蓝海关键词，传参会报错"不支持设置蓝海关键词"
	BlueFlowKeywordName []string `json:"blue_flow_keyword_name,omitempty"`
}
