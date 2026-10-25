package main

import (
	"fmt"

	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer("MCP Name", "1.0.0")

	tool := mcp.NewTool(
		"Name of my tool",
		mcp.WithDescription("My description"),
		mcp.WithString("termo", mcp.Required(), mcp.Description("Descrição da propriedade no schema JSON")),
	)

	s.AddTool(tool, funcaoHandler)

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server with erro: %v\n", err)
	}
}

func funcaoHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	result, err := request.RequireString("termo")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(result), nil
}
