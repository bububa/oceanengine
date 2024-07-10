package ad

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// MaterialDelete 删除广告创意
func MaterialDelete(clt *core.SDKClient, accessToken string, req *ad.MaterialDeleteRequest) error {
	return clt.Post("v1.0/qianchuan/ad/material/delete/", req, nil, accessToken)
}
