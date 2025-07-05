package http

import (
	"bytes"
	"fmt"
	"github.com/Laky-64/http/types"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

func ExecuteRequest(uri string, options ...RequestOption) (*types.HTTPResult, error) {
	var opt types.RequestOptions
	for _, option := range options {
		option.Apply(&opt)
	}
	if (opt.MultiPart != nil) == (opt.Body != nil) && opt.MultiPart != nil {
		return nil, fmt.Errorf("can't use multipart and body at the same time")
	}
	if opt.Method == "" {
		opt.Method = "GET"
	}

	transport := http.DefaultTransport
	if len(opt.Proxy) > 0 && opt.Transport != nil {
		return nil, fmt.Errorf("can't use both proxy and transport at the same time")
	} else if opt.Transport != nil {
		transport = opt.Transport
	} else if len(opt.Proxy) > 0 {
		proxyUrl, err := url.ParseRequestURI(opt.Proxy)
		if err != nil {
			return nil, err
		}
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}
	client := http.Client{
		Timeout:       opt.Timeout,
		Transport:     transport,
		CheckRedirect: opt.HandleRedirect,
	}
	var body io.Reader
	var multiPartWriter *multipart.Writer
	if opt.MultiPart != nil {
		reader := &bytes.Buffer{}
		multiPartWriter = multipart.NewWriter(reader)
		for k, v := range opt.MultiPart.Data {
			_ = multiPartWriter.WriteField(k, v)
		}
		for k, v := range opt.MultiPart.Files {
			file, err := multiPartWriter.CreateFormFile(k, v.FileName)
			if err != nil {
				return nil, err
			}
			_, err = file.Write(v.Content)
			if err != nil {
				return nil, err
			}
		}
		_ = multiPartWriter.Close()
		body = reader
	} else if opt.Body != nil {
		body = bytes.NewBuffer(opt.Body)
	}
	if !requestLacksBody(opt.Method) && body != nil && opt.OverloadReader != nil {
		body = opt.OverloadReader(body)
	}
	req, err := http.NewRequest(opt.Method, uri, body)
	if err != nil {
		return nil, err
	}
	if opt.Headers != nil {
		for k, v := range opt.Headers {
			req.Header.Set(k, v)
		}
	}
	if opt.BearerToken != "" && req.Header.Get("Authorization") == "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", opt.BearerToken))
	}
	if multiPartWriter != nil {
		req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
	}
	req.Header.Add("Accept-Encoding", "identity")
	for k, v := range opt.Cookies {
		req.AddCookie(
			&http.Cookie{
				Name:  k,
				Value: v,
			},
		)
	}
	do, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	var resultBody io.Reader = do.Body
	if requestLacksBody(opt.Method) && opt.OverloadReader != nil {
		resultBody = opt.OverloadReader(resultBody)
	}
	_, err = io.Copy(&buf, resultBody)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(do.Body)
	var errRequest error
	if do.StatusCode != http.StatusOK && do.StatusCode != http.StatusCreated && do.StatusCode != http.StatusNoContent {
		opt.Retries--
		if opt.Retries > 0 {
			time.Sleep(time.Millisecond * 250)
			return ExecuteRequest(uri, options...)
		}
		errRequest = fmt.Errorf("http status code %d", do.StatusCode)
	}
	return &types.HTTPResult{
		Body:       buf.Bytes(),
		Headers:    do.Header,
		Cookies:    do.Cookies(),
		StatusCode: do.StatusCode,
	}, errRequest
}
