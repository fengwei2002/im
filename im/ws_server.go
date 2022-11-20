package im

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

/*
ws_server 用于建立一个 websocket 的连接
*/

// WsServerOptions 代表了可以附加的选项 Host 端口号，写的截止时间，读的截止时间
type WsServerOptions struct {
	Host          string
	Port          int
	ReadDeadLine  time.Duration
	WriteDeadLine time.Duration
}

// WsServe 定义了具体的服务结构体
type WsServer struct {
	options  *WsServerOptions
	upgrader websocket.Upgrader
}

// NewWsServer options can be nil, use default value when nil.
func NewWsServer(options *WsServerOptions) *WsServer {
	if options == nil {
		options = &WsServerOptions{
			Host:          "0.0.0.0",
			Port:          8080,
			ReadDeadLine:  12 * time.Minute,
			WriteDeadLine: 12 * time.Minute,
		}
	}
	ws := new(WsServer)
	ws.options = options
	ws.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 65536,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return ws
}

// HandleWebSocketRequest 用于处理 websocket 的请求
func (ws *WsServer) handleWebSocketRequest(writer http.ResponseWriter, request *http.Request) {

	conn, err := ws.upgrader.Upgrade(writer, request, nil)
	if err != nil {

		return
	}

	con := NewWsConnection(conn, ws.options)
	NewClient(con).Run()
}

// Start 启动一个 WebSocket 的服务
func (ws *WsServer) Start() {

	http.HandleFunc("/ws", ws.handleWebSocketRequest)

	addr := fmt.Sprintf("%s:%d", ws.options.Host, ws.options.Port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}

}
