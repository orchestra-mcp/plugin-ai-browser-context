package cdp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const defaultPort = "9222"

type TabInfo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Type  string `json:"type"`
}

func ListTabs(ctx context.Context) ([]TabInfo, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:"+defaultPort+"/json", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Chrome not running with remote debugging. Start Chrome with: --remote-debugging-port=9222")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var tabs []TabInfo
	json.Unmarshal(body, &tabs)
	return tabs, nil
}
