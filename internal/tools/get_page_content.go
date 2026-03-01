package tools

import (
	"context"
	"fmt"

	pluginv1 "github.com/orchestra-mcp/gen-go/orchestra/plugin/v1"
	"github.com/orchestra-mcp/sdk-go/helpers"
	"github.com/orchestra-mcp/plugin-ai-browser-context/internal/cdp"
	"google.golang.org/protobuf/types/known/structpb"
)

func GetPageContentSchema() *structpb.Struct {
	s, _ := structpb.NewStruct(map[string]any{
		"type": "object",
		"properties": map[string]any{
			"tab_id": map[string]any{
				"type":        "string",
				"description": "Tab ID to get content from. Uses first tab if not specified.",
			},
		},
	})
	return s
}

func GetPageContent() func(context.Context, *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
	return func(ctx context.Context, req *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
		tabID := helpers.GetString(req.Arguments, "tab_id")

		tabs, err := cdp.ListTabs(ctx)
		if err != nil {
			return helpers.ErrorResult("chrome_error", err.Error()), nil
		}

		var matched *cdp.TabInfo
		for i := range tabs {
			if tabID == "" || tabs[i].ID == tabID {
				matched = &tabs[i]
				break
			}
		}

		if matched == nil {
			return helpers.ErrorResult("not_found", fmt.Sprintf("no tab found with ID: %s", tabID)), nil
		}

		msg := fmt.Sprintf(
			"Page content extraction requires Chrome Extension. Install Orchestra Chrome extension to enable page content tools.\n\n**Tab:** %s\n**URL:** %s",
			matched.Title, matched.URL,
		)
		return helpers.TextResult(msg), nil
	}
}
