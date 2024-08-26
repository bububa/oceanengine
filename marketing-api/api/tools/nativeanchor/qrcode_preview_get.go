package nativeanchor

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/nativeanchor"
)

// QrcodePreviewGet 批量获取锚点预览url
// 获取锚点的预览链接，您需将返回的预览url转成二维码，使用抖音APP扫码才可预览
// 预览url在您请求接口时生成，可预览多次，有效期24小时
// 只有当锚点关联广告时，才可查询到预览url
func QrcodePreviewGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *nativeanchor.QrcodePreviewGetRequest) (*nativeanchor.QrcodePreviewGetResult, error) {
	var resp nativeanchor.QrcodePreviewGetResponse
	if err := clt.GetAPI(ctx, "v3.0/native_anchor/qrcode_preview/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
