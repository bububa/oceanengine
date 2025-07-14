package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PlayableListRequest 获取试玩/直玩素材 API Request
type PlayableListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialType 素材类型，可选值:
	// 直玩素材 STRAIGHT_PLAYABLE
	// 试玩素材 PLAYABLE_NEW
	MaterialType enum.MaterialType `json:"material_type,omitempty"`
	// Filtering 可选过滤参数
	Filtering *PlayableListFilter `json:"filtering,omitempty"`
	// Page 页码（默认为1）
	Page int `json:"page,omitempty"`
	// PageSize 分页大小（默认为20）
	PageSize int `json:"page_size,omitempty"`
}

type PlayableListFilter struct {
	// MaterialIDs 素材ID列表
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// AppMaterialIDs 试玩/直玩包Material_id列表
	AppMaterialIDs []uint64 `json:"app_material_ids,omitempty"`
	// Keyword 试玩/直玩素材查询，可通过素材名称/素材ID进行查询
	Keyword string `json:"keyword,omitempty"`
	// Signatures 素材MD5值
	Signatures []string `json:"signatures,omitempty"`
	// StartTime 根据素材上传时间进行过滤的起始时间，与end_time搭配使用
	// 格式：yyyy-mm-dd hh:mm:ss
	StartTime string `json:"start_time,omitempty"`
	// EndTime 根据素材上传时间进行过滤的截止时间，与end_time搭配使用
	// 格式：yyyy-mm-dd hh:mm:ss
	EndTime string `json:"end_time,omitempty"`
	// Status 素材状态，枚举值：
	// ONLINE 素材在线
	// OFFLINE 已下线
	Status string `json:"status,omitempty"`
	// AuditStatus 审核状态，枚举值：
	// STATUS_CONFIRMING 审核中
	// STATUS_CONFIRMED已通过
	// STATUS_CONFIRM_FAIL 已拒绝
	AuditStatus string `json:"audit_status,omitempty"`
	// AppIDs 直玩-小游戏app_id（仅直玩素材类型支持）
	AppIDs []string `json:"app_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r PlayableListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.MaterialType != "" {
		values.Set("material_type", string(r.MaterialType))
	}
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

// PlayableListResponse 获取试玩/直玩素材 API Response
type PlayableListResponse struct {
	model.BaseResponse
	Data *PlayableListResult `json:"data,omitempty"`
}

type PlayableListResult struct {
	// Pagination 分页
	Pagination *model.PageInfo `json:"paginiation,omitempty"`
	// Materials 素材列表
	Materials []PlayableMaterial `json:"mateitrials,omitempty"`
}
