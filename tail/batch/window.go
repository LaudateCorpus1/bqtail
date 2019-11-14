package batch

import (
	"bqtail/tail/config"
	"bqtail/tail/contract"
	"fmt"
	"github.com/viant/afs/url"
	"time"
)

type BatchedWindow struct {
	*Window
	BatchingEventID string
}

//Window represent batching window
type Window struct {
	URL           string    `json:",omitempty"`
	Start         time.Time `json:",omitempty"`
	LostOwnership bool      `json:",omitempty"`
	BaseURL       string    `json:",omitempty"`
	End           time.Time `json:",omitempty"`
	SourceCreated time.Time `json:",omitempty"`
	EventTime     time.Time `json:",omitempty"`
	EventID       string    `json:",omitempty"`
	ScheduleURL       string    `json:",omitempty"`
	Datafiles     []*Datafile `json:",omitempty"`
}

//NewWindow create a stage batch window
func NewWindow(baseURL string, request *contract.Request, windowStarted time.Time, route *config.Rule, sourceCreate time.Time, scheduleURL string) *Window {
	end := windowStarted.Add(route.Batch.Window.Duration)
	return &Window{
		BaseURL:       baseURL,
		SourceCreated: sourceCreate,
		URL:           url.Join(baseURL, fmt.Sprintf("%v%v", end.UnixNano(), windowExtension)),
		EventID:       request.EventID,
		EventTime:     request.Started,
		Start:         windowStarted,
		ScheduleURL: scheduleURL,
		End:           end,
	}
}
