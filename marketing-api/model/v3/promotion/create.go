package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建广告 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// Name 广告名称，长度是1-50个字（两个英文字符占1个字）。名称不可重复
	Name string `json:"name,omitempty"`
	// Operation 广告状态， 允许值: ENABLE开启(默认值）、DISABLE关闭
	Operation enum.OptStatus `json:"operation,omitempty"`
	// MaterialsType  素材类型，直播场景必填，
	// 允许值：LIVE_MATERIALS直播素材;PROMOTION_MATERIALS广告素材;
	// 当auth_type为【抖音号授权】时，广告素材和直播素材两选一；
	// 当auth_type为【直播授权】时，仅支持直播素材；
	// landing_type=LINK&&marketing_goal=LIVE&&materials_type=LIVE_MATERIALS时，promotion_materials仅支持橙子落地页；
	// landing_type=MICRO_GAME&&marketing_goal=LIVE&&materials_type=LIVE_MATERIALS，promotion_materials支持橙子落地页和字节小程序信息；
	// landing_type=APP&&marketing_goal=LIVE&&materials_type=LIVE_MATERIALS时，promotion_materials需传空；
	MaterialsType enum.MaterialsType `json:"materials_type,omitempty"`
	// PromotionRelatedProduct UBP多品广告素材组合
	// 若创建项目为UBP多品项目，广告层级支持设置该参数结构
	// 非必传，不传则对于UBP多品项目系统根据项目层级关联产品自动生成素材
	// 对于UBP在广告层级设置的商品，其关联素材的个数限制：
	// 多个商品总共设置视频素材<=60个
	// 多个商品总共设置图片素材<=50个
	// 多个商品总共设置标题素材<=30个
	// 多个商品总共设置落地页素材<=30个
	PromotionRelatedProduct []RelatedProduct `json:"promotion_related_product,omitempty"`
	// PromotionMaterials 广告素材组合
	PromotionMaterials *PromotionMaterial `json:"promotion_materials,omitempty"`
	// NativeSetting 原生广告设置
	NativeSetting *NativeSetting `json:"native_setting,omitempty"`
	// IsCommentDisable 广告评论，ON为启用，OFF为不启用
	IsCommentDisable enum.OnOff `json:"is_comment_disable,omitempty"`
	// AdDownloadStatus 客户端下载视频功能，ON为启用，OFF为不启用
	AdDownloadStatus enum.OnOff `json:"ad_download_status,omitempty"`
	// Source 广告来源，字数限制：[1-10]
	// 当landing_type = LINK、MICRO_GAME时必填
	Source string `json:"source,omitempty"`
	// BrandInfo 品牌信息
	BrandInfo *BrandInfo `json:"brand_info,omitempty"`
	// BudgetMode 预算类型(创建后不可修改), 允许值: BUDGET_MODE_DAY日预算（默认值）, BUDGET_MODE_TOTAL总预算
	// 适用场景：
	// delivery_mode=MANUAL
	// 项目未开启预算择优分配，即budget_optimize_switch=OFF
	// 项目投放时间类型schedule_type=SCHEDULE_START_END时，支持设置该时间段内的BUDGET_TOTAL总预算
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 预算(出价方式为OCPM时，不少于300元；24小时内修改预算操作，不能超过20次，24小时是指自然天的24小时；单次修改预算幅度不能低于100元（增加或者减少）；修改后预算金额，不能低于当前已消费金额的105%，以整百单位向上取整；取值范围: ≥ 0
	Budget float64 `json:"budget,omitempty"`
	// Bid 点击出价/展示出价，当pricing为CPC、CPM出价方式时必填，
	// 当pricing=CPM时，取值范围：4-100；
	// 当pricing=CPC时，取值范围：0.1-100
	// landing_type=link&&external_action=表单提交，deep_external_action=电话接通时，出价需传空;
	// 当项目成本稳投开启时，不支持该字段
	Bid float64 `json:"bid,omitempty"`
	// CpaBid 目标转化出价/预期成本， 取值范围：0.1-10000元
	// 注意：当 bid_type = NO_BID时，不填写该字段，否则会报错
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// DeepCpaBid 深度优化出价，deep_bid_type = DEEP_BID_MIN 时必填。取值范围：0.1~10000元
	// 注意：当 bid_type = NO_BID时，不填写该字段，否则会报错
	DeepCpaBid float64 `json:"deep_cpabid,omitempty"`
	// RoiGoal 深度转化ROI系数, 范围(0,5]，精度：保留小数点后四位，当 deep_bid_type = ROI_COEFFICIENT 时必填
	// 注意：当 bid_type=NO_BID时，不填写该字段，否则会报错
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// AutoExtendTraffic 智能拓流
	// 允许值：ON开启（默认值）； OFF关闭
	AutoExtendTraffic enum.OnOff `json:"auto_extend_traffic,omitempty"`
	// Keywords 关键词列表，关键词和智能拓流二者必须开启一个，一个广告最多可添加1000个
	Keywords []project.Keyword `json:"keywords,omitempty"`
	// CreativeAutoGenerateSwitch 是否开启自动生成素材
	// 默认值：OFF
	// 枚举值：ON开启、OFF不开启
	CreativeAutoGenerateSwitch enum.OnOff `json:"creative_auto_generate_switch,omitempty"`
	// ConfigID 配置ID，开关打开，不传为黑盒明投派生
	ConfigID uint64 `json:"config_id,omitempty"`
	// 7d_retention 表示7日留存天数，单位：天，取值范围[0.01，7.00]，仅支持最多2位小数。
	// 7d_retention适用创编场景，该场景下有效且必填
	// landing_type = APP 应用推广
	// ad_type = ALL 通投
	// delivery_mode = MANUAL  手动投放
	// external_action = AD_CONVERT_TYPE_ACTIVE 优化目标=激活
	// deep_external_action = AD_CONVERT_TYPE_RETENTION_DAYS深度优化目标 = 留存天数
	// delivery_setting.deep_bid_type = AD_CONVERT_TYPE_RETENTION_DAYS深度优化方式 = 留存天数
	// delivery_range.inventory_catalog = MANUAL  广告位大类 = 首选媒体
	// inventory_type = INVENTORY_UNION_SLOT  投放位置 只选择穿山甲
	SevenDRetention float64 `json:"7d_retention,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建广告 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// PromotionID 广告ID
		PromotionID uint64 `json:"promotion_id,omitempty"`
	} `json:"data,omitempty"`
}
