package aweme

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// VideoGet 获取随心推可投视频列表
// 获取小店可投视频列表
func VideoGet(clt *core.SDKClient, accessToken string, req *aweme.VideoGetRequest) (*aweme.VideoGetResult, error) {
	var resp aweme.VideoGetResponse
	if err := clt.Get("v1.0/qianchuan/aweme/video/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
