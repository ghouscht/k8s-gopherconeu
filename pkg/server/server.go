package server

import (
	"context"
	"net"
	"net/http"
)

// WebServer is ...
type WebServer struct {
	http.Server

	address string
}

// New returns a new server.
func New(host, port string, h http.Handler) *WebServer {
	var ws = new(WebServer)

	ws.Addr = net.JoinHostPort(host, port)
	ws.Handler = h

	return ws
}

// Start starts the server.
func (ws *WebServer) Start() error {
	return ws.ListenAndServe()
}

// Stop stops the server gracefully.
func (ws *WebServer) Stop() error {
	return ws.Shutdown(context.Background())
}
