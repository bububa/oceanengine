package servemarket

import (
	"context"
	"net/http"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/servemarket"
)

// CidVerifyToken 获取App Access Token校验信息
func CidVerifyToken(ctx context.Context, clt *core.SDKClient, token string) (*servemarket.CidVerifyTokenResult, error) {
	cltCopied := clt.Copy()
	cltCopied.AddPreRequests(func(httpReq *http.Request) error {
		httpReq.Header.Set("App-Access-Token", token)
		return nil
	})
	var resp servemarket.CidVerifyTokenResponse
	if err := cltCopied.GetAPI(ctx, "v1.0/serve_market/cid/verify_token/", nil, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
