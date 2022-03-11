package customaudience

import "github.com/bububa/oceanengine/marketing-api/enum"

// CustomAudience 人群包
type CustomAudience struct {
	// ID 人群包ID
	ID uint64 `json:"custom_audience_id,omitempty"`
	// Isdel 人群包是否删除，枚举值："1"：已删除，"0"：未删除
	Isdel int `json:"isdel,omitempty"`
	// DataSourceID 数据源ID
	DataSourceID string `json:"data_source_id,omitempty"`
	// Name 人群包名称
	Name string `json:"name,omitempty"`
	// Source 人群包来源，详见【附录-DMP相关-人群包来源】
	Source enum.DmpSource `json:"source,omitempty"`
	// Status 人群包状态，详见【附录-DMP相关-人群包状态】
	Status int `json:"status,omitempty"`
	// DeliveryStatus 人群包可投放状态
	DeliveryStatus enum.CustomAudienceDeliveryStatus `json:"delivery_status,omitempty"`
	// CoverNum 人群包覆盖人群数目，基于"upload_num"：上传数据源包含的人群数目与uid对应后，再与头条系产品MAU交集后的数量（存在一个设备号/手机号对应多个uid的情况）
	// 实际数量可能多于/少于"upload_num"：上传数据源包含的人群数目
	CoverNum int64 `json:"cover_num,omitempty"`
	// UploadNum 上传数据源包含的人群数目
	UploadNum int64 `json:"upload_num,omitempty"`
	// Tag 人群包标签，通过数据源创建的人群包，标签会默认为"API文件数据源"
	Tag string `json:"tag,omitempty"`
	// PushStatus 推送状态，详见【附录-DMP相关-DMP推送状态】
	PushStatus enum.DmpPushStatus `json:"push_status,omitempty"`
	// CreateTime 人群包创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ModifyTime 人群包修改时间
	ModifyTime string `json:"modify_time,omitempty"`
	// ThirdPartyInfo 是否为三方包，均返回“非三方包”
	ThirdPartyInfo string `json:"third_party_info,omitempty"`
	// ExistPullOffTag 人群包是否包含下线标签，0：不包含，1：包含
	ExistPullOffTag int `json:"exist_pull_off_tag,omitempty"`
}
