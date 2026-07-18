package stdproject

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialGetRequest 获取投放项目下素材 API Request
type MaterialGetRequest struct {
	// AdvertiserID 客户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectID 项目id
	ProjectID uint64 `json:"project_id,omitempty"`
	// Filtering 过滤器
	Filtering *MaterialGetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值：1
	// 当且仅当 material_type 为 VIDEO、IMAGE、CAROUSEL、TRIAL_PLAY_MATERIAL 时支持传入
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：10, 20, 50, 100，默认值：10
	// 仅 material_type 为 VIDEO 时支持传入
	PageSize int `json:"page_size,omitempty"`
}

// MaterialGetFilter 素材过滤器
type MaterialGetFilter struct {
	// MaterialType 素材类型，单次仅支持传入一个，可选值：
	// VIDEO 视频
	// TITLE 标题
	// IMAGE 图片
	// CAROUSEL 图文
	// TRIAL_PLAY_MATERIAL 试玩素材
	// INSTANT_PLAY_MATERIAL 直玩素材
	// MINI_PROGRAM_INFO 字节小程序&小游戏信息
	// OTHER 其他素材类型（包括产品信息、短剧专辑链接、落地页、行动号召等）
	MaterialType enum.StdMaterialType `json:"material_type,omitempty"`
	// SearchKeyword 搜索关键词，支持根据视频mid进行搜索
	// 当且仅当 material_type 为 VIDEO、IMAGE、CAROUSEL 时支持传入
	SearchKeyword string `json:"search_keyword,omitempty"`
	// MaterialStatusFirst 素材一级状态过滤，可选值：
	// STATUS_DELETE 已删除、STATUS_DISABLE 未投放、STATUS_FROZEN 已终止、STATUS_OK 投放中、STATUS_TIME_DONE 已完成
	MaterialStatusFirst enum.MaterialStatusFirst `json:"material_status_first,omitempty"`
	// MaterialStatusSecond 素材二级状态过滤，可选值：
	// MATERIAL_STATUS_OFFLINE_AUDIT、MATERIAL_STATUS_AUDIT、MATERIAL_STATUS_DISABLE、
	// MATERIAL_STATUS_PROCEDURAL_DISABLE、DOUYIN_ITEM_NOT_AVAILABLE_FOR_DELIVERY、GUIDE_VIDEO_NOT_EXIST
	MaterialStatusSecond enum.MaterialStatusSecond `json:"material_status_second,omitempty"`
}

// Encode implement GetRequest interface
func (r MaterialGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("project_id", strconv.FormatUint(r.ProjectID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
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

// MaterialGetResponse 获取投放项目下素材 API Response
type MaterialGetResponse struct {
	model.BaseResponse
	Data *MaterialGetResponseData `json:"data,omitempty"`
}

// MaterialGetResponseData 获取投放项目下素材返回数据
type MaterialGetResponseData struct {
	// VideoMaterialList 视频素材信息
	VideoMaterialList []VideoMaterial `json:"video_material_list,omitempty"`
	// ImageMaterialList 图片素材信息
	ImageMaterialList []ImageMaterial `json:"image_material_list,omitempty"`
	// CarouselMaterialList 图文素材信息
	CarouselMaterialList []CarouselMaterial `json:"carousel_material_list,omitempty"`
	// TrialPlayMaterialList 试玩素材信息
	TrialPlayMaterialList []TrialPlayMaterial `json:"trial_play_material_list,omitempty"`
	// TitleMaterialList 标题素材信息
	TitleMaterialList []TitleMaterial `json:"title_material_list,omitempty"`
	// ProductInfo 产品信息
	ProductInfo *ProductInfo `json:"product_info,omitempty"`
	// PlayletSeriesURLInfoList 短剧专辑链接信息
	PlayletSeriesURLInfoList []PlayletSeriesURLInfo `json:"playlet_series_url_info,omitempty"`
	// ExternalURLMaterialList 普通落地页链接素材
	ExternalURLMaterialList []string `json:"external_url_material_list,omitempty"`
	// WebURLMaterialList 应用下载详情页
	WebURLMaterialList []string `json:"web_url_material_list,omitempty"`
	// ComponentMaterialList 创意组件信息
	ComponentMaterialList []ComponentMaterial `json:"component_material_list,omitempty"`
	// AnchorMaterialList 原生锚点素材
	AnchorMaterialList []AnchorMaterial `json:"anchor_material_list,omitempty"`
	// InstantPlayMaterialList 直玩素材信息
	InstantPlayMaterialList []InstantPlayMaterial `json:"instant_play_material_list,omitempty"`
	// MiniProgramInfo 字节小程序&小游戏信息
	MiniProgramInfo *MiniProgramInfo `json:"mini_program_info,omitempty"`
	// PlayletSeriesURLList 短剧合集链接url
	PlayletSeriesURLList []string `json:"playlet_series_url_list,omitempty"`
	// OpenURLs 多直达链接
	OpenURLs []string `json:"open_urls,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
