package tools

import (
	"context"

	pluginv1 "github.com/orchestra-mcp/gen-go/orchestra/plugin/v1"
	"github.com/orchestra-mcp/sdk-go/helpers"
	"google.golang.org/protobuf/types/known/structpb"
)

func GetPageScreenshotSchema() *structpb.Struct {
	s, _ := structpb.NewStruct(map[string]any{
		"type": "object",
		"properties": map[string]any{
			"tab_id": map[string]any{
				"type":        "string",
				"description": "Tab ID to screenshot. Uses first tab if not specified.",
			},
			"output_path": map[string]any{
				"type":        "string",
				"description": "File path to save the screenshot. Optional.",
			},
		},
	})
	return s
}

func GetPageScreenshot() func(context.Context, *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
	return func(ctx context.Context, req *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error) {
		return helpers.TextResult("Page screenshot requires Orchestra Chrome extension."), nil
	}
}
