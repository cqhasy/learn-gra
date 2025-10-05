package main

import (
	"context"
	"log"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	ctx := context.Background()

	client := mcp.NewClient(&mcp.Implementation{Name: "mcp-client", Version: "v1.0.0"}, nil)

	// 启动服务端 server.go
	transport := &mcp.CommandTransport{Command: exec.Command(".\\sever.exe")}
	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	params := &mcp.CallToolParams{
		Name:      "greet",
		Arguments: map[string]any{"name": "you"},
	}
	res, err := session.CallTool(ctx, params)
	if err != nil {
		log.Fatalf("CallTool failed: %v", err)
	}
	if res.IsError {
		log.Fatal("tool failed")
	}
	for _, c := range res.Content {
		log.Print(c.(*mcp.TextContent).Text)
	}
}
