package aweme

import "github.com/bububa/oceanengine/marketing-api/enum"

// Aweme 抖音号
type Aweme struct {
	// ID 抖音id，用于创建计划，拉取抖音号视频素材时入参
	ID uint64 `json:"aweme_id,omitempty"`
	// Avatar 抖音头像
	Avatar string `json:"aweme_avatar,omitempty"`
	// ShowID 抖音号，即客户在手机端上看到的抖音号，若向客户披露抖音号请使用该字段
	ShowID string `json:"aweme_show_id,omitempty"`
	// Name 抖音号名称
	Name string `json:"aweme_name,omitempty"`
	// Status 抖音号带货状态，返回值：NORMAL 可以正常投放;ANCHOR_FORBID 带货口碑分过低，暂时无法创建计划;ANCHOR_REACH_UPPER_LIMIT_TODAY 带货分过低或暂无带货分，可以创建计划，但无法产生消耗，带货分恢复正常后可正常消耗
	Status enum.AwemeStatus `json:"aweme_status,omitempty"`
	// BindType 抖音号关系类型
	BindType []enum.AwemeBindType `json:"bind_type,omitempty"`
	// HasRoi2DeliveryLimit 该抖音号是否有全域推广计划投放，如果开启了投放，那么该抖音号不可以用来投放其他直播带货广告
	// 允许值：
	// true：存在全域推广计划投放
	// false：暂无全域推广计划投放
	HasRoi2DeliveryLimit bool `json:"has_roi2_delivery_limit,omitempty"`
	// HasRoi2GroupCreate 当前抖音号是否创建过全域推广计划
	// false：未创建过计划, 可以新建
	// true：已经创建过计划，不支持新建
	HasRoi2GroupCreate bool `json:"has_roi2_group_create,omitempty"`
	// HasShopPermission 是否有推商品全域推广权限
	HasShopPermission bool `json:"has_shop_permission,omitempty"`
	// HasLivePermission 是否有直播全域推广权限
	HasLivePermission bool `json:"has_live_permission,omitempty"`
	// CanControlUniprom 当前账户是该抖音号的官方巨量千川账户,对其他账户解除授权即可正常使用直播全域推广,是否需要解除授权
	CanControlUniprom bool `json:"can_control_uniprom,omitempty"`
	// CanApplyUniprom 抖音号是否可申请
	CanApplyUniprom bool `json:"can_apply_uniprom,omitempty"`
	// AnchorForbidden 抖音号是否禁用
	AnchorForbidden bool `json:"anchor_forbidden,omitempty"`
}
