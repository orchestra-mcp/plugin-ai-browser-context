package tools

import (
	"context"
	"fmt"

	pluginv1 "github.com/orchestra-mcp/gen-go/orchestra/plugin/v1"
	"github.com/orchestra-mcp/sdk-go/helpers"
	"google.golang.org/protobuf/types/known/structpb"
)

func NavigateToSchema() *structpb.Struct {
	s, _ := structpb.NewStruct(map[string]any{
		"type": "object",
		"properties": map[string]any{
			"url": map[string]any{
				"type":        "string",
				"description": "URL to navigate to",
			},
			"tab_id": map[string]any{
				"type":        "string",
				"description": "Tab ID to navigate. Uses first tab if not specified.",
			},
		},
		"required": []any{"url"},
	})
	return s
}

func NavigateTo() func(context.Context, *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
	return func(ctx context.Context, req *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
		if err := helpers.ValidateRequired(req.Arguments, "url"); err != nil {
			return helpers.ErrorResult("validation_error", err.Error()), nil
		}

		url := helpers.GetString(req.Arguments, "url")
		msg := fmt.Sprintf("Navigation requires Orchestra Chrome extension. Open %s manually or install the extension.", url)
		return helpers.TextResult(msg), nil
	}
}
