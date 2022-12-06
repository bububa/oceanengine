package datasource

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ReadRequest 数据源详细信息 API Request
type ReadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DataSourceIDList 数据源ID列表
	// 一次最多传400个数据源id
	DataSourceIDList []string `json:"data_source_id_list,omitempty"`
}

// Encode implement GetRequest interface
func (r ReadRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.DataSourceIDList != nil {
		idList, _ := json.Marshal(r.DataSourceIDList)
		values.Set("data_source_id_list", string(idList))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ReadResponse 数据源详细信息 API Response
type ReadResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *ReadResponseData `json:"data,omitempty"`
}

// ReadResponseData json返回值
type ReadResponseData struct {
	// DataList 数据信息列表
	DataList []DataSource `json:"data_list,omitempty"`
}

// DataSource 数据源信息
type DataSource struct {
	// Name 数据源名称
	Name string `json:"name,omitempty"`
	// ID 数据源ID
	ID string `json:"data_source_id,omitempty"`
	// Description 数据源描述
	Description string `json:"description,omitempty"`
	// Status 数据源初次创建状态，枚举值："0"：已创建，"1"：解析中暂不可用，"2"：完成
	Status enum.DataSourceStatus `json:"status,omitempty"`
	// CoverNum 人群包覆盖人群数目，基于"upload_num"：上传数据源包含的人群数目与uid对应后，再与头条系产品MAU交集后的数量（存在一个设备号/手机号对应多个uid的情况）
	// 实际数量可能多于/少于"upload_num"：上传数据源包含的人群数目
	CoverNum int64 `json:"cover_num,omitempty"`
	// UploadNum 上传数据源包含的人群数目
	UploadNum int64 `json:"upload_num,omitempty"`
	// CreateTime 数据源创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ModifyTime 数据源修改时间
	ModifyTime string `json:"modify_time,omitempty"`
	// LatestPublishedChangeLogID 数据源最近一次更新对应的日志id，若未完成更新，返回“-1
	LatestPublishedChangeLogID int64 `json:"latest_published_changelog_id,omitempty"`
	// LatestPublishedTime 数据源最近一次发布完成时间
	LatestPublishedTime string `json:"latest_published_time,omitempty"`
	// DataSourceType 数据源类型，枚举值："UID"：用户ID，"DID" ：设备ID
	DataSourceType string `json:"data_source_type,omitempty"`
	// ChangeLogs 每一次更新，会记录一次日志id
	ChangeLogs []ChangeLog `json:"change_logs,omitempty"`
	// DefaultAudience 数据源对应的人群包，如果还未生成将不会返回
	DefaultAudience *Audience `json:"default_audience,omitempty"`
}

// ChangeLog 每一次更新，会记录一次日志id
type ChangeLog struct {
	// Action 本次更新进行的操作，枚举值："0"：新建，"1"：添加，"2"：删除，"3"：重置
	Action enum.DmpDatasourceOperationType `json:"action,omitempty"`
	// FilePaths 使用数据源文件上传文件后返回的文件路径参数
	FilePaths []string `json:"file_paths,omitempty"`
	// Status 本次更新的状态，枚举值："0"：新建，"1"：处理中，"2"：生效，"3"：失败
	Status enum.DataSourceStatus `json:"status,omitempty"`
	// CreateTime 本次更新创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ModifyTime 本次更新完成时间
	ModifyTime string `json:"modify_time,omitempty"`
	// ID 更新日志id
	ID int64 `json:"change_log_id,omitempty"`
	// PublishTime 数据源本次更新发布完成时间记录
	PublishTime string `json:"publish_time,omitempty"`
}

// Audience 数据源对应的人群包，如果还未生成将不会返回
type Audience struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CustomAudienceID 人群包id
	CustomAudienceID uint64 `json:"custom_audience_id,omitempty"`
	// Name 人群包名称
	Name string `json:"name,omitempty"`
	// CustomType 人群类型，枚举值见【DMP相关-人群类型】
	CustomType enum.DmpCustomType `json:"custom_type,omitempty"`
	// Source 来源，枚举值见【DMP相关-人群包来源】
	Source int `json:"source,omitempty"`
	// Status 人群包状态，枚举值见【DMP相关-人群包状态】
	Status int `json:"status,omitempty"`
	// PushStatus 推送状态，枚举值见【DMP相关-DMP推送状态】
	PushStatus enum.DmpPushStatus `json:"push_status,omitempty"`
	// UploadNum 上传数据源包含的人群数目
	UploadNum int64 `json:"upload_num,omitempty"`
	// CoverNum 人群包覆盖人群数目，基于"upload_num"：上传数据源包含的人群数目与uid对应后，再与头条系产品MAU交集后的数量（存在一个设备号/手机号对应多个uid的情况）
	// 实际数量可能多于/少于"upload_num"：上传数据源包含的人群数目
	CoverNum int64 `json:"cover_num,omitempty"`
	// ExpireDate 人群包过期时间
	ExpireDate string `json:"expire_date,omitempty"`
	// CreateTime 人群包创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ModifyTime 人群包修改时间
	ModifyTime string `json:"modify_time,omitempty"`
	// Isdel 删除标志，枚举值："1"：已删除，"0"：未删除
	IsDel int `json:"isdel,omitempty"`
	// CalculateSubType 计算子类型
	CalculateSubType int `json:"calculate_sub_type,omitempty"`
	// CalculateType 计算类型
	CalculateType int `json:"calculate_type,omitempty"`
	// DataSourceID 数据源id
	DataSourceID string `json:"data_source_id,omitempty"`
	// Tag 人群包标签，通过数据源创建的人群包，标签会默认为“API文件数据源”
	Tag string `json:"tag,omitempty"`
	// ThirdPartyInfo 三方信息，均返回“非三方包”
	ThirdPartyInfo string `json:"third_party_info,omitempty"`
	// DeliveryStatus 人群包可投放状态，只有当状态为CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE时才可进行投放使用
	// 可选值：
	// CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE：可投放，人群包发布完成且推送完成
	// CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH：不可投放，人群包发布完成，但未推送，需要进行推送后再使用
	// CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH：不可投放，群包未发布但已推送，需要进行发布后再使用
	//CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE：不可投放，未发布完成且未推送
	DeliveryStatus enum.CustomAudienceDeliveryStatus `json:"delivery_status,omitempty"`
}
