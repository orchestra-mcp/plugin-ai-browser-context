package tools

import (
	"context"
	"fmt"
	"strings"

	pluginv1 "github.com/orchestra-mcp/gen-go/orchestra/plugin/v1"
	"github.com/orchestra-mcp/sdk-go/helpers"
	"github.com/orchestra-mcp/plugin-ai-browser-context/internal/cdp"
	"google.golang.org/protobuf/types/known/structpb"
)

func GetOpenTabsSchema() *structpb.Struct {
	s, _ := structpb.NewStruct(map[string]any{
		"type":       "object",
		"properties": map[string]any{},
	})
	return s
}

func GetOpenTabs() func(context.Context, *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
	return func(ctx context.Context, req *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
		tabs, err := cdp.ListTabs(ctx)
		if err != nil {
			return helpers.ErrorResult("chrome_error", err.Error()), nil
		}

		if len(tabs) == 0 {
			return helpers.TextResult("No open tabs found."), nil
		}

		var b strings.Builder
		fmt.Fprintf(&b, "## Open Tabs (%d)\n\n", len(tabs))
		for i, tab := range tabs {
			fmt.Fprintf(&b, "%d. **%s**\n   URL: %s\n   ID: %s\n\n", i+1, tab.Title, tab.URL, tab.ID)
		}
		return helpers.TextResult(b.String()), nil
	}
}
