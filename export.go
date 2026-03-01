package aibrowsercontext

import (
	"github.com/orchestra-mcp/plugin-ai-browser-context/internal"
	"github.com/orchestra-mcp/sdk-go/plugin"
)

// Register adds all browser context tools to the builder.
func Register(builder *plugin.PluginBuilder) {
	tp := &internal.ToolsPlugin{}
	tp.RegisterTools(builder)
}
