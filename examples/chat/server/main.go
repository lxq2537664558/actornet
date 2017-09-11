package main

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/gate"
	"github.com/davyxu/golog"
)

var log *golog.Logger = golog.New("main")

var serverDomain *actor.Domain

func main() {

	actor.StartSystem()

	serverDomain = actor.CreateDomain("server")

	lobbyPID := serverDomain.Spawn(actor.NewTemplate().WithID("lobby").WithInstance(newLobby()))

	gate.Listen("127.0.0.1:8081", lobbyPID)

	actor.LoopSystem()
}
