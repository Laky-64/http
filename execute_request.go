package http

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Laky-64/http/types"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

func ExecuteRequest(url string, options ...RequestOption) (*types.HTTPResult, error) {
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
	client := http.Client{
		Timeout: opt.Timeout,
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
	req, err := http.NewRequest(opt.Method, url, body)
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
	var proxyBody io.Reader = do.Body
	if opt.OverloadReader != nil {
		proxyBody = opt.OverloadReader(proxyBody)
	}
	var buf []byte
	for {
		var b = make([]byte, 1024*4)
		n, fErr := io.ReadFull(proxyBody, b)
		buf = append(buf, b[:n]...)
		if fErr != nil {
			if fErr == io.EOF {
				break
			}
			if !errors.Is(fErr, io.ErrUnexpectedEOF) {
				return nil, fErr
			}
		}
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(do.Body)
	var errRequest error
	if do.StatusCode != http.StatusOK && do.StatusCode != http.StatusCreated && do.StatusCode != http.StatusNoContent {
		opt.Retries--
		if opt.Retries > 0 {
			time.Sleep(time.Millisecond * 250)
			return ExecuteRequest(url, options...)
		}
		errRequest = fmt.Errorf("http status code %d", do.StatusCode)
	}
	return &types.HTTPResult{
		Body:       buf,
		StatusCode: do.StatusCode,
	}, errRequest
}
