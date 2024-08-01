package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// MaterialDelete 删除全域推广计划下素材
func MaterialDelete(clt *core.SDKClient, accessToken string, req *unipromotion.MaterialDeleteRequest) error {
	return clt.Post("v1.0/qianchuan/uni_promotion/ad/material/delete/", req, nil, accessToken)
}
