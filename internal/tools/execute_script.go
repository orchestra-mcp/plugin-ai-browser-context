package tools

import (
	"context"

	pluginv1 "github.com/orchestra-mcp/gen-go/orchestra/plugin/v1"
	"github.com/orchestra-mcp/sdk-go/helpers"
	"google.golang.org/protobuf/types/known/structpb"
)

func ExecuteScriptSchema() *structpb.Struct {
	s, _ := structpb.NewStruct(map[string]any{
		"type": "object",
		"properties": map[string]any{
			"script": map[string]any{
				"type":        "string",
				"description": "JavaScript code to execute in the page context",
			},
			"tab_id": map[string]any{
				"type":        "string",
				"description": "Tab ID to execute the script in. Uses first tab if not specified.",
			},
		},
		"required": []any{"script"},
	})
	return s
}

func ExecuteScript() func(context.Context, *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
	return func(ctx context.Context, req *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
		return helpers.TextResult("Script execution requires Orchestra Chrome extension."), nil
	}
}
