package job

type State int

const (
	Pending State = iota
	Scheduled
	Completed
	Running
	Failed
)

var currentStateToAllowedStates = map[State][]State{
	Pending:   []State{Scheduled},
	Scheduled: []State{Scheduled, Running, Failed},
	Running:   []State{Running, Completed, Failed},
	Completed: []State{},
	Failed:    []State{},
}
