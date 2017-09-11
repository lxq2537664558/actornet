package proto

import (
	_ "github.com/davyxu/cellnet/codec/binary"
)

type PID struct {
	Domain string
	Id     string
}

func (self PID) IsValid() bool {
	return self.Domain != "" || self.Id != ""
}

// ============================================
// Actor
// ============================================

// 一个Actor启动时
// any -> any
type Start struct {
}

// 一个Actor停止时
// any -> any
type Stop struct {
}

// 整个物理进程退出
// any -> localhost/system
type SystemExit struct {
	Code int32
}

// ============================================
// Nexus
// ============================================

// 进程互联通道打开
// nexus -> any
type NexusOpen struct {
	Domain string
}

// 进程互联通道关闭
// nexus -> any
type NexusClose struct {
	Domain string
}

// 路由到另外一个进程
type RouteACK struct {
	Source PID
	Target PID

	MsgID   uint32
	MsgData []byte
	CallID  int64
}

// 领域标识
type DomainSyncACK struct {
	DomainNames []string
}

// ============================================
// Gate
// ============================================

// 客户端请求后台服务器绑定
// client -> gate -> gate_assit
type BindClientREQ struct {
	ClientSessionID int64 // 网关上的id (透传)

}

// gate_assit -> gate_receiptor -> client
type BindClientACK struct {
	ClientSessionID int64  // 网关上的id (透传)
	ID              string // 用户后台的网关用户pid.ID
}

// 将消息固定转发到某个PID
// backend -> gate
type RouteToPIDACK struct {
	Target  PID
	MsgName string
}
