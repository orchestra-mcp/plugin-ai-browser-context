package tools

import (
	"context"

	pluginv1 "github.com/orchestra-mcp/gen-go/orchestra/plugin/v1"
	"github.com/orchestra-mcp/sdk-go/helpers"
	"google.golang.org/protobuf/types/known/structpb"
)

func GetPageDOMSchema() *structpb.Struct {
	s, _ := structpb.NewStruct(map[string]any{
		"type": "object",
		"properties": map[string]any{
			"tab_id": map[string]any{
				"type":        "string",
				"description": "Tab ID to get DOM from. Uses first tab if not specified.",
			},
			"depth": map[string]any{
				"type":        "number",
				"description": "Maximum DOM depth to traverse. Optional.",
			},
		},
	})
	return s
}

func GetPageDOM() func(context.Context, *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
	return func(ctx context.Context, req *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
		return helpers.TextResult("DOM access requires Orchestra Chrome extension."), nil
	}
}
