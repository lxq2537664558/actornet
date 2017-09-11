package main

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/examples/chat/proto"
	"github.com/davyxu/actornet/gate"
	"github.com/davyxu/actornet/proto"
)

type Lobby struct {
	actor.LocalProcess

	userByOuboundID map[string]*actor.PID
}

func (self *Lobby) addUser(outboundidstr string, userPID *actor.PID) {
	self.userByOuboundID[outboundidstr] = userPID
}

func (self *Lobby) getUser(outboundid string) *actor.PID {
	if v, ok := self.userByOuboundID[outboundid]; ok {
		return v
	}

	return nil
}

func (self *Lobby) OnRecv(c actor.Context) {

	switch msg := c.Msg().(type) {
	case *proto.Start:
		gate.RouteMessageToPID("chatproto.PublicChatREQ", self.PID())
	case *proto.BindClientREQ:

		u := newUser(msg.ClientSessionID)

		pid := serverDomain.Spawn(actor.NewTemplate().WithInstance(u).WithParent(c.Self()))

		self.addUser(u.clientpid.String(), pid)

		c.Reply(&proto.BindClientACK{
			ClientSessionID: msg.ClientSessionID,
			ID:              pid.Id,
		})
	case *chatproto.PublicChatREQ:

		log.Debugln("chat", c.Source())

		u := self.getUser(c.Source().String())
		if u != nil {
			name := u.Call(&chatproto.GetNameREQ{}, c.Self()).(*chatproto.GetNameACK).Name

			// TODO 封装起来，然user中不用对每个消息进行透传
			c.Self().Broadcast(&chatproto.PublicChatACK{
				User:    u.ToProto(),
				Name:    name,
				Content: msg.Content,
			})

		}

	}
}

func newLobby() actor.Actor {
	return &Lobby{
		userByOuboundID: make(map[string]*actor.PID),
	}
}
