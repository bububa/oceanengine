package ad

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// KeywordPackageGet 获取词包推荐关键词
func KeywordPackageGet(clt *core.SDKClient, accessToken string, req *ad.KeywordPackageGetRequest) ([]ad.WordPackage, error) {
	var resp ad.KeywordPackageGetResponse
	if err := clt.Get("v1.0/qianchuan/keyword_package/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.WordPackageInfos, nil
}
