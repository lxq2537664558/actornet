package main

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/examples/chat/proto"
	"github.com/davyxu/actornet/gate"
	"github.com/davyxu/actornet/proto"
)

type user struct {
	actor.LocalProcess
	name      string
	clientpid *actor.PID
}

func (self *user) OnRecv(c actor.Context) {

	switch msg := c.Msg().(type) {
	case *proto.Start:
		self.name = "noname"
	case *chatproto.PublicChatREQ:

		log.Debugln("chat", c.Source())

		self.ParentPID().Broadcast(&chatproto.PublicChatACK{
			User:    self.PID().ToProto(),
			Name:    self.name,
			Content: msg.Content,
		})

	case *chatproto.RenameACK:

		log.Debugf("[%s] rename '%s' -> '%s'", c.Self().String(), self.name, msg.NewName)

		self.name = msg.NewName
	case *chatproto.GetNameREQ:

		c.Reply(&chatproto.GetNameACK{
			Name: self.name,
		})
	case *chatproto.PublicChatACK:
		self.clientpid.TellBySender(msg, c.Self())
	}
}

func newUser(clientSesID int64) *user {
	return &user{
		clientpid: gate.MakeOutboundPID(clientSesID),
	}

}
