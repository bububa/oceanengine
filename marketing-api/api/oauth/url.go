package oauth

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// Url 生成授权链接
func Url(clt *core.SDKClient, redirectUrl string, state string, materialAuth bool) string {
	values := util.GetUrlValues()
	values.Set("app_id", strconv.FormatUint(clt.AppID, 10))
	if state != "" {
		values.Set("state", state)
	}
	values.Set("redirect_uri", redirectUrl)
	if materialAuth {
		values.Set("material_auth", "1")
	}
	rawQuery := values.Encode()
	util.PutUrlValues(values)
	builder := util.GetStringsBuilder()
	builder.WriteString("https://ad.oceanengine.com/openapi/audit/oauth.html?")
	builder.WriteString(rawQuery)
	ret := builder.String()
	util.PutStringsBuilder(builder)
	return ret
}
