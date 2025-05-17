package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoAgentGetRequest 代理商获取视频素材 API Request
type VideoAgentGetRequest struct {
	// AgentID 代理商ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// Filtering 视频过滤条件
	Filtering *VideoAgentGetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值20
	PageSize int `json:"page_size,omitempty"`
}

// VideoAgentGetFilter 视频过滤条件
type VideoAgentGetFilter struct {
	// VideoIDs 视频ids，示例: ["86adb23eaa21229fc04ef932b5089bb8"]
	// 数量限制：<=100
	// 注意：video_ids、material_ids、signatures只能选择一个进行过滤
	VideoIDs []string `json:"video_ids,omitempty"`
	// MaterialIDs 素材id列表，可以根据material_ids（素材报表使用的id，一个素材唯一对应一个素材id）进行过滤
	// 数量限制：<=100
	// 注意：video_ids、material_ids、signatures只能选择一个进行过滤
	MaterialIDs []string `json:"material_ids,omitempty"`
	// Signatures md5值列表，可以根据素材的md5进行过滤
	// 数量限制：<=100
	// 注意：video_ids、material_ids、signatures只能选择一个进行过滤
	Signatures []string `json:"signatures,omitempty"`
	// StartTime 根据视频上传时间进行过滤的起始时间，与end_time搭配使用，格式：yyyy-mm-dd
	StartTime string `json:"start_time,omitempty"`
	// EndTime 根据视频上传时间进行过滤的截止时间，与start_time搭配使用，格式：yyyy-mm-dd
	EndTime string `json:"end_time,omitempty"`
	// Source 素材来源枚举，可选值：
	// AD_SITE : ad后台本地上传
	// CREATIVE_CENTER : 创意中心（巨量创意）
	// OPEN_API : 开放平台
	// SUPPLIER : 即合视频
	// VIDEO_CAPTURE : 易拍视频
	// ACCOUNT_PUSH : 推送视频
	// STAR : 星图视频
	// CEWEBRITY_VIDEO : 达人视频（抖音主页视频）
	// BP : 巨量纵横
	// E_COMMERCE : 巨量千川
	// DPA_MERSCHANT_CENTER : 行业产品中心
	// QF_FUWU : 群峰服务市场
	// AIC : 即创
	Source []enum.MaterialSource `json:"source,omitempty"`
}

// Encode implements GetRequest interface
func (r VideoAgentGetRequest) Encode() string {
	values := util.GetUrlValues()
	defer util.PutUrlValues(values)
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// VideoAgentGetResponse 代理商获取视频素材 API Response
type VideoAgentGetResponse struct {
	model.BaseResponse
	Data *VideoAgentGetResult `json:"data,omitempty"`
}

type VideoAgentGetResult struct {
	// List 视频列表
	List []Video `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
