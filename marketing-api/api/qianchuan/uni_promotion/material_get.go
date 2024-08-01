package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// MaterialGet 获取全域推广计划下素材
func MaterialGet(clt *core.SDKClient, accessToken string, req *unipromotion.MaterialGetRequest) (*unipromotion.MaterialGetResult, error) {
	var resp unipromotion.MaterialGetResponse
	if err := clt.Get("v1.0/qianchuan/uni_promotion/ad/material/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
