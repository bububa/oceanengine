package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// DeliveryQualificationSubmit 投放资质提交
// ## 提交接口使用必看
// 该接口用于广告主投放资质提交，请注意以下使用规则：
// - 需要将资质上传至对应的同名资质类型中，资质类型选择错误将会被审核拒绝。若找不到对应的资质类型，可以上传至“其他资质”，详细的资质类型可选值见下方接口入参。
// - 需要上传多份资质时，例如要上传三份肖像授权书，请将每一份肖像授权书分开上传，多份资质合并上传将会被审核拒绝。
// - 对于一份完整资质的多张图片请上传至一个资质id中，例如一份经销授权书的3张图片内容
//  1. 错误方式：分3次调用接口，会导致资质接收不全而被拒绝
//  2. 正确方式：将一份完整的资质调一次接口上传至一个资质中，保证审核平台能够一次收到完整的资质内容
func DeliveryQualificationSubmit(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.DeliveryQualificationSubmitRequest) ([]uint64, error) {
	var resp advertiser.DeliveryQualificationSubmitResponse
	if err := clt.PostAPI(ctx, "v3.0/advertiser/delivery_qualification/submit/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.QualificationIDs, nil
}
