package tools

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
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
