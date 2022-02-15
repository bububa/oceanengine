package smartphone

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/smartphone"
)

// Delete 删除智能电话
func Delete(clt *core.SDKClient, accessToken string, req *smartphone.DeleteRequest) error {
	return clt.Post("2/clue/smartphone/delete/", req, nil, accessToken)
}
