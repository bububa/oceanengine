package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// AttachmentUpload 批量上传资质附件
// 附件上传接口，在通过【投放资质提交】接口提交资质前，需要将资质附件从本接口上传至认证中心，获取附件id，将附件id作为入参传入投放资质提交接口
func AttachmentUpload(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.AttachmentUploadRequest) (uint64, error) {
	var resp advertiser.AttachmentUploadResponse
	if err := clt.Upload(ctx, "v3.0/advertiser/attachment/upload/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.AttachmentID, nil
}
