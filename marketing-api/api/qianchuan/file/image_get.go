package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/file"
)

// ImageGet 获取千川素材库图片
// 该接口用于获取千川素材库图片
// 获取图片素材数据目前仅支持10000个
// 对素材图片进行过滤的时候，image_ids（图片ID）、material_ids（素材ID）、signatures（图片的md5值）只能选择一个进行过滤
// 为保证接口使用的安全性，避免调取他人的图片信息，该接口针对素材URL的字段仅查询自己广告主下的素材才会返回，即需查询的广告主账号的主体需与APPID对应开发者的主体保持一致，才可获取到素材的预览URL的信息，否则会提示：“素材所属主体与开发者主体不一致无法获取URL”（第三方获取敏感物料信息可在授权时申请广告主授权敏感物料权限，可参考常见问题【敏感物料授权】）
// 由于素材库存在分钟级延迟，上传素材后请勿立即获取并创建计划。
func ImageGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.ImageGetRequest) (*file.ImageGetResult, error) {
	var resp file.ImageGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/image/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
