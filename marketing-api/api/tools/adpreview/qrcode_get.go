package adpreview

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adpreview"
)

// QrcodeGet 获取广告预览二维码
func QrcodeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *adpreview.QrcodeGetRequest) (*adpreview.QrcodeGetResponseData, error) {
	var resp adpreview.QrcodeGetResponse
	err := clt.Get(ctx, "2/tools/ad_preview/qrcode_get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
