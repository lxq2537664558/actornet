package gate

import (
	"fmt"
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

var (
	acceptorPeer cellnet.Peer

	// 后台的辅助actor
	backendAssitPID *actor.PID

	gateDomain *actor.Domain
)

const gateDomainName = "gate"

func Listen(address string, backendAssit *actor.PID) {

	gateDomain = actor.CreateDomain(gateDomainName)

	backendAssitPID = backendAssit

	acceptorPeer = socket.NewAcceptor(nil)

	// 添加客户端消息侦听
	acceptorPeer.AddChainRecv(
		cellnet.NewHandlerChain(
			newInboundHandler(),
		),
	)

	// 客户端断开
	cellnet.RegisterMessage(acceptorPeer, "coredef.SessionClosed", func(ev *cellnet.Event) {

		// TODO 通知到后台
		pid := removeClient(ev.Ses)
		if pid != nil {
			gateDomain.Kill(pid)
		}
	})

	acceptorPeer.Start(address)

	initBackendManager()
}

func makeOutboundID(clientSessionID int64) string {
	return fmt.Sprintf("sid:%d", clientSessionID)
}

func MakeOutboundPID(clientSessionID int64) *actor.PID {
	return actor.NewPID(gateDomainName, makeOutboundID(clientSessionID))
}
