package tools

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Post(url string, params any) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	body, _ := json.Marshal(params)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost,
		url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	resp, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetWithHeader(path string, m map[string]string, proxy string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		httpReq.Header.Set(k, v)
	}

	client := http.DefaultClient
	if proxy != "" {
		proxyAddress, _ := url.Parse(proxy)
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyAddress),
			},
		}
	}
	httpRsp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	return io.ReadAll(httpRsp.Body)
}
