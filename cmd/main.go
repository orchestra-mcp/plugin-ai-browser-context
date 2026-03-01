package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/orchestra-mcp/sdk-go/plugin"
	"github.com/orchestra-mcp/plugin-ai-browser-context/internal"
)

func main() {
	builder := plugin.New("ai.browser-context").
		Version("0.1.0").
		Description("Browser context tools for reading tabs, page content, and browser state via Chrome DevTools Protocol").
		Author("Orchestra").
		Binary("ai-browser-context")

	tp := &internal.ToolsPlugin{}
	tp.RegisterTools(builder)

	p := builder.BuildWithTools()
	p.ParseFlags()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		cancel()
	}()

	if err := p.Run(ctx); err != nil {
		log.Fatalf("ai.browser-context: %v", err)
	}
}
