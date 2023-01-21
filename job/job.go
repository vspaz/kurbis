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

func hasState(states []State, state State) bool {
	for _, aState := range states {
		if aState == state {
			return true
		}
	}
	return false
}

func isStateAllowed(currentState State, targetState State) bool {
	return hasState(currentStateToAllowedStates[currentState], targetState)
}
