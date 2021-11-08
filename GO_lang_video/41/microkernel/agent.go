package microkernel

import "context"

type EventReceiver interface {
	OnEvent(evt Event)
}

type Collector interface {
	Init(evtReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destory() error
}

type Agent struct {
	colletors map[string]Collector
	evtBuf    chan Event
	cancel    context.CancelFunc
	ctx       context.Context
	state     int
}

func (agt *Agent) EventProcessGroutine() {

}

func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}
