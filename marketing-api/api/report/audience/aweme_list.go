package audience

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/audience"
)

// 抖音达人数据
// 此接口用于获取：抖音达人数据，如果你在计划的受众定向中设置了抖音达人，那么在这能看到所选的抖音达人指标所对应的数据，包括花费、展示、点击等。
// 抖音达人数据不支持获取当天的数据；
func AwemeList(clt *core.SDKClient, accessToken string, req *audience.ListRequest) (*audience.ListResponse, error) {
	var resp audience.ListResponse
	err := clt.Get("2/report/audience/aweme/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
