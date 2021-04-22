package socket

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

// Connection 支持并发调用的websocket客户端
type Connection struct {
	webConn   *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	isClose   bool
}

func InitConnection(webConn *websocket.Conn) *Connection {
	conn := &Connection{
		webConn:   webConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte),
		mutex:     sync.Mutex{},
		isClose:   false,
	}
	//循环读
	go conn.readLoop()

	//循环写
	go conn.writeLoop()

	return conn
}

// 对外API
// ReadMessage 从inchan中读取数据返回给上传调用者
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return data, err

}

// WriteMessage 将数据写入到outChan中
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return err
}

// Close 关闭连接 并发安全，可重入
func (conn *Connection) Close() {
	// 1.关闭底层websocket
	conn.webConn.Close()

	// 2. 标记关闭,并关闭底层通道
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	if !conn.isClose {
		conn.isClose = true
	}
	close(conn.closeChan)
}

// 内部封装
// readLoop 读取web数据到 inchan中
func (conn *Connection) readLoop() {
	// 避免在for中创建数据
	var (
		data []byte
		err  error
	)
	for {
		_, data, err = conn.webConn.ReadMessage()
		if err != nil {
			conn.Close()
		}
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			conn.Close()

		}
	}
}

// writeLoop 从outchan中读出数据返回web端
func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			conn.Close()
		}

		if err = conn.webConn.WriteMessage(websocket.TextMessage, data); err != nil {
			conn.Close()
		}
	}
}
