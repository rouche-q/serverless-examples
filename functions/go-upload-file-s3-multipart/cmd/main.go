package main

import (
	handler "file-read"

	"github.com/scaleway/serverless-functions-go/local"
)

func main() {
	// Replace "Handle" with your function handler name if necessary
	local.ServeHandler(handler.Handle, local.WithPort(8080))
}
