package debug

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/bububa/oceanengine/marketing-api/util"
)

// PrintError print error with debug
func PrintError(err error, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [ERROR]", err)
}

// PrintStringResponse print string response with debug
func PrintStringResponse(str string, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [RESPONSE]", str)
}

// PrintGetRequest print get request with debug
func PrintGetRequest(url string, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [API] GET", url)
}

// PrintPostJSONRequest print json request with debug
func PrintPostJSONRequest(url string, body []byte, debug bool) {
	if !debug {
		return
	}
	const format = "[DEBUG] [API] JSON POST %s\n" +
		"http request body:\n%s\n"

	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	json.Indent(buf, body, "", "    ")
	log.Printf(format, url, buf.String())
}

// PrintJSONRequest print json request with debug
func PrintJSONRequest(method string, url string, header http.Header, body []byte, debug bool) {
	if !debug {
		return
	}

	const format = "[DEBUG] [API] JSON %s %s\n" +
		"http request header:\n%s\n" +
		"http request body:\n%s\n"

	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	json.Indent(buf, body, "", "\t")
	headers := make([]string, 0, len(header))
	for k := range header {
		headers = append(headers, util.StringsJoin(k, ": ", header.Get(k)))
	}

	log.Printf(format, method, url, strings.Join(headers, "\n"), buf.String())
}

// PrintPostMultipartRequest print multipart/form-data post request with debug
func PrintPostMultipartRequest(url string, mp map[string]string, debug bool) {
	if !debug {
		return
	}
	const format = "[DEBUG] [API] multipart/form-data POST %s\n" +
		"http request body:\n%s\n"

	bs, _ := json.MarshalIndent(mp, "", "\t")
	log.Printf(format, url, bs)
}

// DecodeJSONHttpResponse decode json response with debug
func DecodeJSONHttpResponse(r io.Reader, v interface{}, debug bool) ([]byte, error) {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	tee := io.TeeReader(r, buf)
	if err := json.NewDecoder(tee).Decode(v); err != nil {
		return buf.Bytes(), err
	}
	if !debug {
		return nil, nil
	}
	bs := buf.Bytes()
	debugBuf := util.GetBufferPool()
	defer util.PutBufferPool(debugBuf)
	if err := json.Indent(debugBuf, bs, "", "\t"); err != nil {
		return bs, err
	}

	log.Println(util.StringsJoin("[DEBUG] [API] http response body:\n", debugBuf.String()))
	return nil, nil
}
