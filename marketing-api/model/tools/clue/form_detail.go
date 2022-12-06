package clue

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// FormDetailRequest 建站工具——查询表单详情 API Request
type FormDetailRequest struct {
	// AdvertiserID 广告主id，范围：1 <= advertiser_id <= MAX_INT64
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InstanceID 表单ID
	InstanceID uint64 `json:"instance_id,omitempty"`
}

// Encode implement GetRequest interface
func (r FormDetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.InstanceID > 0 {
		values.Set("instance_id", strconv.FormatUint(r.InstanceID, 10))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// FormDetailResponse 建站工具——查询表单详情 API Response
type FormDetailResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *FormDetail `json:"data,omitempty"`
}

// FormDetail 表单详情
type FormDetail struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InstanceID 表单ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// Name 表单名
	Name string `json:"name,omitempty"`
	// Title 表单标题
	Title string `json:"title,omitempty"`
	// SubmitText 提交按钮文案
	SubmitText string `json:"submit_text,omitempty"`
	// SubTitle 副标题
	SubTitle string `json:"sub_title,omitempty"`
	// Elements 表单元素
	Elements []FormElement `json:"elements,omitempty"`
}

// FormElementType 表单元素类型
type FormElementType string

const (
	// FormElementType_NAME 姓名
	FormElementType_NAME FormElementType = "NAME"
	// FormElementType_TELEPHONE 电话
	FormElementType_TELEPHONE FormElementType = "TELEPHONE"
	// FormElementType_EMAIL 邮箱
	FormElementType_EMAIL FormElementType = "EMAIL"
	// FormElementType_NUMBER 数值
	FormElementType_NUMBER FormElementType = "NUMBER"
	// FormElementType_SEX 性别
	FormElementType_SEX FormElementType = "SEX"
	// FormElementType_DATE 日期
	FormElementType_DATE FormElementType = "DATE"
	// FormElementType_CITY 城市
	FormElementType_CITY FormElementType = "CITY"
	// FormElementType_TEXT 文本
	FormElementType_TEXT FormElementType = "TEXT"
	// FormElementType_TEXTAREA 文本域
	FormElementType_TEXTAREA FormElementType = "TEXTAREA"
	// FormElementType_SELECT 下拉框
	FormElementType_SELECT FormElementType = "SELECT"
	// FormElementType_RADIO 单选框
	FormElementType_RADIO FormElementType = "RADIO"
	// FormElementType_CHECKBOX 多选框
	FormElementType_CHECKBOX FormElementType = "CHECKBOX"
	// FormElementType_CALCULATOR 计算器
	FormElementType_CALCULATOR FormElementType = "CALCULATOR"
)

// FormElement 表单元素
type FormElement struct {
	// Label 元素标签
	Label string `json:"label,omitempty"`
	// AllowEmpty 是否允许为空
	// 返回值：
	// 0（必填），
	// 1（可为空）
	AllowEmpty int `json:"allow_empty,omitempty"`
	// Type 表单元素类型
	// 返回值：
	// NAME（姓名）
	// TELEPHONE（电话）
	// EMAIL（邮箱）
	// NUMBER（数值）
	// SEX（性别）
	// DATE（日期）
	// CITY（城市）
	// TEXT（文本）
	// TEXTAREA（文本域）
	// SELECT（下拉框）
	// RADIO（单选框）
	// CHECKBOX（多选框）
	// CALCULATOR（计算器
	Type FormElementType `json:"type,omitempty"`
	// IsUnique 是否可重复
	// 返回值：
	// 0（唯一）
	// 1（可重复）
	IsUnique int `json:"is_unique,omitempty"`
}
