
package main

import (
    "github.com/jcburns-vt/http-study/htstud"
)

func main() {
    server := htstud.NewHTTPServer(69)
    server.Start()
}

