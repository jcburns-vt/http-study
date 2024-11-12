
package htstud

import (
    "net"
)

type Placeholder struct {
    // use as placeholders for types which have not yet been identified
    // from the python version
}

type Address struct {
    Host net.IP
    Port int
}

type Request struct {

}

type RequestHandler struct {

}

type BaseServer struct {

    Timeout Placeholder
    AddressFamily Placeholder
    SocketType Placeholder
    AllowReuseAddress Placeholder
    AllowReusePort Placeholder

    RequestHandlerClass RequestHandler
    ServerAddress Address
    Socket Placeholder

    isShutdown Placeholder
    isShutdownRequest Placeholder

    // Overridable functions
    ServerBindFunction func(*BaseServer)
    ServerActivateFunction func(*BaseServer)
    GetRequestFunction func(*BaseServer) (Request, Address)
    HandleTimeoutFunction func(*BaseServer)
    VerifyRequestFunction func(*BaseServer, Request, Address)
    ServerCloseFunction func(*BaseServer)
    ProcessRequestFunction func(*BaseServer, Request, Address)
    ShutdownRequestFunction func(*BaseServer, Request)
    CloseRequestFunction func(*BaseServer, Request)
    ServiceActionsFunction func(*BaseServer)
    HandleErrorFunction func(*BaseServer)
}

func NewBaseServer (serverAddress Address,
    requestHandlerClass RequestHandler) *BaseServer {

    return &BaseServer{
        ServerAddress: serverAddress,
        RequestHandlerClass: requestHandlerClass,
        // TODO: self.__is_shut_down = threading.Event()
        // TODO: self.__shutdown_request = False
    }
}

func (b *BaseServer) ServerBind() {
    // Overridable
    b.ServerBindFunction(b)
}

func (b *BaseServer) ServerActivate() {
    // Overridable
    b.ServerActivateFunction(b)
}

// ServeForever handles one request at a time until shutdown.
//
// Polls for shutdown every pollInterval seconds. Ignores
// b.Timeout. If you need to do periodic tasks, do them
// in another thread.
func (b *BaseServer) ServeForever(pollInterval float32) {
    // TODO: self.__is_shut_down.clear()
    // TODO: rest of server forever
}

func (b *BaseServer) GetRequest() (Request, Address){
    // Overridable
    request, clientAddress := b.GetRequestFunction(b)
    return request, clientAddress
}

func (b *BaseServer) HandleTimeout() {
    // Overridable
    b.HandleTimeoutFunction(b)
}

func (b *BaseServer) VerifyRequest(request Request, clientAddress Address) {
    // Overridable
    b.VerifyRequestFunction(b, request, clientAddress)
}

func (b *BaseServer) ServerClose() {
    // Overridable
    b.ServerCloseFunction(b)
}

func (b *BaseServer) ProcessRequest(request Request, clientAddress Address) {
    // Overridable
    b.ProcessRequestFunction(b, request, clientAddress)
}

func (b *BaseServer) ShutdownRequest(request Request) {
    // Overridable
    b.ShutdownRequestFunction(b, request)
}

func (b *BaseServer) CloseRequest(request Request) {
    // Overridable
    b.CloseRequestFunction(b, request)
}

func (b *BaseServer) ServiceActions() {
    // Overridable
    b.ServiceActionsFunction(b)
}

func (b *BaseServer) HandleError() {
    // Overridable
    b.HandleErrorFunction(b)
}

func (b *BaseServer) HandleRequest() {

}

func (b *BaseServer) Fileno() int {
    return 0
}



