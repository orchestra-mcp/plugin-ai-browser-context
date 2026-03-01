package tools

// Tests for the ai.browser-context plugin tool handlers.
//
// All tools that call cdp.ListTabs will return chrome_error when Chrome is not
// running with remote debugging (which is always the case in CI). We test:
//   1. Validation errors (missing required args) — no Chrome needed.
//   2. Chrome-unavailable paths — cdp.ListTabs fails → chrome_error response.
//   3. Stub tools (execute_script, navigate_to, etc.) that return setup hints.

import (
	"context"
	"testing"

	pluginv1 "github.com/orchestra-mcp/gen-go/orchestra/plugin/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

// ---------- helpers ----------

func callTool(t *testing.T, handler func(context.Context, *pluginv1.ToolRequest) (*pluginv1.ToolResponse, error), args map[string]any) *pluginv1.ToolResponse {
	t.Helper()
	var s *structpb.Struct
	if args != nil {
		var err error
		s, err = structpb.NewStruct(args)
		if err != nil {
			t.Fatalf("NewStruct: %v", err)
		}
	}
	resp, err := handler(context.Background(), &pluginv1.ToolRequest{Arguments: s})
	if err != nil {
		t.Fatalf("handler returned Go error: %v", err)
	}
	return resp
}

func isError(resp *pluginv1.ToolResponse) bool {
	return resp != nil && !resp.Success
}

func errorCode(resp *pluginv1.ToolResponse) string {
	if resp == nil {
		return ""
	}
	return resp.GetErrorCode()
}

// ---------- get_open_tabs ----------

func TestGetOpenTabs_NoChromeRunning(t *testing.T) {
	// Chrome is not running in CI — cdp.ListTabs fails → chrome_error.
	resp := callTool(t, GetOpenTabs(), map[string]any{})
	if !isError(resp) {
		t.Skip("Chrome appears to be running; skipping chrome_error path")
	}
	if errorCode(resp) != "chrome_error" {
		t.Errorf("expected chrome_error, got %q", errorCode(resp))
	}
}

// ---------- get_page_content ----------

func TestGetPageContent_NoChromeRunning(t *testing.T) {
	resp := callTool(t, GetPageContent(), map[string]any{})
	if !isError(resp) {
		t.Skip("Chrome appears to be running; skipping chrome_error path")
	}
	if errorCode(resp) != "chrome_error" {
		t.Errorf("expected chrome_error, got %q", errorCode(resp))
	}
}

func TestGetPageContent_WithTabID_NoChromeRunning(t *testing.T) {
	resp := callTool(t, GetPageContent(), map[string]any{"tab_id": "abc123"})
	// Either chrome_error (no Chrome) or not_found (Chrome running but no such tab).
	if !isError(resp) {
		t.Skip("Chrome appears to be running with matching tab; skipping error paths")
	}
	code := errorCode(resp)
	if code != "chrome_error" && code != "not_found" {
		t.Errorf("expected chrome_error or not_found, got %q", code)
	}
}

// ---------- get_page_dom ----------

func TestGetPageDOM_NoChromeRunning(t *testing.T) {
	resp := callTool(t, GetPageDOM(), map[string]any{})
	if !isError(resp) {
		t.Skip("Chrome appears to be running; skipping chrome_error path")
	}
	if errorCode(resp) != "chrome_error" {
		t.Errorf("expected chrome_error, got %q", errorCode(resp))
	}
}

// ---------- get_selected_text ----------

func TestGetSelectedText_NoArgs(t *testing.T) {
	// get_selected_text may call cdp or return a stub — either is acceptable.
	resp := callTool(t, GetSelectedText(), map[string]any{})
	_ = resp
}

// ---------- get_page_screenshot ----------

func TestGetPageScreenshot_NoChromeRunning(t *testing.T) {
	resp := callTool(t, GetPageScreenshot(), map[string]any{})
	if !isError(resp) {
		t.Skip("Chrome appears to be running; skipping chrome_error path")
	}
	code := errorCode(resp)
	if code != "chrome_error" && code != "not_found" {
		t.Errorf("expected chrome_error or not_found, got %q", code)
	}
}

// ---------- navigate_to ----------

func TestNavigateTo_MissingURL(t *testing.T) {
	resp := callTool(t, NavigateTo(), map[string]any{})
	if !isError(resp) {
		t.Error("expected validation_error for missing url")
	}
	if errorCode(resp) != "validation_error" {
		t.Errorf("expected validation_error, got %q", errorCode(resp))
	}
}

func TestNavigateTo_ValidURL(t *testing.T) {
	resp := callTool(t, NavigateTo(), map[string]any{"url": "https://example.com"})
	// Returns a setup-hint message (no Chrome needed).
	if isError(resp) {
		t.Errorf("unexpected error: %s", errorCode(resp))
	}
}

// ---------- execute_script ----------

func TestExecuteScript_MissingScript(t *testing.T) {
	// execute_script currently ignores args and returns a stub message.
	// Still verify it doesn't panic and returns a response.
	resp := callTool(t, ExecuteScript(), map[string]any{})
	_ = resp
}

func TestExecuteScript_WithScript(t *testing.T) {
	resp := callTool(t, ExecuteScript(), map[string]any{"script": "document.title"})
	// Returns a setup-hint message — not an error.
	if isError(resp) {
		t.Errorf("unexpected error: %s", errorCode(resp))
	}
}
