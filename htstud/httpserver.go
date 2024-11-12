
package htstud

import (
    "fmt"
)

type HTTPServer struct {
    Value int
}

func NewHTTPServer(value int) *HTTPServer {
    return &HTTPServer{
        Value: value,
    }
}

func (h *HTTPServer) Start() {
    fmt.Println(h.Value)
}
