package chatproto

import (
	"github.com/davyxu/actornet/proto"
)

// 聊天消息
// client -> server
type ChatREQ struct {
	To      proto.PID // 玩家要发给目标
	Content string
}

// 聊天消息
// server -> client
type ChatACK struct {
	User    proto.PID
	Name    string
	Content string
}

// 改名
// client <-> server
type RenameACK struct {
	NewName string
}
