package request

import (
	"bytes"
	"singapore/src/utils/log"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

type (
	// Headers 消息头
	Headers struct {
		UserAgent   string
		ContentType string
		Cookies     map[string]string
		Others      map[string]string // Refer, Authorization
	}

	// Method 请求方式
	Method string

	// Client
	Client struct {
		Url    string
		Method Method
		Params map[string]interface{}
	}
)

const (
	MethodForGet  Method = "GET"
	MethodForPost Method = "POST"

	DefaultUserAgent string = "Mozilla/5.0 (SF) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.6.6666.66 Safari/537.36"
)

// RequestBodyFormat 请求消息内容格式
type RequestBodyFormat int

const (
	RequestBodyFormatForFormData RequestBodyFormat = iota + 1
	RequestBodyFormatForXWWWFormUrlencoded
	RequestBodyFormatForRaw
	RequestBodyFormatForJson
	RequestBodyFormatForMultipartFormData
	RequestBodyFormatNull
)

const (
	RequestContentTypeForFormData           string = "application/form-data"
	RequestContentTypeForXWWWFormUrlencoded string = "application/x-www-form-urlencoded"
	RequestContentTypeForJson               string = "application/json"
	RequestContentTypeForMultipartFormData  string = "multipart/form-data"
)

// Request 发起请求
func (c *Client) Request(format RequestBodyFormat, http_proxy string, headers ...Headers) ([]byte, error) {
	client := new(http.Client)

	// http client设置代理
	if http_proxy != "" {
		proxyURL, err := url.Parse(http_proxy)
		if err != nil {
			// http_proxy url.Parse error
			return nil, err
		}

		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	var reqBody io.Reader
	if c.Method == MethodForGet {
		_params := make([]string, 0)

		for k, v := range c.Params {
			_params = append(_params, fmt.Sprintf("%s=%v", k, v))
		}
		c.Url += "?" + strings.Join(_params, "&")
	} else {
		if format == RequestBodyFormatForFormData || format == RequestBodyFormatForXWWWFormUrlencoded {
			_params := make([]string, 0)
			for k, v := range c.Params {
				_params = append(_params, fmt.Sprintf("%s=%v", k, v))
			}
			reqBody = strings.NewReader(strings.Join(_params, "&"))
		} else if format == RequestBodyFormatForRaw || format == RequestBodyFormatForJson {
			_bytes, _ := json.Marshal(c.Params)
			reqBody = bytes.NewReader(_bytes)
		} else if format == RequestBodyFormatForMultipartFormData {
			var err error
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			// 表单写入数据
			// err = writer.WriteField("mode", "automatic")
			// if err != nil {
			// 	panic(err)
			// }

			for k, v := range c.Params {
				err = writer.WriteField(k, fmt.Sprintf("%v", v))
				if err != nil {
					log.Error("Request Type:FormData WriteField Failed", zap.String("field", k), zap.String("value", fmt.Sprintf("%v", v)))
					return nil, errors.New("Request Type:FormData WriteField Failed")
				}
			}

			// 关闭表单写入
			err = writer.Close()
			if err != nil {
				log.Error("Request Type:FormData Writer Close Failed", zap.Error(err))
				return nil, err
			}
			// log.Debug("Request body.Bytes()", zap.ByteString("body", body.Bytes()))
			reqBody = bytes.NewReader(body.Bytes())

			// 获取Headers 数据内容类型，需要boundary=<calculated when request is sent>信息
			headers[0].ContentType = writer.FormDataContentType()
		}
	}
	req, err := http.NewRequest(string(c.Method), c.Url, reqBody)
	if err != nil {
		log.Error("Request http.NewRequest error", zap.Error(err))
		return nil, err
	}
	for _, v := range headers {
		if v.UserAgent != "" {
			req.Header.Add("User-Agent", v.UserAgent)
		}
		if v.ContentType != "" && req.Header.Get("Content-Type") == "" {
			req.Header.Add("Content-Type", v.ContentType)
		}
		if len(v.Cookies) > 0 {
			for key, val := range v.Cookies {
				req.AddCookie(&http.Cookie{Name: key, Value: val})
			}
		}
		if len(v.Others) > 0 {
			for key, val := range v.Others {
				req.Header.Add(key, val)
			}
		}
	}
	resp := new(http.Response)

	log.Debug("[Request Detail]", zap.Any("URL", req.URL), zap.Any("Header", req.Header), zap.Any("reqBody", reqBody), zap.Any("Body", req.Body))
	if resp, err = client.Do(req); err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Error("Request ioutil.ReadAll error", zap.Error(err))
		return nil, err
	}
	// log.Debugf("[Response Detail] status_code:%v, headers:%+v, body:%s", resp.StatusCode, resp.Header, string(bytes))

	if resp.StatusCode != 200 {
		log.Errorf("Request Failed, status_code:%v, headers:%+v, body:%s", resp.StatusCode, resp.Header, string(bytes))
		return bytes, fmt.Errorf("request failed, status_code:%v", resp.StatusCode)
	}

	return bytes, err
}

// NewClient
func NewClient(url string, method Method, params map[string]interface{}) *Client {
	return &Client{
		Url:    url,
		Method: method,
		Params: params,
	}
}
