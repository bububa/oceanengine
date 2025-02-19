package tools

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeAuthListRequest 获取抖音授权关系
type AwemeAuthListRequest struct {
	// Filtering 筛选条件
	Filtering *AwemeAuthListFilter `json:"filtering,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

// AwemeAuthListFilter 筛选条件
type AwemeAuthListFilter struct {
	// AuthType 授权类型，可选值:
	// AWEME_ACCOUNT 抖音号授权
	// LIVE_ACCOUNT 直播授权
	// VIDEO_ITEM  单视频授权
	// AWEME_HOMEPAGE主页作品授权 new
	// 说明：创建广告时，如果选用的aweme_id（抖音号）为「抖音号授权」，此时可以在广告下添加新的视频素材并发布到所选抖音号下进行推广；而当所选抖音号的授权类型为「主页作品授权」时，不具备此能力
	// 抖音号授权：授权使用抖音号发布作品并推广、使用主页全部视频推广、直播间引流
	// 主页作品授权：授权使用抖音号主页全部视频推广、直播间引流
	AuthType []enum.AwemeAuthType `json:"auth_type,omitempty"`
	// AuthStatus 授权状态， 可选值:
	// AUTHRIZED: 授权中、AUTHRIZING: 待授权确认、INVALID: 授权失效
	AuthStatus []enum.AwemeAuthStatus `json:"auth_status,omitempty"`
	// AwemeIDs 按抖音号id过滤结果，最长传入50个
	AwemeIDs []string `json:"aweme_ids,omitempty"`
	// ItemIDs 按抖音视频id过滤结果，一次最多允许查询50个
	ItemIDs []uint64 `json:"item_ids,omitempty"`
	// WarningTypes 抖音授权关系警告信息类型，您可使用此筛选查询所有可能存在问题的授权关系。支持枚举，一条授权关系可能存在多种警告类型：
	// UNQUALIFIED 不达要求，表示发起授权的抖音号未达到要求，详细未达门槛信息可通过应答参数中的auth_threshold_info获取
	// UNQUALIFIED_PENDING_EXPIRE 不达要求-即将解除，表示平台定期巡检并解除不符合授权要求的【个人抖音号授权关系】您可按需将您的个人抖音号升级为抖音企业机构账号或督促账号满足要求
	// 仅当auth_type=AWEME_ACCOUNT抖音号授权时，可能会有此警告
	// AUTHOR_REUQUEST_PENDING_EXPIRE 申请解除中，表示抖音号作者已发起解除授权申请，您需要及时联系作者或同意解除
	WarningTypes []enum.AwemeAuthWarningType `json:"warning_types,omitempty"`
}

// Encode implement GetRequest interface
func (r AwemeAuthListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		buf, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(buf))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AwemeAuthListResponse 获取抖音授权关系
type AwemeAuthListResponse struct {
	// Data json返回值
	Data *AwemeAuthListData `json:"data,omitempty"`
	model.BaseResponse
}

// AwemeAuthListData .
type AwemeAuthListData struct {
	// PageInfo .
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List .
	List []AwemeAuthItem `json:"list,omitempty"`
}

type AwemeAuthItem struct {
	// AuthType 授权类型
	AuthType enum.AwemeAuthType `json:"auth_type,omitempty"`
	// AwemeID 抖音号
	AwemeID string `json:"aweme_id,omitempty"`
	// AwemeName 抖音账号名称
	AwemeName string `json:"aweme_name,omitempty"`
	// AwemeUserType 可选值:
	// ENTERPRISE 企业机构号
	// SINGLE 个人号
	AwemeUserType enum.AwemeUserType `json:"aweme_user_type,omitempty"`
	// AuthStatus 授权状态
	AuthStatus enum.AwemeAuthStatus `json:"auth_status,omitempty"`
	// SubStatus 授权子状态，返回值
	// RENEWING 续期待确认
	// RENEW_FAIL 续期失效
	// RENEW_SUCCESS 续期成功
	// INVALID_TIME_OUT 超时未确认
	// INVALID_EXPIRED 授权期限已到
	// INVALID_CANCEL 主动操作解除授权
	// INVALID_REJECT 抖音号作者拒绝授权
	// INVALID_FAILED_BY_AWEME 授权申请失败
	// RENEW_FAILED_BY_AWEME 授权续期申请失败
	// AWEME_REVOKE_REQUEST 创作者发起解除申请
	// INVALID_PROCESS_TIME_OUT 超时未处理自动解除
	// CONFIRM_REVOKE_REQUEST 同意解除授权申请
	// UNQUALIFIED_AUTO_RELEASE不达要求自动解除
	// ENTERPRISE_AUTH_RELEASE 身份变更，不达要求
	SubStatus enum.AwemeAuthSubStatus `json:"sub_status,omitempty"`
	// StartTime 授权开始时间，格式为yyyy-MM-dd HH:mm:ss
	StartTime string `json:"start_time,omitempty"`
	// EndTime 授权结束时间，格式为yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty"`
	// ShareType 授权共享类型 可选值:
	// SHARE_BY_ONESELF 广告账户自主授权
	// SHARE_BY_SAME_ENTITY 客户共享授权
	// SHARE_FROM_BP 组织共享授权
	ShareType enum.AwemeAuthShareType `json:"share_type,omitempty"`
	// VideoInfo 授权视频信息，若为单视频授权会返回
	VideoInfo *AwemeAuthVideo `json:"video_info,omitempty"`
	// Note 备注信息，发起抖音授权申请时填写的希望展示给抖音号创作者的备注，创作者可在授权邀请页&授权详情页查看
	Note string `json:"note,omitempty"`
	// AwemeCancelReason 抖音号作者发起解除授权的原因，仅当抖音号作者发起解除授权时有值，100字以内，可能包括：
	// 不知道该授权是怎么建立的，申请解除授权
	// 联系不到对方，无法进行合作沟通，申请解除授权
	// 与对方合作到期或者有纠纷，申请解除授权
	// 其他情况（作者会填写其他文案给到）
	AwemeCancelReason string `json:"aweme_cancel_reason,omitempty"`
	// AwemeCancelImageList 抖音号作者发起解除授权时上传的凭证信息（选填项，抖音号作者可能不填，此时该参数返回为null）
	AwemeCancelImageList []string `json:"aweme_cancel_image_list,omitempty"`
	// AwemeCancelNote 抖音号作者发起解除授权时填写的联系方式（选填项，抖音号作者可能不填，此时该参数返回为null）
	AwemeCancelNote string `json:"aweme_cancel_note,omitempty"`
	// WarningContent 抖音授权关系警告信息，您可根据该信息及时处理，可能返回
	// 不达门槛：表示发起授权的抖音号未达到要求，详细未达门槛信息可通过auth_threshold_info获取
	// 即将解除：表示抖音号作者已发起解除授权申请，您需要及时联系作者或同意解除
	WarningContent []string `json:"warning_content,omitempty"`
	// WarningTypes 抖音授权关系警告信息类型支持枚举，一条授权关系可能存在多种警告类型：
	// UNQUALIFIED 不达要求，表示发起授权的抖音号未达到要求，详细未达门槛信息可通过应答参数中的auth_threshold_info获取
	// UNQUALIFIED_PENDING_EXPIRE 不达要求-即将解除，表示平台定期巡检并解除不符合授权要求的【个人抖音号授权关系】您可按需将您的个人抖音号升级为抖音企业机构账号或督促账号满足要求
	// 仅当auth_type=AWEME_ACCOUNT抖音号授权时，可能会有此警告
	// AUTHOR_REUQUEST_PENDING_EXPIRE 申请解除中，表示抖音号作者已发起解除授权申请，您需要及时联系作者或同意解除
	WarningTypes []enum.AwemeAuthWarningType `json:"warning_types,omitempty"`
	// AuthAutoExpireDate 授权不达要求后自动解除的日期，日期格式yyyy-mm-dd
	// 2025年2月中起，平台将分批解除不符合授权要求的【个人抖音号授权关系】，此抖音号未满足授权要求，授权关系将于auth_auto_expire_date被解除。如您仍希望保持授权关系，可按需将个人抖音号升级为抖音企业机构账号或督促账号达到要求，不达要求的具体信息可查看auth_threshold_info参数。说明详见此文档：【个人抖音号】商业合作授权定期、分批次复核启动通知
	// 当抖音号正常时，该结构体下信息会返回为空
	AuthAutoExpireDate string `json:"auth_auto_expire_date,omitempty"`
	// AuthThresholdInfo 当抖音号不达授权门槛时，您可通过此结构体获取具体不达门槛的原因并及时处理
	// 当抖音号正常时，该结构体下信息会返回为空
	AuthThresholdInfo *AwemeAuthThresholdInfo `json:"auth_threshold_info,omitempty"`
	// HasVideoHpVisibilityLimit 发布新视频素材到该抖音号下时，视频主页可见性只能设置「仅单次展示可见」，枚举值：
	// true：是
	// false：否，表示无此限制
	// 当值返回true时，代表在创建广告时添加新视频素材到该抖音号下推广，视频的主页可见性设置只允许HIDE_VIDEO_ON_HP「仅单次展示可见」
	HasVideoHpVisibilityLimit bool `json:"has_video_hp_visibility_limit,omitempty"`
}

// AwemeAuthVideo 授权视频信息，若为单视频授权会返回
type AwemeAuthVideo struct {
	// ItemID 抖音视频ID
	ItemID uint64 `json:"item_id,omitempty"`
	// ImageMode 素材类型
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// Title 视频名称
	Title string `json:"title,omitempty"`
	// AwemePlayURL 视频播放链接
	AwemePlayURL string `json:"aweme_play_url,omitempty"`
	// Duration 视频时长，单位为秒
	Duration float64 `json:"duration,omitempty"`
	// VideoCoverID 视频封面ID
	VideoCoverID string `json:"video_cover_id,omitempty"`
	// VideoCoverURL 视频封面链接
	VideoCoverURL string `json:"video_cover_url,omitempty"`
	// Mid 视频素材ID
	// （仅抖音视频会有此字段，抖音图文素材没有此字段）
	Mid string `json:"mid,omitempty"`
}

// AwemeAuthThresholdInfo 当抖音号不达授权门槛时，您可通过此结构体获取具体不达门槛的原因并及时处理
type AwemeAuthThresholdInfo struct {
	// IsAudit 是否已成年
	// true表示 是
	// false 表示 否，表示未成年，会导致不达门槛A
	IsAudit bool `json:"is_audit,omitempty"`
	// IsRealNameCert 是否已实名
	// true表示 是
	// false 表示 否，表示未实名，会导致不达门槛
	IsRealNameCert bool `json:"is_real_name_cert,omitempty"`
	// IsSeriousViolation 是否严重违规，
	// true表示 严重违规，会导致不达门槛
	// false 表示 没有严重违规
	IsSeriousViolation bool `json:"is_serious_violation,omitempty"`
	// IsReachedRiseFansCount 有效涨粉数是否达到1000
	// true表示 达到
	// false 表示 否，表示没有达到1000，会导致不达门槛
	IsReachedRiseFansCount bool `json:"is_reached_rise_fans_count,omitempty"`
}
