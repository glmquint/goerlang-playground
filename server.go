package main

import (
	"fmt"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

type demo struct {
	gen.Server
}

func (d *demo) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	fmt.Printf("[%s] HandleCast: %#v\n", process.Name(), message)
	switch message {
	case etf.Atom("stop"):
		return gen.ServerStatusStopWithReason("stop they said")
	}
	return gen.ServerStatusOK
}

func (d *demo) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	fmt.Printf("[%s] HandleCall: %#v, From: %s\n", process.Name(), message, from.Pid)

	switch message.(type) {
	case etf.Atom:
		return "hello", gen.ServerStatusOK

	default:
		return message, gen.ServerStatusOK
	}
}

// HandleInfo
func (d *demo) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	fmt.Printf("[%s] HandleInfo: %#v\n", process.Name(), message)
	if message == etf.Atom("ping") {
		mailbox := "mailbox"
		dest_node := "server@127.0.0.1"
		err := process.Send(gen.ProcessID{Name: mailbox, Node: dest_node}, "pong")
		if err != nil {
			panic(err)
		}
	}
	return gen.ServerStatusOK
}

func (d *demo) Terminate(process *gen.ServerProcess, reason string) {
	fmt.Printf("[%s] Terminating process with reason %q", process.Name(), reason)
}
