package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// KeywordPackageGetRequest 获取词包推荐关键词 API Request
type KeywordPackageGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音id，与计划创建时设置的抖音ID一致
	// 注：PC侧创建「直播带货」搜索广告时，默认根据aweme_id获取推荐关键
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// ProductID 商品ID，此参数用于获取与商品更为相关的关键词推荐
	// 注：PC侧创建「短视频带货」搜索广告时，默认根据product_id获取推荐关键
	ProductID uint64 `json:"product_id,omitempty"`
}

// Encode implement GetRequest interface
func (r KeywordPackageGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	if r.ProductID > 0 {
		values.Set("product_id", strconv.FormatUint(r.ProductID, 10))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// KeywordPackageGetResponse 获取词包推荐关键词 API Respons
type KeywordPackageGetResponse struct {
	model.BaseResponse
	Data struct {
		// WordPackageInfos 词包列表信息
		WordPackageInfos []WordPackage `json:"word_package_infos"`
	} `json:"data,omitempty"`
}

// WordPackage 词包列表信息
type WordPackage struct {
	// ID 词包id
	ID uint64 `json:"id,omitempty"`
	// Name 词包名称
	Name string `json:"name,omitempty"`
	// SearchSum 月搜索量
	SearchSum int64 `json:"search_sum,omitempty"`
	// Degree 竞争程度，枚举值：
	// 1、2、3、4、5（对应pc页面上点亮的格数）
	Degree int `json:"degree,omitempty"`
	// KeywordInfos 关键词列表信息
	KeywordInfos []KeywordInfo `json:"keyword_infos,omitempty"`
}

// KeywordInfo 关键词列表信息
type KeywordInfo struct {
	// KeywordID 关键词id
	KeywordID uint64 `json:"keyword_id,omitempty"`
	// KeywordName 关键词名称
	KeywordName string `json:"keyword_name,omitempty"`
	// KeywordMatchType 关键词匹配方式，可选值:
	// EXTENSIVE
	// PHRASE
	// PRECISION
	KeywordMatchType enum.KeywordMatchType `json:"keyword_match_type,omitempty"`
	// WordPackageName 词包名称
	WordPackageName string `json:"word_package_name,omitempty"`
	// SearchSum 月搜索量
	SearchSum int64 `json:"search_sum,omitempty"`
	// Degree 竞争程度，枚举值：
	// 1、2、3、4、5（对应pc页面上点亮的格数）
	Degree int `json:"degree,omitempty"`
}
