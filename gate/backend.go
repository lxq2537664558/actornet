package gate

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"sync"
)

var (
	backendMgrPID *actor.PID

	backendPIDByMsgID sync.Map

	cachedRoutePIDList []*proto.RouteToPIDACK // 当后台和网关连接还没建立时，路由消息会被缓冲
)

func initBackendManager() {

	// 前台接收 后面的服务器发过来的消息
	backendMgrPID = gateDomain.Spawn(actor.NewTemplate().WithID("backendmgr").WithFunc(func(c actor.Context) {
		switch msg := c.Msg().(type) {
		case *proto.BindClientACK:

			clientSes := acceptorPeer.GetSession(msg.ClientSessionID)
			if clientSes != nil {

				backendPID := actor.NewPID(c.Source().Domain, msg.ID)

				outboundID := makeOutboundID(clientSes.ID())

				outboundPID := gateDomain.Spawn(actor.NewTemplate().WithID(outboundID).WithCreator(newOutboundClient(clientSes)))

				addClient(outboundPID, backendPID, clientSes)

				// 回应客户端
				clientSes.Send(&proto.BindClientACK{})

			} else {
				log.Warnln("BindClinet: client session not found: ", msg.ClientSessionID)
			}
		case *proto.RouteToPIDACK:

			meta := cellnet.MessageMetaByName(msg.MsgName)
			if meta != nil {

				backendPIDByMsgID.LoadOrStore(meta.ID, actor.NewPIDFromProto(msg.Target))
			} else {
				log.Errorf("Route message not found: %s", msg.MsgName)
			}

		}
	}))

	for _, ack := range cachedRoutePIDList {
		backendMgrPID.Tell(ack)
	}

	cachedRoutePIDList = cachedRoutePIDList[0:0]
}

func targetRoutePID(msgid uint32) *actor.PID {

	if raw, ok := backendPIDByMsgID.Load(msgid); ok {
		return raw.(*actor.PID)
	}

	return nil
}

func RouteMessageToPID(msgName string, pid *actor.PID) {

	ack := &proto.RouteToPIDACK{
		MsgName: msgName,
		Target:  pid.ToProto(),
	}

	if backendMgrPID == nil {
		cachedRoutePIDList = append(cachedRoutePIDList, ack)
	} else {
		backendMgrPID.Tell(ack)
	}

}
