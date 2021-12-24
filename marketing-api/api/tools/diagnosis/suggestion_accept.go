package diagnosis

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/diagnosis"
)

// SuggestionAccept 采纳计划诊断建议
// 通过采纳诊断建议接口，广告主可以采纳【获取诊断建议】接口所获得的所有建议
func SuggestionAccept(clt *core.SDKClient, accessToken string, req *diagnosis.SuggestionAcceptRequest) (*diagnosis.SuggestionAcceptResponseData, error) {
	var resp diagnosis.SuggestionAcceptResponse
	err := clt.Post("2/tools/diagnosis/suggestion/accept/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
