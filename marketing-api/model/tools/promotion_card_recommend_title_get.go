package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PromotionCardRecommendTitleGetRequest 查询推广卡片推荐内容（新版） API Request
type PromotionCardRecommendTitleGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划id，需要传入"advertiser_id"广告主id名下的计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// IndustryID 行业id，可通过【工具——查询工具——获取行业列表】来获取行业id，任何级别行业id都可以传入
	IndustryID uint64 `json:"industry_id,omitempty"`
	// TextType 推荐文案类型
	// 允许值："PROMOTION": 推广卖点"CARD_TITLE": 卡片标题"CALL_TO_ACTION": 行动号召
	TextType string `json:"text_type,omitempty"`
	// ExternalURL 落地页链接,可通过【建站管理】模块中的落地页获取接口来获取落地页链接
	ExternalURL string `json:"external_url,omitempty"`
	// ContentType 推广类型
	// 允许值："UNKNOWN"：未知 "DOWNLOAD"： 应用下载类 "LANDING"： 落地页类 "TABLE"：附加创意-表单 "CARD"： 附加创意-卡券 "CONSULT"： 附加创意-咨询 "PHONE"： 附加创意-电话 "GAME_PACKAGE"： 游戏礼包码 "GAME_FORM"： 游戏表单 "GAME_SUBSCRIBE"： 游戏预约
	// 默认值："UNKNOWN": 未知
	// 只对行动号召文案产生影响，系统会根据用户传入的推广类型枚举值来在返回文案中添加和推广类型风格一致的文案。
	ContentType string `json:"content_type,omitempty"`
}

// Encode implement GetRequest interface
func (r PromotionCardRecommendTitleGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Add("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Add("text_type", r.TextType)
	if r.AdID > 0 {
		values.Add("ad_id", strconv.FormatUint(r.AdID, 10))
	}
	if r.IndustryID > 0 {
		values.Add("industry_id", strconv.FormatUint(r.IndustryID, 10))
	}
	if r.ExternalURL != "" {
		values.Add("external_url", r.ExternalURL)
	}
	if r.ContentType != "" {
		values.Add("content_type", r.ContentType)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// PromotionCardRecommendTitleGetResponse 查询推广卡片推荐内容（新版） API Response
type PromotionCardRecommendTitleGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List 推荐文案列表
		List []string `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
