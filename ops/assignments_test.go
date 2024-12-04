package ops

import (
	"fmt"
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/test"
)

func ExampleInitialize_Error() {
	notifier := test.NewNotifier()
	agent := newAgent(Class, notifier, test.DefaultTracer, nil, nil)

	createAssignments(agent, nil)
	fmt.Printf("test: initialize() -> [status:%v]\n", notifier.Status())

	notifier.Reset()
	createAssignments(agent, func(origin core.Origin, handler messaging.OpsAgent, dispatcher messaging.TraceDispatcher) messaging.OpsAgent {
		return test.NewAgent("")
	})
	fmt.Printf("test: initialize() -> [status:%v]\n", notifier.Status())

	notifier.Reset()
	a := test.NewAgent("agent:test")
	err := agent.caseOfficers.Register(a)
	if err != nil {
		fmt.Printf("test: Register() -> [err:%v]\n", err)
	}
	createAssignments(agent, func(origin core.Origin, handler messaging.OpsAgent, dispatcher messaging.TraceDispatcher) messaging.OpsAgent {
		return a
	})
	fmt.Printf("test: initialize() -> [status:%v]\n", notifier.Status())

	//Output:
	//test: initialize() -> [status:Invalid Argument [error: initialize newAgent is nil]]
	//test: initialize() -> [status:Invalid Argument [error: exchange.Register() agent Uri is empty]]
	//test: initialize() -> [status:Invalid Argument [error: exchange.Register() agent already exists: [agent:test]]]

}
