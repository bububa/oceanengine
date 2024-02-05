package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialAttributesListRequest 获取视频素材评估标签（新版） API Request
type MaterialAttributesListRequest struct {
	// AccountID 各平台账户id，必须同时选择account_type类型
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型， 可选值:
	// BP 巨量引擎工作台（纵横）
	// AGENT 巨量方舟
	// AD 巨量广告
	// QIANCHUAN 巨量千川
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// Filtering 过滤条件
	// 不同条件之间为「且」的关系，只有同时满足传入的多个条件，才会返回信息
	Filtering *MaterialAttributesListFilter `json:"filtering,omitempty"`
	// ReturnLowqualitySuggestions 应答参数是否返回低质原因，不传或传入false时，应答参数不返回低质原因
	// 当素材为低质素材时，可查询素材低质原因，允许值：
	//  返回：true
	//  不返回：false
	ReturnLowqualitySuggestions bool `json:"return_lowquality_suggestions,omitempty"`
	// Page 页码，默认值: 1，page*page_size 最大1000
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量，默认值: 100，page*page_size 最大1000
	PageSize int `json:"page_size,omitempty"`
}

type MaterialAttributesListFilter struct {
	// MaterialIDs 按素材ID过滤，范围为1-1000
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// MaterialProperties 素材标签过滤项，如果不传，则默认返回广告主ID下所有素材。
	// 传入多个允许值，该参数多个值之间为「或」的关系，只要素材存在某一种类型的标签，都会在应答参数中返回。
	// 允许值：
	// FIRST_PUBLISH_MATERIAL 首发素材
	// AD_HIGH_QUALITY_MATERIAL  AD优质素材
	// ECP_HIGH_QUALITY_MATERIAL 千川优质素材
	// AD_LOW_QUALITY_MATERIAL AD低质素材
	// ECP_LOW_QUALITY_MATERIAL 千川低质素材
	// INEFFICIENT_MATERIAL 低效素材
	// SIMILAR_MATERIAL 同质化挤压素材
	// SIMILAR_QUEUE_MATERIAL 同质化排队素材
	// CARRY_MATERIAL存在搬运打压风险
	MaterialProperties []enum.MaterialProperty `json:"material_properties,omitempty"`
	// StartTime 起始日期，表示按照素材上传到账户下的时间过滤，格式为yyyy-mm-dd HH:MM:SS，最早允许传入时间：2022-01-01 00:00:00
	// （时间是账号和素材绑定的时间，删除重绑时间会更新）
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束日期，表示过滤出一段时间内上传的素材，格式为yyyy-mm-dd HH:MM:SS，如果传入起始日期，未传此参数，则默认为当前时间
	EndTime string `json:"end_time,omitempty"`
	// AttributesModifyTime 「存在搬运打压风险」属性最后一次更新时间，格式为yyyy-mm-dd ，筛选传入代表筛选出当天「存在搬运打压风险」发生过变化的素材信息
	AttributesModifyTime string `json:"attributes_modify_time,omitempty"`
}

// Encode implement GetRequest interface
func (r MaterialAttributesListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.ReturnLowqualitySuggestions {
		values.Set("return_lowquality_suggestions", "true")
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// MaterialAttributesListResponse 获取视频素材评估标签（新版） API Response
type MaterialAttributesListResponse struct {
	model.BaseResponse
	Data *MaterialAttributesListResult `json:"data,omitempty"`
	// Materials
}

type MaterialAttributesListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Materials 素材及素材属性列表
	Materials []MaterialAttribute `json:"materials,omitempty"`
}

// 素材及素材属性
type MaterialAttribute struct {
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// AdLowQualitySuggestions 当该素材为AD低质素材时，返回低质原因，仅当return_lowquality_suggestions = true时，会返回此参数
	AdLowQualitySuggestions []string `json:"ad_low_quality_suggestions,omitempty"`
	// EcpLowQualitySuggestions 当该素材为千川低质素材时，返回低质原因，仅当return_lowquality_suggestions = true时，会返回此参数
	EcpLowQualitySuggestions []string `json:"ecp_low_quality_suggestions,omitempty"`
	// IsAdHighQualityMaterial 是否AD优质素材
	IsAdHighQualityMaterial bool `json:"is_ad_high_quality_material,omitempty"`
	// IsAdLowQualityMaterial 是否AD低质素材
	IsAdLowQualityMaterial bool `json:"is_ad_low_quality_material,omitempty"`
	// IsEcpHighQualityMaterial 是否千川优质素材
	IsEcpHighQualityMaterial bool `json:"is_ecp_high_quality_material,omitempty"`
	// IsEcpLowQualityMaterial 是否千川低质素材
	IsEcpLowQualityMaterial bool `json:"is_ecp_low_quality_material,omitempty"`
	// IsFirstPublishMaterial 是否是首发素材
	IsFirstPublishMaterial bool `json:"is_first_publish_material,omitempty"`
	// IsInefficientMaterial 是否低效素材
	IsInefficientMaterial bool `json:"is_inefficient_material,omitempty"`
	// IsCarryMaterial 是否存在搬运风险，建议入参account_type = AD 或 QIANCHUAN查询
	IsCarryMaterial bool `json:"is_carry_material,omitempty"`
	// IsSimilarMaterial 是否同质化挤压严重素材
	// 方舟/工作台账户不支持
	IsSimilaryMaterial bool `json:"is_similary_material,omitempty"`
	// IsSimilaryQueueMaterial 是否同质化素材风险-排队投放素材
	// 方舟/工作台账户不支持
	IsSimilaryQueueMaterial bool `json:"is_similary_queue_material,omitempty"`
	// IsSimilaryExpectedQueueMaterial 是否同质化素材风险-未投放预计排队素材
	// 方舟/工作台账户不支持
	IsSimilaryExpectedQueueMaterial bool `json:"is_similary_expected_queue_material,omitempty"`
	// AttributesModifyTime 存在搬运打压风险」属性最后一次更新时间，如素材未被标记为搬运，则不会返回该时间。格式为yyyy-mm-dd HH:MM:SS
	AttributesModifyTime string `json:"attributes_modify_time,omitempty"`
}
