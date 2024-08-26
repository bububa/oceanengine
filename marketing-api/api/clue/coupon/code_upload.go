package coupon

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/coupon"
)

// CodeUpload 上传券码
// 通过此接口，你可以上传优惠券的券码。调用此接口前，需要先创建卡券，获取优惠券id，优惠券的券码类型需要为商家上传券码类型。
func CodeUpload(ctx context.Context, clt *core.SDKClient, accessToken string, req *coupon.CodeUploadRequest) (*coupon.CodeUploadResponseData, error) {
	var resp coupon.CodeUploadResponse
	if err := clt.Upload(ctx, "2/clue/coupon/upload/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
