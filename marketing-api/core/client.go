package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/bububa/oceanengine/marketing-api/core/internal/debug"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// SDKClient sdk client
type SDKClient struct {
	AppID   string
	Secret  string
	debug   bool
	sandbox bool
}

// NewSDKClient 创建SDKClient
func NewSDKClient(appID string, secret string) *SDKClient {
	return &SDKClient{
		AppID:  appID,
		Secret: secret,
	}
}

// SetDebug 设置debug模式
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// UseSandbox 启用sandbox
func (c *SDKClient) UseSandbox() {
	c.sandbox = true
}

// DisableSandbox 禁用sandbox
func (c *SDKClient) DisableSandbox() {
	c.sandbox = false
}

// Post post api
func (c *SDKClient) Post(gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	var reqBytes []byte
	if req != nil {
		reqBytes = req.Encode()
	}
	reqUrl := fmt.Sprintf("%s%s", BASE_URL, gw)
	debug.PrintPostJSONRequest(reqUrl, reqBytes, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	httpReq.Header.Add("Content-Type", "application/json")
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}

// Get get api
func (c *SDKClient) Get(gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	var reqQuery string
	if req != nil {
		reqQuery = req.Encode()
	}
	reqUrl := fmt.Sprintf("%s%s?%s", BASE_URL, gw, reqQuery)
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}

// Upload multipart/form-data post
func (c *SDKClient) Upload(gw string, req model.UploadRequest, resp model.Response, accessToken string) error {
	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
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
			mp[v.Key] = fmt.Sprintf("@%s", v.Value)
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
	reqUrl := fmt.Sprintf("%s%s", BASE_URL, gw)
	debug.PrintPostMultipartRequest(reqUrl, mp, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, &buf)
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}

// AnalyticsPost 转化回传API专用
func (c *SDKClient) AnalyticsPost(gw string, req model.PostRequest, resp model.Response) error {
	reqBytes := req.Encode()
	reqBuf := bytes.NewBuffer(reqBytes)
	reqBuf.WriteString(c.Secret)
	sign := fmt.Sprintf("%x", sha256.Sum256(reqBuf.Bytes()))
	reqUrl := fmt.Sprintf("%s%s", ANALYTICS_URL, gw)
	debug.PrintPostJSONRequest(reqUrl, reqBuf.Bytes(), c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("x-signature", sign)
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}
