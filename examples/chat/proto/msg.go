package chatproto

import (
	"github.com/davyxu/actornet/proto"
)

// 聊天消息
// client -> 'lobby'
type PublicChatREQ struct {
	Content string
}

// 聊天消息
// 'lobby' -> client
type PublicChatACK struct {
	User    proto.PID
	Name    string
	Content string
}

// 改名
// client <-> server
type RenameACK struct {
	NewName string
}

type GetNameREQ struct {
}

type GetNameACK struct {
	Name string
}
