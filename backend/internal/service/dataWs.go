package service

import (
	"sync"

	"github.com/gorilla/websocket"
)

var (
	UserConnections = make(map[int][]*websocket.Conn)
	UserConnMu      sync.RWMutex

	ConvSubscribers = make(map[int][]int)
	ConvSubMu       sync.RWMutex
)
