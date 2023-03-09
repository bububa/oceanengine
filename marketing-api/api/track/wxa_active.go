package track

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/model/track"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// WxaActive 微信小程序转化回传
// gw 回传url
// token 密钥
// req 回传payload
func WxaActive(gw string, token string, req *track.WxaActiveRequest, debug bool) error {
	values := util.GetUrlValues()
	values.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	values.Set("nonce", strconv.Itoa(rand.Intn(1000)))
	strs := make([]string, 0, 3)
	strs = append(strs, token)
	strs = append(strs, values.Get("timestamp"))
	strs = append(strs, values.Get("nonce"))
	sort.Strings(strs)
	b := util.GetBufferPool()
	for _, s := range strs {
		b.WriteString(s)
	}
	h := sha1.New()
	h.Write(b.Bytes())
	values.Set("siganture", hex.EncodeToString(h.Sum(nil)))
	util.PutBufferPool(b)
	link := util.StringsJoin(gw, "?", values.Encode())
	util.PutUrlValues(values)
	rw := util.GetBufferPool()
	defer util.PutBufferPool(rw)
	bs, err := json.Marshal(req)
	if err != nil {
		return err
	}
	if debug {
		printPostJSONRequest(link, bs)
	}
	postReq, err := http.NewRequest("post", link, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	postReq.Header.Add("content-type", "application/json")
	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return errors.New(string(body))
}

// printPostJSONRequest print json request with debug
func printPostJSONRequest(url string, body []byte) {
	const format = "[DEBUG] [API] JSON POST %s\n" +
		"http request body:\n%s\n"

	buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
	if err := json.Indent(buf, body, "", "    "); err == nil {
		body = buf.Bytes()
	}
	log.Printf(format, url, body)
}
