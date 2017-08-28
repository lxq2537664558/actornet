package main

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/examples/chat/proto"
	"github.com/davyxu/actornet/gate"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/golog"
)

var log *golog.Logger = golog.New("main")

type user struct {
	actor.LocalProcess
	name      string
	clientpid *actor.PID
}

func (self *user) OnRecv(c actor.Context) {

	switch msg := c.Msg().(type) {
	case *proto.Start:
		self.name = "noname"
	case *chatproto.ChatREQ:

		log.Debugln("chat", c.Source())

		self.ParentPID().Broadcast(&chatproto.ChatACK{
			User:    self.PID().ToProto(),
			Name:    self.name,
			Content: msg.Content,
		})

	case *chatproto.RenameACK:

		log.Debugf("[%s] rename '%s' -> '%s'", c.Self().String(), self.name, msg.NewName)

		self.name = msg.NewName

	case *chatproto.ChatACK:
		self.clientpid.Tell(msg)
	}
}

func newUser(clientSesID int64) actor.ActorCreator {
	return func() actor.Actor {
		return &user{
			clientpid: gate.MakeOutboundPID(clientSesID),
		}
	}

}

func main() {

	actor.StartSystem()

	domain := actor.CreateDomain("server")

	lobbyPID := domain.Spawn(actor.NewTemplate().WithID("lobby").WithFunc(func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.BindClientREQ:

			pid := domain.Spawn(actor.NewTemplate().WithCreator(newUser(msg.ClientSessionID)).WithParent(c.Self()))

			c.Reply(&proto.BindClientACK{
				ClientSessionID: msg.ClientSessionID,
				ID:              pid.Id,
			})

		}

	}))

	gate.Listen("127.0.0.1:8081", lobbyPID)

	actor.LoopSystem()
}
