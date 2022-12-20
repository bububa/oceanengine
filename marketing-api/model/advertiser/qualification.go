package advertiser

import "github.com/bububa/oceanengine/marketing-api/enum"

// Qualification 广告主资质
type Qualification struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Status 状态，枚举值：STATUS_NOT_SUBMIT 未提交、STATUS_WAIT_CONFIRM 待审核、STATUS_PENDING_CONFIRM 审核中、STATUS_CONFIRM 审核通过、STATUS_CONFIRM_FAIL 审核不通过
	Status enum.QualificationStatus `json:"status,omitempty"`
	// RejectReason 拒绝理由
	RejectReason string `json:"reject_reason,omitempty"`
	// Subject 主体资质
	Subject *QualificationSubject `json:"subject,omitempty"`
	// Industries 行业资质
	Industries []QualificationIndustry `json:"industries,omitempty"`
}

// BaseQualification 资质基本信息
type BaseQualification struct {
	// QualificationID 资质id
	QualificationID uint64 `json:"qualification_id,omitempty"`
	// AttachmentID 资质图片附件id
	AttachmentID string `json:"attachment_id,omitempty"`
	// PictureURL 资质图片地址
	PictureURL string `json:"picture_url,omitempty"`
	// Status 状态，枚举值：STATUS_NOT_SUBMIT 未提交、STATUS_WAIT_CONFIRM 待审核、STATUS_PENDING_CONFIRM 审核中、STATUS_CONFIRM 审核通过、STATUS_CONFIRM_FAIL 审核不通过
	Status enum.QualificationStatus `json:"status,omitempty"`
	// RejectReason 拒绝理由
	RejectReason string `json:"reject_reason,omitempty"`
}

// QualificationSubject 主体资质
type QualificationSubject struct {
	BaseQualification
	// CompanyName 公司名称
	CompanyName string `json:"company_name,omitempty"`
	// CompanyType 公司类型，枚举值:
	// COMPANY 企业、INDIVIDUAL 个人
	CompanyType string `json:"company_type,omitempty"`
	// CheckType 对公验证类型，枚举值:
	// COMPANY 企业、GOVERNMENT 政府组织机构/事业单位、HK_MACAO_TAIWAN 港澳台、INDIVIDUAL 个人、OTHERS 其他机构（如民办机构）、OVERSEA 海外、SELF_EMPLOY 个体工商户
	CheckType string `json:"check_type,omitempty"`
	// QualificationType 资质类型，枚举值:
	// ASSOCIATION_REGISTER_CODE 社会团体法人登记证书编号
	// COMMERCIAL_REGISTER_NUMBER 商业登记证号码
	// COMPANY_CREDIT_CODE 企业营业执照统一社会信用代码
	// COMPANY_REGISTER_CODE 企业营业执照注册号
	// CREDIT_CODE 统一社会信用代码证书编号
	// HK_MACAO_TAIWAN_ID 港澳/台湾居民往来大陆通行证证件号码
	// HK_REGISTER_CODE 香港公司注册证书编号
	// ID 身份证号
	// INDIVIDUAL_CREDIT_CODE 个体工商户营业执照统一社会信用代码
	// INDIVIDUAL_REGISTER_CODE 个体工商户营业执照注册号
	// LAWYER_CERTIFICATE_NUMBER 律师执业证书执业证号
	// LAWYER_PERMIT_CODE 律师事务所执业许可证编号
	// LEGAL_PERSON_CREDIT_CODE 事业单位法人证书统一社会信用代码
	// ORGANIZATION_REGISTER_CODE 组织机构代码证代号
	// OTHER 其他编号
	// OVERSEA_REGISTER_CODE 外国（地区）企业常驻代表机构登记证编号
	// PASSPORT_ID 护照号
	// SCHOOL_PERMIT_CODE 民办学校办学许可证编号
	// SOCIAL_REGISTER_CODE 民办非企业单位登记证书编号
	QualificationType string `json:"qualification_type,omitempty"`
	// QualificationCode 资质编号
	QualificationCode string `json:"qualification_code,omitempty"`
	// RegisteredNationName 注册国家
	RegisteredNationName string `json:"registered_nation_name,omitempty"`
	// RegisteredProvinceName 注册省份
	RegisteredProvinceName string `json:"registered_province_name,omitempty"`
	// RegisteredCityName 注册城市
	RegisteredCityName string `json:"registered_city_name,omitempty"`
	// HasEffectiveDate 是否有有效日期
	HasEffectiveDate bool `json:"has_effective_date,omitempty"`
	// EffectiveDate 过期时间，格式yyyy-mm-dd
	EffectiveDate string `json:"effective_date,omitempty"`
	// ProprietorName 法人
	ProprietorName string `json:"proprietor_name,omitempty"`
	// Address 详细地址
	Address string `json:"address,omitempty"`
}

// QualificationIndustry 行业资质
type QualificationIndustry struct {
	// IndustryID 行业id
	IndustryID uint64 `json:"industry_id,omitempty"`
	// Industry1ID 一级行业ID
	Industry1ID uint64 `json:"industry1_id,omitempty"`
	// Industry1Name 一级行业名称
	Industry1Name string `json:"industry1_name,omitempty"`
	// Industry2ID 二级行业ID
	Industry2ID uint64 `json:"industry2_id,omitempty"`
	// Industry2Name 二级行业名称
	Industry2Name string `json:"industry2_name,omitempty"`
	// Industry3ID 三级行业ID
	Industry3ID uint64 `json:"industry3_id,omitempty"`
	// Industry3Name 三级行业名称
	Industry3Name string `json:"industry3_name,omitempty"`
	// Promotion 推广资质
	Promotion *QualificationPromotion `json:"promotion,omitempty"`
	// Others 开户资质列表
	Others []BaseQualification `json:"others,omitempty"`
	// IsHistory 是否是历史的补充资质及推广资质，如果是，则行业相关字段无值，重新提交时需归档到具体行业下
	IsHistory bool `json:"is_history,omitempty"`
}

// QualificationPromotion 推广资质
type QualificationPromotion struct {
	BaseQualification
	// Content 推广内容
	Content string `json:"content,omitempty"`
}
