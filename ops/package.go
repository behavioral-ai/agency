package ops

import (
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/log/timeseries"
	"github.com/behavioral-ai/resiliency/guidance"
)

const (
	PkgPath = "github/behavioral-ai/agency/ops"
)

var (
	westOrigin    = core.Origin{Region: "us-west", Host: "www.west-host1.com"}
	centralOrigin = core.Origin{Region: "us-central", Host: "www.central-host1.com"}
)

func StartAgents() {
	opsAgent.Message(messaging.NewControlMessage(opsAgent.Uri(), opsAgent.Uri(), startAgentsEvent))
	timeseries.Reset()
}

func StopAgents() {
	opsAgent.Message(messaging.NewControlMessage(opsAgent.Uri(), opsAgent.Uri(), stopAgentsEvent))
}

func SendCalendar() {
	msg := messaging.NewControlMessage(opsAgent.Uri(), opsAgent.Uri(), messaging.DataChangeEvent)
	msg.SetContent(guidance.ContentTypeCalendar, guidance.NewProcessingCalendar())
	opsAgent.Message(msg)
}
