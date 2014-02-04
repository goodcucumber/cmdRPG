package Event

type Event struct {
	X     int
	Y     int
	Run   func(e *Event) int
	Valid bool
	Var   []int
}
