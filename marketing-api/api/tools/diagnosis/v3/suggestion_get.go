package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	diagnosis "github.com/bububa/oceanengine/marketing-api/model/tools/diagnosis/v3"
)

// SuggestionGet 获取计划诊断建议
// 计划诊断是帮助广告主判断计划投放效果，并定位投放问题给出相关建议的一款优化类工具
// 通过获取计划诊断建议接口，广告主可指定计划ID和场景类型获取相应的诊断建议，针对每一计划ID，本接口可返回计划对应场景下的所有建议，以及建议对应下的所有工具。广告主可通过【采纳诊断建议】接口采纳本接口获得的所有建议，对广告计划进行优化。
func SuggestionGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *diagnosis.SuggestionGetRequest) (*diagnosis.SuggestionGetResponseData, error) {
	var resp diagnosis.SuggestionGetResponse
	err := clt.GetAPI(ctx, "v3.0/tools/promotion_diagnosis/suggestion/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
