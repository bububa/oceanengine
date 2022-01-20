package star

import "github.com/bububa/oceanengine/marketing-api/enum"

// Clue 星图订单投后线索
type Clue struct {
	// StarID 星图id
	StarID uint64 `json:"star_id,omitempty"`
	// DemandID 任务id
	DemandID uint64 `json:"demand_id,omitempty"`
	// OrderID 订单id
	OrderID uint64 `json:"order_id,omitempty"`
	// ItemID 视频id
	ItemID uint64 `json:"item_id,omitempty"`
	// AuthorUid 作者抖音uid
	AuthorUid uint64 `json:"author_uid,omitempty"`
	// AuthorName 作者昵称
	AuthorName string `json:"author_name,omitempty"`
	// ComponentType 组件类型
	ComponentType enum.StarComponentType `json:"component_type,omitempty"`
	// AnchorName 锚点名称
	AnchorName string `json:"anchor_name,omitempty"`
	// Phone 手机号码
	Phone string `json:"phone,omitempty"`
	// CarSeries 车系
	CarSeries string `json:"car_series,omitempty"`
	// CarType 车型
	CarType string `json:"car_type,omitempty"`
	// Address 地址
	Address string `json:"address,omitempty"`
	// Name 姓名
	Name string `json:"name,omitempty"`
	// Province 省份
	Province string `json:"province,omitempty"`
	// City 城市
	City string `json:"city,omitempty"`
	// CreateTime 创建时间，时间格式“%Y-%m-%d %H:%M:%S”
	CreateTime string `json:"create_time,omitempty"`
}
