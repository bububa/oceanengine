package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// PromotionCardRecommendTitleGet 查询推广卡片推荐内容（新版）
// 用于查询推广卡片推荐内容，可以查询的文案包括：推广卖点，卡片标题，行动号召。
// advertiser_id，ad_id，industry_id 三个字段的优先级按照所写顺序从低到高排列。例如：用户同时上传了"ad_id"广告计划id和 "industry_id"行业id ，两者信息不一致时，以"industry_id"行业id内信息为准。
// 请求参数中"ad_id"广告主id，"industry_id"行业id，"external_url"落地页链接需要与查询到的文案所使用在的创意对应的广告主，广告主行业，计划中包含的落地页链接保持一致，否则可能会导致获取到的文案效果不好，影响广告投放效果。
// "content_type"推广类型，只对行动号召文案产生影响，系统会根据用户传入的推广类型枚举值来在返回文案中添加和推广类型风格一致的文案。例如：用户请求行动号召文案时，推广类型传入“应用下载类”，则返回文案会包含“立即下载”，“免费下载”等文案。
func PromotionCardRecommendTitleGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.PromotionCardRecommendTitleGetRequest) ([]string, error) {
	var resp tools.PromotionCardRecommendTitleGetResponse
	if err := clt.Get(ctx, "2/tools/promotion_card/recommend_title/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
