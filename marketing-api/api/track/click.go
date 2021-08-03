package track

import (
	"net/url"
)

// DEFAULT_CLICK_FIELDS 默认点击检测字段
var DEFAULT_CLICK_FIELDS = []string{
	"request_id",
	"aid",
	"advertiser_id",
	"cid",
	"campaign_id",
	"ctype",
	"csite",
	"imei",
	"idfa",
	"android_id",
	"oaid",
	"os",
	"mac",
	"mac1",
	"ip",
	"ua",
	"geo",
	"ts",
	"callback",
	"callback_url",
	"model",
	"union_site",
	"caid1",
}

// Click 生成击检测链接
func Click(baseUrl string, fields []string) string {
	if fields == nil {
		fields = DEFAULT_CLICK_FIELDS
	}
	parsedUrl, _ := url.Parse(baseUrl)
	values := parsedUrl.Query()
	for _, field := range fields {
		switch field {
		case "request_id":
			values.Set("request_id", "__REQUEST_ID__")
		case "aid":
			values.Set("aid", "__AID__")
		case "advertiser_id":
			values.Set("advertiser_id", "__ADVERTISER_ID__")
		case "cid":
			values.Set("cid", "__CID__")
		case "campaign_id":
			values.Set("campaign_id", "__CAMPAIGN_ID__")
		case "ctype":
			values.Set("ctype", "__CTYPE__")
		case "csite":
			values.Set("csite", "__CSITE__")
		case "imei":
			values.Set("imei", "__IMEI__")
		case "idfa":
			values.Set("idfa", "__IDFA__")
		case "android_id":
			values.Set("android_id", "__ANDROIDID__")
		case "oaid":
			values.Set("oaid", "__OAID__")
		case "os":
			values.Set("os", "__OS__")
		case "mac":
			values.Set("mac", "__MAC__")
		case "mac1":
			values.Set("mac1", "__MAC1__")
		case "ip":
			values.Set("ip", "__IP__")
		case "ua":
			values.Set("ua", "__UA__")
		case "geo":
			values.Set("geo", "__GEO__")
		case "ts":
			values.Set("ts", "__TS__")
		case "callback":
			values.Set("callback", "__CALLBACK_PARAM__")
		case "callback_url":
			values.Set("callback_url", "__CALLBACK_URL__")
		case "model":
			values.Set("model", "__MODEL__")
		case "union_site":
			values.Set("union_site", "__UNION_SITE__")
		case "caid1":
			values.Set("caid1", "__CAID1__")
		case "caid2":
			values.Set("caid2", "__CAID2__")
		}
	}
	parsedUrl.RawQuery = values.Encode()
	return parsedUrl.String()
}
