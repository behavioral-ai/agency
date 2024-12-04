package ops

import (
	"fmt"
	"github.com/advanced-go/resiliency/guidance"
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/test"
)

var (
	shutdownMsg   = messaging.NewControlMessage("", "", messaging.ShutdownEvent)
	dataChangeMsg = messaging.NewControlMessage("", "", messaging.DataChangeEvent)
	startMsg      = messaging.NewControlMessage("", "", startAgentsEvent)
	stopMsg       = messaging.NewControlMessage("", "", stopAgentsEvent)
)

func init() {
	dataChangeMsg.SetContent(guidance.ContentTypeCalendar, guidance.NewProcessingCalendar())
}

func officer(origin core.Origin, handler messaging.OpsAgent, dispatcher messaging.TraceDispatcher) messaging.OpsAgent {
	return test.NewAgent("officer:" + origin.Region)
}

func ExampleEmissary() {
	ch := make(chan struct{})
	traceDispatcher := messaging.NewTraceDispatcher(nil, "")
	agent := newAgent(Class, messaging.OutputErrorNotifier, test.DefaultTracer, traceDispatcher, newTestDispatcher())

	go func() {
		go emissaryAttend(agent, officer)
		agent.Message(dataChangeMsg)
		agent.Message(startMsg)
		agent.Message(stopMsg)
		agent.Message(shutdownMsg)
		fmt.Printf("test: emissaryAttend() -> [finalized:%v]\n", agent.IsFinalized())
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//test: Trace() -> agency-ops : emissary event:data-change Broadcast() -> calendar data change event
	//test: dispatch(event:start-agents) -> [count>0:true]
	//test: dispatch(event:stop-agents) -> [count:0]
	//test: emissaryAttend() -> [finalized:true]

}
