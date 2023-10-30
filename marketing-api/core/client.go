package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/bububa/oceanengine/marketing-api/core/internal/debug"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SDKClient sdk client
type SDKClient struct {
	AppID      uint64
	Secret     string
	debug      bool
	sandbox    bool
	operatorIP string
	limiter    RateLimiter
	client     *http.Client
}

// NewSDKClient 创建SDKClient
func NewSDKClient(appID uint64, secret string) *SDKClient {
	return &SDKClient{
		AppID:  appID,
		Secret: secret,
		client: http.DefaultClient,
	}
}

// SetDebug 设置debug模式
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// SetHttpClient 设置http.Client
func (c *SDKClient) SetHttpClient(client *http.Client) {
	c.client = client
}

// UseSandbox 启用sandbox
func (c *SDKClient) UseSandbox() {
	c.sandbox = true
}

// DisableSandbox 禁用sandbox
func (c *SDKClient) DisableSandbox() {
	c.sandbox = false
}

// SetOperatorIP 设置操作者IP, 支持ipv4/ipv6
func (c *SDKClient) SetOperatorIP(ip string) {
	c.operatorIP = ip
}

// SetRateLimiter 设置限流
func (c *SDKClient) SetRateLimiter(limiter RateLimiter) {
	c.limiter = limiter
}

// Copy 复制SDKClient
func (c *SDKClient) Copy() *SDKClient {
	return &SDKClient{
		AppID:      c.AppID,
		Secret:     c.Secret,
		debug:      c.debug,
		sandbox:    c.sandbox,
		operatorIP: c.operatorIP,
		client:     c.client,
	}
}

// Post post api
func (c *SDKClient) Post(gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	var reqBytes []byte
	if req != nil {
		reqBytes = req.Encode()
	}
	reqUrl := util.StringsJoin(BASE_URL, gw)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.operatorIP != "" {
		httpReq.Header.Add("Operator-Ip", c.operatorIP)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	return c.fetch(httpReq, resp)
}

// Get get api
func (c *SDKClient) Get(gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	reqUrl := util.StringsJoin(BASE_URL, gw)
	if req != nil {
		reqUrl = util.StringsJoin(reqUrl, "?", req.Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.operatorIP != "" {
		httpReq.Header.Add("Operator-Ip", c.operatorIP)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	return c.fetch(httpReq, resp)
}

// GetBytes get bytes api
func (c *SDKClient) GetBytes(gw string, req model.GetRequest, accessToken string) ([]byte, error) {
	reqUrl := util.StringsJoin(BASE_URL, gw)
	if req != nil {
		reqUrl = util.StringsJoin(reqUrl, "?", req.Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.operatorIP != "" {
		httpReq.Header.Add("Operator-Ip", c.operatorIP)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	return io.ReadAll(httpResp.Body)
}

// Upload multipart/form-data post
func (c *SDKClient) Upload(gw string, req model.UploadRequest, resp model.Response, accessToken string) error {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	mw := multipart.NewWriter(buf)
	params := req.Encode()
	mp := make(map[string]string, len(params))
	for _, v := range params {
		var (
			fw  io.Writer
			r   io.Reader
			err error
		)
		if v.Reader != nil {
			if fw, err = mw.CreateFormFile(v.Key, v.Value); err != nil {
				return err
			}
			r = v.Reader
			builder := util.GetStringsBuilder()
			builder.WriteString("@")
			builder.WriteString(v.Value)
			mp[v.Key] = builder.String()
			util.PutStringsBuilder(builder)
		} else {
			if fw, err = mw.CreateFormField(v.Key); err != nil {
				return err
			}
			r = strings.NewReader(v.Value)
			mp[v.Key] = v.Value
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}
	}
	mw.Close()
	reqUrl := util.StringsJoin(BASE_URL, gw)
	debug.PrintPostMultipartRequest(reqUrl, mp, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, buf)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.operatorIP != "" {
		httpReq.Header.Add("Operator-Ip", c.operatorIP)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	return c.fetch(httpReq, resp)
}

// TrackActive 转化回传API专用
func (c *SDKClient) TrackActive(req model.TrackRequest, resp model.Response) error {
	var (
		reqUrl   = req.RequestURI()
		reqBytes = req.Encode()
	)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if token, err := req.Sign(httpReq, nil); err == nil {
		httpReq.Header.Add("x-rs256-token", token)
	}
	if token := req.GetAppAccessToken(); token != "" {
		httpReq.Header.Add("App-Access-Token", token)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	if resp != nil {
		return c.fetch(httpReq, resp)
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != 200 {
		body, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return err
		}
		return model.BaseResponse{Code: httpResp.StatusCode, Message: string(body)}
	}
	return nil
}

// AnalyticsPost 转化回传API专用
func (c *SDKClient) AnalyticsPost(gw string, req model.ConversionRequest, resp model.Response) error {
	reqBytes := req.Encode()
	reqUrl := util.StringsJoin(ANALYTICS_URL, gw)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if token := req.GetAppAccessToken(); token != "" {
		httpReq.Header.Add("App-Access-Token", token)
	}
	if token, err := req.Sign(httpReq, reqBytes); err == nil {
		httpReq.Header.Add("x-rs256-token", token)
	} else {
		reqBuf := util.GetBufferPool()
		reqBuf.Grow(len(reqBytes) + len(c.Secret))
		reqBuf.Write(reqBytes)
		reqBuf.WriteString(c.Secret)
		sha256Sum := sha256.Sum256(reqBuf.Bytes())
		util.PutBufferPool(reqBuf)
		sign := hex.EncodeToString(sha256Sum[:])
		httpReq.Header.Add("x-signature", sign)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	return c.fetch(httpReq, resp)
}

// AnalyticsV1Post 电话转化回传API专用
func (c *SDKClient) AnalyticsV1Post(gw string, req model.PostRequest, resp model.Response) error {
	reqBytes := req.Encode()
	reqUrl := util.StringsJoin(ANALYTICSV1_URL, gw)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	reqBuf := util.GetBufferPool()
	reqBuf.Grow(len(reqBytes) + len(c.Secret))
	reqBuf.Write(reqBytes)
	reqBuf.WriteString(c.Secret)
	sha256Sum := sha256.Sum256(reqBuf.Bytes())
	util.PutBufferPool(reqBuf)
	sign := hex.EncodeToString(sha256Sum[:])
	httpReq.Header.Add("x-signature", sign)
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	return c.fetch(httpReq, resp)
}

// OpenGet get api
func (c *SDKClient) OpenGet(gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	reqUrl := util.StringsJoin(OPEN_URL, gw)
	if req != nil {
		reqUrl = util.StringsJoin(reqUrl, "?", req.Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if accessToken != "" {
		httpReq.Header.Add("App-Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	return c.fetch(httpReq, resp)
}

// fetch execute http request
func (c *SDKClient) fetch(httpReq *http.Request, resp model.Response) error {
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		resp = &model.BaseResponse{}
	}
	if body, err := debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug); err != nil {
		debug.PrintError(err, c.debug)
		return errors.Join(model.BaseResponse{
			Code:    httpResp.StatusCode,
			Message: string(body),
		}, err)
	}
	if resp.IsError() {
		return resp
	}
	return nil
}
