package dmp

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AudienceListGetRequest 获取人群管理列表 API Request
type AudienceListGetRequest struct {
	// AdvertiserID 千川广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *AudienceListFilter `json:"filtering,omitempty"`
	// Page 页码，默认为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值20，取值范围：1-100
	PageSize int `json:"page_size,omitempty"`
}

// AudienceListFilter 过滤条件
type AudienceListFilter struct {
	// AudienceID 人群ID
	AudienceID []uint64 `json:"audience_id,omitempty"`
	// AudienceName 人群名称
	AudienceName string `json:"audience_name,omitempty"`
	// AudienceType 人群类型，允许值：
	// BASIC 基础
	// EXTEND 扩展
	// UPLOAD 上传
	AudienceType qianchuan.AudienceType `json:"audience_type,omitempty"`
	// Status 人群状态，允许值：
	// 1 计算完成
	// 2 等待计算
	// 3 计算中
	// 4 计算失败
	// 5 即将过期
	// 6 已过期
	// 7 推送中
	// 8 推送完成
	// 9 推送失败
	Status int `json:"status,omitempty"`
	// StartTime 创建开始时间，例如“2022-08-02”
	StartTime string `json:"start_time,omitempty"`
	// EndTime 创建结束时间，例如“2022-08-02”
	EndTime string `json:"end_time,omitempty"`
}

// Encode implement GetRequest interface
func (r AudienceListGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// AudienceListGetResponse 获取人群管理列表 API Response
type AudienceListGetResponse struct {
	model.BaseResponse
	Data *AudienceListGetResult `json:"data,omitempty"`
}

type AudienceListGetResult struct {
	// AudienceList 人群列表
	AudienceList []Audience `json:"audience_list,omitempty"`
	// PageInfo 页码信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
