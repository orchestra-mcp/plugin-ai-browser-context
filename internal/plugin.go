package internal

import (
	"github.com/orchestra-mcp/sdk-go/plugin"
	"github.com/orchestra-mcp/plugin-ai-browser-context/internal/tools"
)

// ToolsPlugin registers all ai.browser-context tools with the plugin builder.
type ToolsPlugin struct{}

// RegisterTools registers all 7 browser context tools on the given plugin builder.
func (tp *ToolsPlugin) RegisterTools(builder *plugin.PluginBuilder) {
	builder.RegisterTool("get_open_tabs",
		"List all open browser tabs with their titles and URLs",
		tools.GetOpenTabsSchema(), tools.GetOpenTabs())

	builder.RegisterTool("get_page_content",
		"Get the text content of the current or specified browser tab",
		tools.GetPageContentSchema(), tools.GetPageContent())

	builder.RegisterTool("navigate_to",
		"Navigate the browser to a specified URL",
		tools.NavigateToSchema(), tools.NavigateTo())

	builder.RegisterTool("get_selected_text",
		"Get the currently selected text in the browser",
		tools.GetSelectedTextSchema(), tools.GetSelectedText())

	builder.RegisterTool("get_page_dom",
		"Get the DOM structure of the current or specified browser tab",
		tools.GetPageDOMSchema(), tools.GetPageDOM())

	builder.RegisterTool("get_page_screenshot",
		"Take a screenshot of the current or specified browser tab",
		tools.GetPageScreenshotSchema(), tools.GetPageScreenshot())

	builder.RegisterTool("execute_script",
		"Execute JavaScript in the current or specified browser tab",
		tools.ExecuteScriptSchema(), tools.ExecuteScript())
}
